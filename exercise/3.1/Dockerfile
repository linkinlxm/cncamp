FROM golang:1.16-alpine as build

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY . /app

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && \
    go build -ldflags="-w -s" -o /bin/http http.go

FROM alpine

COPY --from=build /bin/http /bin/http

ENTRYPOINT [ "/bin/http" ]