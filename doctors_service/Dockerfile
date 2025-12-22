FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Copy workspace file and only the modules needed to build this service.
# This keeps the build context small and avoids copying unrelated services.
RUN printf 'go 1.25.1\n\nuse (\n\t./doctors_service\n\t./libs/json\n\t./libs/utils_sql\n\t./libs/kafka\n)\n' > /app/go.work

# Module files (helps `go` resolve before copying sources)
COPY doctors_service/go.mod doctors_service/
COPY libs/json/go.mod libs/json/
COPY libs/utils_sql/go.mod libs/utils_sql/
COPY libs/kafka/go.mod libs/kafka/

# Source for the modules we need
COPY doctors_service/ doctors_service/
COPY libs/json/ libs/json/
COPY libs/utils_sql/ libs/utils_sql/
COPY libs/kafka/ libs/kafka/

WORKDIR /app/doctors_service

# Download dependencies (will use go.work and local modules copied above)
RUN go mod download

# Build the service
RUN go build -trimpath -x -o /out/app ./cmd/app

FROM gcr.io/distroless/static:nonroot AS runtime

WORKDIR /

COPY --from=builder /out/app /app

    # Непривилегированный пользователь
USER nonroot:nonroot

EXPOSE 8080

CMD ["/app"]
