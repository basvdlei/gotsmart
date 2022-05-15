GotSmart
========

GotSmart collects data from the Dutch SlimmeMeter (translates as Smart Meter)
to export as Prometheus metrics.

Setup
-----

### Build

```sh
go install github.com/basvdlei/gotsmart
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

By default gotsmart listens on port 8080 and exposes the metrics on `/metrics`.


Build for Raspberry Pi
----------------------

GotSmart can run on Raspberry Pi. To cross compile use:

```
GOARCH=arm go build
```
