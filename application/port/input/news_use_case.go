package input

import (
	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(subject string, from string) (*domain.NewsDomain, *rest_err.RestErr)
}
