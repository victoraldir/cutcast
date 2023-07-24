FROM ubuntu:latest

RUN set -x && \
    apt update && \
    apt install -y git && \
    apt install -y ffmpeg && \
    apt install -y python3 && \
    apt install -y python3-pip && \
    pip install --upgrade --force-reinstall "git+https://github.com/ytdl-org/youtube-dl.git"

WORKDIR /data

ENTRYPOINT ["youtube-dl"]