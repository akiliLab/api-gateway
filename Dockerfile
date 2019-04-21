############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache openssh-client
COPY . /home/api-gateway
# for private repo pulling
ARG SSH_KEY
ARG SSH_KEY_PASSPHRASE
RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan github.com > /root/.ssh/known_hosts && \
    echo "${SSH_KEY}" > /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa

WORKDIR /home/api-gateway

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/api-gateway

RUN rm -rf /root/.ssh/
############################
# STEP 2 build a small image
############################

FROM alpine:3.4

RUN apk --no-cache --update upgrade && apk add --no-cache ca-certificates && update-ca-certificates

EXPOSE 5001

WORKDIR /root
# Copy our static executable.
COPY --from=builder /go/bin/api-gateway .


# Run the payment binary.
ENTRYPOINT ["./api-gateway"]