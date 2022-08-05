package services

import (
	"encoding/gob"
	"fmt"
	"github.com/Nkez/check/internal/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type RequestService struct{}

func NewRequestService() *RequestService {
	return &RequestService{}
}

var Statistics = make(map[uuid.UUID]*models.ReqResMap)

func (s *RequestService) PostRequest(request *models.Request) (*models.Response, error) {
	reqRes := new(models.ReqResMap)
	reqRes.ID = uuid.New()
	fmt.Println(reqRes.ID)
	reqRes.RequestMap = *request
	reqRes.ResponseMap = models.Response{
		ID:      reqRes.ID,
		Status:  "OK",
		Headers: request.Headers,
	}
	Statistics[reqRes.ID] = reqRes
	defer func() {
		if err := s.Save(); err != nil {
			logrus.Info(err.Error())
		}
	}()
	return &reqRes.ResponseMap, nil
}

func (s *RequestService) GetRequest(id string) (*models.ReqResMap, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	if val, ok := Statistics[u]; ok {
		return val, nil
	} else {
		return nil, err
	}
}

func (s *RequestService) Save() error {
	encodeFile, err := os.OpenFile(viper.GetString("path"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(encodeFile)
	if err := encoder.Encode(Statistics); err != nil {
		return err
	}
	encodeFile.Close()
	return nil
}

func (s *RequestService) GetStatus() *models.Status {
	return &models.Status{Answer: "Старичок, идём на пикничок!\n" +
		"Всё готово!\n" +
		"Все в сборе!\n" +
		"Тебя не хвотает!\n" +
		"Тебя не хвотает!\n\n" +
		"Старичок, идём на пикничок!"}
}
