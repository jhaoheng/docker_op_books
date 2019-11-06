FROM golang:1.13.1

ADD ./app/ossignalbin /go/bin

RUN ln -s /go/bin/ossignalbin /usr/bin/app

RUN ls -al /usr/bin/app

RUN apt-get update && apt-get install -y tini

COPY entrypoint.sh /home
RUN chmod +x /home/entrypoint.sh
ENTRYPOINT ["tini", "--", "/home/entrypoint.sh"]

CMD ["app"]