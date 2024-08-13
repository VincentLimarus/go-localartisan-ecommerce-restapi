# docker buildx build --platform linux/amd64 -t vincentlim27/localartisansv1 . --push
FROM golang:1.22-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

# docker pull vincentlim27/localartisans
# docker image ls
# docker run -p 8080:3000 [image_id]