package response

type NewsClientResponse struct {
	Status       string
	TotalResults string
	Articiles    []ArticleResponse
}

type ArticleResponse struct {
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
