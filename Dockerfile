FROM golang:alpine
ADD . /go/src/github.com/vico1993/Yoshi
RUN go install github.com/vico1993/Yoshi
CMD ["/go/bin/Yoshi"]
EXPOSE 3000
