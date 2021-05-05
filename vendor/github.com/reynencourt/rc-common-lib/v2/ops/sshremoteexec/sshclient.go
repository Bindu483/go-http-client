package sshremoteexec

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"time"
)

type SshConnection struct {
	User        string
	Ip          string
	KeyLocation string
	Timeout     time.Duration
}

func (s *SshConnection) Execute(cmdLine []string) error {

	d, err := ioutil.ReadFile(s.KeyLocation)
	if err != nil {
		return err
	}

	sshKey, err := ssh.ParsePrivateKey(d)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:    s.User,
		Timeout: s.Timeout,
		Auth:    []ssh.AuthMethod{ssh.PublicKeys(sshKey)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := ssh.Dial("tcp", s.Ip+":22", config)
	if err != nil {
		return err
	}

	defer client.Close()

	for _, ln := range cmdLine {

		session, err := client.NewSession()
		if err != nil {
			session.Close()
			return err
		}

		writeCloudInfoOut, err := session.Output(fmt.Sprintf("sh -c '%v'", ln))
		if err != nil {
			session.Close()
			return errors.Wrap(err, string(writeCloudInfoOut))
		}

		session.Close()
	}

	return nil
}

func (s *SshConnection) IsMachineReachable() error {

	d, err := ioutil.ReadFile(s.KeyLocation)
	if err != nil {
		return err
	}

	sshKey, err := ssh.ParsePrivateKey(d)
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:    s.User,
		Timeout: s.Timeout,
		Auth:    []ssh.AuthMethod{ssh.PublicKeys(sshKey)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := ssh.Dial("tcp", s.Ip+":22", config)
	if err != nil {
		return err
	}

	defer client.Close()

	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		return err
	}

	return nil
}
