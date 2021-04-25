# Banking

## Table of contents
1. [Description](#Description)
2. [Documentation](#Documentation)
3. [Technologies](#Technologies)
4. [Deployment](#Deployment)
6. [Testing](#Testing)
7. [Linting](#Linting)
8. [Future_Work](#Future_Work)

## 1. Description

banks provides certain basic features for financial institution

## 2. Documentation

> documentation resides at ./docs folder.  

## 3. Technologies

Project is created with:

* Go
* Echo
* Postgres
* Docker
* Mockery

## 4. Deployment

> Note: The app will start with a seeder for customers. 

```
$ docker-compose up -d
```

## 6. Testing

> 1\. First get a bash shell in the container

```
$ docker-compose exec bank_go bash
```

> 2\. Execute all test cases with coverage

```
$ go test -cover -v ./...
```

> 3\. Execute test cases for a specific package [optional]

```
$ go test -cover -v ./internal/server/... -v
$ go test -cover -v ./internal/service/... -v
```

## 7. Linting
Project uses [golangci-lint](https://golangci-lint.run/). It is a go linter aggregator that can enable up to 48 linters.

#### 7.1 Configuartion

golanci-lint configuration is found in .golangci.yml file.

#### 7.2 Installation

```
# binary installation for linux and Windows, binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.35.2
```

Check if the tool is correctly installed

```
golangci-lint --version
```

#### 7.3 Run the tool with the enabled linter

```
golangci-lint run
```

golangci-lint print out warning messages in command-line related to the enabled linters in order to fix them.

#### 7.4 Linters commands to automatically fix warning messages provided

To format all files in the project based on the gofmt linter. [Ref](https://stackoverflow.com/a/13333931/5486622)

```
gofmt -s -w -l .
```

To fix go import packages linting warnings based on goimport linter. [Ref](https://stackoverflow.com/a/59964885/5486622)

```
goimports -local ./ -w .
```
[Guide](https://stackoverflow.com/a/38714480/5486622) How you should group your package based on golang structure.

## 8. Future_Work

* Implement more coverage with Unit Test
* Implement intergration tests
* Implement functional tests
