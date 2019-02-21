sudo docker stop working
sudo docker rm working

sudo docker run -v $PWD/myFlix:/go -i -p 80:8080 --name working  golang

