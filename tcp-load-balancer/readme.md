## TCP (layer4) load balancer in Go

This is just a reverse proxy that forwards the tcp traffic to another port.

No fancy load balancing logic (just round robin), no liveness


## Usage

`$ app -path=config.json -demo`

* **path** - path to configuration file in json format. Accepted format:
```
{
    "main_port": 8000,
    "backends": [4000, 4001, 4002, 4003]
}
```
* **demo** - demo mode flag - if provided then example backends will be generated