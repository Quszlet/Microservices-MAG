FROM golang:1.25.1-alpine AS builder

# Ускоряем и делаем сборку воспроизводимой
# RUN apk add --no-cache ca-certificates tzdata git build-base
# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Кэшируем зависимости
# COPY go.mod go.sum ./
COPY . .
RUN go mod download 

# Сборка
RUN go build -trimpath -x -o /out/app ./cmd/app

FROM gcr.io/distroless/static:nonroot AS runtime

WORKDIR /

COPY --from=builder /out/app /app

    # Непривилегированный пользователь
USER nonroot:nonroot

EXPOSE 8080

CMD ["/app"]