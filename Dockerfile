FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

EXPOSE 4567

ENV GIN_MODE=release

CMD ./portfolio-golang-api