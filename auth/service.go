package auth
 
import(
	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int, email string) (string, error)
}

type jwtService struct{
}

var SECRET_KEY = []byte("BWASTARTUP_secret_k3y")

func NewService() *jwtService{
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int, email string) (string, error){
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["email"] = email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken , err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}