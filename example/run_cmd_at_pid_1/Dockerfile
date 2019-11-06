FROM golang:1.13.1

ADD ./app/ossignalbin /go/bin

RUN ln -s /go/bin/ossignalbin /usr/bin/app

RUN ls -al /usr/bin/app

CMD exec app