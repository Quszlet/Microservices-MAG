FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Workspace includes gateway and shared libs
RUN printf 'go 1.25.1\n\nuse (\n\t./gateway_service\n\t./libs/json\n)\n' > /app/go.work

COPY gateway_service/go.mod gateway_service/go.sum gateway_service/gqlgen.yml /app/gateway_service/
COPY gateway_service/graph /app/gateway_service/graph
COPY gateway_service/server.go /app/gateway_service/

# shared libs used by gateway
COPY libs/json /app/libs/json

WORKDIR /app/gateway_service

RUN go mod download
RUN go build -trimpath -o /out/app ./server.go

FROM gcr.io/distroless/static:nonroot AS runtime

WORKDIR /

COPY --from=builder /out/app /app

USER nonroot:nonroot

EXPOSE 8070

CMD ["/app"]
