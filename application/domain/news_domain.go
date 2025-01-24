package domain

type NewsReqDomain struct {
	Subject string
	From    string
}

type NewsDomain struct {
	Status       string
	TotalResults int
	Articles     []Article
}

type Article struct {
	Source      ArticleSource
	Author      string
	Title       string
	Description string
	URL         string
	UrlToImage  string
	PublishedAt string
	Content     string
}

type ArticleSource struct {
	ID   *string
	Name string
}
