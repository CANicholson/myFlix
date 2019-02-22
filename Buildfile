cd ..
sudo docker stop working
sudo docker rm working

sudo docker run -v "$PWD/myFlix:/go" -p 80:8080 --name working golang sh -c "export GOPATH=${PWD} && cd src && go get gopkg.in/mgo.v2 && go build main.go"
