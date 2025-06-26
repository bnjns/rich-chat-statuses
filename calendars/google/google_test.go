package google

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithCredentialsJson(t *testing.T) {
	t.Run("should set the credentials json", func(t *testing.T) {
		expectedCredentialsJson := `{"type": "service_account"}`
		options := &Options{}

		fn := WithCredentialsJson(expectedCredentialsJson)

		fn(options)

		assert.Equal(t, expectedCredentialsJson, options.CredentialsJson)
	})
}
