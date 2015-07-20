FROM daocloud.io/golang:1.3-onbuild

RUN go get github.com/astaxie/beego &&\ 
go get github.com/beego/bee && \ 
go get github.com/ckeyer/dbc

RUN cd $GOPATH/src/github.com/ckeyer/dbc && \
bee run dbc
