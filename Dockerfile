FROM golang:latest AS builder

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main -ldflags '-w -s' src/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /app/main /usr/bin
COPY --from=builder /app /app
EXPOSE 3000