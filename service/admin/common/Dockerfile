FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /app/common common.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/common /app/common
COPY ./etc /app/etc

CMD ["./common", "-f", "etc/common.yaml"]
