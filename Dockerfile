############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . /home/api-gateway


WORKDIR /home/api-gateway

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/api-gateway

############################
# STEP 2 build a small image
############################

FROM alpine:3.4

RUN apk --no-cache --update upgrade

EXPOSE 8080

WORKDIR /root
# Copy our static executable.
COPY --from=builder /go/bin/api-gateway .


# Run the payment binary.
ENTRYPOINT ["./api-gateway"]