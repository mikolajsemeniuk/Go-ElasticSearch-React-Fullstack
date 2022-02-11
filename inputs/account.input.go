package inputs

type Account struct {
	Username string `json:"username" binding:"min=4,max=50"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"min=6,max=50"`
}
