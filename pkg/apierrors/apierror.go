package apierrors

//APIError Define a error
type APIError struct {
	Msg    string `json:"msg,omitempty"`
	Status int    `json:"status,omitempty"`
	Code   string `json:"code,omitempty"`
}

func (e APIError) Error() string {
	return e.Msg
}
