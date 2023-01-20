# Default to Go 1.18
ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src
# Import the code from the context.
COPY ./ ./
# Build the Go app
RUN CGO_ENABLED=0 GOFLAGS=-mod=vendor GOOS=linux go build -a -o /app .


######## Start a new stage from scratch #######
# Final stage: the running container.
FROM scratch AS final
# Import the compiled executable from the first stage.
COPY --from=builder /app /app

EXPOSE 3000

# Run the compiled binary.
ENTRYPOINT ["/app"]
