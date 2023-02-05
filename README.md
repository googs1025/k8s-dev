# K8s-Dev

## 简介

原生 k8s 应用指通过 kube-apiserver 进行交互，可直接查询、更新资源的状态的应用。这类应用强依赖 k8s，并在一定程度上可以直接调用 k8s 的 API 及相关资源。

<img src="figures/image-20220725092124197.png" alt="image-20220725092124197" style="zoom:50%;" />

### k8s扩展模式

扩展 K8s 的方式可以分为以下几类：

- 二进制插件：一段可执行的二进制程序，主要由 kubelet、kubectl 使用，例如 Flex 卷插件、网络插件。
- Webhook模式：另一种扩展模式是 K8s 作为客户端来访问外部服务。这种模式叫 Webhook 模式，被访问的外部服务叫做 Webhook Backend。
- Controller 控制器模式：还有一种扩展模式叫做控制器模式，控制器负责执行例行性任务来保证集群尽可能接近其期望状态。典型情况下控制器读取`.spec`字段，运行一些逻辑，然后修改`.status`字段。K8s 自身提供了大量的控制器，并由控制器管理器统一管理。

## 基本概念

- [基本概念](10_concepts/README.md)


## client-go

- [client-go](20_client-go/README.md)

## CRD Operator

- [CRD Operator](30_crd-operator/README.md)


## aggregate

- [aggregate](40_aaserver/README.md)

## Ref

1. [通过自定义资源扩展Kubernetes](https://blog.gmem.cc/crd)
