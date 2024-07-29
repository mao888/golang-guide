![https://mao888.github.io/picx-images-hosting/soft-learn/image.26ldszahww.webp](https://mao888.github.io/picx-images-hosting/soft-learn/image.26ldszahww.webp)
# 生命周期管理
## 1. 创建
### 1. 创建资源

- **kubectl run**
- 

   - 创建并运行一个或多个容器镜像。
   - ***创建一个deployment或job来管理容器***。

**语法**：kubectl run NAME --image=image [--env="key=value"] [--port=port] [--replicas=replicas] [--dry-run=bool] [--overrides=inline-json] [--command] -- [COMMAND] [args...]
```bash
kubectl run nginx --replicas=3 --labels="app-nginx-example" --image=nginx:1.10 --port=80
```

- **kubectl create**
```bash
kubectl create deployment nginx --image=nginx
# 根据yaml配置文件创建资源对象
kubectl create -f zookeeper.yaml
# 根据yaml配置文件一次创建Service和RC
kubectl create -f my-service.yaml -f my-rc.yaml
# 创建名称空间
kubectl create namespace bigdata
```

- **kubectl apply**
```bash
kubectl apply deployment nginx --image=nginx
# 使用yaml文件创建资源
kubectl apply -f zookeeper.yaml
```
### 2. 标签操作

- **查询标签**
```bash
kubectl get nodes --show-labels
```

- **添加****标签**
```bash
# 为指定节点添加标签
kubectl label nodes nodeName labelName=value
# 为指定Pod添加标签
kubectl label pod podName -n nsName labelName=value
```

- **修改****标签**
```bash
# 修改节点标签值
kubectl label nodes nodeName
# 修改Pod标签值（需要overwrite参数）
kubectl label pod podName -n nsName labelName=value --overwrite
```

- **删除****标签**
```bash
# 为指定节点删除标签
kubectl label nodes nodeName labelName-
# 删除Pod标签
kubectl label pod podName -n nsName labelName-
```
## 2. 查看
```bash
# 查看集群状态
kubectl get cs

# 查看Pod
kubectl get pods
kubectl get pod
kubectl get po

# 查看指定名称Pod
kubectl get pod mynginx
kubectl get pod/mynginx

# 同时查看多个资源
kubectl get deploy,pods

# 查看Pod端口信息
kubectl get pod,svc

# 特定命名空间资源查看
kubectl get pods -n bigdata

# 查看所有命名空间下的pod信息
kubectl get pod --all-namespaces
kubectl get pods --A

# 获取Pod运行在哪个节点上的信息
kubectl get pod -o wide

# 显示Pod标签信息
kubectl get pods --show-labels

# 查看特定标签的Pod
kubectl get pods -l app=example

# 以JSON格式显示Pod的详细信息
kubectl get pod podName -o json

# 查看RS
kubectl get replicasets -o wide

# 查看Deployments
kubectl get deployments -o wide

# 查看ip和端口，也叫端点
kubectl get ep

# 查看事件
kubectl get ev
```

- **yaml方式**
```bash
# 以yaml格式显示Pod的详细信息
kubectl get pod podName -o yaml
kubectl get pod -f pod.yaml
kubectl get pod -f pod1.yaml -f pod2.yaml

# 用get生成yaml文件
kubectl get deploy/nginx --export -o yaml > my-deploy2.yaml

# 查看资源子节点详情
kubectl explain pods.spec.containers

# 用run命令生成yaml文件，dry-run尝试运行，但不会生成，可用于检查语法错误
kubectl run nginx --image=nginx:latest --port=80 --replicas=3 --dry-run
# 尝试运行，并生成yaml文件
kubectl run nginx --image=nginx:latest --port=80 --replicas=3 --dry-run -o yaml > my-deploy.yaml
```
## 3. 发布
```bash
# 暴露端口
kubectl expose deployment nginx --port=80 --type=NodePort
kubectl expose deployment nginx --port=80 --type=NodePort --target-port=80 --name=nginx-service

# 输出为yaml文件（推荐）
kubectl expose deployment nginx --port=80 --type=NodePort --target-port=80 --name=web1 -o yaml > web1.yaml
kubectl expose deployment nginx -n bigdata --port=80 --type=NodePort
```
## 4. 故障排查
### 1. 资源详情排查
```bash
# 显示Node的详细信息
kubectl describe nodes nodeNamePrefix
# 显示Pod的详细信息
kubectl describe pods  podNamePrefix
# 显示由RC管理的Pod的信息
kubectl describe pods  rcNamePrefix
```
### 2. 资源日志排查
```bash
# 容器日志查看
kubectl logs zk-0
kubectl logs zk-0 -n bigdata
# 跟踪查看容器的日志，相当于tail -f命令的结果
kubectl logs -f <pod-name> -c <container-name>
```
### 3. 进入资源容器
```bash
# 进入容器
kubectl exec -it podName -n nsName /bin/sh    
kubectl exec -it podName -n nsName /bin/bash
```
## 5. 更新
### 1. 版本更新
```bash
kubectl set image deployment/nginx nginx=nginx:1.15
# 记录更新操作命令以便后续查看变更历史
kubectl set image deployment/nginx nginx=nginx:1.15 --record
```
### 2. 编辑更新
```bash
kubectl edit deployment/nginx
```
### 3. 滚动更新
```bash
kubectl rolling-update frontend-v1 frontend-v2 --image=image:v2
kubectl rolling-update frontend --image=image:v2
kubectl rolling-update frontend-v1 frontend-v2 --rollback
```
### 4. 替换更新
```bash
kubectl replace -f zookersts.yaml
```
### 5. 扩缩容
```bash
kubectl scale deployment nginx --replicas=10
```
## 6. 回滚
```bash
# 查看更新过程
kubectl rollout status deployment/nginx --namespace=nsName
# 如果更新成功, 返回值为0 
kubectl rollout status deployment nginx-deployment --watch=false | grep -ic waiting

# 查看变更历史版本信息
kubectl rollout history deployment/nginx
kubectl rollout history deployment/nginx --revision=3 --namespace=nsName

# 终止升级
kubectl rollout pause deployment/nginx --namespace=nsName

# 继续升级
kubectl rollout resume deployment/review-demo --namespace=nsName

# 回滚版本
kubectl rollout undo deployment/nginx --namespace=nsName
kubectl rollout undo deployment/nginx --to-revision=3  --namespace=nsName
```
## 7. 清理
```bash
# 删除资源
kubectl delete deploy/nginx
kubectl delete svc/nginx-service

# 删除所有Pod
kubectl delete pods --all

# 删除所有包含某个label的Pod和Service
kubectl delete pod,service -l name=labelName

# 基于yaml定义的名称删除
kubectl delete -f pod.yaml

# 删除指定命名空间
kubectl delete ns nsName

# 删除指定命名空间的资源
kubectl delete pod zk-0 -n bigdata
kubectl delete pod --all -n bigdata

# 删除计时（观察删除总耗时）
time -p kubectl delete pod podName

# 强制删除（默认：30s）
# 指定删除延迟时间：0s，整体删除时间会明显降低
kubectl delete pod podName -n nsName --grace-period=0 --force
# 以下两行命令功能相同（grace-period=1，等价于now，立即执行）
kubectl delete pod podName -n nsName --grace-period=1
kubectl delete pod podName -n nsName --now
# 删除所有Pods
kubectl delete pods --all --force --grace-period=0
```
# 常用操作命令
| **类型** | **命令** | **描述** |
| --- | --- | --- |
| **基础命令** | **create** | 通过文件名或标准输入创建资源 |
| **expose** | 将一个资源公开为一个新的Service |  |
| **run** | 在集群中运行一个特定的镜像 |  |
| **set** | 在对象上设置特定的功能 |  |
| **get** | 显示一个或多个资源 |  |
| **explain** | 文档参考资料 |  |
| **edit** | 使用默认的编辑器编辑资源 |  |
| **delete** | 通过文件名、标准输入、资源名称或标签选择器来删除资源 |  |
| **部署命令** | **rollout** | 管理资源的发布 |
| **rolling-update** | 对给定的复制控制器滚动更新 |  |
| **scale** | 扩容或缩容Pod、Deployment、ReplicaSet、RC或Job |  |
| **autoscale** | 创建一个自动选择扩容或缩容并设置Pod数量 |  |
| **集群管理命令** | **certificate** | 修改证书资源 |
| **cluster-info** | 显示集群信息 |  |
| **top** | 显示资源（CPU、Memory、Storage）使用。需要Heapster运行 |  |
| **cordon** | 标记节点不可调度 |  |
| **uncordon** | 标记节点可调度 |  |
| **drain** | 维护期间排除节点（驱除节点上的应用，准备下线维护） |  |
| **taint** | 设置污点属性 |  |
| **故障诊断和调试命令** | **describe** | 显示特定资源或资源组的详细信息 |
| **logs** | 在一个Pod中打印一个容器日志。如果Pod只有一个容器，容器名称是可选的 |  |
| **attach** | 附加到一个运行的容器 |  |
| **exec** | 执行命令到容器 |  |
| **port-forward** | 转发一个或多个本地端口到一个Pod |  |
| **proxy** | 运行一个proxy到Kubernetes API Server |  |
| **cp** | 拷贝文件或目录到容器 |  |
| **auth** | 检查授权 |  |
| **高级命令** | **apply** | 通过文件名或标准输入对资源应用配置 |
| **patch** | 使用补丁修改、更新资源的字段 |  |
| **replace** | 通过文件名或标准输入替换一个资源 |  |
| **convert** | 不同的API版本之间转换配置文件 |  |
| **设置命令** | **label** | 更新资源上的标签 |
| **annotate** | 更新资源上的注释 |  |
| **completion** | 用于实现kubectl工具自动补全 |  |
| **其他命令** | **api-versions** | 打印支持的API版本 |
| **config** | 修改kubeconfig文件（用于访问API，比如配置认证信息） |  |
| **help** | 所有命令帮助 |  |
| **plugin** | 运行一个命令行插件 |  |
| **version** | 打印客户端和服务版本信息 |  |

## 1. 获取帮助
```bash
# 检查kubectl是否安装
rpm -qa | grep kubectl
# 获取kubectl及其子命令帮助方法
kubectl --help
kubectl create --help
```
## 1. Worker上执行kubectl
```bash
# Worker节点上执行
mkdir -p ~/.kube
scp master1:/root/.kube/config ~/.kube/
# 验证（查看K8s集群节点列表）
kubectl get nodes
```
## 2. api相关操作命令
```bash
# 查看api版本信息
kubectl api-versions
# 查看api资源列表
kubectl api-resources
```
## 3. K8s相关进程操作命令
```bash
netstat -lntp | grep kube-proxy
netstat -tnlp | grep kubelet
```
## 4. 节点操作命令

- **加入新节点**
```bash
# 加入新节点，在master节点上执行，将输出再到新节点上执行
kubeadm token create --print-join-command
```

- **驱逐节点**
```bash
# 驱逐节点的Pod
kubectl drain nodeName
```

- **节点下线**
```bash
# 将节点标记为不可调度，不影响现有Pod（注意daemonSet不受影响）
kubectl cordon nodeName
```

- **节点上线**
```bash
# 维护结束，节点重新投入使用
kubectl uncordon nodeName
```

- **污点设置**
```bash
# 设置污点
kubectl taint nodes nodeName key1=value1:NoSchedule
kubectl taint nodes nodeName key1=value1:NoExecute
kubectl taint nodes nodeName key2=value2:NoSchedule
# 删除污点
kubectl taint nodes nodeName key1:NoSchedule-
kubectl taint nodes nodeName key1:NoExecute-
kubectl taint nodes nodeName key2:NoSchedule-
# 查看污点详情
kubectl describe nodes nodeName
```
# 资源创建实例
## Namespace
### 命令行方式
```bash
kubectl create namespace bigdata
```
### yaml方式
```bash
vi ns-test.yaml
```
编排文件如下：
```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: bigdata
```
执行yaml文件：
```bash
kubectl apply -f ns-test.yaml
```
验证：
```bash
kubectl get namespaces
kubectl get namespace
kubectl get ns
```
清除：
```bash
kubectl delete -f ns-test.yaml
kubectl delete ns bigdata
```
## Pod
### 命令行方式
未提供直接创建Pod的命令，命令行方式一般通过创建Deployment、RC、RS等资源间接创建Pod。
### yaml方式
```bash
vi pod-test.yaml
```
编排文件如下：
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: pod1
spec:
  containers:
  - name: nginx-containers
    image: nginx:latest
```
执行yaml文件：
```bash
kubectl apply -f pod-test.yaml
```
验证：
```bash
kubectl get pods
kubectl get pod
kubectl get po
kubectl describe pod pod1
kubectl get pods -o wide
curl http://172.16.189.68
```
清除：
```bash
kubectl delete -f pod-test.yaml
kubectl delete pod pod1
```
## Service
### 命令行方式
```bash
kubectl run nginx-app --image=nginx:latest --image-pull-policy=IfNotPresent --replicas=2
kubectl expose deployment.apps nginx-app --type=ClusterIP --target-port=80 --port=83
```
参数说明：

- **expose**：创建service。
- **deployment.apps**：控制器类型。
- **nginx-app**：应用名称，也是service名称。
- **--type=****ClusterIP**：指定service类型。
- **--target-port=80**：指定Pod中容器端口。
- **--port=80**：指定service端口。

验证：
```bash
kubectl get service
kubectl get svc
kubectl get endpoints
kubectl get ep
curl http://10.104.173.230:83
kubectl get all
```
清除：
```bash
kubectl delete service nginx-app
kubectl delete svc nginx-app
```
### yaml方式
```bash
vi nginx-service.yaml
```
编排文件如下：
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata: 
  name: nginx-app
  labels:
    app: nginx
spec: 
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginxapp
        image: nginx:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 80

---
apiVersion: v1
kind: Service
metadata:
  name: nginx-app-svc
  labels:
    name: nginx-app-svc
spec:
  type: ClusterIP
  ports: 
  - protocol: TCP
    port: 83
    targetPort: 80
  selector:
    app: nginx

---
apiVersion: v1
kind: Service
metadata:
  name: nginx-app-svc2
  labels:
    name: nginx-app-svc2
spec:
  type: NodePort
  ports: 
  - protocol: TCP
    port: 83
    targetPort: 80
    nodePort: 30083
  selector:
    app: nginx
```
执行yaml文件：
```bash
kubectl apply -f nginx-service.yaml
```
验证：
```bash
kubectl describe deployment nginx-app
kubectl describe svc nginx-app-svc
kubectl get service
kubectl get svc
kubectl get endpoints
kubectl get ep
# nginx-app-svc
curl http://10.107.141.109:83
# nginx-app-svc2
curl http://192.168.216.100:30083
# 查看k8s集群指定端口的侦听状态
ss -anput | grep ":30083"
kubectl get all
```
清除：
```bash
kubectl delete -f nginx-service.yaml
kubectl delete service nginx-app-svc
kubectl delete svc nginx-app-svc
```
# 常用控制器
## 1. Deployment
### 命令行方式
```bash
kubectl run nginx-app --image=nginx:latest --image-pull-policy=IfNotPresent --replicas=2
```
参数说明：

- **nginx-app**：Deployment控制器类型的应用名称。
- **--image=nginx****:****latest**：应用运行的Pod中的Container所使用的镜像。
- **IfNotPresent**：Container容器镜像下载策略，如果本地有镜像，使用本地，如果本地没有镜像，下载镜像。
- **--**replicas**=****2**：是指应用运行的Pod共计2个副本，这是用户的期望值，Deployment控制器中的ReplicaSet控制器会一直监控此应用运行的Pod副本状态，如果数量达不到用户期望，就会重新拉起一个新的Pod，会让Pod数量一直维持在用户期望值数量。

验证：
```bash
kubectl get deployment.apps
kubectl get deployment
kubectl get deploy
kubectl get replicaset
kubectl get rs
kubectl get all
```
清除：
```bash
kubectl delete deployment nginx-app
```
### yaml方式
```bash
vi nginx-deployment.yaml
```
编排文件如下：
```yaml
apiVersion: apps/v1
kind: Deployment
metadata: 
  name: nginx-app
  labels:
    app: nginx
spec: 
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginxapp
        image: nginx:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 80
```
执行yaml文件：
```bash
kubectl apply -f nginx-deployment.yaml
```
验证：
```bash
kubectl get deployment.apps
kubectl get deployment
kubectl get deploy
kubectl get replicaset
kubectl get rs
kubectl get all
kubectl describe deployment nginx-app
kubectl get pods -o wide
curl http://172.16.189.77
curl http://172.16.235.138
```
清除：
```bash
kubectl delete -f nginx-deployment.yaml
kubectl delete deployment nginx-app
```
## 2. ReplicaSet
### 命令行方式
### yaml方式
## 3. StatefulSet
### 命令行方式
### yaml方式
## 4. DaemonSet
### 命令行方式
### yaml方式
## 5. Job
### 命令行方式
### yaml方式
## 6. CronJob
### 命令行方式
### yaml方式
# 操作命令补充说明
## 1. create和apply的异同点

- **create**

先删除所有现有的东西，重新根据yaml文件生成新的。所以要求yaml文件中的配置必须是完整的。

- **apply**

根据配置文件里面列出来的内容，升级现有的。所以yaml文件的内容可以只写需要升级的属性。apply命令将配置应用于资源。 如果资源不在那里，那么它将被创建。
![](https://cdn.nlark.com/yuque/0/2020/png/788484/1604280290859-1809687a-bf36-48db-9d92-297a240228cb.png?x-oss-process=image%2Fwatermark%2Ctype_d3F5LW1pY3JvaGVp%2Csize_19%2Ctext_cG9sYXJpcw%3D%3D%2Ccolor_FFFFFF%2Cshadow_50%2Ct_80%2Cg_se%2Cx_10%2Cy_10#averageHue=%23f0ece9&errorMessage=unknown%20error&id=LtD4O&originHeight=337&originWidth=671&originalType=binary&ratio=1&rotation=0&showTitle=false&status=error&style=none)
从执行的角度来看，如上所示，在kubectl create和kubectl apply之间第一次创建资源时没有区别。 但是，第二次kubectl create会抛出错误。简单来说，如果在单个文件上运行操作以创建资源，则create和apply基本相同。 但是， apply允许您在目录下的多个文件上同时创建和修补。
