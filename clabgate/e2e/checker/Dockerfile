FROM alpine:3.24.2
COPY entrypoint.sh /usr/local/bin/checker
RUN chmod 0755 /usr/local/bin/checker
USER 65532:65532
ENTRYPOINT ["/usr/local/bin/checker"]
