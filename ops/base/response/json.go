package response

type JSONResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJsonResponse(code int, msg string, data interface{}) *JSONResponse {
	return &JSONResponse{code, msg, data}
}

var (
	Unauthorization = NewJsonResponse(401, "unauthorization", nil)
	Ok              = NewJsonResponse(200, "ok", nil)
	BadResquest     = NewJsonResponse(400, "bad request", nil)
	InvalidToken    = NewJsonResponse(403, "token is invalid", nil)
	NotAcceptable   = NewJsonResponse(406, "Not Acceptable", nil)
)
