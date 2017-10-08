FROM armv7/armhf-ubuntu:latest

MAINTAINER Eugene Goncharov NikeLambert@gmail.com

RUN apt-get install -y wget \
                       git \
  && wget https://storage.googleapis.com/golang/go1.9.linux-armv6l.tar.gz \
  && tar -C /usr/local -xzf go1.9.linux-armv6l.tar.gz \
  && rm go1.9.linux-armv6l.tar.gz \
  && export PATH=$PATH:/usr/local/go/bin \
  && go get github.com/stianeikeland/go-rpio


ADD . /home

RUN rm /etc/localtime && \
 ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

# signal SIGTERM
STOPSIGNAL 15

WORKDIR /home
