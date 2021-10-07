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
		log.Fatalf("failed to clean up organization before integration tests: %v", err)
	}
	exitVal := m.Run()
	err = sweep()
	if err != nil {
		log.Fatalf("failed to clean up organization after integration tests: %v", err)
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
	charSetAlpha = "abcdefghijklmnopqrstuvwxyz"
	testPrefix   = "SDK_TEST_"
)

func randomEmail() string {
	return randomTestString() + "@example.com"
}

// randStringFromCharSet generates a random string by selecting characters from
// the charset provided
func randomTestString() string {
	result := make([]byte, 15)
	for i := 0; i < 15; i++ {
		result[i] = charSetAlpha[rand.Intn(len(charSetAlpha))]
	}
	return testPrefix + string(result)
}

func testName(name string) string {
	s := fmt.Sprintf("%s %s", randomTestString(), name)
	if len(s) > 50 {
		s = s[:50]
	}
	return s
}
