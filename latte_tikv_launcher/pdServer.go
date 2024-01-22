package latte_tikv_launcher

import (
	"os/exec"
	"strconv"
)

type PdServer struct {
	Port int
}

func (pd *PdServer) start() {
	exec.Command("sh", "./scripts/start_pd.sh", strconv.Itoa(pd.Port))
}

func (pd *PdServer) stop() {
	exec.Command("sh", "./scripts/stop_pd.sh", strconv.Itoa(pd.Port))
}

type TiKVServer struct {
	PdPort int
	Port   int
}

func (kv *TiKVServer) start() {
	exec.Command("sh", "./scripts/start_tikv_server.sh", strconv.Itoa(kv.PdPort), strconv.Itoa(kv.Port))
}

func (kv *TiKVServer) stop() {
	exec.Command("sh", "./scripts/stop_tikv_server.sh", strconv.Itoa(kv.Port))
}
