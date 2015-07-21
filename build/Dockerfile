FROM daocloud.io/golang:1.3-onbuild

#RUN go get github.com/astaxie/beego &&\ 
#go get github.com/beego/bee && \ 
#go get github.com/ckeyer/dbc

#CMD cd $GOPATH/src/github.com/ckeyer/dbc && \
#bee run dbc

ADD /build/test.go /test.go

EXPOSE 8080

CMD go run /test.go


