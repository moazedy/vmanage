# syntax=docker/dockerfile:1

FROM golang:1.22.3 as builder

WORKDIR /vmanage

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vmanage ./cmd/vmanage/

FROM alpine:latest

LABEL maintainer="moazedy@gmail.com"

WORKDIR /root/

COPY --from=builder /vmanage/vmanage .
COPY --from=builder /vmanage/configs/ ./configs/ 

EXPOSE 4853

CMD ./vmanage
