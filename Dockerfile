FROM golang as builder
#copy in all source files
COPY ./utilities /utilities
WORKDIR /utilities
#run the utilities unit tests
RUN go test

COPY ./cmd /cmd
WORKDIR /cmd
#build and install source files into binary including c lib into the binary
RUN CGO_ENABLED=0 GOOS=linux go install -a
#run the go app tests
RUN go test

#------------------------------------
FROM alpine
#Expose host:container
EXPOSE 9000
WORKDIR /root/
#copy binary across
COPY --from=builder /go/bin/ /cmd
#copy html across
COPY views /views
WORKDIR /cmd/
#run server
CMD ["./server"]