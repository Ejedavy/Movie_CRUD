FROM golang:1.15.3-alpine3.12

EXPOSE 8000

RUN apk update 
  
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build cmd/main.go
RUN mv main /usr/local/bin/

CMD ["main"]