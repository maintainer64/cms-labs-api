FROM alpine:3.24.2
COPY build/jupyter-smoke /usr/local/bin/jupyter-smoke
USER 1000:100
ENTRYPOINT ["/usr/local/bin/jupyter-smoke"]
