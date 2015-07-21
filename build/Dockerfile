FROM daocloud.io/golang:1.3-onbuild

WORKDIR /gopath/app
ENV GOPATH /gopath/app

RUN go get github.com/astaxie/beego &&\ 
go get github.com/beego/bee && \ 
go get github.com/ckeyer/dbc


ADD /build/test.go /test.go

EXPOSE 8080

CMD cd $GOPATH/src/github.com/ckeyer/dbc && \
bee run dbc
# CMD cd /gopath/app/bin/github.com/ckeyer/dbc/
# go run /test.go


