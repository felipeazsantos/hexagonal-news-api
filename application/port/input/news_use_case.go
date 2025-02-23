package input

import (
	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(newsReqDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
