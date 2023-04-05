package requests

type CreateNodeRequest struct {
	Key   string `json:"key" form:"key" binding:"required"`
	Value string `json:"value" form:"value" binding:"required"`
	Prev  string `json:"prev" form:"prev" binding:"required"`
}
