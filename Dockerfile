FROM balenalib/rpi-raspbian

RUN apt-get update \
 && apt-get install -y wget \
                       git \
 && wget https://storage.googleapis.com/golang/go1.11.linux-armv6l.tar.gz \
 && tar -C /usr/local -xzf go1.11.linux-armv6l.tar.gz \
 && rm go1.11.linux-armv6l.tar.gz \
 && export PATH=$PATH:/usr/local/go/bin \
 && go get github.com/stianeikeland/go-rpio

ADD . /home

RUN rm /etc/localtime && \
 ln -s /usr/share/zoneinfo/Europe/Moscow /etc/localtime

# signal SIGTERM
STOPSIGNAL 15

WORKDIR /home
