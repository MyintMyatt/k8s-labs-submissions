# Log Output Sample With Go

## Build Docker Image
```shell
docker build -t log-output:1.1 .
```

## Import the image into k3d (If use k3d)
```shell
k3d import image log-output:1.1
```

## Deploy to Kubernates
```shell
kubectl apply -f manifests/deployment.yaml
```
You will get:
```text
deployment.apps/log-output created
```

## Check the Deployment
```shell
kubectl get deployments
```
The result is:
```text
NAME                READY   UP-TO-DATE AVAILABLE AGE
log-output          1/1      1           1        6s
```

## Check the Pod
```shell
kubectl get pods
```
The result is like that:
```
NAME                           READY   STATUS RESTARTS   AGE
log-output-xxxxxxxxxx-xxxxx    1/1     Running   0          10s
```


## Check the application logs
```shell
kubectl logs deployment/log-output
```

You will see logs like that:
```text
2026-08-29T16:23:22.111686224Z: 84fbe6fc-d4ec-45a9-ba05-126acd6cf917
2026-08-29T16:23:27.114474749Z: 84fbe6fc-d4ec-45a9-ba05-126acd6cf917
2026-08-29T16:23:31.915225496Z: 84fbe6fc-d4ec-45a9-ba05-126acd6cf917
2026-08-29T16:23:36.910629236Z: 84fbe6fc-d4ec-45a9-ba05-126acd6cf917
```