**_<p style="text-align: center;">Welcome to Admin Mock API</p>_**

<p align="center">
  <img src="/.github/media/go-clorian.png" />
</p>

[![Terraform](https://img.shields.io/badge/Terraform-1s?style=flat&logo=terraform&labelColor=white&color=back)](https://www.terraform.io/)
[![Go version](https://img.shields.io/badge/Go%201.22-1s?style=flat&logo=go&labelColor=white)](https://tip.golang.org/doc/go1.22)
[![Docker](https://img.shields.io/badge/Docker-1s?style=flat&logo=Docker&labelColor=white&color=blue)](https://www.docker.com/)
[![Deploy to Amazon ECS](https://github.com/tiqueteo/adminv2-mock-api/actions/workflows/deploy.yml/badge.svg?branch=infra)](https://github.com/tiqueteo/adminv2-mock-api/actions/workflows/deploy.yml)

# Table of Contents

1. [🐋 Docker Compose](#run-it-with-docker-compose)
2. [📦 Run it on Devbox](#run-it-with-devbox)
3. [🧪 Pipeline](#pipeline)
4. [📄 Notes](#notes)

Mock Admin v2 API implementation for frontend testing purposes, by using [air](https://github.com/cosmtrek/air).

The following project is developed on Go 1.22 the localhost run are set by docker compose and [devbox](https://www.jetify.com/devbox/docs/quickstart/).

🛂 **NOTE for Developers**: Your workdir it's **"/admin"** folder. Do not change any other directory out of your workdir. If you need to make some changes (improve/update) out of your workdir, create a new branch and open a PR. [WIP]

## Run it with Docker Compose

1. compose.yml
   On this compose it will use the local.Dockerfile build as image.

### START

```
docker compose -f compose.yml up -d
```

**compose-name**: compose.yml

### STOP

```
docker compose down
docker compose stop
```

## Requirements to launch:

- Docker

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

- Redefine a new task-definition
- Build and push a new container iamge to the ECR registry
- Deploy a new service at ECS by using new image and task definition.

## Notes

Testing project by using a cosmtrek/air as image.

```
docker run -it --rm -w "$HOME/projects/github/adminv2-mock-api/admin" -v $(pwd)/admin:$HOME/projects/github/adminv2-mock-api/admin -p 8080:8080 cosmtrek/air
```

**No healcheck implemented** on this solution due we have alredy defined the Healcheck on ALB TG [admin-mock-api](arn:aws:elasticloadbalancing:eu-west-1:227320912480:targetgroup/admin-mock-api/1c0ea9a901028fa6). further info: [aws question](https://repost.aws/questions/QUdmR0oMn2Spa61RpKGWyPfg/ecs-should-i-use-alb-healthchecks-container-healthchecks-or-both)
