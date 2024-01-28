for pid in `ps -ef | grep tikv-server | grep -v grep | grep $1 | awk '{print $2}'`
do
    kill -9 $pid
done
echo "tikv $1 stopped successfully!"