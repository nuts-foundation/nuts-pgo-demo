package accounts

import (
	nutsApi "github.com/nuts-foundation/nuts-node/vdr/api/v1"
	"github.com/nuts-foundation/nuts-pgo-demo/domain/types"
)

type Service struct {
	VDRClient  nutsApi.HTTPClient
	Repository Repository
}

func (s Service) CreateAccount(username, password string) (*types.Account, error) {
	selfControl := true
	capabilityInvocation := true

	didDoc, err := s.VDRClient.Create(nutsApi.DIDCreateRequest{
		SelfControl:          &selfControl,
		CapabilityInvocation: &capabilityInvocation,
	})
	if err != nil {
		return nil, err
	}

	newAccount := types.Account{
		Username: username,
		Password: password,
		DID:      didDoc.ID,
	}
	return s.Repository.AddAccount(newAccount)
}
