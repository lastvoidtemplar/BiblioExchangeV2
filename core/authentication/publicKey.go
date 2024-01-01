package authentication

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	authoptions "github.com/lastvoidtemplar/BiblioExchangeV2/core/authentication/auth_options"
)

func LoadPublicKey(opt authoptions.AuthOptions) (*rsa.PublicKey, error) {
	res, err := http.Get(opt.RealmInfoUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch public key from Keycloak - %s", err.Error())
	}

	body := map[string]interface{}{}
	json.NewDecoder(res.Body).Decode(&body)

	if _, ok := body["public_key"]; !ok {
		return nil, fmt.Errorf("missing public_ key field in the payload - %s", err.Error())
	}

	publicKey, ok := (body["public_key"]).(string)

	if !ok {
		return nil, fmt.Errorf("the public key field isn`t a string - %s", err.Error())
	}

	publicKeyPem := fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKey)

	rsaKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPem))

	if err != nil {
		return nil, fmt.Errorf("cannot parse public key to RSA public key - %s", err.Error())
	}

	return rsaKey, nil
}
