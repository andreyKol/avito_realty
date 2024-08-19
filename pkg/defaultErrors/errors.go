package defaultErrors

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func NewError(statusCode int, description string) *fiber.Error {
	return &fiber.Error{
		Code:    statusCode,
		Message: description,
	}
}

type Response struct {
	Status      *string `json:"status"`
	Description *string `json:"description"`
}

// HandleError Method to handle fiber error type to send readable error response
func HandleError(errorToHandle *fiber.Error, ctx fiber.Ctx) error {
	status := strconv.Itoa(errorToHandle.Code)
	return ctx.Status(errorToHandle.Code).JSON(Response{
		Status:      &status,
		Description: &errorToHandle.Message,
	})
}

// 400 STATUS CODES

func ForbiddenError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Forbidden"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusForbidden, message)
}

func UnauthorizedError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Unauthorized"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusForbidden, message)
}

func BadRequestError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Bad Request"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusBadRequest, message)
}

func NotFoundError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Not Found"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusNotFound, message)
}

func ConflictError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Conflict"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusConflict, message)
}

func UnprocessableEntityError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Unprocessable Entity"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusUnprocessableEntity, message)
}

func TooManyRequestsError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Too Many Requests"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusTooManyRequests, message)
}

func NetworkConnectTimeoutError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Failed Dependency"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusFailedDependency, message)
}

func MethodNotAllowedError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Method Not Allowed"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusMethodNotAllowed, message)
}

func NotAcceptableError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Not Acceptable"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusNotAcceptable, message)
}

func GoneError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Gone"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusGone, message)
}

func ImTeapotError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "I'm a teapot"
	} else {
		message = description[0]
	}
	return NewError(418, message)
}

// 500 STATUS CODES

func InternalServerError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Internal Server Error"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusInternalServerError, message)
}

func ServiceUnavailableError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Service Unavailable"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusServiceUnavailable, message)
}

func GatewayTimeoutError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Gateway Timeout"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusGatewayTimeout, message)
}

func HTTPVersionNotSupportedError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "HTTP Version Not Supported"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusHTTPVersionNotSupported, message)
}

func NotImplemented(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Not Implemented"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusNotImplemented, message)
}

func BadGatewayError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Bad Gateway"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusBadGateway, message)
}

func NetworkAuthenticationRequiredError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Network Authentication Required"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusNetworkAuthenticationRequired, message)
}

func InsufficientStorageError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Insufficient Storage"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusInsufficientStorage, message)
}

func LoopDetectedError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Loop Detected"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusLoopDetected, message)
}

func NotExtendedError(description ...string) *fiber.Error {
	message := ""
	if len(description) == 0 {
		message = "Not Extended"
	} else {
		message = description[0]
	}
	return NewError(fiber.StatusNotExtended, message)
}
