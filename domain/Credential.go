package domain

type Credential struct {
	Username     string `json:"username" gorm:"primaryKey"`
	Email        string `json:"email"`
	UserPassword string `json:"user_password"`
}

//func Signup(db *gorm.DB, req *http.Request) (*http.Request, error) {
//	//passwordHash, err :=  passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	return &req
//}
