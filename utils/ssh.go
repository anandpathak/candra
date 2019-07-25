package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"github.com/shiena/ansicolor"
	"golang.org/x/crypto/ssh"
)

type password string

func (p password) Password(user string) (password string, err error) {
	return string(p), nil
}

// PublicKeyFile reads pem file for ssh
func PublicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
}

// Login creates psedo terminal for ssh, and can handle event. this will allow us to add more functionality over ssh.
// but currently as it Login does not have terminal feature, it is not used for ssh session
func Login(pemKey string, user string, server string, port string) {
	publicKey, err := PublicKeyFile(pemKey)
	if err != nil {
		log.Println(err)
		return
	}

	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			//ssh.Password(p),
			publicKey,
		},
	}
	conn, err := ssh.Dial("tcp", server+":"+port, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	defer conn.Close()
	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := conn.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Set IO
	session.Stdout = ansicolor.NewAnsiColorWriter(os.Stdout)
	session.Stderr = ansicolor.NewAnsiColorWriter(os.Stderr)
	in, _ := session.StdinPipe()

	// Set up terminal modes
	// https://net-ssh.github.io/net-ssh/classes/Net/SSH/Connection/Term.html
	// https://www.ietf.org/rfc/rfc4254.txt
	// https://godoc.org/golang.org/x/crypto/ssh
	// THIS IS THE TITLE
	// https://pythonhosted.org/ANSIColors-balises/ANSIColors.html
	modes := ssh.TerminalModes{
		ssh.ECHO:  0, // Disable echoing
		ssh.IGNCR: 1, // Ignore CR on input.
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		//if err := session.RequestPty("xterm-256color", 80, 40, modes); err != nil {
		//if err := session.RequestPty("vt100", 80, 40, modes); err != nil {
		//if err := session.RequestPty("vt220", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	// Start remote shell
	if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}

	// Handle control + C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for {
			<-c
			fmt.Println("^C")
			fmt.Fprint(in, "\n")
			//fmt.Fprint(in, '\t')
		}
	}()

	var _ = make([]byte, 1)

	// Accepting commands
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		fmt.Fprint(in, str)
	}
}

// Commando executes shell session
func Commando(pemFileName string, serverIP string, user string) {
	cmd := exec.Command("ssh", "-i", pemFileName, user+"@"+serverIP)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
