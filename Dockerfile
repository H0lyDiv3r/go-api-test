FROM golang:1.23-alpine3.20

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

# Remove the RUN command for migration
# RUN go run ./migrate/migrate.go

CMD ["sh", "-c", "go run ./migrate/migrate.go && air"]
