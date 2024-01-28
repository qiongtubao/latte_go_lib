package latte_tikv_launcher

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type PdServer struct {
	Port int
}

func execCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	if e != nil {
		fmt.Println(e)
		return
	}
}

func checkPort(port string, duration time.Duration) bool {
	startTime := time.Now()
	for {
		conn, err := net.DialTimeout("tcp", "localhost:"+port, 100*time.Millisecond)
		if err != nil {
			if time.Since(startTime) >= duration {
				return false
			}
			time.Sleep(100 * time.Millisecond) // 每次查询间隔100毫秒
		} else {
			conn.Close()
			return true
		}
	}
}

func (pd *PdServer) Start() bool {
	execCommand("sh", "./scripts/start_pd.sh", strconv.Itoa(pd.Port))
	result := checkPort(strconv.Itoa(pd.Port), 3*time.Second)
	return result
}

func (pd *PdServer) Stop() bool {
	execCommand("sh", "./scripts/stop_pd.sh", strconv.Itoa(pd.Port))
	result := checkPort(strconv.Itoa(pd.Port), 3*time.Second)
	return !result
}

type TiKVServer struct {
	PdPort int
	Port   int
}

func (kv *TiKVServer) Start() bool {
	execCommand("sh", "./scripts/start_tikv_server.sh", strconv.Itoa(kv.Port), strconv.Itoa(kv.PdPort))
	result := checkPort(strconv.Itoa(kv.Port), 3*time.Second)
	return result
}

func (kv *TiKVServer) Stop() bool {
	execCommand("sh", "./scripts/stop_tikv_server.sh", strconv.Itoa(kv.Port))
	result := checkPort(strconv.Itoa(kv.Port), 3*time.Second)
	return !result
}
