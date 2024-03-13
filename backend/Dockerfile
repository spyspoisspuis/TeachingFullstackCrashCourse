# Stage 1: Build the application
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Run the application
FROM scratch

WORKDIR /app

COPY app.env /app
COPY --from=builder /app/main /app

CMD ["/app/main"]