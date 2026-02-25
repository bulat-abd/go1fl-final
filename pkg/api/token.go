package api

func CreateToken() string {
	return "STUBTOKEN"
}

func ValidateToken(token string) bool {
	return token == "STUBTOKEN"
}
