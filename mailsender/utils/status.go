package utils

import (
	"mailsender/models"
	"math/rand"
	"time"
)

func GetRandomStatus() string {
	rand.Seed(time.Now().Unix())
	if rand.Intn(2) == 1 {
		return models.SuccessStatus
	}
	return models.FailStatus
}
