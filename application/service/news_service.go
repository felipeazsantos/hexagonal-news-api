package service

import (
	"fmt"

	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/logger"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/rest_err"
)

type newsService struct {
}

func NewNewsService() *newsService {
	return &newsService{}
}

func (ns *newsService) GetNewsService(subject string, from string) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("Init GetNewsService function, subject=%s, from=%s", subject, from))

	return nil, nil
}
