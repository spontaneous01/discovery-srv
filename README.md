# Discovery Service

The discovery service is a micro service which layers on the [go-micro.Registry](https://godoc.org/github.com/micro/go-micro/registry#Registry) to 
provide service discovery. It includes heartbeating, in memory caching and maintains a registry based on liveness.

## Features

- Zero dependencies
- In-memory registry
- Clustering
- Heartbeating
- RPC interface

## Usage

### Run Standalone

Discovery can be run as a single instance. It just needs a server address and the same as the registry address.

```shell
discovery-srv --server_address=10.0.0.1:8001 --registry_address=10.0.0.1:8001
```

### Cluster Instances

Clustering instances of discovery is simply done by specifying the addresses of other instances.

```shell
discovery-srv --server_address=10.0.0.1:8001 --registry_address=10.0.0.1:8001,10.0.0.2:8001,10.0.0.3:8001
```

Use via [go-os/discovery](https://github.com/micro/go-os/tree/master/discovery) client

## The API

The discovery service implements the [go-micro.Registry](https://godoc.org/github.com/micro/go-micro/registry#Registry) as RPC methods as well as the following

### Discovery.Heartbeats
```shell
micro query go.micro.srv.discovery Discovery.Heartbeats
{
	"heartbeats": [
		{
			"id": "foo-123",
			"interval": 1,
			"service": {
				"endpoints": [
					{
						"metadata": {
							"index": "Handles index requests"
						},
						"name": "/index",
						"request": {
							"name": "request",
							"type": "Request"
						},
						"response": {
							"name": "response",
							"type": "Response"
						}
					}
				],
				"metadata": {
					"foo": "bar"
				},
				"name": "go.micro.srv.foo",
				"nodes": [
					{
						"address": "localhost",
						"id": "foo-123",
						"metadata": {
							"bar": "baz"
						},
						"port": 8080
					}
				],
				"version": "latest"
			},
			"timestamp": 1.451177551e+09,
			"ttl": 5
		}
	]
}
```

### Discovery.Endpoints

```shell
micro query go.micro.srv.discovery Discovery.Endpoints
{
	"endpoints": [
		{
			"endpoint": {
				"metadata": {
					"index": "Handles index requests"
				},
				"name": "/index",
				"request": {
					"name": "request",
					"type": "Request"
				},
				"response": {
					"name": "response",
					"type": "Response"
				}
			},
			"service": "go.micro.srv.foo",
			"version": "latest"
		}
	]
}
```

### Sending Heartbeats

Heartbeats are sent to the discovery service using [go-os/discovery](https://github.com/micro/go-os/tree/master/discovery)

