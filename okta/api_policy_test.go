package okta

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupAccessPolicy(name string) (*ListPolicies200ResponseInner, *APIResponse, error) {
	configuration, err := NewConfiguration()
	if err != nil {
		return nil, nil, err
	}
	configuration.Debug = true
	proxyClient := NewAPIClient(configuration)
	req := proxyClient.PolicyAPI.CreatePolicy(apiClient.cfg.Context)
	req = req.Policy(ListPolicies200ResponseInner{AccessPolicy: testFactory.NewValidAccessPolicy(name)})
	return req.Execute()
}

func cleanUpPolicy(policyId string) error {
	_, err := apiClient.PolicyAPI.DeactivatePolicy(apiClient.cfg.Context, policyId).Execute()
	if err != nil {
		return err
	}
	_, err = apiClient.PolicyAPI.DeletePolicy(apiClient.cfg.Context, policyId).Execute()
	if err != nil {
		return err
	}
	return err
}

func cleanUpPolicyRule(policyId, policyRuleId string) (err error) {
	_, err = apiClient.PolicyAPI.DeactivatePolicyRule(apiClient.cfg.Context, policyId, policyRuleId).Execute()
	if err != nil {
		return err
	}
	_, err = apiClient.PolicyAPI.DeletePolicyRule(apiClient.cfg.Context, policyId, policyRuleId).Execute()
	if err != nil {
		return err
	}
	return err
}

func Test_Get_Policy(t *testing.T) {
	createdPolicy, _, err := setupAccessPolicy(randomTestString())
	require.NoError(t, err, "Creating a new policy should not error")
	t.Run("get policy by id", func(t *testing.T) {
		policy, _, err := apiClient.PolicyAPI.GetPolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
		require.NoError(t, err, "Could not get policy by ID")
		assert.Equal(t, createdPolicy.AccessPolicy.GetId(), policy.AccessPolicy.GetId())
	})
	err = cleanUpPolicy(createdPolicy.AccessPolicy.GetId())
	require.NoError(t, err, "Clean up policy should not error")
}

func Test_Get_List_Policies(t *testing.T) {
	createdPolicy, _, err := setupAccessPolicy(randomTestString())
	require.NoError(t, err, "Creating a new policy should not error")
	t.Run("get all policy", func(t *testing.T) {
		policies, _, err := apiClient.PolicyAPI.ListPolicies(apiClient.cfg.Context).Type_("ACCESS_POLICY").Execute()
		require.NoError(t, err, "Could not get list policy")
		var createPolicyInList bool
		for _, p := range policies {
			if (p.AccessPolicy != nil) && (p.AccessPolicy.GetId() == createdPolicy.AccessPolicy.GetId()) {
				createPolicyInList = true
			}
		}
		assert.True(t, createPolicyInList, "Could not find policy from list")
	})
	err = cleanUpPolicy(createdPolicy.AccessPolicy.GetId())
	require.NoError(t, err, "Clean up policy should not error")
}

func Test_Update_Policies(t *testing.T) {
	createdPolicy, _, err := setupAccessPolicy(randomTestString())
	require.NoError(t, err, "Creating a new policy should not error")
	t.Run("update policy", func(t *testing.T) {
		newName := randomTestString()
		payload := testFactory.NewValidAccessPolicy(newName)
		policy, _, err := apiClient.PolicyAPI.ReplacePolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Policy(ListPolicies200ResponseInner{AccessPolicy: payload}).Execute()
		require.NoError(t, err, "Could not update policy")
		require.NotNil(t, policy.AccessPolicy)
		assert.Equal(t, newName, policy.AccessPolicy.GetName())
	})
	err = cleanUpPolicy(createdPolicy.AccessPolicy.GetId())
	require.NoError(t, err, "Clean up policy should not error")
}

