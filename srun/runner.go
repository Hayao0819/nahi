package srun

import (
	"os"
	"os/exec"

	"errors"
)

type Runner interface {
	IsCorrect() (bool, error)
	Build(tmp string) error
	Cmd(tmp string) *exec.Cmd
}

var ErrNotExecutable = errors.New("not executable")

func Execute(r Runner) (*exec.Cmd, error) {
	if correct, err := r.IsCorrect(); err != nil {
		return nil, err
	} else if !correct {
		return nil, ErrNotExecutable
	}

	tmpDir, err := os.MkdirTemp("", "srun")
	if err != nil {
		return nil, err
	}

	if err := r.Build(tmpDir); err != nil {
		return nil, err
	}

	return r.Cmd(tmpDir), nil

}
