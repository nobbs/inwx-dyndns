FROM golang:1.17-alpine AS builder

WORKDIR /go/src/app

# copy go.mod/sum files and install dependencies
COPY go.* ./
RUN go mod download

# copy everything else
COPY . .

# compile the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o inwx-dyndns ./cmd

# and now put the compiled binary into a bare alpine image
FROM alpine:3.15

LABEL org.opencontainers.image.authors "Alexej Disterhoft <code@disterhoft.de>"
LABEL org.opencontainers.image.source "https://github.com/nobbs/inwx-dyndns"

WORKDIR /usr/local/bin
COPY --from=builder /go/src/app/inwx-dyndns /usr/local/bin/

# use an unprivileged user
USER nobody

# and run the application
CMD ["inwx-dyndns"]
