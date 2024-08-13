package auth_service

import "github.com/golang-jwt/jwt/v5"

type SignMethod string

const (
	MethodRSA  SignMethod = "RSA"
	MethodHMAC SignMethod = "HMAC"
)

type SigningMethod interface {
	GetSigningKey() interface{}
	KeyFunc(token *jwt.Token) (interface{}, error)
	GetMethod() SignMethod
}

func NewRsa512Signing(loadPrivateKey func() string,
	loadPublicKey func() string) RSA512Signing {
	return RSA512Signing{
		privateKey: loadPrivateKey(),
		publicKey:  loadPublicKey(),
	}
}

type RSA512Signing struct {
	privateKey interface{}
	publicKey  interface{}
}

func (rsa RSA512Signing) GetSigningKey() interface{} {
	return rsa.privateKey
}

func (rsa RSA512Signing) KeyFunc(token *jwt.Token) (interface{}, error) {

	return rsa.publicKey, nil
}

func (rsa RSA512Signing) GetMethod() SignMethod {
	return MethodRSA
}
