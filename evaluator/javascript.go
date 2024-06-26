package runner

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"os/exec"
)

type JavascriptEvaluator struct {
}

func NewJavaScriptEvaluator() *JavascriptEvaluator {
	return new(JavascriptEvaluator)
}

func (p *JavascriptEvaluator) Evaluate(code string, testLocation string) (err error) {
	return errors.New("not implemented")
}

func (p *JavascriptEvaluator) Execute(code string) (stdOut string, err error) {
	err = os.WriteFile("index.js", []byte(code), os.FileMode(os.O_RDWR))

	if err != nil {
		return "", err
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("sh", "-c", "/root/.nvm/versions/node/v20.15.0/bin/node index.js")

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		log.Error(err)
	}

	fmt.Println(stdout.String())
	fmt.Println(stderr.String())

	return stdout.String() + stderr.String(), nil
}
