package auth
 
import(
	"github.com/dgrijalva/jwt-go"
	"errors"
)

type Service interface {
	GenerateToken(userID int, email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error){
	//validate Token
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		//check token algorithm
		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil{
		return token, err
	}

	return token, nil
}