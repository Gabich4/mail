//go:build integra
// +build integra

package integratest_test

import (
	"bytes"
	"encoding/json"
	"mailsender/common"
	"mailsender/models"
	"mailsender/service"
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
	server, err := service.New()
	if err != nil {
		s.T().Error(err)
	}

	s.service = *server
	go s.service.Server.ListenAndServe()
}

func (s *integraTestSuite) TestSendMessage() {
	// creating payload of incoming request to mailsender
	var inRequest models.IncomingSendRequest
	inRequest.Receivers = []string{"x", "y", "z"}
	inRequest.Message = "<div><h1>{{ .paramValue1}}</h1><p>{{ .paramValue2}}</p></div>"

	payload, err := json.Marshal(inRequest)
	s.Nil(err)

	// creating and sending post request to mailsender
	client := http.DefaultClient
	resp, err := client.Post(
		"http://"+common.ServiceConfig.Host+"/api/v1/send",
		"application/json",
		bytes.NewBuffer(payload),
	)
	s.Nil(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	// unmarshaling response into response structure
	var response models.Response
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	s.Nil(err)

	s.True(response.Success)
	s.Equal("successfully inserted message", response.Data)
	s.Nil(response.Error)
}
