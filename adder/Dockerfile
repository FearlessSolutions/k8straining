FROM golang:1.12 AS builder
WORKDIR /app

COPY . .
RUN go build -o app

EXPOSE 3000

ENTRYPOINT ["./app"]


# # Copy in the full app
# COPY go.mod .
# RUN go get

# COPY main.go main.go

# # Disable CGO to create statically linked binary.
# # Without this, the go compiler will attempt to dynamically link to libc
# ENV CGO_ENABLED 0
# # Build / install app
# RUN go build -o app

# # Begin Deployed Stage:
# FROM scratch
# WORKDIR /
# COPY --from=builder /app/app /app

# EXPOSE 3000

# ENTRYPOINT ["./app"]
