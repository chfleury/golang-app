FROM golang:1.21 AS generate-binary

COPY main.go .

RUN go build -o main main.go

FROM scratch

WORKDIR /bin

COPY --from=generate-binary /go/main /main

CMD [ "/main" ]
