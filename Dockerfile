FROM golang:alpine AS builder
ENV CGO_ENABLE=0
ENV GOOS=linux
ENV GO111MODULE=auto
WORKDIR /build
COPY main.go ./
RUN go build -o hello_world_go

FROM golang:alpine
WORKDIR /app
RUN adduser --disabled-password app
COPY --from=builder --chown=app:app /build/hello_world_go ./
USER app
CMD [ "./hello_world_go" ]