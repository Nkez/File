package services

import "github.com/Nkez/check/internal/models"

type Request interface {
	PostRequest(request *models.Request) (*models.Response, error)
	GetRequest(id string) (*models.ReqResMap, error)
	GetStatus() *models.Status
}

type Service struct {
	Request
}

func NewService() *Service {
	return &Service{
		Request: NewRequestService(),
	}
}
