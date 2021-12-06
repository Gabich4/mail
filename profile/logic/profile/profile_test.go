//go:build !integra
// +build !integra

package profile_test

import (
	"profile/logic/profile"
	"testing"

	"github.com/stretchr/testify/suite"
)

type unitTestSuite struct {
	suite.Suite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &unitTestSuite{})
}

func (s *unitTestSuite) TestSuccessCreateNotExistedUser() {
	p := profile.New()
	err := p.Create("username1", "receivers")
	s.Nil(err, "error must be nil")
}

func (s *unitTestSuite) TestFailCreateExistedUser() {
	p := profile.New()
	username := "username2"
	p.Create(username, "receivers")
	err := p.Create(username, "receivers")
	s.Equal("user already exists", err.Error())
}

func (s *unitTestSuite) TestSuccessReadReceiversByExistedUser() {
	p := profile.New()
	username := "username3"
	receivers := "receivers"
	p.Create(username, receivers)
	result, err := p.Read(username)
	s.Equal(receivers, result)
	s.Nil(err)
}

func (s *unitTestSuite) TestFailReadReceiversByNotExistedUser() {
	p := profile.New()
	username := "username4"
	result, err := p.Read(username)
	s.Equal("", result)
	s.Equal("user not found", err.Error())
}

func (s *unitTestSuite) TestSuccessUpdateReceiversByExistedUser() {
	p := profile.New()
	username := "username5"
	receivers := "receivers"
	newReceivers := "new_receivers"
	p.Create(username, receivers)

	err := p.Update(username, newReceivers)
	s.Nil(err)

	result, err := p.Read(username)
	s.Equal(newReceivers, result)
	s.Nil(err)
}

func (s *unitTestSuite) TestFailUpdateReceiversByNotExistedUser() {
	p := profile.New()
	username := "username6"
	newReceivers := "new_receivers"

	err := p.Update(username, newReceivers)
	s.Equal("user not found", err.Error())
}

func (s *unitTestSuite) TestSuccessDeleteReceiversByExistedUser() {
	p := profile.New()
	username := "username7"
	receivers := "receivers"
	p.Create(username, receivers)

	err := p.Delete(username)
	s.Nil(err)

	result, err := p.Read(username)
	s.Equal("user not found", err.Error())
	s.Equal("", result)
}

func (s *unitTestSuite) TestFailDeleteReceiversByNotExistedUser() {
	p := profile.New()
	username := "username8"

	err := p.Delete(username)
	s.Equal("user not found", err.Error())
}

func (s *unitTestSuite) TestSuccessReadAllReceivers() {
	p := profile.New()
	username := "username9"
	receivers := "receivers"
	p.Create(username, receivers)
	userReceivers, err := p.ReadAll()
	s.Nil(err)
	s.NotNil(userReceivers)
	s.NotEmpty(userReceivers)
}
