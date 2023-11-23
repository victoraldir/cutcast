FROM golang:1.20.6-alpine3.18 AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/cutcast

FROM alpine:3.18.2

RUN apk add --no-cache ca-certificates && \
    apk add git && \
    apk add ffmpeg && \
    apk add python3 && \
    apk add py3-pip && \
    apk add curl && \
    curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp && \
    chmod a+rx /usr/local/bin/yt-dlp  # Make executable
     
WORKDIR /app

COPY --from=builder /bin/cutcast /bin/cutcast

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["/bin/cutcast"]