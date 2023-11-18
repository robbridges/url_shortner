FROM --platform=linux/amd64 golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/api

EXPOSE 8080

CMD [ "./main" ]