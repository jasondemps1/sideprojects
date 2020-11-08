FROM golang:1.15.4 as build
ADD . /server
RUN cd /server && CGO_ENABLED=0 go build -o server

FROM alpine
WORKDIR /app
COPY --from=build /server/server /app
ENTRYPOINT ["./server"]