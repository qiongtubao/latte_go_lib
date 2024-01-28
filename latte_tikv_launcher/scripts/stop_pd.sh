for pid in $(ps -ef | grep pd-server | grep -v grep | grep $1 | awk '{print $2}')
do
    kill -9 $pid
done