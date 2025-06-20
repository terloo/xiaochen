kill `cat pid`
go build -o xiaochen cmd/xiaochen/main.go
nohup ./xiaochen > xiaochen.log &
echo $! > pid