package runner

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"os/exec"
)

type PythonEvaluator struct {
}

func NewPythonEvaluator() *PythonEvaluator {
	return new(PythonEvaluator)
}

func (p *PythonEvaluator) Evaluate(code string, testLocation string) (err error) {
	return errors.New("not implemented")
}

func (p *PythonEvaluator) Execute(code string) (stdOut string, err error) {
	err = os.WriteFile("script.py", []byte(code), os.FileMode(os.O_RDWR))

	if err != nil {
		return "", err
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", "python3 script.py")

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
