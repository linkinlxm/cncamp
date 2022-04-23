#FROM golang:1.17-alpine3.13 AS builder
FROM golang:1.17 AS builder
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
#  apk update && \
#  apk add --no-cache ca-certificates tzdata
RUN apt update && apt upgrade -y && apt install -y ca-certificates tzdata
WORKDIR /app

COPY ./app/ ./
RUN ls -lh ./
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go build -ldflags="-w -s" -o /bin/http http.go && \
    ls -lh /bin/ /app ./
#FROM krallin/centos-tini:trusty
FROM krallin/centos-tini:centos7
WORKDIR /app
RUN yum update -y && yum upgrade -y
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/http /app/http

ENTRYPOINT ["/usr/local/bin/tini", "--", "/app/http"]

