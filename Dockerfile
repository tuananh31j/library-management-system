FROM golang:1.22.2 AS build

WORKDIR /app
COPY . .
RUN go clean --modcache
RUN go mod tidy
RUN go build ./main.go


EXPOSE 3000
CMD ["./main"]
