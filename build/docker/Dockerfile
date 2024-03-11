#===================================================================================================#
#                                         GO BUILDER
#
# * Golang-alpine docker stage to compile presidium-oapi3 for alpine
#===================================================================================================#

FROM golang:1.20-alpine AS builder

WORKDIR /presidium-oapi3
COPY go.mod .
COPY go.sum .

RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
    update-ca-certificates

# Create appuser.
RUN adduser -D -g '' presidium

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o presidium-oapi3

#===================================================================================================#
#                                      ALPINE PRODUCTION IMAGE
#===================================================================================================#
FROM scratch

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copying presidium compiled version
COPY --from=builder /presidium-oapi3/presidium-oapi3 /app

USER presidium

ENTRYPOINT ["./presidium-oapi3", "convert"]
