# Digit Go

**compile go files**
```bash
go build main.go tm1637.go
go build clock.go tm1637.go stop.go
go build day.go tm1637.go stop.go
go build ds18b20.go tm1637.go stop.go
```

**build image**
```bash
docker build -t digit_go:1.0 -f Dockerfile .
```

**run docker container**
```bash
docker run -d --device /dev/gpiomem digit_go:1.0 ./clock
docker run -d --device /dev/gpiomem digit_go:1.0 ./day
docker run -d --device /dev/gpiomem -v /sys/bus/w1/devices/28-05170143ccff/w1_slave:/sys/bus/w1/devices/28-05170143ccff/w1_slave digit_go:1.0 ./ds18b20
```

**build via docker-compose**
```bash
docker-compose -f docker-compose.yml build
```

**run via docker-compose**
```bash
docker-compose -f docker-compose.yml up -d
```