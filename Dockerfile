FROM golang:1.9.2 
MAINTAINER Dennis Kleber
RUN mkdir /go/src/project
WORKDIR /go/src/project
VOLUME /go/src/project
# install go packages
RUN go get -u github.com/labstack/echo/...
RUN go get github.com/tkanos/gonfig
#install godep package manager 
RUN go get github.com/tools/godep
ADD main.go /go/src/project
ADD dispenser.go /go/src/project
ADD public  /go/src/project/public
ADD conf  /go/src/project/conf
CMD ../../bin/godep save
CMD go run main.go dispenser.go
EXPOSE 8080
