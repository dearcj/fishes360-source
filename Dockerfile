FROM golang
ADD ./ /go/

#COPY ./wb.crt /wbserv/bin
#COPY ./wb.key /wbserv/bin

WORKDIR /go/src/github.com/dearcj/golangproj
RUN go get -u github.com/golang/dep/...
RUN  dep ensure -v
ENV OPT {$OPT}
RUN go install github.com/dearcj/golangproj
EXPOSE 80 443
ENTRYPOINT "golangproj" -$OPT