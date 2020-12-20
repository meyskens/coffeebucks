# Start by building the application.
FROM golang:1.15-buster as build

ADD . /go/src/github.com/meyskens/coffebucks
WORKDIR /go/src/github.com/meyskens/coffebucks/backend

RUN go build -o /go/bin/coffebucks

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/coffebucks /
CMD ["/coffebucks"]