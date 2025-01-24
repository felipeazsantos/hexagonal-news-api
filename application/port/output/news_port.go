package output

import (
	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/rest_err"
)

type NewsPort interface {
	GetNews(domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
