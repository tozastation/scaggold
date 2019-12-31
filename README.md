# sgoffold
Goで開発するときのサポートツールです．
## Command list
- [ ] : `sgoffold init {package name}`
    - プロジェクトの初期化をします． (golang-standards/project-layout)
- [ ] : `sgoffold addService {service name}`
    - アプリケーションの初期化をします．
### Description

#### sgoffold init

#### sgoffold addService
- architecture
    - generate clean architecture package construction
- Dockerfile : [Dockerfile.dev,Dockerfile.prod]
- docker-compose: [docker-compose.dev.yaml]
    - with MySQL or MongoDB option
- kubernetes: [deployment.yaml,service.yaml,namespace.yaml]
## Getting Start
### Install
`go install github.com/tozastation/sgoffold`