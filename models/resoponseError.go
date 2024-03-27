package models

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (r ResponseError) Error() string {
	return r.Message
}
