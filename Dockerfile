FROM golang
ADD . /go/src/playlistclient
RUN go get github.com/golang/protobuf/proto
RUN go get google.golang.org/grpc
RUN go install playlistclient
ENTRYPOINT ["/go/bin/playlistclient"]
EXPOSE 8000