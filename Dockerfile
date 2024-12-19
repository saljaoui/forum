FROM golang:1.22.3-alpine

WORKDIR /app

RUN apk update && apk add --no-cache \
bash \
build-base \
sqlite-dev \
gcc \
libc-dev \
&& rm -rf /var/cache/apk/*

COPY . .
COPY go.mod go.sum ./

RUN go build -o ourForum .

EXPOSE 8080

CMD ["./ourForum"]