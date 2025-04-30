package okta

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Fix: remove *testing.T from here
func ListFactors() ([]ListFactors200ResponseInner, *APIResponse, error) {
	req := apiClient.UserFactorAPI.ListFactors(apiClient.cfg.Context, "00unr46bcnSPtbcVj5d7")
	return apiClient.UserFactorAPI.ListFactorsExecute(req)
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
	for _, factor := range factors {
		var result map[string]interface{}
		marshalJSON, err := factor.MarshalJSON()
		json.Unmarshal(marshalJSON, &result)
		if err != nil {
			return
		}
		if factorType, ok := result["factorType"].(string); ok {
			fmt.Println("factorType:", factorType)
		} else {
			fmt.Println("factorType key not found or not a string")
		}

	}
}
