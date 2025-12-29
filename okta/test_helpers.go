package okta

import (
	"math/rand"
	"time"
)

const (
	charSetAlphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charSetAlphaLower = "abcdefghijklmnopqrstuvwxyz"
	charSetNumeric    = "0123456789"
	charSetAlpha      = charSetAlphaLower + charSetAlphaUpper + charSetNumeric
	testPrefix        = "SDK_TEST_"
)

func randomEmail() string {
	return randomTestString() + "@example.com"
}

// randStringFromCharSet generates a random string of 15 lower case letters
func randomTestString() string {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	result := make([]byte, 15)
	for i := range 15 {
		result[i] = charSetAlphaLower[rng.Intn(len(charSetAlphaLower))]
	}
	return testPrefix + string(result)
}

// testPassword generates a random string of at least 4 characters in length
func testPassword(length int) string {
	if length < 8 {
		length = 8
	}
	result := make([]byte, length)
	result[0] = charSetAlphaLower[rand.Intn(len(charSetAlphaLower))]
	result[1] = charSetAlphaUpper[rand.Intn(len(charSetAlphaUpper))]
	result[2] = charSetNumeric[rand.Intn(len(charSetNumeric))]
	for i := 3; i < length; i++ {
		result[i] = charSetAlpha[rand.Intn(len(charSetAlpha))]
	}
	return string(result)
}

type TestFactory struct{}

var testFactory TestFactory

func (t *TestFactory) NewValidTestUserProfile() UserProfile {
	email := randomEmail()
	firstName := "John"
	lastName := "Doe"
	login := email

	profile := NewUserProfile()
	profile.SetFirstName(firstName)
	profile.SetLastName(lastName)
	profile.SetEmail(email)
	profile.SetLogin(login)

	return *profile
}

func (t *TestFactory) NewTestUserProfileUpdate() UserProfile {
	email := randomEmail()
	firstName := "Jane"
	lastName := "Smith"
	login := email

	profile := NewUserProfile()
	profile.SetFirstName(firstName)
	profile.SetLastName(lastName)
	profile.SetEmail(email)
	profile.SetLogin(login)

	return *profile
}

func (t *TestFactory) NewValidTestUserCredentialsWithPassword() *UserCredentialsWritable {
	pc := t.NewValidTestPasswordCredential()
	creds := NewUserCredentialsWritable()
	creds.SetPassword(*pc)
	return creds
}

func (t *TestFactory) NewValidTestPasswordCredential() *PasswordCredential {
	p := testPassword(10)
	pc := NewPasswordCredential()
	pc.SetValue(p)
	return pc
}

func (t *TestFactory) NewValidTestGroupProfile() OktaUserGroupProfile {
	name := randomTestString()
	description := "Test group created by SDK tests"

	profile := NewOktaUserGroupProfile()
	profile.SetName(name)
	profile.SetDescription(description)

	return *profile
}

func (t *TestFactory) NewTestGroupProfileUpdate() OktaUserGroupProfile {
	name := randomTestString() + "_updated"
	description := "Updated test group description"

	profile := NewOktaUserGroupProfile()
	profile.SetName(name)
	profile.SetDescription(description)

	return *profile
}

func (t *TestFactory) NewValidTestAddGroupRequest() AddGroupRequest {
	profile := t.NewValidTestGroupProfile()

	request := NewAddGroupRequest()
	request.SetProfile(profile)

	return *request
}

func (t *TestFactory) NewValidTestAccessPolicy() AccessPolicy {
	name := randomTestString() + "_policy"
	description := "Test access policy created by SDK tests"
	policyType := "ACCESS_POLICY"

	policy := NewAccessPolicy(name, policyType)
	policy.SetDescription(description)

	return *policy
}

func (t *TestFactory) NewTestAccessPolicyUpdate() AccessPolicy {
	name := randomTestString() + "_updated_policy"
	description := "Updated test policy description"
	policyType := "ACCESS_POLICY"

	policy := NewAccessPolicy(name, policyType)
	policy.SetDescription(description)

	return *policy
}

func (t *TestFactory) NewValidTestCreatePolicyRequest() CreatePolicyRequest {
	accessPolicy := t.NewValidTestAccessPolicy()
	return AccessPolicyAsCreatePolicyRequest(&accessPolicy)
}

func (t *TestFactory) NewValidTestPolicyRule() ListPolicyRules200ResponseInner {
	name := randomTestString() + "_rule"

	rule := NewAccessPolicyRule()
	rule.SetName(name)
	rule.SetType("ACCESS_POLICY")

	peopleCondition := NewPolicyPeopleCondition()

	conditions := NewAccessPolicyRuleConditions()
	conditions.SetPeople(*peopleCondition)

	rule.SetConditions(*conditions)

	return AccessPolicyRuleAsListPolicyRules200ResponseInner(rule)
}

func (t *TestFactory) NewValidTestBookmarkApplication() BookmarkApplication {
	label := "Test Bookmark App"
	name := "bookmark"
	url := "https://example.com/bookmark"

	appSettings := NewBookmarkApplicationSettingsApplication(url)
	appSettings.SetRequestIntegration(false)

	settings := NewBookmarkApplicationSettings()
	settings.SetApp(*appSettings)

	app := NewBookmarkApplication(name, *settings, label, "BOOKMARK")

	return *app
}

func (t *TestFactory) NewValidTestCreateApplicationRequest() ListApplications200ResponseInner {
	bookmarkApp := t.NewValidTestBookmarkApplication()
	return BookmarkApplicationAsListApplications200ResponseInner(&bookmarkApp)
}

func (t *TestFactory) NewTestBookmarkApplicationUpdate() BookmarkApplication {
	label := "Updated Test Bookmark App"
	name := "bookmark"
	url := "https://example.com/updated-bookmark"

	appSettings := NewBookmarkApplicationSettingsApplication(url)
	appSettings.SetRequestIntegration(false)

	settings := NewBookmarkApplicationSettings()
	settings.SetApp(*appSettings)

	app := NewBookmarkApplication(name, *settings, label, "BOOKMARK")

	return *app
}
