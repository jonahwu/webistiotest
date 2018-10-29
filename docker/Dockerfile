
# build stage 1
#FROM golang:alpine AS unittest-stage
FROM golang:1.9.0 AS unittest-stage
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   

ADD ./ $GOPATH/src/src
ADD ./vendor $GOPATH/src/src/vendor
#ADD ./src /src
RUN cd $GOPATH/src/src  && ls -alh
RUN ls $GOPATH/src/src/vendor -alh
RUN go get -u github.com/kardianos/govendor
RUN cd $GOPATH/src/src && govendor sync && go test ./...

# build stage 2
#FROM golang:alpine AS build-stage
FROM golang:1.9.0 AS build-stage
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   
COPY --from=unittest-stage $GOPATH/src/src/ $GOPATH/src/src/
RUN ls $GOPATH/src/src -alh
RUN ls $GOPATH/src/src/vendor -alh
RUN go get -u github.com/kardianos/govendor
# if use golang image to run
#RUN cd $GOPATH/src/src && govendor sync &&  make
# if use alpine image to run 
RUN cd $GOPATH/src/src && govendor sync && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gsw

# final stage
FROM alpine
#FROM golang:1.9.0
WORKDIR /app
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH   
COPY --from=build-stage $GOPATH/src/src/gsw ./gsw
RUN chmod +x ./gsw
#CMD ["sh","./gsw"]
ENTRYPOINT ["./gsw"]
