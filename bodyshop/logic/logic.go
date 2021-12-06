package logic

import (
	"bodyshop/models"
)

type Logic interface {
	FormMessage(templateID string, params []models.TemplateParameter) (string, error)
}
