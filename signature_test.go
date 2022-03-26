package fasapay

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Signature_GenerateSignature(t *testing.T) {
	dt := time.Date(2011, time.Month(7), 20, 15, 30, 0, 0, time.UTC)
	result := generateSignature("11123548cd3a5e5613325132112becf", "kata rahasia", dt)
	assert.Equal(t, "e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b", result)
}
