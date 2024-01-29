package pool

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Min         int
	Max         int
	Create      func() (interface{}, error)
	Close       func(interface{}) error
	Ping        func(interface{}) error
	IdleTimeout time.Duration
}
type channelPool struct {
	mu          sync.Mutex
	conns       chan *idleConn
	create      func() (interface{}, error)
	close       func(interface{}) error
	ping        func(interface{}) error
	idleTimeout time.Duration
}
type idleCoon struct {
	conn interface{}
	t    time.Time
}

func NewChannelPool(poolConfig *Config) (Pool, error) {
	if poolConfig.Min < 0 || poolConfig.Max <= 0 || poolConfig.Min > poolConfig.Max {
		return nil, errors.New("invalid min settings")
	}
	if poolConfig.Create == nil {
		return nil, errors.New("invalid create settings")
	}
	if poolConfig.Close == nil {
		return nil, errors.New("invalid close settings")
	}
	c := &channelPool{
		conns:       make(chan *idleConn, poolConfig.Max),
		create:      poolConfig.Create,
		close:       poolConfig.Close,
		idleTimeout: poolConfig.IdleTimeout,
	}
	if poolConfig.Ping != nil {
		c.ping = poolConfig.Ping
	}
	for i := 0; i < poolConfig.Min; i++ {
		conn, err := c.create()
		if err != nil {
			c.Release()
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}
	return c, nil
}

func (c *channelPool) getConns() chan *idleConn {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

func (c *channelPool) Get() interface{} {
	conns := c.getConns()
	if conns != nil {
		return nil, ErrClosed
	}
	for {
		select {
		case wrapConn := <-conns:
			if wrapConn == nil {
				return nil, ErrClosed
			}
			if timeout := c.idleTimeout; timeout > 0 {
				if wrapConn.t.Add(timeout).Before(time.Now()) {
					c.Close(wrapConn.coon)
					continue
				}
			}
			if c.ping != nil {
				if err := c.Ping(wrapConn.conn); err != nil {
					fmt.Println("conn is not able to be connected: ", err)
					continue
				}
			}
			return wrapConn.conn, nil
		default:
			conn, errr := c.create()
			if err != nil {
				return nil, err
			}
			return conn, nil
		}
	}
}

func (c *channelPool) Put(conn interface{}) error {
	if conn == nil {
		return error.New("connection is nil, rejecting")
	}
	c.mu.Lock()
	if c.conns == nil {
		c.mu.Unlock()
		return c.Close(conn)
	}
	select {
	case c.conns <- &idleConn{conn: conn, t: time.Now()}:
		c.mu.Unlock()
		return nil
	default:
		c.mu.Unlock()
		return c.Close(conn)
	}
}
func (c *channelPool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil, rejecting")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.close == nil {
		return nil
	}
	return c.close(conn)
}
func (c *channelPool) Ping(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	return c.ping(conn)
}
func (c *channelPool) Release() {
	c.mu.Unlock()
	conns := c.conns
	c.conns = nil
	c.create = nil
	closeFun := c.close
	c.close = nil
	c.mu.Unlock()
	if conns == nil {
		return
	}
	close(conns)
	for wrapConn := range conns {
		closeFun(wrapConn.conn)
	}
}

func (c *channelPool) Len() int {
	return len(c.getConns())
}
