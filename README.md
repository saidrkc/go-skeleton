## GO-SKELETON API (Golang + Cqrs + Grafana + Prometheus) 
#### Created by [@saidrkc]("https://github.com/saidrkc")

Go-skeleton is a basic api with most commonly tools for a multipurpose API with observability.
```
- Gin Web http route handling
- CQRS (Command Query Responsibility Segregation)
- Grafana
- Prometheus
``` 

## Installation

Use instructions below for a correct stand up

```bash
make up ## create docker containers
make bash ## enter to basic go-skeleton
make go-test ## run initial tests (unit and end 2 end)
make go-test-unit ## run only unit tests
make go-test-e2e ## run only e2e tests
make go-test-coverage ## run tests and create coverage files
```

## Usage

### Basic Endpoints

```
GET http://127.0.0.1:8080/ping
```
```
'{"Resp": "test"}'
```
```
POST http://127.0.0.1:8080/pong
```
```
'{}'
```

### Basic Dashboards
```bash
http://127.0.0.1:9090 - Prometheus
http://127.0.0.1:3000 - Grafana dashboards
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://mit.com/licenses/mit/)


### Basic Observability Dashboard

<a href="http://127.0.0.1:3000/d/1JNOL0aGz/golang-http?orgId=1" target="_blank">Grafana Dashboard</a>