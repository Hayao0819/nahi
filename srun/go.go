package srun

type Golang struct {
	Name  string
	Path  string
	Flags []string
}

func (g Golang) IsCorrect() (bool, error) {
	return true, nil
}

// TODO: Implement Runner
