package bodyshop

import (
	"bodyshop/logic"
	"bodyshop/models"
	"errors"
	"strings"
)

// Bodyshop is a structure for generating messages.
// It implements logic.Logic interface.
type Bodyshop struct {
}

// New returns new Bodyshop instance.
func New() logic.Logic {
	return new(Bodyshop)
}

// FormMessage replaces parameters with values.
func (b *Bodyshop) FormMessage(id string, params []models.TemplateParameter) (string, error) {
	template, ok := models.Templates[id]
	if !ok {
		return "", errors.New("cannot find template with such id")
	}

	for _, param := range params {
		template = strings.ReplaceAll(template, param.ParameterName, param.ParameterValue)
	}

	return template, nil
}
