# GitLab Config

## Objective

Configure global variables at the root groups level in GitLab

## Why not use Gitlabform ?

Gitlabform does not properly support nested groups amongst other things

## Running locally

```bash
docker-compose up -d
# run a local GitLab instance for testing

cd src/gitlabconf
go test
# run unit tests

go run main.go
# run main program from source
```

## Configuration

See etc/gitlab-config-sample.yml for a configuration file sample

## Usage

### Compile Go binary

```bash
cd src/gitlabconf
go install
```

### Run command

```bash
$GOPATH/bin/gitlabconf <configfile>
```

configfile can be omitted, will default to ./gitlab-config.yml
