# gogoenv
よしなに，Goのパッケージを作成する内製ツールを作ってみようと思いはじめました．

## Command list

### create essential file
`$ gogoenv init {your application name}`  

#### file list
```
├── Dockerfile: for production
├── Dockerfile.dev: for development has hot reload
├── application: application layer
├── docker-compose.yaml: docker-compose for dev app
├── domain: domain layer
├── infrastructure: infra layer
└── k8s: Deployment.yaml and Service.yaml
```

### create template file for layer architecture
`$ gogoenv create {domain_name}`  
#### file list
- ./domain/services/{domain_name}_service.go
- ./infrastructure/persistence/repositories/{domain_name}_repository.go
### Install
##### Set the app to $GOPATH
`$ go get -u github.com/tozastation/gogoenv`  
##### RUN this command in root directory of gogoenv  
`$ go install`  
