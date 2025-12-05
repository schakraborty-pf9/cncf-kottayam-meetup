****Build Application Image Using Cloud Native Buildpack****
=====================
This repository has instruction to build image without Dockerfile (a sample go application) using [Cloud Native Buildpack](https://buildpacks.io/).

Instruction
-------------------------------------------

* Install required tools stated below.
* Clone repository `git clone https://github.com/schakraborty-pf9/cncf-kottayam-meetup.git`.
* Build and publish image.
```sh
cd cncf-kottayam-meetup
pack build <registry_name>/<repo_name>/kottayam-view:v1 --path . --builder paketobuildpacks/builder-jammy-full  --publish --clear-cache
```
* Quick test.
```sh
docker run -d -p 9090:8080 --name kottayam-view <registry_name>/<repo_name>/kottayam-view:v1
curl -s http://0.0.0.0:9090
```

*****Deploy to Kubernetes*****

**Create Deployment**
```sh
kubectl create deployment kottayam-view --image=docker.io/sumanpf9/kottayam-view:v7 --replicas=1
```
**Expose as Service**
```sh
kubectl expose deploy/kottayam-view-7 --name kottayam-view-svc --type=LoadBalancer --port 80 --target-port=8080
```

Note: Use command `pack suggest-builders` to see available builders to use

Tools:
* Installed locally:
[git](https://www.atlassian.com/git/tutorials/install-git), [docker](https://hub.docker.com/search?type=edition&offering=community), [pack](https://buildpacks.io/docs/install-pack/)

