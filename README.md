**Build Image Using Cloud Native Buildpack**
=====================
This repository has instruction to build image without Dockerfile (a sample go app) using [Cloud Native Buildpack](https://buildpacks.io/).

Instruction
-------------------------------------------

* Install required tools stated below.
* Clone repository `git clone https://github.com/suchakra012/cnb-go-app.git`.
* Build and publish image.
```sh
cd cnb-sample-go
pack build --publish <image-repo> --builder cloudfoundry/cnb:bionic
```
* Quick test.
```sh
docker run -d -p 9090:8080 --name sample-go <image-repo>
curl -s http://0.0.0.0:9090
```

Note: Use command `pack suggest-builders` to see available builders to use

Tools:
* Installed locally:
[git](https://www.atlassian.com/git/tutorials/install-git), [docker](https://hub.docker.com/search?type=edition&offering=community), [pack](https://buildpacks.io/docs/install-pack/)

