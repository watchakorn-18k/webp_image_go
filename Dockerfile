FROM golang:1.22 AS builder
WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download
RUN go get github.com/nickalie/go-webpbin

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

FROM alpine:3.19
WORKDIR /root/

RUN apk add --no-cache libwebp-tools libpng libjpeg-turbo giflib tiff

RUN mkdir -p /root/.bin/webp && \
    ln -s /usr/bin/cwebp /root/.bin/webp/cwebp

COPY --from=builder /app/myapp .

EXPOSE 3000

# Run the application
CMD ["./myapp"]