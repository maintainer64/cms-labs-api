FROM docker.io/library/alpine:latest

WORKDIR /app
# Copy binary and config files from /build to root folder of scratch container.
COPY /build /app/

RUN chmod +x /app/start.sh /app/clabgate

# Command to run when starting the container.
ENTRYPOINT ["/app/start.sh"]
EXPOSE 5000
