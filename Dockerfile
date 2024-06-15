FROM golang:1.22-bookworm

WORKDIR /app

COPY . .
RUN go build -o main

CMD [ "/app/main" ]
