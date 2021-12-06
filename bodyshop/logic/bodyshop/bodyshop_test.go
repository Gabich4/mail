//go:build !integra
// +build !integra

package bodyshop_test

import (
	"bodyshop/logic/bodyshop"
	"bodyshop/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type unitTestSuite struct {
	suite.Suite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &unitTestSuite{})
}

func (s *unitTestSuite) TestInsertMessageForExistingTemplate() {
	l := bodyshop.New()

	message, err := l.FormMessage(
		"testTemplateId1",
		[]models.TemplateParameter{
			{"{{ .testParamName1}}", "paramValue1"},
			{"{{ .testParamName2}}", "paramValue2"},
		})
	s.Equal("<div><h1>paramValue1</h1><p>paramValue2</p></div>", message)
	s.Nil(err, "error must be nil")
}

func (s *unitTestSuite) TestInsertMessageForNotExistingTemplate() {
	l := bodyshop.New()

	message, err := l.FormMessage("templateId", []models.TemplateParameter{{"paramName", "paramValue"}})
	s.Equal("", message, "message must be empty")
	s.Equal(errors.New("cannot find template with such id"), err)
}
