package integration

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/okta/okta-sdk-golang/v2/tests"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestMain(m *testing.M) {
	err := sweep()
	if err != nil {
		log.Printf("failed to clean up organization before integration tests: %v", err)
	}
	exitVal := m.Run()
	err = sweep()
	if err != nil {
		log.Printf("failed to clean up organization after integration tests: %v", err)
	}
	os.Exit(exitVal)
}

// sweep the resources before running integration tests
func sweep() error {
	log.Println("[INFO] sweeping test users, groups and rules")
	ctx, client, err := tests.NewClient(context.Background())
	if err != nil {
		return err
	}
	err = sweepUsers(ctx, client)
	if err != nil {
		return err
	}
	err = sweepGroups(ctx, client)
	if err != nil {
		return err
	}
	return sweepGroupRules(ctx, client)
}

func sweepGroups(ctx context.Context, client *okta.Client) error {
	groups, _, err := client.Group.ListGroups(ctx, &query.Params{Q: "SDK_TEST"})
	if err != nil {
		return err
	}
	for _, g := range groups {
		_, err = client.Group.DeleteGroup(ctx, g.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func sweepGroupRules(ctx context.Context, client *okta.Client) error {
	groupRules, _, err := client.Group.ListGroupRules(ctx, &query.Params{Q: "SDK_TEST"})
	if err != nil {
		return err
	}
	for _, g := range groupRules {
		if g.Status == "ACTIVE" {
			_, err = client.Group.DeactivateGroupRule(ctx, g.Id)
			if err != nil {
				return err
			}
		}
		_, err = client.Group.DeleteGroupRule(ctx, g.Id, &query.Params{})
		if err != nil {
			return err
		}
	}
	return nil
}

func sweepUsers(ctx context.Context, client *okta.Client) error {
	users, resp, err := client.User.ListUsers(ctx, &query.Params{Q: "SDK_TEST", Limit: 200})
	if err != nil {
		return err
	}
	for _, u := range users {
		if err := ensureUserDelete(ctx, client, u.Id, u.Status); err != nil {
			return err
		}
	}
	for resp.HasNextPage() {
		var users []*okta.User
		resp, err = resp.Next(ctx, &users)
		if err != nil {
			return err
		}
		for _, u := range users {
			if err := ensureUserDelete(ctx, client, u.Id, u.Status); err != nil {
				return err
			}
		}
	}
	return nil
}

func ensureUserDelete(ctx context.Context, client *okta.Client, id, status string) error {
	// only deprovisioned users can be deleted fully from okta
	// make two passes on the user if they aren't deprovisioned already to deprovision them first
	passes := 2
	if status == "DEPROVISIONED" {
		passes = 1
	}
	for i := 0; i < passes; i++ {
		_, err := client.User.DeactivateOrDeleteUser(ctx, id, nil)
		if err != nil {
			return fmt.Errorf("failed to deprovision or delete user: %w", err)
		}
	}
	return nil
}

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
