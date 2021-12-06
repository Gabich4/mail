//go:build !integra
// +build !integra

package mailsender_test

import (
	"mailsender/logic/mailsender"
	"mailsender/models"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockRepository struct {
	mock.Mock
}

func (r *mockRepository) CreateMessage(msg models.Message) (string, error) {
	args := r.Called(msg)

	return args.Get(0).(string), args.Error(1)
}

func (r *mockRepository) UpdateMessage(id, status string) error {
	args := r.Called(id, status)
	return args.Error(0)
}

type unitTestSuite struct {
	suite.Suite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &unitTestSuite{})
}

func (s *unitTestSuite) TestInsertMessage() {
	// preparing data for test
	var (
		msg models.Message
		req models.IncomingSendRequest
	)
	msg.Receivers = []string{"x", "y", "z"}
	msg.Text = "test message"
	msg.Status = models.WaitingStatus
	req.Message = "test message"
	req.Receivers = []string{"x", "y", "z"}

	r := new(mockRepository)

	r.On("CreateMessage", msg).Return(
		msg.ID.Hex(),
		nil,
	)

	l := mailsender.New(r)

	id, err := l.InsertMessage(req)
	s.Nil(err, "error must be nil")
	s.Equal(msg.ID.Hex(), id)

	r.AssertExpectations(s.T())
}

func (s *unitTestSuite) TestUpdateMessage() {
	objID := primitive.NewObjectID()
	r := new(mockRepository)

	r.On("UpdateMessage", objID.Hex(), models.SuccessStatus).Return(
		nil,
	)

	l := mailsender.New(r)
	err := l.UpdateMessage(objID.Hex(), models.SuccessStatus)
	s.Nil(err, "error must be nil")

	r.AssertExpectations(s.T())
}
