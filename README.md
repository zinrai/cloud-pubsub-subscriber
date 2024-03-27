# Cloud Pub/Sub Subscriber

This is a synchronous pull client for Google Cloud Pub/Sub.

https://cloud.google.com/pubsub/docs/overview

I created this program because I wanted to easily check if the expected data was stored in Google Cloud Pub/Sub.

## Configuration

The configuration for this client is done through a YAML file named `config.yaml`.

- `projectID`: The ID of your Google Cloud project.
- `topicName`: The name of the Cloud Pub/Sub topic from which messages will be pulled.
- `subscriptionName`: The name of the Cloud Pub/Sub subscription to pull messages from.
- `sleepIntervalSeconds`: The duration in seconds for which the program will sleep before attempting to pull messages again.
- `maxOutstandingMessages`: The maximum number of outstanding (unacknowledged) messages that the subscription can have at any given time.

Here is an example `config.yaml` file:

```
cat << EOF > config.yaml
projectID: your-project-id
topicName: your-topic-name
subscriptionName: your-subscription-name
sleepIntervalSeconds: 10
maxOutstandingMessages: 1
EOF
```

## Usage

```
$ go run main.go
2024/03/27 04:11:41 Got message from topic k8s-test: {"insertId":"kouohof73vq2t3yp","labels":{"compute.googleapis.com/resource_name":"gke-k8s-test-k8s-test-0ab12345-c6de","k8s-pod/app_kubernetes_io/component":"gateway","k8s-pod/app_kubernetes_io/instance":"loki","k8s-pod/app_kubernetes_io/name":"loki","k8s-pod/pod-template-hash":"79fd9fbdbd"},"logName":"projects/project-01234567/logs/stderr","receiveTimestamp":"2024-03-27T04:09:19.869384123Z","resource":{"labels":{"cluster_name":"k8s-test","container_name":"nginx","location":"asia-northeast1-b","namespace_name":"test-ns","pod_name":"loki-gateway-79fd9fbdbd-rhhhg","project_id":"project-01234567"},"type":"k8s_container"},"severity":"ERROR","textPayload":"192.0.2.10 - - [27/Mar/2024:04:09:18 +0000] 200 \"GET / HTTP/1.1\" 2 \"-\" \"kube-probe/1.27\" \"-\"","timestamp":"2024-03-27T04:09:18.053557179Z"}
```

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
