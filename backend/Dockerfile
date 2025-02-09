FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
COPY /build/* .

# Command to run when starting the container.
ENTRYPOINT ["/apiserver"]
