package config

// BadRequest 400 StatusBadRequest
var BadRequest = map[string]string{
	"error": "BadRequest",
}

// Unauthorized 401 StatusUnauthorized
var Unauthorized = map[string]string{
	"error": "Unauthorized",
}

// NotFound 404 StatusNotFound
var NotFound = map[string]string{
	"error": "NotFound",
}

// NotAcceptable 406 StatusNotAcceptable
var NotAcceptable = map[string]string{
	"error": "NotAcceptable",
}

// ValdError 409 StatusConflict
var ValdError = map[string]string{
	"error": "ValdError",
}

// ServerError 500 StatusInternalServerError
var ServerError = map[string]string{
	"error": "InternalServerError",
}
