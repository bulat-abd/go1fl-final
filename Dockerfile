# builder
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o  task_app .

# production

FROM scratch

ARG PORT=5000
ARG DBFILE=scheduler.db
ARG PASSWORD=12345

WORKDIR /app

COPY --from=builder /app/task_app ./

COPY web ./web

EXPOSE $PORT

ENV TODO_PORT=$PORT
ENV TODO_DBFILE=$DBFILE
ENV TODO_PASSWORD=$PASSWORD

CMD ["./task_app"]

