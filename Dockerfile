FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /echo-server cmd/app/main.go

CMD ["/echo-server"]

# Запуск с:
# docker run -d \
#   -e DB_HOST=postgres \
#   -e DB_PASSWORD=$DB_PASSWORD \
#   -e JWT_SECRET=$JWT_SECRET \
#   -p 8882:8882 \
#   echo-server:latest