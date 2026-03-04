# syntax=docker/dockerfile:1

FROM golang:1.26-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .  
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/taskmanager ./cmd/taskmanager

FROM alpine:latest AS runner
COPY --from=build /bin/taskmanager /taskmanager
EXPOSE 8080
CMD ["/taskmanager"]