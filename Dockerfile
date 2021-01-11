FROM golang:1.15-alpine AS builder

LABEL maintainer="sshit+zhilyaev.dmitriy@gmail.com"
LABEL name="sshit"

# enable Go modules support
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/zhilyaev/sshit

COPY go.mod go.sum ./
RUN go mod download

# Copy src code from the host and compile it
COPY cmd cmd
COPY pkg pkg
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /sshit github.com/zhilyaev/sshit/cmd/sshit

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
COPY --from=builder /sshit /bin
ENTRYPOINT ["/bin/sshitgo"]
