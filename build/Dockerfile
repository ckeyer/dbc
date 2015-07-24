FROM daocloud.io/golang:1.3-onbuild

WORKDIR /gopath/app
ENV GOPATH /gopath/app

RUN go get github.com/astaxie/beego &&\ 
go get github.com/beego/bee && \ 
go get github.com/hoisie/redis && \ 
go get github.com/go-sql-driver/mysql && \
go get github.com/ckeyer/go-log

EXPOSE 8080

RUN go get github.com/ckeyer/dbc

ADD /build/test.go /test.go

CMD cd $GOPATH/src/github.com/ckeyer/dbc && \
go run main.go
# CMD cd /gopath/app/bin/github.com/ckeyer/dbc/
# go run /test.go


