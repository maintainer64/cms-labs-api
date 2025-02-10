FROM alpine:latest

WORKDIR /app
# Copy binary and config files from /build to root folder of scratch container.
COPY /build /app/

RUN chmod +x /app/migrate.sh /app/start.sh /app/apiserver

# Command to run when starting the container.
ENTRYPOINT ["/app/start.sh"]
EXPOSE 5000
