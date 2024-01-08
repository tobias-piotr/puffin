FROM golang:1.21.3-alpine AS builder

EXPOSE 8008

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Local development
FROM builder as dev

ENV DEBUG=true

# Install air for hot reloading
RUN go install github.com/cosmtrek/air@latest
# Install templ to hot generate templates
RUN go install github.com/a-h/templ/cmd/templ@latest

ENTRYPOINT ["air"]

# Production
FROM gcr.io/distroless/base-debian11 AS prod

WORKDIR /

COPY --from=builder /main /main

USER nonroot:nonroot

ENTRYPOINT ["/main"]
