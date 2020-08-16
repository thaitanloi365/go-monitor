package errs

import "net/http"

var (
	// ErrRoleInvalid invalid role
	ErrRoleInvalid = New(10000, "Role is invalid", http.StatusUnauthorized)
	// ErrTokenInvalid invalid token
	ErrTokenInvalid = New(10001, "Token is invalid", http.StatusUnauthorized)
	// ErrTokenExpired expired token
	ErrTokenExpired = New(10002, "Token is expired", http.StatusUnauthorized)
	// ErrTokenMissing missing token
	ErrTokenMissing = New(10003, "Token is missing or malformed", http.StatusUnauthorized)
	// ErrSessionTimeout session timeout
	ErrSessionTimeout = New(10004, "Session timeout", http.StatusUnauthorized)
	// ErrPasswordExpiration session timeout
	ErrPasswordExpiration = New(10005, "Your current password was expired, please update new password.", http.StatusUnauthorized)
	// ErrPasswordOrUsernameIncorrect session timeout
	ErrPasswordOrUsernameIncorrect = New(10006, "Your password or username is incorrect.", http.StatusBadRequest)
)

var (
	// ErrUserNotFound user not found
	ErrUserNotFound = New(20000, "User not found", http.StatusNotFound)
	// ErrPasswordIncorrect password is incorrect
	ErrPasswordIncorrect = New(20001, "Password is incorrect", http.StatusBadRequest)
	// ErrPasswordRequired password is required
	ErrPasswordRequired = New(20002, "Password is required", http.StatusBadRequest)
)

var (
	// ErrContainerNotFound container not found
	ErrContainerNotFound = New(30000, "The container is not found", http.StatusNotFound)
	// ErrStreamingUnsupported stream unsupported
	ErrStreamingUnsupported = New(30000, "The streaming unsupported", http.StatusNotFound)
)
