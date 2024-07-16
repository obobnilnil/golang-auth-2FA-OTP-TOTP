package utility

import (
	"fmt"
	"net/url"

	"github.com/pquerna/otp"
)

func GenerateQRCodeURL(key *otp.Key) string {
	issuer := url.QueryEscape(key.Issuer())
	accountName := url.QueryEscape(key.AccountName())
	secret := url.QueryEscape(key.Secret())
	return fmt.Sprintf("otpauth://totp/%s:%s?secret=%s&issuer=%s", issuer, accountName, secret, issuer)
}
