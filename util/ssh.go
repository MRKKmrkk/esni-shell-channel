package util

import (
	"time"

	"golang.org/x/crypto/ssh"
)

func GetSSHConf(user string, password string) *ssh.ClientConfig {

	return &ssh.ClientConfig{User: user, Auth: []ssh.AuthMethod{ssh.Password(password)}, HostKeyCallback: ssh.InsecureIgnoreHostKey(), ClientVersion: "", Timeout: 10 * time.Second}

}
