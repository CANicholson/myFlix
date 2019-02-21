sudo docker stop working
sudo docker rm working

sudo docker run -v "$PWD/myFlix:/go" -it -p 80:8080 --name working  golang
sleep 20
cd myFlix
cd src
sudo go run main.go
