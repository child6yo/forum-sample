package forum

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required" validation:"email"`
	Username string `json:"username" binding:"required" validation:"username,min_len=2,max_len=15"`
	Password string `json:"password" binding:"required"`
}

type SignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}