FROM golang:1.20-alpine3.19 AS builder


ENV GOPROXY https://goproxy.cn
ENV GOSUMDB sum.golang.google.cn
ENV SRC_PATH ${GOPATH}/gin-swagger

WORKDIR ${SRC_PATH}

COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk update && apk upgrade

RUN set -ex \
&& apk add git tar \
&& go mod tidy \
&& go mod download \
&& CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o gin-swagger .

FROM alpine:3.19 AS final

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV LANGUAGE en_US:en

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk update && apk upgrade

RUN set -ex \
&& apk add bash tzdata \
&& ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
&& echo ${TZ} > /etc/timezone \
&& rm -rf /var/cache/apk/*

RUN mkdir -p /app/conf

COPY --from=builder /go/gin-swagger/gin-swagger /app/
COPY --from=builder /go/gin-swagger/conf/app.ini /app/conf/

EXPOSE 8000

ENTRYPOINT ["/app/gin-swagger"]

