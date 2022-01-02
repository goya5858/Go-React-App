FROM golang:1.17.5-alpine
WORKDIR /go/src/app

RUN apk add mysql-client

#>> mysql -h mysql_container -P 3306 -uroot -proot
#>> mysql -h mysql_container -P 3306 -udocker -pdocker