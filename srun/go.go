package srun

import (
	"os/exec"
	"path"
)

type Golang struct {
	Name  string
	Path  string
	Flags []string
}

func (g Golang) IsCorrect() (bool, error) {
	return path.Ext(g.Path) == ".go", nil
}

func (g Golang) Build(tmp string) error {
	args := []string{"build",
		"-o",
		path.Join(tmp, g.Name),
	}
	args = append(args, g.Flags...)
	args = append(args, "--", g.Path)
	cmd := exec.Command("go", args...)
	cmd.Dir = tmp

	return cmd.Run()
}

func (g Golang) Cmd(tmp string) *exec.Cmd {
	cmd := exec.Command(path.Join(tmp, g.Name))
	return cmd
}
