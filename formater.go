package curlx

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/miolini/jsonf/jsonflib"
)

var (
	// ErrNotJSON is returned when http body is not json string
	ErrNotJSON = errors.New("not json")

	// ErrHighlight is returned when json string highlight error
	ErrHighlight = errors.New("json highlight error")
)

// Formater is a util to format http body
type Formater struct {
	colorize bool
}

// NewFormater is used to create formater
func NewFormater(colorize bool) *Formater {
	f := &Formater{
		colorize: colorize,
	}
	return f
}

// ExportJSON is used to format and highlight json string
func (f *Formater) ExportJSON(data string) (string, error) {
	strings.Replace(data, "\"", "", -1)

	var v interface{}

	// Json Decoding
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		return "", ErrNotJSON
	}

	// Json Encode
	jsonData, err := json.MarshalIndent(v, "", "  ")

	// Syntax Highlight
	if f.colorize {
		jsonData, err = jsonflib.Highlight(
			jsonData,
			jsonflib.HighlightFlags{
				Colorize: true,
				Verbose:  false,
				Debug:    false,
			},
		)
		if err != nil {
			return "", ErrHighlight
		}
	}

	return string(jsonData), nil
}
