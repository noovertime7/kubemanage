# kubemanage
kubemanage是一个简单易用的K8S管理平台，前端使用vue3，后端使用gin+gorm,对于初学k8s开发的同学来说，是一个很方便练手的项目，也可以作为企业二次开发的模板

前端项目地址 https://github.com/noovertime7/kubemanage-web

## 开始部署
### 初始化数据库
需要手动创建数据库，数据表与数据会通过`DBInitializer`自动初始化

```sql
CREATE DATABASE kubemanage;
```

### 运行工程

前端

```shell
git clone https://github.com/noovertime7/kubemanage-web.git

cd kubemanage-web

npm install

npm run serve
```
后端

注意：请确保用户名/./kube  文件夹下存在k8s的kubeconfig文件，后面会改成使用crd，容器部署

开始前请设置配置文件环境变量 KubeManageConfigFile="配置文件位置"，配置文件优先级: 默认配置 < 环境变量< 命令行

```
git clone https://github.com/noovertime7/kubemanage.git

cd kubemanage

go mod tidy

go run cmd/main.go
```
默认用户名密码 admin/chenteng

## Roadmap

- 支持RBAC的权限管理
- 支持多集群管理
- 支持应用一键发布
- 支持资产管理

## 效果演示
首页
![首页](https://github.com/noovertime7/kubemanage/blob/master/img/dashboard.jpg?raw=true)

工作流
![工作流](https://github.com/noovertime7/kubemanage/blob/master/img/wordflow.jpg?raw=true)

deployment
![deployment](https://github.com/noovertime7/kubemanage/blob/master/img/deployment.jpg?raw=true)

pod
![首页](https://github.com/noovertime7/kubemanage/blob/master/img/pod.jpg?raw=true)

POD日志
![POD 日志](https://github.com/noovertime7/kubemanage/blob/master/img/pod_log.jpg?raw=true)

POD终端
![POD 终端](https://github.com/noovertime7/kubemanage/blob/master/img/pod_ter.jpg?raw=true)

service
![service](https://github.com/noovertime7/kubemanage/blob/master/img/service.jpg?raw=true)

configmap
![configmap](https://github.com/noovertime7/kubemanage/blob/master/img/cm_detail.jpg?raw=true)

node
![node](https://github.com/noovertime7/kubemanage/blob/master/img/node.jpg?raw=true)
