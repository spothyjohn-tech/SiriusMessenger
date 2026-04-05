package user

import (
	"MesEdge/internal/models"
	"MesEdge/internal/repository"
)

// Struct service
type UserService struct {
	repo repository.UserRepository
}

// For service in main.go
func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

// JSON from Frontend
type RegistrationRequest struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Salt       []byte `json:"salt"`
	Verifier   []byte `json:"verifier"`
	PublicKey  []byte `json:"public_key"`
	EncPrivKey []byte `json:"enc_priv_key"`
}

// Method registation
func (s *UserService) Register(req RegistrationRequest) error {
	newUser := &models.User{
		Email:      req.Email,
		Username:   req.Username,
		Salt:       req.Salt,
		Verifier:   req.Verifier,
		PublicKey:  req.PublicKey,
		EncPrivKey: req.EncPrivKey,
	}

	return s.repo.Create(newUser)
}

// // Registration new user
// func RegistrationUser(db *gorm.DB, email string, password string){
// 	reader := bufio.NewReader(os.Stdin);
// 	fmt.Println("Введите почту:")
// 	email, erre := reader.ReadString('\n')
// 	if erre != nil {
// 		fmt.Println(erre)
// 		return
// 	}
// 	email = strings.TrimSpace(email)
// 	fmt.Println("Введите пароль:")
// 	password, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	password = strings.TrimSpace(password)
// 	argon2.hashPassword(password)
// }

// // Login user
// func Login(){

// }
