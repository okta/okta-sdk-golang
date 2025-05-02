package okta

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ListFactors() ([]ListFactors200ResponseInner, *APIResponse, error) {
	req := apiClient.UserFactorAPI.ListFactors(apiClient.cfg.Context, "00unr46bcnSPtbcVj5d7")
	return req.Execute()
}

func TestListFactors(t *testing.T) {
	factors, resp, err := ListFactors()
	if err != nil {
		t.Fatalf("Error listing factors: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %d", resp.StatusCode)
	}
	if len(factors) == 0 {
		t.Fatal("Expected at least one factor")
	}
	factorTypes := make([]string, 0)
	for _, factor := range factors {
		var result map[string]interface{}
		marshalJSON, err := factor.MarshalJSON()
		json.Unmarshal(marshalJSON, &result)
		if err != nil {
			return
		}
		if factorType, ok := result["factorType"].(string); ok {
			//fmt.Println("factorType:", factorType)
			factorTypes = append(factorTypes, factorType)
		}
	}
	found := false
	for f := range factorTypes {
		if factorTypes[f] == "signed_nonce" {
			found = true
		}
	}
	assert.Equal(t, found, true)
}
