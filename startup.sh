kill `cat pid`
go build -o xiaochen
nohup ./xiaochen > xiaochen.log &
echo $! > pid