FROM golang:1.16

MAINTAINER xueqi
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/go
COPY . $GOPATH/src/go
RUN go get -t . && \
    go build .

EXPOSE 8888
CMD ./go_project