# proxy-NG

Proxy-Next-Generation项目基于[goproxy](https://github.com/ouqiang/goproxy.git) ，增加了分流功能

## 架构

![架构图](./imgs/架构图.jpg)

## Feature
- [x] 可以将单个流量分流到其他上游代理，而不阻塞Proxy-NG主线程代理，常见于代理黑盒扫描器等
- [x] 支持代理流量预处理
- [x] 支持多节点分发