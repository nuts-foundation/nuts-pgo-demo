package types

import (
	"errors"
	"github.com/nuts-foundation/go-did/did"
)

var NotFoundErr = errors.New("not found")

type Account struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	DID      did.DID `json:"did"`
}
