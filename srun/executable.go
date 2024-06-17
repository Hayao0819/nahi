package srun

import (
	"os/exec"

	"github.com/Hayao0819/nahi/futils"
)

type Executable string

func (e Executable) IsCorrect() (bool, error) {
	return futils.Executable(string(e))
}

func (e Executable) Cmd(tmp string) *exec.Cmd {
	return exec.Command(string(e))
}

func (e Executable) Build(tmp string) error {
	return nil
}
