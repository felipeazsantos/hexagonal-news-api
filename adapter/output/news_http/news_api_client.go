package news_http

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/output/model/response"
	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/env"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/rest_err"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type newsClient struct{}

var client *resty.Client

func NewNewsClient() *newsClient {
	client = resty.New().SetBaseURL("https://newsapi.org/v2")
	return &newsClient{}
}

func (nc *newsClient) GetNews(newsReqDomain domain.NewsReqDomain) (*domain.NewsDomain, *rest_err.RestErr) {
	newsResponse := &response.NewsClientResponse{}

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q":      newsReqDomain.Subject,
			"from":   newsReqDomain.From,
			"apiKey": env.GetNewsTokenAPI(),
		}).
		SetResult(newsResponse).
		Get("/everything")

	if err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to call NewsAPI with params")
	}

	newsDomain := &domain.NewsDomain{}
	copier.Copy(newsDomain, newsResponse)

	return newsDomain, nil
}