func Test_Activate_Policy(t *testing.T) {
	createdPolicy, _, err := setupAccessPolicy(randomTestString())
	require.NoError(t, err, "Creating a new policy should not error")
	t.Run("deactivate policy", func(t *testing.T) {
		_, err = apiClient.PolicyAPI.DeactivatePolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
		require.NoError(t, err, "Could not deactivate the policy")
		policy, _, err := apiClient.PolicyAPI.GetPolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
		require.NoError(t, err, "Could not get policy by ID")
		assert.Equal(t, createdPolicy.AccessPolicy.GetId(), policy.AccessPolicy.GetId())
		assert.Equal(t, LIFECYCLESTATUS_INACTIVE, policy.AccessPolicy.GetStatus())
	})
	t.Run("activate policy", func(t *testing.T) {
		_, err = apiClient.PolicyAPI.ActivatePolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
		require.NoError(t, err, "Could not activate the policy")
		policy, _, err := apiClient.PolicyAPI.GetPolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
		require.NoError(t, err, "Could not get policy by ID")
		assert.Equal(t, createdPolicy.AccessPolicy.GetId(), policy.AccessPolicy.GetId())
		assert.Equal(t, LIFECYCLESTATUS_ACTIVE, policy.AccessPolicy.GetStatus())
	})
	err = cleanUpPolicy(createdPolicy.AccessPolicy.GetId())
	require.NoError(t, err, "Clean up policy should not error")
}

// ACCESS/AUTHENTICATION POLICY ONLY
// TODO Used to work, now fail with 401
func Test_Clone_Policy(t *testing.T) {
	createdPolicy, _, err := setupAccessPolicy(randomTestString())
	require.NoError(t, err, "Creating a new policy should not error")
	var policyID string
	t.Run("clone policy", func(t *testing.T) {
		policy, _, err := apiClient.PolicyAPI.ClonePolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
		require.NoError(t, err, "Could not clone policy")
		policyID = policy.AccessPolicy.GetId()
		assert.NotEqual(t, createdPolicy.AccessPolicy.GetId(), policy.AccessPolicy.GetId())
		assert.Equal(t, fmt.Sprintf("[cloned] %v", createdPolicy.AccessPolicy.GetName()), policy.AccessPolicy.GetName())
		assert.Equal(t, createdPolicy.AccessPolicy.GetDescription(), policy.AccessPolicy.GetDescription())
		assert.Equal(t, createdPolicy.AccessPolicy.GetPriority(), policy.AccessPolicy.GetPriority())
	})
	err = cleanUpPolicy(createdPolicy.AccessPolicy.GetId())
	require.NoError(t, err, "Clean up policy should not error")
	err = cleanUpPolicy(policyID)
	require.NoError(t, err, "Clean up policy should not error")
}

