FROM golang:1.12-alpine as builder

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" src/app/go-hpa.go

FROM scratch

WORKDIR /go/src/app

COPY --from=builder /go/src/app .

CMD ["./go-hpa"]