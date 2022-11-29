# esni-shell-channel

## 简介  
esni-shell-channel是一个基于服务器端口转发的内网ssh连接方案。  
当使用虚拟机或服务器部署在内网且又没有公网IP时，我们无法远程连接这些主机，esni-shell-channel可以通过在云端部署端口转发服务，实现远程主机和内网主机的互访。  

## 使用教程

**公网服务器部署端口转发服务**  
[下载mserver](https://github.com/MRKKmrkk/esni-shell-channel/releases/download/v0.0.1/mserver)  **启动端口转发服务**  
```shell
nohup ./mserver &
```

**内网主机部署监听服务**  

[下载mserver](https://github.com/MRKKmrkk/esni-shell-channel/releases/download/v0.0.1/mclient)
**启用内网主机监听服务**  
```shell
./mclient ssh用户名 
```
启用后需要输入ssh用户对应的密码  

  
当云主机的端口转发服务和内网主机的监听服务全部启动后，即可使用ssh服务连接云主机的9657端口，输入用户密码后即可访问到内网主机。经测试XShell等终端连接工具均可正常使用，
sftp服务也可正常使用。

## 注意事项
1. 需要保证云主机的9656、9657、9658端口畅通且能穿过防火墙
2. 需要保证内网主机的9658、9656端口畅通且能穿过防火墙
