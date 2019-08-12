# Prometheus Demo

A simple demo project which demonstrates how to instrument Go service using Prometheus.  

## Requirements
* Running Prometheus on minikube Kubernetes cluster - [prometheus-grafana-demo](https://github.com/martinsirbe/prometheus-grafana-demo).

## Run example
1. Simply run `make start` - this will start an HTTP server on port `1337` and expose `metrics` endpoint.    
2. Check https://localhost:1337/metrics  

You should see something like this in your console:
```bash
Open http://localhost:1337/metrics in your browser to view metrics.
deleted - 038ae333-e583-458f-9a6d-49622a7508ec
inserted - e9e9806a-4935-456f-ada3-9c83a143e2c2
inserted - 95c23f65-cb46-42c4-9e2d-23ffdd28a4b1
inserted - e0e91a82-fbcd-4600-b0c6-9026ff95f076
inserted - 4891fc26-02c6-4776-8c5c-ec34e59a85a1
inserted - 6a8b7e54-1f99-4e83-97a2-559b06e2eb08
deleted - c4a412da-891b-40ea-8385-37d5bc2821b1
inserted - b1f98f60-05e1-4f00-8917-21d70a47253a
```

## Run tests
Simply run `make test`. If you are playing around with tests and you have changed the mocked `Storage` interface, then 
you will need to install [go-mock](https://github.com/golang/mock#installation)8 and run `make generate` to regenerate the Storage mock.  

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENCE.md) file for details.  

[1]: https://github.com/rcrowley/go-metrics
[2]: https://github.com/rcrowley/go-metrics#publishing-metrics
[3]: https://github.com/golang/mock#installation
