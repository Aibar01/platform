package response

type Error struct {
	Detail string `json:"detail"`
}

func NewError(detail string) *Error {
	return &Error{Detail: detail}
}

var (
	PermissionDeniedError = NewError("permission denied")
)
