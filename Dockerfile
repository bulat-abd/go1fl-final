# builder
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o  task_app .

# production

FROM scratch

ARG PORT=7540

WORKDIR /app

COPY --from=builder /app/task_app ./

COPY web ./web

COPY scheduler.db .

EXPOSE $PORT

ENV TODO_PORT=$PORT

CMD ["./task_app"]

