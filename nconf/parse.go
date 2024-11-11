package nconf

import (
	"io"
	"strings"
)

func parseKeyValue(r io.Reader) (map[string]string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	env := make(map[string]string)
	lines := strings.Split(string(b), "\n")

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
