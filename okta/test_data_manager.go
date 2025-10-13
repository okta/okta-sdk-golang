package okta

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// TestDataManager manages test resources and cleanup
type TestDataManager struct {
	client       *APIClient
	createdUsers []string
	mutex        sync.Mutex
	ctx          context.Context
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
			client:       client,
			createdUsers: make([]string, 0),
			ctx:          context.Background(),
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
