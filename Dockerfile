FROM golang:latest AS compiling_stage
RUN mkdir -p /go/src/tasks
WORKDIR /go/src/tasks
RUN go env -w GO111MODULE=auto
ADD main.go .
ADD go.mod .
RUN go install .
 
FROM alpine:latest
LABEL version="1.0.0"
LABEL maintainer="Task Student<test@test.ru>"
WORKDIR /root/
COPY --from=compiling_stage /go/bin/tasks .
ENTRYPOINT /go/bin/tasks
EXPOSE 8080