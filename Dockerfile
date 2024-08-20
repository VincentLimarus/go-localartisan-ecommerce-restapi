FROM golang:1.22-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .env .env
COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

EXPOSE 3000

CMD ["./main"]

# docker buildx build --platform linux/amd64 -t vincentlim27/localartisanv1 . --push

# docker pull vincentlim27/localartisans
# docker image ls
# docker run -p 3000:3000 [image_id]

# docker push vincentlim27/localartisansv1:latest
 
# docker buildx build --platform linux/amd64 \
#   --build-arg PORT=3000 \
#   --build-arg DB_URL="host=localhost user=postgres password=Vincent27 dbname=localArtisans port=5432 sslmode=disable" \
#   --build-arg TIMEOUT="3s" \
#   -t vincentlim27/localartisansv1 . --push
   