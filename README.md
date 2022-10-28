## GO-SKELETON API (Grafana + Prometheus) 
#### Create by [@saidrkc]("https://github.com/saidrkc")

Go-skeleton is a basic api with most commonly tools for a multipurpose API with observability.
```
- Gin Web http route handling
- Grafana
- Prometheus
``` 

## Installation

Use instructions below for a correct stand up

```bash
docker-compose up -d
```

## Usage

```bash
http://127.0.0.1:9090 - Prometheus
http://127.0.0.1:3000 - Grafana dashboards
http://127.0.0.1:8080/ping - Basic ping endpoint
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://mit.com/licenses/mit/)


### Basic Observability Dashboard

<a href="http://127.0.0.1:3000/d/1JNOL0aGz/golang-http?orgId=1" target="_blank">Grafana Dashboard</a>