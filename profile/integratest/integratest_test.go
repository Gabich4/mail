//go:build integra
// +build integra

package integratest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"profile/models"
	"profile/service"
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
	service, err := service.New()
	if err != nil {
		s.T().Error(err)
	}

	s.service = *service
	go s.service.Serve()
}

func (s *integraTestSuite) TeardownSuite() {
	s.service.Shutdown()
}

func (s *integraTestSuite) TestUploadTemplate() {
	template := `
	<html lang='\''ru'\''>
		<head>
			<meta charset='\''utf-8'\''>
			<title>{{.Title}}</title>
		</head>
		
		<body>
			<header>
				<h1>{{.Title}}</h1>
			</header>
			
			<main>
				<h2>Добрый день {{.FIO}}</h2>
				<p>Материалы для ознакомления: <a href={{.ContentURL}}>{{.ContentTitle}}</a></p>
			</main>
		</body>
	</html>
	`
	client := http.Client{}
	resp, err := client.Post(
		"http://localhost:3000/api/v1/upload_template",
		"text/html",
		bytes.NewBuffer([]byte(template)),
	)
	s.Nil(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	template = ""
	resp, err = client.Post(
		"http://localhost:3000/api/v1/upload_template",
		"text/html",
		bytes.NewBuffer([]byte(template)),
	)
	s.Nil(err)
	s.Equal(http.StatusBadRequest, resp.StatusCode)
}

func (s *integraTestSuite) TestReceiveStatus() {
	successPayload, err := json.Marshal(models.SendRequestStatus{
		Receivers: []string{"x", "y", "z"},
		Message:   "test_message",
		Status:    "test_status",
	})
	s.Nil(err)

	client := http.Client{}
	resp, err := client.Post(
		"http://localhost:3000/api/v1/status",
		"application/json",
		bytes.NewBuffer(successPayload),
	)

	s.Nil(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	fooStruct := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		"admin",
		"admin",
	}
	failedPayload, err := json.Marshal(fooStruct)
	s.Nil(err)

	resp, err = client.Post(
		"http://localhost:3000/api/v1/status",
		"application/json",
		bytes.NewBuffer(failedPayload),
	)

	s.Nil(err)
	s.Equal(http.StatusBadRequest, resp.StatusCode)
}

func (s *integraTestSuite) TestCreateReceiversOnUser() {
	client := http.Client{}
	payload := "z, x, c, v"

	resp, err := client.Post(
		"http://localhost:3000/api/v1/receivers/test_user",
		"text/plain",
		bytes.NewBuffer([]byte(payload)),
	)
	s.Nil(err)
	s.Equal(http.StatusOK, resp.StatusCode)
	receivers, err := s.service.App.Logic.Read("test_user")
	s.Nil(err)
	s.Equal("z, x, c, v", receivers)
}

func (s *integraTestSuite) TestReadReceiversOnUser() {
	client := http.Client{}
	resp, err := client.Get(
		"http://localhost:3000/api/v1/receivers/admin",
	)
	s.Nil(err)
	s.Equal(http.StatusNotFound, resp.StatusCode)
}

func (s *integraTestSuite) TestReaddAllReceivers() {
	client := http.Client{}
	resp, err := client.Get(
		"http://localhost:3000/api/v1/receivers",
	)
	s.Nil(err)
	s.Equal(http.StatusNotFound, resp.StatusCode)
}
