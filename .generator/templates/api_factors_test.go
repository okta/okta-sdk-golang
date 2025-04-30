package okta

import "testing"

func ListFactors(t *testing.T) ([]ListFactors200ResponseInner, *APIResponse, error) {
	req := apiClient.UserFactorAPI.ListFactors(apiClient.cfg.Context, "00unr46bcnSPtbcVj5d7")
	return req.Execute()
}

func TestListFactors(t *testing.T) {
	factors, resp, err := ListFactors(t)
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
		if factor.Id == "" {
			t.Fatal("Expected factor ID to be non-empty")
		}
	}
}
