FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Create a minimal go.work that references only the modules we copy below.
RUN printf 'go 1.25.1\n\nuse (\n\t./schedule_service\n\t./libs/json\n\t./libs/utils_sql\n\t./libs/kafka\n)\n' > /app/go.work

# Module files
COPY schedule_service/go.mod schedule_service/
COPY libs/json/go.mod libs/json/
COPY libs/utils_sql/go.mod libs/utils_sql/
COPY libs/kafka/go.mod libs/kafka/

# Source for needed modules
COPY schedule_service/ schedule_service/
COPY libs/json/ libs/json/
COPY libs/utils_sql/ libs/utils_sql/
COPY libs/kafka/ libs/kafka/

WORKDIR /app/schedule_service

# Download dependencies
RUN go mod download

# Build the service
RUN go build -trimpath -x -o /out/app ./cmd/app

FROM gcr.io/distroless/static:nonroot AS runtime

WORKDIR /

COPY --from=builder /out/app /app

    # Непривилегированный пользователь
USER nonroot:nonroot

EXPOSE 8090

CMD ["/app"]
