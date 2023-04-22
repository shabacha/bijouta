FROM golang:alpine as builder

RUN apk update && apk upgrade \
    && apk add --no-cache git

RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bijouta cmd/app/*.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /build/bijouta .
COPY --from=builder /build/pkg/config /app/config   

CMD ["./bijouta"]
