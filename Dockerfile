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
    pip install --upgrade --force-reinstall "git+https://github.com/ytdl-org/youtube-dl.git"
     
WORKDIR /app

COPY --from=builder /bin/cutcast /bin/cutcast

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["/bin/cutcast"]