FROM golang:1.21.4 as builder

WORKDIR /app

ENV CGO_ENABLED=0

COPY /go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base-debian11

COPY --from=builder /app /app
WORKDIR /app

CMD ["./main"]