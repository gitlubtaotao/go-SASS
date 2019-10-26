package test

import (
	"quickstart/models"
	"testing"
)

func TestModelCount(T *testing.T)  {
	models.ModelCount("user")
}
