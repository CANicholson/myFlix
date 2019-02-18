FROM golang
COPY myFlix /go
EXPOSE 8080
CMD ["/go/main"]