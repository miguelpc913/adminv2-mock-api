

***<p style="text-align: center;">Welcome to Admin Mock API</p>***

<p align="center">
  <img src="/.github/media/go-clorian.png" />
</p>

# Table of Contents
1. [🐋 Docker Compose](#run-it-with-docker-compose)
2. [📦 Run it on Devbox](#run-it-with-devbox)
3. [🧪 Pipeline](#pipeline)
4. [📄 Notes](#notes)

[![Terraform](https://img.shields.io/badge/Terraform-1s?style=flat&logo=terraform&labelColor=white&color=back)](https://www.terraform.io/)
[![Go version](https://img.shields.io/badge/Go%201.22-1s?style=flat&logo=go&labelColor=white)](https://tip.golang.org/doc/go1.22)
[![Static Badge](https://img.shields.io/badge/Docker-1s?style=flat&logo=Docker&labelColor=white&color=blue)](https://www.docker.com/)
[![Deploy to Amazon ECS](https://github.com/tiqueteo/adminv2-mock-api/actions/workflows/deploy.yml/badge.svg?branch=infra)](https://github.com/tiqueteo/adminv2-mock-api/actions/workflows/deploy.yml)


Mock Admin v2 API implementation for frontend testing purposes.

The following project is developed on Go 1.22 the localhost run are set by docker compose and [devbox](https://www.jetify.com/devbox/docs/quickstart/).

## Run it with Docker Compose

### START

```
docker-compose up -d
```
### STOP
```
docker compose down
docker compose stop
```
## Requirements to launch:

- Docker
- Docker compose
- [Devbox](https://www.jetify.com/devbox/docs/quickstart/)

## Run it with Devbox:

1. install [Devbox](https://www.jetify.com/devbox/docs/quickstart/) on your computer.
2. pull the project.
3. issue the following command in order:

### START app:
```
devbox init
devbox shell
task db
task air
```

- **devbox init**: initialize devbox tool
- **devbox shell**: go into devbox sandbox
- **task db**: start database by using [Task](https://taskfile.dev/)
- **task air**: start Air by using [Task](https://taskfile.dev/)


### STOP app:
```
devbox run stop_db
```

## Requirements to launch:

- [Devbox](https://www.jetify.com/devbox/docs/quickstart/)


## Pipeline

![](/.github/media/workflow.png)

The workflow designed does the following steps:

* Redefine a new task-definition
* Build and push a new container iamge to the ECR registry
* Deploy a new service at ECS by using new image and task definition.

## Notes