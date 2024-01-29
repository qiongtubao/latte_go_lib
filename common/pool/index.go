package pool

import "errors"
var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")
)
type Pool iterface {
	Get() (interface{}, error)
	Put(interface{}) error
	Close(interface{}) error
	Release()
	Len() int
}
