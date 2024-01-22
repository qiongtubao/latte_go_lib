#!/bin/bash
WORKSPACE=$(dirname $(readlink -f "$0") || (cd "$(dirname "$0")";pwd))
cd $WORKSPACE
pd_port=$1
peer_port=$(expr $pd_port + 1)
chmod 777 pd-server
nohup ./pd-server --name=pd  --config=./pd.toml --data-dir=./pd_data$pd_port --client-urls="http://127.0.0.1:$pd_port" --peer-urls="http://127.0.0.1:$peer_port" --initial-cluster="pd=http://127.0.0.1:$peer_port" --log-file=pd$pd_port.log > start_pd.out 2>&1 &
