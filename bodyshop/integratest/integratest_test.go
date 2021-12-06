//go:build integra
// +build integra

package integratest_test

import (
	"bodyshop/models"
	"bodyshop/service"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type integraTestSuite struct {
	suite.Suite
	service service.Service
}

func TestIntegraTestSuite(t *testing.T) {
	suite.Run(t, &integraTestSuite{})
}

func (s *integraTestSuite) SetupSuite() {
	server := service.New()

	s.service = *server
	go s.service.Server.ListenAndServe()
}

func (s *integraTestSuite) TestSendMessage() {
	var inRequest models.IncomingRequest = models.IncomingRequest{
		Receivers:  []string{"receiver1", "receiver2"},
		TemplateId: "testTemplateId1",
		Parameters: []models.TemplateParameter{
			{ParameterName: "{{ .testParamName1}}", ParameterValue: "paramValue1"},
			{ParameterName: "{{ .testParamName2}}", ParameterValue: "paramValue2"},
		},
	}
	payload, err := json.Marshal(inRequest)
	s.NoError(err)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/api/v1/send", s.service.Server.Addr), bytes.NewBuffer(payload))
	s.NoError(err)
	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)
}
