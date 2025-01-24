package news_http

import (
	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/rest_err"
)

type newsClient struct{}

func NewNewsClient() *newsClient {
	return &newsClient{}
}

func (nc *newsClient) GetNews(newsReqDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {

	return nil, nil
}
