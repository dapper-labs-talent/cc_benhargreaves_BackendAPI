package utils

const (
	UnableToParseJWTClaimsErr = "could not parse JWT claims"
	ExpiredTokenErr           = "this token has expired. please log in again"
	EmptyJWTErr               = "a valid token is required for this request"
	InvalidJWTErr             = "unable to verify provided Token"
	InvalidPasswordErr        = "password does not match"
	MalformedRequestBodyErr   = "the provided json is invalid or malformed"
	InsertNewUserErr          = "an error occurred while creating new user"
	CreateNewJWTErr           = "an error occurred while generating a new token"
	GetUserByEmailErr         = "an error occurred while fetching user info"
	GetAllUsersErr            = "an error occurred while fetching user list"
	UpdateUserErr             = "an error occurred while updating user info"
	UserAlreadyExistsErr      = "a user with this email already exists"
	UnknownErr                = "Something went wrong"
)
