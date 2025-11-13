FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# файлы для резолва модулей
COPY go.work ./go.work
COPY libs/json/go.mod libs/json/go.mod
COPY libs/utils_sql/go.mod libs/utils_sql/go.mod
COPY doctors_service/go.mod doctors_service/go.mod


# исходники только нужных модулей
COPY libs/json/ libs/json/
COPY libs/utils_sql/ libs/utils_sql/
COPY doctors_service/ doctors_service/

RUN go mod download 

WORKDIR /app/doctors_service

# Сборка
RUN go build -trimpath -x -o /out/app ./cmd/app

FROM gcr.io/distroless/static:nonroot AS runtime

WORKDIR /

COPY --from=builder /out/app /app

    # Непривилегированный пользователь
USER nonroot:nonroot

EXPOSE 8080

CMD ["/app"]