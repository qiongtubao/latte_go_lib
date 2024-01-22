WORKSPACE=$(dirname $(readlink -f "$0") || (cd "$(dirname "$0")";pwd))
cd $WORKSPACE
chmod 777 tikv-server
nohup ./tikv-server --labels zone=z1 --pd-endpoints="127.0.0.1:$2" --addr="127.0.0.1:$1" --config=./tikv_server_config.toml --data-dir=./tikv$1_data --log-file=tikv$1.log > start_tikv_server$1.out 2>&1 &
