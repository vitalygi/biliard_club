package jwt

import "github.com/golang-jwt/jwt"

type Data struct {
	Phone string
}
type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}
func (j *JWT) Create(data Data) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": data.Phone,
	})
	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(token string) (*Data, bool) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, false
	}
	email := t.Claims.(jwt.MapClaims)["phone"]
	return &Data{Phone: email.(string)}, t.Valid
}
