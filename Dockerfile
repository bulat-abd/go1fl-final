# builder
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o  task_app .

# production

FROM scratch

WORKDIR /app

COPY --from=builder /app/task_app ./

COPY web ./web

EXPOSE 7540

CMD ["./task_app"]

