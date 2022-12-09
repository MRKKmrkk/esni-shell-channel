# esni-shell-channel  

[中文文档](https://github.com/MRKKmrkk/esni-shell-channel/blob/main/docs/CN-README.md)   

## Brief
esni-shell-channel is a program dedicated to establish LAN-to-LAN ssh connection.  
we can not establish ssh connection when two hosts in diffrent network(LAN).you can deploy esni-shell-channel on your server which has public IP address, and it's able to transmit your ssh connection on specific port. Then those hosts can establish ssh connection through the port.

## Geting Started

### Deploy 
#### Public Server
[Download mserver](https://github.com/MRKKmrkk/esni-shell-channel/releases/download/v0.0.1/mserver)  **enable transmit service on your server which obsess public IP address**  
```shell
$ nohup ./mserver $LocalAddress &
```

#### Controlled Node 
[Download mclient](https://github.com/MRKKmrkk/esni-shell-channel/releases/download/v0.0.1/mclient)
**enable listener on your internal host**  
```shell
$ ./mclient $ssh-user
```
Enter your password after the listener is started  

"You are able to connect controlled node via ssh(third-party ssh tool, like Xshell, supported as well) on public server's 9657 port then enter your username and password. Plus, sftp is also available. Make sure the listeners on public server and controlled node are enabled."

## Warm prompt
1. Ensure that port 9656, 9657 and 9658 of the public server is allowed by the firewall
2. Ensure that port 9658 and 9656 of the controlled node is allowed by the firewall

