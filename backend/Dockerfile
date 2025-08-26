FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o notes

FROM alpine:edge
WORKDIR /app
COPY --from=build /app/notes .
RUN apk --no-cache add ca-certificates tzdata
RUN export GIN_MODE=release
ENTRYPOINT ["/app/notes"]