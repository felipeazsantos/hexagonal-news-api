package request

type NewsRequest struct {
	Subject string `form:"q" binding:"required,min=2"`
	From    string `form:"from" binding:"required,datetime" time_format:"2006-01-02"`
}
