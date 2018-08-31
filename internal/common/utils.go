package common

import (
	"bytes"
	"html/template"
)

// ParseTmpl ...
func ParseTmpl(tmpl string, data interface{}) ([]byte, error) {
	parsedTmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := parsedTmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
