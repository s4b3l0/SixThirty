package domain

type Credential struct {
	Username string
	Email    string
	Password string
}

//func Signup(db *gorm.DB, req *http.Request) (*http.Request, error) {
//	//passwordHash, err :=  passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
//	return &req
//}
