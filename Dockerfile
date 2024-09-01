# Build Stage
FROM golang:1.23.0-alpine AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
    go build -a \
    -mod=vendor \
    -installsuffix cgo \
    -o url-shortener \
    cmd/server/main.go

# Final stage
FROM scratch
COPY --from=build /app/url-shortener /url-shortener
ENTRYPOINT ["/url-shortener"]
