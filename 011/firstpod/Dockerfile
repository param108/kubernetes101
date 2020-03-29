FROM ubuntu:18.04

ADD web/web /web

RUN apt-get update && apt-get install -y wget

RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz

EXPOSE 8080

CMD /web
