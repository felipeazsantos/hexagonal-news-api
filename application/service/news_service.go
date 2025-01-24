package service

import (
	"fmt"

	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/application/port/output"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/logger"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/rest_err"
)

type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(newsReqDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("Init GetNewsService function, subject=%s, from=%s", newsReqDomain.Subject, newsReqDomain.From))

	newsDomainResponse, err := ns.newsPort.GetNews(newsReqDomain)
	return newsDomainResponse, err
}
