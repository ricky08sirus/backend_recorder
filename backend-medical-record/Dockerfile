# Build stage
FROM golang:1.19-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]

