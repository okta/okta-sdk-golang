package okta

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// TestDataManager manages test resources and cleanup
type TestDataManager struct {
	client              *APIClient
	createdUsers        []string
	createdGroups       []string
	createdPolicies     []string
	createdApplications []string
	mutex               sync.Mutex
	ctx                 context.Context
}

var testDataManager *TestDataManager
var testDataOnce sync.Once

// GetTestDataManager returns singleton instance of test data manager
func GetTestDataManager() *TestDataManager {
	testDataOnce.Do(func() {
		config, err := NewConfiguration()
		if err != nil {
			log.Fatalf("Failed to create test configuration: %v", err)
		}

		client := NewAPIClient(config)
		testDataManager = &TestDataManager{
			client:              client,
			createdUsers:        make([]string, 0),
			createdGroups:       make([]string, 0),
			createdPolicies:     make([]string, 0),
			createdApplications: make([]string, 0),
			ctx:                 context.Background(),
		}
	})
	return testDataManager
}

// CreateTestUser creates a user for testing and tracks it for cleanup
func (tdm *TestDataManager) CreateTestUser() (*User, error) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	profile := testFactory.NewValidTestUserProfile()
	credentials := testFactory.NewValidTestUserCredentialsWithPassword()

	createReq := NewCreateUserRequest(profile)
	createReq.SetCredentials(*credentials)

	req := tdm.client.UserAPI.CreateUser(tdm.ctx).Body(*createReq).Activate(false)
	createdUser, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create test user: %w", err)
	}

	if createdUser.Id != nil {
		tdm.createdUsers = append(tdm.createdUsers, *createdUser.Id)
	}

	return createdUser, nil
}

// CreateAndActivateTestUser creates and activates a user for testing
func (tdm *TestDataManager) CreateAndActivateTestUser() (*User, error) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	profile := testFactory.NewValidTestUserProfile()
	credentials := testFactory.NewValidTestUserCredentialsWithPassword()

	createReq := NewCreateUserRequest(profile)
	createReq.SetCredentials(*credentials)

	req := tdm.client.UserAPI.CreateUser(tdm.ctx).Body(*createReq).Activate(true)
	createdUser, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create and activate test user: %w", err)
	}

	if createdUser.Id != nil {
		tdm.createdUsers = append(tdm.createdUsers, *createdUser.Id)
	}

	return createdUser, nil
}

// CleanupTestUser safely removes a specific test user
func (tdm *TestDataManager) CleanupTestUser(userID string) error {
	// First deactivate the user if not already deactivated using UserLifecycleAPI
	_, deactivateErr := tdm.client.UserLifecycleAPI.DeactivateUser(tdm.ctx, userID).Execute()
	if deactivateErr != nil {
		log.Printf("Warning: Failed to deactivate user %s: %v", userID, deactivateErr)
	}

	// Then delete the user
	_, deleteErr := tdm.client.UserAPI.DeleteUser(tdm.ctx, userID).Execute()
	if deleteErr != nil {
		return fmt.Errorf("failed to delete user %s: %w", userID, deleteErr)
	}

	return nil
}

// TrackUser adds a user ID to the tracking list for cleanup
func (tdm *TestDataManager) TrackUser(userID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()
	tdm.createdUsers = append(tdm.createdUsers, userID)
}

// RemoveUserFromTracking removes a user ID from the tracking list
func (tdm *TestDataManager) RemoveUserFromTracking(userID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for i, id := range tdm.createdUsers {
		if id == userID {
			// Remove the element by slicing
			tdm.createdUsers = append(tdm.createdUsers[:i], tdm.createdUsers[i+1:]...)
			break
		}
	}
}

// CleanupAllTestUsers removes all tracked test users
func (tdm *TestDataManager) CleanupAllTestUsers() {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for _, userID := range tdm.createdUsers {
		if err := tdm.CleanupTestUser(userID); err != nil {
			log.Printf("Failed to cleanup user %s: %v", userID, err)
		}
	}

	tdm.createdUsers = tdm.createdUsers[:0]
}

// CreateTestGroup creates a group for testing and tracks it for cleanup
func (tdm *TestDataManager) CreateTestGroup() (*Group, error) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	groupRequest := testFactory.NewValidTestAddGroupRequest()

	req := tdm.client.GroupAPI.AddGroup(tdm.ctx).Group(groupRequest)
	createdGroup, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create test group: %w", err)
	}

	if createdGroup.Id != nil {
		tdm.createdGroups = append(tdm.createdGroups, *createdGroup.Id)
	}

	return createdGroup, nil
}

// CleanupTestGroup safely removes a specific test group
func (tdm *TestDataManager) CleanupTestGroup(groupID string) error {
	_, deleteErr := tdm.client.GroupAPI.DeleteGroup(tdm.ctx, groupID).Execute()
	if deleteErr != nil {
		return fmt.Errorf("failed to delete group %s: %w", groupID, deleteErr)
	}

	return nil
}

// TrackGroup adds a group ID to the tracking list for cleanup
func (tdm *TestDataManager) TrackGroup(groupID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()
	tdm.createdGroups = append(tdm.createdGroups, groupID)
}

// RemoveGroupFromTracking removes a group ID from the tracking list
func (tdm *TestDataManager) RemoveGroupFromTracking(groupID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for i, id := range tdm.createdGroups {
		if id == groupID {
			tdm.createdGroups = append(tdm.createdGroups[:i], tdm.createdGroups[i+1:]...)
			break
		}
	}
}

