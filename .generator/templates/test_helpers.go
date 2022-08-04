package okta

import (
	"fmt"
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
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 15)
	for i := 0; i < 15; i++ {
		result[i] = charSetAlphaLower[rand.Intn(len(charSetAlphaLower))]
	}
	return testPrefix + string(result)
}

// testPassword generates a random string of at least 4 characters in length
func testPassword(length int) string {
	if length < 5 {
		length = 4
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

func testName(name string) string {
	s := fmt.Sprintf("%s %s", randomTestString(), name)
	if len(s) > 50 {
		s = s[:50]
	}
	return s
}

type TestFactory struct{}

var testFactory TestFactory

func (t *TestFactory) NewValidTestUserProfile() UserProfile {
	email := randomEmail()
	firstName := "John"
	lastName := "Doe"
	return UserProfile{
		FirstName: NullableString{value: &firstName, isSet: true},
		LastName:  NullableString{value: &lastName, isSet: true},
		Email:     &email,
		Login:     &email,
	}
}

func (t *TestFactory) NewValidTestUserCredentialsWithPassword() *UserCredentials {
	pc := t.NewValidTestPasswordCredential()
	return &UserCredentials{Password: pc}
}

func (t *TestFactory) NewValidTestPasswordCredential() *PasswordCredential {
	p := testPassword(10)
	return &PasswordCredential{Value: &p}
}

func (t *TestFactory) NewValidTestRecoveryQuestionCredential() *RecoveryQuestionCredential {
	question := "How many roads must a man walk down?"
	answer := "forty two"
	return &RecoveryQuestionCredential{Question: &question, Answer: &answer}
}
