FROM golang:1.15.6

WORKDIR /go/src/app

COPY . .

CMD [ "go", "run", "math.go" ]