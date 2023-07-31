package handler

type PasswordRequest struct {
	Length     int
	UseUpper   bool
	UseLower   bool
	UseNumbers bool
	UseSpecial bool
	Filter     string
}

type PageVariables struct {
	GeneratedPassword string
}
