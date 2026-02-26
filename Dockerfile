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
ARG DBFILE=scheduler.db
ARG PASSWORD=12345
ARG TOKENSECRET=secret
ARG MAXTASKS=50

WORKDIR /app

COPY --from=builder /app/task_app ./

COPY web ./web

COPY scheduler.db .

EXPOSE $PORT

ENV TODO_PORT=$PORT
ENV TODO_DBFILE=$DBFILE
ENV TODO_PASSWORD=$PASSWORD
ENV TODO_TOKENSECRET=$TOKENSECRET
ENV TODO_MAXTASKS=$MAXTASKS

CMD ["./task_app"]

