package dto

type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`				
	UserInfo UserInfo `json:"userInfo"`
}

type UserInfo struct {
	FirstName string `json:"firstName"`	
	LastName string `json:"lastName"`	
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}


