# Rarelly changed first
FROM ubuntu:latest
EXPOSE 8000
WORKDIR /app
ENV HOST=localhost PORT=5434 NAME=root
ENV USER=root PASSWORD=root

COPY ./main main

# ENTRYPOINT == CMD
CMD [ "./main" ]