// CleanupAllTestGroups removes all tracked test groups
func (tdm *TestDataManager) CleanupAllTestGroups() {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for _, groupID := range tdm.createdGroups {
		if err := tdm.CleanupTestGroup(groupID); err != nil {
			log.Printf("Failed to cleanup group %s: %v", groupID, err)
		}
	}

	tdm.createdGroups = tdm.createdGroups[:0]
}

// CreateTestPolicy creates a policy for testing and tracks it for cleanup
func (tdm *TestDataManager) CreateTestPolicy() (*CreatePolicyRequest, error) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	policyRequest := testFactory.NewValidTestCreatePolicyRequest()

	req := tdm.client.PolicyAPI.CreatePolicy(tdm.ctx).Policy(policyRequest)
	createdPolicy, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create test policy: %w", err)
	}

	if createdPolicy.AccessPolicy != nil && createdPolicy.AccessPolicy.Id != nil {
		tdm.createdPolicies = append(tdm.createdPolicies, *createdPolicy.AccessPolicy.Id)
	}

	return createdPolicy, nil
}

// CleanupTestPolicy safely removes a specific test policy
func (tdm *TestDataManager) CleanupTestPolicy(policyID string) error {
	_, deleteErr := tdm.client.PolicyAPI.DeletePolicy(tdm.ctx, policyID).Execute()
	if deleteErr != nil {
		return fmt.Errorf("failed to delete policy %s: %w", policyID, deleteErr)
	}

	return nil
}

// TrackPolicy adds a policy ID to the tracking list for cleanup
func (tdm *TestDataManager) TrackPolicy(policyID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()
	tdm.createdPolicies = append(tdm.createdPolicies, policyID)
}

// RemovePolicyFromTracking removes a policy ID from the tracking list
func (tdm *TestDataManager) RemovePolicyFromTracking(policyID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for i, id := range tdm.createdPolicies {
		if id == policyID {
			tdm.createdPolicies = append(tdm.createdPolicies[:i], tdm.createdPolicies[i+1:]...)
			break
		}
	}
}

// CleanupAllTestPolicies removes all tracked test policies
func (tdm *TestDataManager) CleanupAllTestPolicies() {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for _, policyID := range tdm.createdPolicies {
		if err := tdm.CleanupTestPolicy(policyID); err != nil {
			log.Printf("Failed to cleanup policy %s: %v", policyID, err)
		}
	}

	tdm.createdPolicies = tdm.createdPolicies[:0]
}

// CreateTestApplication creates an application for testing and tracks it for cleanup
func (tdm *TestDataManager) CreateTestApplication() (*ListApplications200ResponseInner, error) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	appRequest := testFactory.NewValidTestCreateApplicationRequest()

	req := tdm.client.ApplicationAPI.CreateApplication(tdm.ctx).Application(appRequest)
	createdApp, _, err := req.Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to create test application: %w", err)
	}

	if createdApp.BookmarkApplication != nil && createdApp.BookmarkApplication.Id != nil {
		tdm.createdApplications = append(tdm.createdApplications, *createdApp.BookmarkApplication.Id)
	}

	return createdApp, nil
}

// CleanupTestApplication safely removes a specific test application
func (tdm *TestDataManager) CleanupTestApplication(appID string) error {
	// First deactivate the application if not already deactivated
	_, deactivateErr := tdm.client.ApplicationAPI.DeactivateApplication(tdm.ctx, appID).Execute()
	if deactivateErr != nil {
		log.Printf("Warning: Failed to deactivate application %s: %v", appID, deactivateErr)
	}

	// Then delete the application
	_, deleteErr := tdm.client.ApplicationAPI.DeleteApplication(tdm.ctx, appID).Execute()
	if deleteErr != nil {
		return fmt.Errorf("failed to delete application %s: %w", appID, deleteErr)
	}

	return nil
}

// TrackApplication adds an application ID to the tracking list for cleanup
func (tdm *TestDataManager) TrackApplication(appID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()
	tdm.createdApplications = append(tdm.createdApplications, appID)
}

// RemoveApplicationFromTracking removes an application ID from the tracking list
func (tdm *TestDataManager) RemoveApplicationFromTracking(appID string) {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for i, id := range tdm.createdApplications {
		if id == appID {
			tdm.createdApplications = append(tdm.createdApplications[:i], tdm.createdApplications[i+1:]...)
			break
		}
	}
}

// CleanupAllTestApplications removes all tracked test applications
func (tdm *TestDataManager) CleanupAllTestApplications() {
	tdm.mutex.Lock()
	defer tdm.mutex.Unlock()

	for _, appID := range tdm.createdApplications {
		if err := tdm.CleanupTestApplication(appID); err != nil {
			log.Printf("Failed to cleanup application %s: %v", appID, err)
		}
	}

	tdm.createdApplications = tdm.createdApplications[:0]
}

// ValidateTestEnvironment checks if test environment is properly configured
func ValidateTestEnvironment() error {
	config, err := NewConfiguration()
	if err != nil {
		return fmt.Errorf("failed to create configuration: %w", err)
	}

	if config.Okta.Client.OrgUrl == "" || config.Okta.Client.Token == "" {
		return fmt.Errorf(`OKTA_CLIENT_ORGURL or OKTA_CLIENT_TOKE not configured. Set it via:
			1. Environment variable: export OKTA_CLIENT_ORGURL="https://your-domain.okta.com"
			2. System config file: ~/.okta/okta.yaml
			3. Project config file: .okta.yaml
			See TESTING.md for complete configuration examples`)
	}

	return nil
}

// CleanupAll removes all tracked test resources
func (tdm *TestDataManager) CleanupAll() {
	tdm.CleanupAllTestUsers()
	tdm.CleanupAllTestGroups()
	tdm.CleanupAllTestPolicies()
	tdm.CleanupAllTestApplications()
}
