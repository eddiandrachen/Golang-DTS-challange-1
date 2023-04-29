package response

const (
	InvalidParam       = "invalid param request"
	InvalidBody        = "invalid body request"
	InvalidQuery       = "invalid query request"
	SomethingWentWrong = "Something went wrong"
)

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message  string `json: "message"`
	ErrorMsg string `json: "errorMsg"`
}
