package webshell

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"

	"github.com/noovertime7/kubemanage/pkg/logger"
	"github.com/noovertime7/kubemanage/pkg/utils"
)

type Config struct {
	UserName      string
	IpAddress     string //IP地址
	Port          string
	Password      string // 密码连接
	PrivateKey    string // 私钥连接
	KeyPassphrase string // 私钥密码
	Width         int    // pty width
	Height        int    // pty height
}

type Terminal struct {
	Client  *ssh.Client
	TERM    string
	session *ssh.Session
	config  Config
	stdout  io.Reader
	stdin   io.Writer
	stderr  io.Reader
	closed  bool
}

func (t *Terminal) SetWinSize(h int, w int) {
	if err := t.session.WindowChange(h, w); err != nil {
		logger.LG.Error(fmt.Sprintf("ssh pty change windows size failed:\t %v", err.Error()))
		return
	}
}

// IsClosed 终端是否已关闭
func (t *Terminal) IsClosed() bool {
	return t.closed
}

func (t *Terminal) Close() (err error) {
	if t.IsClosed() {
		return nil
	}
	if err = t.session.Close(); err != nil {
		return
	}

	if err = t.Client.Close(); err != nil {
		return
	}

	return
}
func getTerm() (term string) {
	if term = os.Getenv("xterm"); term == "" {
		term = "xterm-256color"
	}
	return
}
func (t *Terminal) Connect(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	var err error
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}

	if err = t.session.RequestPty(t.TERM, t.config.Height, t.config.Width, modes); err != nil {
		return err
	}
	t.session.Stdin = stdin
	t.session.Stderr = stderr
	t.session.Stdout = stdout

	if err = t.session.Shell(); err != nil {
		return err
	}

	quit := make(chan int)
	go func() {
		i := 0
		_ = t.session.Wait()
		_ = t.Close()
		quit <- i + 1
	}()
	return nil
}

func NewTerminal(config Config) (*Terminal, error) {
	var authMethods []ssh.AuthMethod

	sshConfig := &ssh.ClientConfig{
		User:            config.UserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback:  ssh.BannerDisplayStderr(),
		Timeout:         time.Second * 15,
	}

	if config.PrivateKey != "" {
		if pk, err := getPrivateKey(config.PrivateKey, utils.AesDecryptCBC2Hex(config.KeyPassphrase)); err != nil {
			return nil, err
		} else {
			authMethods = append(authMethods, pk)
		}
	} else {
		authMethods = append(authMethods, ssh.Password(config.Password))
	}

	sshConfig.Auth = authMethods

	addr := net.JoinHostPort(config.IpAddress, config.Port)

	client, err := ssh.Dial("tcp", addr, sshConfig)

	if err != nil {
		logger.LG.Error("Failed to connect to remote terminal", zap.Error(err))
		return nil, err
	}

	session, err := client.NewSession()

	if err != nil {
		return nil, err
	}

	s := Terminal{
		TERM:    getTerm(),
		Client:  client,
		config:  config,
		session: session,
	}

	return &s, nil
}

func getPrivateKey(privateKeyPath string, privateKeyPassphrase string) (ssh.AuthMethod, error) {
	if !utils.FileExist(privateKeyPath) {
		privateKeyPath = filepath.Join(os.Getenv("HOME"), ".ssh/id_rsa")
	}
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}
	var signer ssh.Signer
	if privateKeyPassphrase != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(key, []byte(privateKeyPassphrase))
	} else {
		signer, err = ssh.ParsePrivateKey(key)
	}
	if err != nil {
		return nil, fmt.Errorf("parse private key failed: %v", err)
	}
	return ssh.PublicKeys(signer), nil
}
