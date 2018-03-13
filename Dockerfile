FROM golang:1.10.0 AS build

ADD . /go/src/github.com/dlish/helloworld-auth
WORKDIR /go/src/github.com/dlish/helloworld-auth
RUN go get -d -v -t
RUN go test --cover -v ./...
RUN go build -v -o auth-service


FROM alpine:3.6

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
CMD ["auth-service"]
COPY --from=build /go/src/github.com/dlish/helloworld-auth/auth-service /usr/local/bin/auth-service
RUN chmod +x /usr/local/bin/auth-service
