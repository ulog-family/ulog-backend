FROM golang:alpine
LABEL authors="uli"
WORKDIR /home/workspace

RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /home/workspace

RUN sh build.sh

EXPOSE 443

ENTRYPOINT ["output/bin/ulog_service"]