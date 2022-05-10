## Сontainerized golang hello world application, ready for deploy to k8s and run with docker-compose. Ready for building image with docker and [werf  1.1.*](werf.io) .

## Build with docker
#
```bash
docker build -t cassizx/hello_world_go:docker .
docker push cassizx/hello_world_go:docker
```


## Build with werf. Required werf version = 1.1.*
#
### Build and publish image
```bash
werf build-and-publish --stages-storage=cassizx/hello_world_go --images-repo=cassizx/hello_world_go --tag-custom=werf --tag-custom=0.1.0
```

## Local run
#
### werf
```bash
werf run --docker-options="-p 9090:9090" --stages-storage=:local
```
### docker-compose
```bash
docker-compose up
docker-compose up -d
```

Go to http://localhost:9090


## Deploy to k8s
#
### Deploy with [helm](https://helm.sh/docs/intro/install/)
```bash
helm upgrade --install --wait dev deployments/hello_world_go/ --namespace dev --create-namespace
```

### Deploy with [helmfile](https://github.com/helmfile/helmfile)
```bash
helmfile apply
```

Go to application URL.