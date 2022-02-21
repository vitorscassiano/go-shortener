package wire

type CreateShortener struct {
	From string `json:"from" binding:"required"`
}
