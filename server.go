package rpc

import (
	"github.com/abxuz/b-tools/v2/bcrypt"
	"github.com/abxuz/b-tools/v2/brpc"
)

func NewServer(serverPrivateKey string, clientPublicKeys ...string) (*brpc.Server, error) {
	var sPrivKey bcrypt.NoisePrivateKey
	if err := sPrivKey.FromString(serverPrivateKey); err != nil {
		return nil, err
	}

	e := brpc.NewServer()
	e.SetServerPrivateKey(sPrivKey)

	for _, clientPublicKey := range clientPublicKeys {
		var cPubKey bcrypt.NoisePublicKey
		if err := cPubKey.FromString(clientPublicKey); err != nil {
			return nil, err
		}
		e.AddClientPublicKey(cPubKey)
	}

	return e, nil
}
