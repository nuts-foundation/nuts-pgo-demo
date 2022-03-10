package api

import (
	"crypto/ecdsa"
	"github.com/nuts-foundation/nuts-pgo-demo/domain/accounts"
	"log"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/lestrrat-go/jwx/jwt/openid"
)

type UserAccount struct {
	Username string
	Password string
}

type auth struct {
	sessionKey *ecdsa.PrivateKey
	accounts   accounts.Repository
}

func NewAuth(key *ecdsa.PrivateKey, accountRepo accounts.Repository) auth {
	return auth{
		sessionKey: key,
		accounts:   accountRepo,
	}
}

func (auth auth) CheckCredentials(username, password string) bool {
	account, err := auth.accounts.FindByUserName(username)
	if err != nil {
		return false
	}
	return account.Password == password
}

func (auth auth) CreateJWT(email string) ([]byte, error) {
	t := openid.New()
	t.Set(jwt.IssuedAtKey, time.Now())
	// session is valid for 20 minutes
	t.Set(jwt.ExpirationKey, time.Now().Add(20*time.Minute))
	t.Set(openid.EmailKey, email)

	signed, err := jwt.Sign(t, jwa.ES256, auth.sessionKey)
	if err != nil {
		log.Printf("failed to sign token: %s", err)
		return nil, err
	}
	return signed, nil
}

func (auth auth) ValidateJWT(token []byte) (jwt.Token, error) {
	pubKey := auth.sessionKey.PublicKey
	t, err := jwt.Parse(token, jwt.WithVerify(jwa.ES256, pubKey), jwt.WithValidate(true))
	if err != nil {
		log.Printf("unable to parse token: %s", err)
		return nil, err
	}
	return t, nil
}
