package jinja2

import (
	"os"
	"strings"

	"github.com/noirbizarre/gonja"
)

// ExpandTemplate takes a blob of text and parses it with J2, filling values as applicable with env
func ExpandTemplate(template string) (string, error) {
	tpl, err := gonja.FromString(template)
	if err != nil {
		return "", err
	}

	// Now you can render the template with the given
	// gonja.Context how often you want to.
	out, err := tpl.Execute(getEnvAsMap())
	if err != nil {
		return "", err
	}
	return out, nil
}

func getEnvAsMap() map[string]interface{} {
	envMap := make(map[string]interface{})
	envVars := os.Environ()

	for _, envVar := range envVars {
		parts := strings.SplitN(envVar, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}

	return envMap
}
