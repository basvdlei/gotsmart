GotSmart
========

GotSmart collects information from the Dutch SlimmeMeter (translated as Smart
Meter) and exports them as Prometheus metrics.

Originally forked from: https://github.com/basvdlei/gotsmart

Setup
-----

### Build

```sh
go get github.com/robbertnoordzij/gotsmart
```

### Run

Specify the serial device that is connected with the Smart Meter.

```sh
gotsmart -device /dev/ttyS0
```

Setup with Docker
-----------------

### Build

```sh
docker build -t gotsmart .
```

### Run

Make sure to add the serial device that is connected with the Smart Meter to
the container. (ttyS0, ttyUSB0, ttyAMA0, etc)

```sh
docker run -d -p 8080:8080 --device /dev/ttyS0:/dev/ttyS0 gotsmart
```

Usage
-----

By default gotsmart listens on port 8080 and exposes the metrics under
`/metrics`.
