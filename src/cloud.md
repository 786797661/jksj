#Kubernetes
## Docker
### 原理
docker\
docker 其实是宿主机的一个进程，只不过是通过了namespace进行了参数的过滤，
他共享了宿主机的CPU等资源。\
docker是进程隔离，虚拟机是物理隔离\
对于 Docker 等大多数 Linux 容器来说，
Cgroups 技术是用来制造约束的主要手段，而 Namespace 技术则是用来修改进程视图的主要方法。
###优点
“敏捷”和“高性能”是容器相较于虚拟机最大的优势，也是它能够在 PaaS 这种更细粒度的资源管理平台上大行其道的重要原因。\
敏捷是指快速打包，快速发布 高性能是指容器进程本来就是系统里的普通进程，没有虚拟化层的性能损失。 但缺点是没有虚拟化层的隔离性
###隔离
####Namespace 可见性隔离
docker 通过linux的Nmaespace实现了docker的进程隔离
也就是说docker内部的进程编号PID是1但是在宿主机上的进程PID是100.\
进程隔离并不是彻底的隔离?\
1. docker作为一个进程在宿主机中和其他进程是平等，也就是说其他进程可能会吃掉docker的资源，反过来一个docker也会吃掉其他进程的资源
2. Linux中Namespace并不能对所有的文件都做隔离，比如说时间。如果我在一个容器中修改了时间，那么就会影响到宿主机以及宿主机上运行的其他容器
3. 容器的对外安全也是一个问题，恶意攻击可以通过容器攻击宿主机以及宿主机上的其他容器
####Cgroups （限制）资源隔离
Linux Cgroups 的全称是 Linux Control Group。它最主要的作用，就是限制一个进程组能够使用的资源上限，包括 CPU、内存、磁盘、网络带宽等等。
#### Chanroot
更改进程的根目录
#### 问题
1. 同生命周期
   容器是『单进程』模型，只有 PID=1 的进程才会被 Dockerd 控制，即 pid=1 的进程挂了 Dockerd 能够感知到，但是其它的进程却不受 Dockerd 的管理，当出现孤儿进程的时候，管理和调度是个问题
2. /proc 文件系统的问题
   容器中 top 等命令看到的CPU和内存等是宿主机的信息，lxcfs 可解决此问题
#### 核心原理
对 Docker 项目来说，它最核心的原理实际上就是为待创建的用户进程：
1. 启用 Linux Namespace 配置；
2. 设置指定的 Cgroups 参数；
3. 切换进程的根目录（Change Root）。
## Kubernetes 一键部署利器之 kubeadm
### 安装

1. 使用 kubeadm 的第一步，是在机器上手动安装 kubeadm、kubelet 和 kubectl 这三个二进制文件。
2. 使用“kubeadm init”部署 Master 节点
   1) Preflight Checks:当你执行 kubeadm init 指令后，kubeadm 首先要做的，是一系列的检查工作，以确定这台机器可以用来部署 Kubernetes
   2) 生成 Kubernetes 对外提供服务所需的各种证书和对应的目录:Kubernetes 对外提供服务时，除非专门开启“不安全模式”，
   否则都要通过 HTTPS 才能访问 kube-apiserver。这就需要为 Kubernetes 集群配置好证书文件。
   kubeadm 为 Kubernetes 项目生成的证书文件都放在 Master 节点的 /etc/kubernetes/pki 目录下。
   在这个目录下，最主要的证书文件是 ca.crt 和对应的私钥 ca.key。
   3) 证书生成后，kubeadm 接下来会为其他组件生成访问 kube-apiserver 所需的配置文件。
   这些文件的路径是：/etc/kubernetes/xxx.conf
   4) kubeadm 会为 Master 组件生成 Pod 配置文件:Kubernetes 有三个 Master 组件 kube-apiserver、kube-controller-manager、kube-scheduler，而它们都会被使用 Pod 的方式部署起来。
   5) kubeadm 还会再生成一个 Etcd 的 Pod YAML 文件，用来通过同样的 Static Pod 的方式启动 Etcd
   6) kubeadm 就会为集群生成一个 bootstrap token。在后面，只要持有这个 token，任何一个安装了 kubelet 和 kubadm 的节点，都可以通过 kubeadm join 加入到这个集群当中。
   7) 在 token 生成之后，kubeadm 会将 ca.crt 等 Master 节点的重要信息，通过 ConfigMap 的方式保存在 Etcd 当中，供后续部署 Node 节点使用。这个 ConfigMap 的名字是 cluster-info。
   8) kubeadm init 的最后一步，就是安装默认插件。Kubernetes 默认 kube-proxy 和 DNS 这两个插件是必须安装的。它们分别用来提供整个集群的服务发现和 DNS 功能。其实，这两个插件也只是两个容器镜像而已，所以 kubeadm 只要用 Kubernetes 客户端创建两个 Pod 就可以了。
      

