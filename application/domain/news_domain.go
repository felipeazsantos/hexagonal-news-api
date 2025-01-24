package domain

type NewsReqDomain struct {
	Subject string
	From string
}

type NewsDomain struct {
	Status       string
	TotalResults string
	Articiles    []Article
}

type Article struct {
	Source      string
	ID          string
	Name        string
	Title       string
	Description string
	URL         string
	UrlToImage  string
	PublishedAt string
	Content     string
}
