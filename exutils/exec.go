package exutils

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"

	"github.com/samber/lo"
)

func EvalSh(env map[string]string, code string) (string, string, int, error) {
	cmdArg := []string{"-c"}
	cmdArg = append(cmdArg, code)
	// fmt.Println(cmdArg)
	cmd := exec.Command("sh", cmdArg...)
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env,
		lo.MapToSlice(env, func(k string, v string) string {
			return k + "=" + env[k]
		})...,
	)

	// new string writer
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return stdout.String(), stderr.String(), cmd.ProcessState.ExitCode(), err
	// return stdout.String(), stderr.String(), err
}

func EvalString(env map[string]string, str string) (string, error) {
	stdout, stdin, exit, err := EvalSh(env, "printf \"%s\" "+"\""+str+"\"")
	slog.Debug("run", "stdout", stdout, "stdin", stdin, "exit", exit)
	return stdout, err
}
