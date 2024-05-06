package nahi

import (
	"strings"

	"github.com/Hayao0819/nahi/osutils"
)

func LoadEnvFile(path string) (map[string]string, error) {
	lines, err := osutils.ReadFileLine(path)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)

	for _, line := range lines {
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		value := strings.TrimPrefix(parts[1], "\"")
		value = strings.TrimSuffix(value, "\"")
		env[parts[0]] = value
	}

	return env, nil

}