###kubeadm join
kubeadm init 生成 bootstrap token 之后，你就可以在任意一台安装了 kubelet 和 kubeadm 的机器上执行 kubeadm join 了
可是，为什么执行 kubeadm join 需要这样一个 token 呢？\
因为，任何一台机器想要成为 Kubernetes 集群中的一个节点，就必须在集群的 kube-apiserver 上注册。可是，要想跟 apiserver 打交道，这台机器就必须要获取到相应的证书文件（CA 文件）。可是，为了能够一键安装，我们就不能让用户去 Master 节点上手动拷贝这些文件。所以，kubeadm 至少需要发起一次“不安全模式”的访问到 kube-apiserver，从而拿到保存在 ConfigMap 中的 cluster-info（它保存了 APIServer 的授权信息）。
而 bootstrap token，扮演的就是这个过程中的安全验证的角色。只要有了 cluster-info 里的 kube-apiserver 的地址、端口、证书，kubelet 就可以以“安全模式”连接到 apiserver 上，这样一个新的节点就部署完成了。

```shell
$ kubeadm join 10.168.0.2:6443 --token 00bwbx.uvnaa2ewjflwu1ry --discovery-token-ca-cert-hash sha256:00eb62a2a6020f94132e3fe1ab721349bbcd3e9b94da9654cfe15f2985ebd711
```
### 污点
默认情况下 Master 节点是不允许运行用户 Pod 的。而 Kubernetes 做到这一点，依靠的是 Kubernetes 的 Taint/Toleration 机制。它的原理非常简单：一旦某个节点被加上了一个 Taint，即被“打上了污点”，那么所有 Pod 就都不能在这个节点上运行，因为 Kubernetes 的 Pod 都有“洁癖”。除非，有个别的 Pod 声明自己能“容忍”这个“污点”，即声明了 Toleration，它才可以在这个节点上运行。其中，为节点打上“污点”（Taint）的命令是：
```shell
$ kubectl taint nodes node1 foo=bar:NoSchedule
```

## 安装集群
1. 安装kuberadm
配置yum
```shell
cat > /etc/yum.repos.d/kubernetes.repo <<EOF 
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

```
安装
```shell

 yum install -y kubelet-1.23.5 kubeadm-1.23.5 kubectl-1.23.5
systemctl start kubelet   # 此时kubelet没有起来，因为还没有init

```
2. init
```shell
kubeadm init --image-repository=registry.cn-qingdao.aliyuncs.com/k8s-images1  --kubernetes-version=v1.22.2 --service-cidr=10.96.0.0/16 --pod-network-cidr=10.244.0.0/16

  --image-repository：选择用于拉取镜像的镜像仓库（默认为“k8s.gcr.io” ）
  --kubernetes-version：选择特定的Kubernetes版本 
  --service-cidr：为服务的VIP指定使用的IP地址范围（默认为“10.96.0.0/12”）
  --pod-network-cidr：指定Pod网络的IP地址范围。如果设置，则将自动为每个节点分配CIDR。

```

