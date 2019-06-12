package Initialize

var (
	Dockerfile                   string
	DockerfileDev                string
	DockerCompose                string
	KubernetesDeploymentManifest string
	KubernetesServiceManifest    string
)

func init() {
	Dockerfile = `# Build Phase
FROM golang:latest as builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o app gogoenv.go
# Runtime Phase
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]`

	DockerfileDev = `FROM golang:1.12-stretch
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/pilu/fresh
RUN go get -u github.com/go-delve/delve/cmd/dlv
COPY . .
EXPOSE 3000
CMD fresh -c runner.conf gogoenv.go`

	DockerCompose = `version: '3'
services:
  dev-app:
    build:
      context: .
      dockerfile: Dockerfile.dev
    volumes:
      - .:/app
    ports:
      - "8080:8080"
    environment:
      - "PORT=8080"`

	KubernetesDeploymentManifest = `apiVersion: apps/v1beta1
kind: Deployment
metadata:
  name: # please write service name 
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: # please write service name 
        version: v1
    spec:
      containers:
        - name: # please write service name 
          image: # please write docker image path
          ports:
            - name: # please write service name 
              containerPort: 8080
              protocol: TCP
      imagePullSecrets:
        - name: # please write your docker registry credential`

	KubernetesServiceManifest = `apiVersion: v1
kind: Service
metadata:
  name: # please write service name 
  labels:
    app: # please write service name 
spec:
  type: LoadBalancer
  selector:
    app: # please write service name 
  ports:
    - port: 8080
      name: # please write service name 
      protocol: TCP`
}
