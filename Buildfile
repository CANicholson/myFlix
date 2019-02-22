cd ..
sudo docker stop working
sudo docker rm working

sudo docker run -v "$PWD/myFlix:/go" -it -p 80:8080 --name working golang 
export GOPATH=${PWD} 
cd src 
go get gopkg.in/mgo.v2 
go run main.go
