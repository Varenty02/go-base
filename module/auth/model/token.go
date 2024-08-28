package authmodel

//dto ra register,login,generate
type TokenResponse struct{
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//dto vào generate token 
type TokenRequest struct{
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}