func Test_Policy_Rules_Operation(t *testing.T) {
	createdPolicy, _, err := setupAccessPolicy(randomTestString())
	require.NoError(t, err, "Creating a new policy should not error")
	configuration, err := NewConfiguration()
	require.NoError(t, err, "Creating a new config should not error")
	configuration.Debug = true
	proxyClient := NewAPIClient(configuration)
	accessPolicyRule := &AccessPolicyRule{}
	accessPolicyRule.SetType(POLICYRULETYPE_ACCESS_POLICY)
	name := randomTestString()
	accessPolicyRule.SetName(name)
	payload := ListPolicyRules200ResponseInner{AccessPolicyRule: accessPolicyRule}
	createdPolicyRule, _, err := proxyClient.PolicyAPI.CreatePolicyRule(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).PolicyRule(payload).Execute()
	require.NoError(t, err, "Creating a new policy rule should not error")
	t.Run("get policy rule by id", func(t *testing.T) {
		rpolicyRule, _, err := apiClient.PolicyAPI.GetPolicyRule(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId(), createdPolicyRule.AccessPolicyRule.GetId()).Execute()
		require.NoError(t, err, "Could not get policy rule by ID")
		assert.Equal(t, name, rpolicyRule.AccessPolicyRule.GetName())
	})
	t.Run("list policy rule", func(t *testing.T) {
		rpolicyRules, _, err := apiClient.PolicyAPI.ListPolicyRules(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
		require.NoError(t, err, "Could not listing policy rule by ID")
		found := false
		for _, pr := range rpolicyRules {
			if pr.AccessPolicyRule.GetId() == createdPolicyRule.AccessPolicyRule.GetId() {
				found = true
			}
		}
		assert.True(t, found, "Found policy rule in list")
	})
	t.Run("update policy rule", func(t *testing.T) {
		newName := randomTestString()
		createdPolicyRule.AccessPolicyRule.SetName(newName)
		rpolicyRule, _, err := apiClient.PolicyAPI.ReplacePolicyRule(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId(), createdPolicyRule.AccessPolicyRule.GetId()).PolicyRule(*createdPolicyRule).Execute()
		require.NoError(t, err, "Could not update policy rule")
		assert.NotEqual(t, name, rpolicyRule.AccessPolicyRule.GetName())
		assert.Equal(t, newName, rpolicyRule.AccessPolicyRule.GetName())
	})
	t.Run("deactivate policy rule", func(t *testing.T) {
		_, err = apiClient.PolicyAPI.DeactivatePolicyRule(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId(), createdPolicyRule.AccessPolicyRule.GetId()).Execute()
		require.NoError(t, err, "Could not deactivate policy rule")
		rpolicyRule, _, err := apiClient.PolicyAPI.GetPolicyRule(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId(), createdPolicyRule.AccessPolicyRule.GetId()).Execute()
		require.NoError(t, err, "Could not get policy rule by ID")
		assert.Equal(t, LIFECYCLESTATUS_INACTIVE, rpolicyRule.AccessPolicyRule.GetStatus())
	})

	t.Run("activate policy rule", func(t *testing.T) {
		_, err = apiClient.PolicyAPI.ActivatePolicyRule(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId(), createdPolicyRule.AccessPolicyRule.GetId()).Execute()
		require.NoError(t, err, "Could not activate policy rule")
		rpolicyRule, _, err := apiClient.PolicyAPI.GetPolicyRule(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId(), createdPolicyRule.AccessPolicyRule.GetId()).Execute()
		require.NoError(t, err, "Could not get policy rule by ID")
		assert.Equal(t, LIFECYCLESTATUS_ACTIVE, rpolicyRule.AccessPolicyRule.GetStatus())
	})
	err = cleanUpPolicyRule(createdPolicy.AccessPolicy.GetId(), createdPolicyRule.AccessPolicyRule.GetId())
	require.NoError(t, err, "Clean up policy rule should not error")
	err = cleanUpPolicy(createdPolicy.AccessPolicy.GetId())
	require.NoError(t, err, "Clean up policy should not error")
}

// TODU
// func Test_Policy_Mapping_Operations(t *testing.T) {
// 	createdPolicy, _, err := setupAccessPolicy(randomTestString())
// 	require.NoError(t, err, "Creating a new policy should not error")
// 	createdApp, _, err := setupBasicAuthApplication(randomTestString())
// 	require.NoError(t, err, "Creating a new application should not error")
// 	t.Run("assign app to policy", func(t *testing.T) {
// 		_, err = apiClient.ApplicationApi.UpdateApplicationPolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
// 		require.NoError(t, err, "Could not deactivate the policy")
// 		policy, _, err := apiClient.PolicyApi.ClonePolicy(apiClient.cfg.Context, createdPolicy.AccessPolicy.GetId()).Execute()
// 		require.NoError(t, err, "Could not get policy by ID")
// 		assert.Equal(t, createdPolicy.AccessPolicy.GetId(), policy.AccessPolicy.GetId())
// 		assert.Equal(t, "INACTIVE", policy.AccessPolicy.GetStatus())
// 	})
// 	err = cleanUpPolicy(createdPolicy.AccessPolicy.GetId())
// 	require.NoError(t, err, "Clean up policy should not error")
// 	err = cleanUpApplication(createdApp.BasicAuthApplication.GetId())
// 	require.NoError(t, err, "Clean up app should not error")
// }
