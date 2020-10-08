# Haproxy short config yml

[github](https://github.com/kiselev-nikolay/haproxy-short-config-yml)

Configuration generator for HAProxy.


## Usage

[More examples here](https://github.com/kiselev-nikolay/haproxy-short-config-yml/tree/v1/example)

Generator used for convert to haproxy configuration this:

```yaml
backends:
  machineand.me:
    cluster:
      - host: 10.114.0.1
        port: 80
  influx.machineand.me:
    balance: leastconn
    cluster:
      - host: 10.114.0.3
        port: 8086
      - host: 10.114.0.4
        port: 8086
  stats.machineand.me:
    cluster:
      - host: 10.114.0.2
        port: 80
```

[view result of this example](https://github.com/kiselev-nikolay/haproxy-short-config-yml/blob/v1/example/shortconf/haproxy.cfg)

#### Command:

```bash
haproxy-short-config-yml conf.yaml
```

## Installation

```bash
wget https://github.com/kiselev-nikolay/haproxy-short-config-yml/releases/download/v1/haproxy-short-config-yml_${OS}_${ARCH}
```

[all versions](https://github.com/kiselev-nikolay/haproxy-short-config-yml/releases/tag/v1)

## Build

```bash
go build
```