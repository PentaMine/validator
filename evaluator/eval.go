package runner

type Evaluator interface {
	Evaluate(code string, testLocation string) (err error)
	Execute(code string) (stdOut string, err error)
}
