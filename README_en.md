English | [简体中文](./README.md)
# kubemanage
Kubemanage is a simple and easy to use K8S management platform. The front end uses vue3 and the back end uses gin+gorm. It is a very convenient project for beginners of K8S development, and can also be used as a template for enterprise secondary development

Front end project address https://github.com/noovertime7/kubemanage-web
## Start Deployment
### Initialize database
The database needs to be created manually, and the data table and data will be initialized automatically through the 'DBInitializer'
```sql
CREATE DATABASE kubemanage;
```
### Operation engineering
front
```shell
git clone  https://github.com/noovertime7/kubemanage-web.git
cd kubemanage-web
npm install
npm run serve
```
back-end
Note: Please ensure the username// The kubeconfig file of k8s exists in the kube folder. Later, it will be changed to use crd and container deployment

Before starting, please set the configuration file environment variable KubeManageConfigFile="configuration file location", configuration file priority: default configuration<environment variable<command line
```
git clone  https://github.com/noovertime7/kubemanage.git
cd kubemanage
go mod tidy
go run cmd/main.go
```
Default username password admin/kubemanage
## Roadmap
- Support RBAC permission management
- Support multi cluster management
- Support application one click publishing
- Support asset management
## Issue Specification
- Issue is only used to submit bugs or features and design related content. Other content may be closed directly.
- Before submitting the issue, please search whether the relevant content has been proposed.
## Pull Request Specification
- Please fork a copy to your own project and create a new branch under your own project.
- The commit information should be filled in in the form of `feat (model): description information`, for example, `fix (user): fix xxx bug/fat (user): add xxx`.
- If you want to fix the bug, please give a description in PR.
- Two maintenance personnel are required to participate in merging code: one person reviews and approves, and the other person reviews again. After approval, the code can be merged.
## Generate APi Document
Generate api documents using swag

PS: Please use the latest version of the swag tool. It is recommended to pull the latest code and compile it by yourself, otherwise 'swag init' initialization will fail
```shell
swag init   --pd  -d ./cmd,docs
```
Access after successful generation `http://127.0.0.1:6180/swagger/index.html`
## Effect demonstration
home page
![Home Page](./img/dashboard.jpg?raw=true )

Workflow
![Workflow](./img/wordflow.jpg?raw=true )
deployment
![deployment](./img/deployment.jpg?raw=true )
pod
![Home Page](./img/pod.jpg?raw=true )
POD Log
![POD Log](./img/pod_log.jpg?raw=true )
POD terminal
![POD terminal](./img/pod_ter.jpg?raw=true )
service
![service]( ./img/service.jpg?raw=true )
configmap
![configmap]( ./img/cm_detail.jpg?raw=true )
node
![node]( ./img/node.jpg?raw=true )