# Okta SDK Testing Guide

## Overview

The Okta SDK now includes functional unit tests that replace the previous skipped tests. These tests use real API calls to validate SDK functionality.

## Environment Setup

Before running the tests, you need to set up your Okta environment:

### Required values

#### Set the values as Environment variables

```bash
export OKTA_CLIENT_ORGURL="https://your-okta-domain.okta.com"
export OKTA_CLIENT_TOKEN="your-okta-api-token"
```

#### Set the values in ~/.okta/okta.yaml or .okta.yaml

```yaml
okta:
  client:
    orgUrl: "https://your-okta-domain.okta.com"
    token: "your-okta-api-token"
```

### Getting Your Okta Credentials

1. **Okta Organization URL**: Your Okta domain (e.g., `https://dev-123456.okta.com`)
2. **API Token**: Generate an API token in your Okta Admin Console:
   - Go to Security > API > Tokens
   - Click "Create Token"
   - Give it a descriptive name
   - Copy the token value

## Running Tests

### Run User API Tests
```bash
# Run all user tests
go test -v ./okta/test -run Test_okta_UserAPIService

# Run specific test
go test -v ./okta/test -run TestUserAPIService_CreateUser
```

### Run All Tests
```bash
# Run all API tests
go test -v ./okta/test
```

## Test Features

### Automatic Resource Cleanup
- Tests automatically track and clean up created users
- No manual cleanup required
- Resources are cleaned up even if tests fail

### Test Data Management
- Singleton TestDataManager handles resource lifecycle
- Automatic tracking of created resources
- Environment validation before test execution

### Test Helpers
- Factory methods for creating test data
- Consistent test user profiles
- Password and credential generation

## Test Structure

### TestMain Pattern
Each test file uses TestMain for setup/teardown:
- Environment validation
- Resource cleanup
- Centralized configuration

### Test Data Factory
The TestFactory provides methods for creating test data:
- `NewValidTestUserProfile()` - Creates user profiles
- `NewValidTestUserCredentialsWithPassword()` - Creates credentials
- `NewTestUserProfileUpdate()` - Creates update data

### Real API Testing
Tests use actual Okta API calls to:
- Validate SDK functionality
- Test against real API responses
- Ensure compatibility with Okta service

## Example Test

```go
func TestCreateUser(t *testing.T) {
    // Create test data
    testFactory := &okta.TestFactory{}
    profile := testFactory.NewValidTestUserProfile()
    credentials := testFactory.NewValidTestUserCredentialsWithPassword()

    // Create user via API
    createReq := okta.NewCreateUserRequest(profile)
    createReq.SetCredentials(*credentials)

    resp, httpRes, err := apiClient.UserAPI.CreateUser(testContext).
        Body(*createReq).Activate(false).Execute()

    // Assertions
    require.NoError(t, err)
    assert.Equal(t, 200, httpRes.StatusCode)
    assert.NotNil(t, resp.Id)

    // Automatic cleanup via TestDataManager
}
```

## Security Notes

- Never commit API tokens to the repository
- Use environment variables for credentials
- Tests create and delete real resources in your Okta org
- Run tests in a development/test environment only
