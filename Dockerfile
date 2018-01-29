FROM golang:1.9.2 
MAINTAINER Dennis Kleber
RUN mkdir /go/src/project
WORKDIR /go/src/project
VOLUME /go/src/project
# install iris framework 
RUN go get -u github.com/kataras/iris
#install godep package manager 
RUN go get github.com/tools/godep
ADD main.go /go/src/project
ADD public  /go/src/project/public
CMD ../../bin/godep save
CMD go run main.go
EXPOSE 8080
