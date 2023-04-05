package requests

type CreateHeadRequest struct {
	Key string `json:"key" form:"key" binding:"required"`
}