### 安装踩过的坑
1. kubelet 无法启动
原因是 docker的驱动和 kubelet 的不一致
docker查看方法：

```shell
#查看docker的驱动
docker info|grep Drive
#修改配置 增加"exec-opts": ["native.cgroupdriver=systemd"]
vim /etc/docker/daemon.json
#加载配置
systemctl daemon-reload
#重启docker
systemctl restart docker

```
kubelet 
```shell
##查看kubelet状态
systemctl status kubelet
##查看kubelet证书到期时间，拷贝前面的密钥
vim /etc/kubernetes/kubelet.conf
##对密钥解码
echo -n "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvakNDQWVhZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeU1EUXlNREE1TURjeU9Gb1hEVE15TURReE56QTVNRGN5T0Zvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBT2hpCmFtM0syd0pVYURmOVVqRUYzMHN1ZGU0T2FlTFNTNCtFZ3EvMzk2OTlyRnE0MEpPUmF1T1RGZXRzQk12aGNaT2YKK0hSbWtUd0h6QzRnMmdWUTc2MzA2L1RzRndSd0MvNU4wN21jb3hVeXd2QmJUQmxBYlF5blJ5TWEybW4rajdnKwpNdE52L2duUHU0ajBUSGFZNzI2WWVSSkRrY3pLRGE4dnVDOVo2NWFLNGMyMmNyZDVacHFkSnJIV3VkRFhrK3JQCmpKcHNTRnJnZFlQYXJHdmxrUEhYekJhcVN0dEdwVGNvQXcwTWRpKzBMa0hNZGFxRnJpNVVaV3RvSkQvT1dTTCsKNXI1QVlqYzRwS3pLQTVDV3JKUTBMNjhMWDVZL1VlV0Fjc2xjck0zcUFHT1FBcXhJams0SklCTkNxcVRpNm9wUQptbWRJRHVxZE5RK2NwQWQvS2JVQ0F3RUFBYU5aTUZjd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZHZDF2K2M1SE5ydEdWK2E5ZHBweElCcjRPWklNQlVHQTFVZEVRUU8KTUF5Q0NtdDFZbVZ5Ym1WMFpYTXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBR2FYQXZxOWNvaWNNR1pmYStBWAp3ZWdCL0FUMWtZVzN6NUhiVVhvTVBFQ1VWbmprOW1xK1paQUl5YUxrU2w1K1JYUll5ZmpqT2VwRTBBUVhLei9XClFUMFBZT3RrWkd3cENhRHYyOEtmazdYOEtydDNMb0hXQkYxR2lIYmJtM1Q4YXhsbUE0NWhKQzR5ZlFZcmxPRjcKRnViV0N1Q1A3Z2FkbG9yN2FheWkrZ1QzT1JacGpXbm0vV0JvemVQY3hBM0NTekFYUkJnNVZabEZ0OU95a2k5QgptWlhkZTRFc25NdWJiM2U3YjNaRSswRXM1OWZybWVqQlJFVjlsRGxsUEFuWWRZZU5TY1BRdHlUeU1NOVpOakt2CndyODNmZDFYVG56WWR3U0QwSFRqVTRnZkw1RjUvUmZ1d0xRNUMrbzR2N1pCQitaeWxleEdBbzYzWjFkRmJ4Z3IKdUxZPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==" | base64 --decode > kubelet.crt 
##查看日期
openssl x509 -in kubelet.crt -text -noout
##查看kubelet配置信息
vim var/lib/kubelet/config.yaml

```
安装成功
```shell

Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:
##执行这段就可以加入集群了哈哈哈
kubeadm join 192.168.1.121:6443 --token i5nxqy.vthuh1amh4gqujbh \
	--discovery-token-ca-cert-hash sha256:73d57faba46f4647eaf7e557cb704e94efe481a52f731c84dec764ddba8ee5b3 

```