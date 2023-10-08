package core

type AuthSignupPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPaylaod struct {
	Action string `json:"action"`
	Auth   AuthSignupPayload
}

type AuthResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type ResponsePayload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}
