package request

type HTTPMethod int8

const (
	POST = iota
	PUT
	PATCH
	DELETE
	GET
	HEAD
)

func (v HTTPMethod) String() string {
	switch v {
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	case GET:
		return "GET"
	case HEAD:
		return "HEAD"
	default:
		return "POST"
	}
}
