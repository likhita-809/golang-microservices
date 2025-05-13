FROM golang::1.22 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o user-service ./cmd/main.go

FROM grc.io/distroless/base
WORKDIR /
COPY --from=builder /app/user-service /
ENTRYPOINT ["/user-service"]