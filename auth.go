package fasapay

import (
	"crypto/sha256"
	"fmt"
	"time"
)

//GenerateAuthToken method
func generateAuthToken(apiKey string, apiSecret string, dt time.Time) string {
	h := sha256.New()
	str := apiKey + ":" + apiSecret + ":" + dt.Format("2006010215")
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
