package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/okta/okta-sdk-golang/v2/tests"
)

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
	ctx, client, err := tests.NewClient(context.Background())
	if err != nil {
		return err
	}
	return sweepUsers(ctx, client)
}

func sweepUsers(ctx context.Context, client *okta.Client) error {
	users, _, err := client.User.ListUsers(ctx, &query.Params{Q: "admin_role"})
	if err != nil {
		return err
	}
	for _, u := range users {
		if err := ensureUserDelete(ctx, client, u.Id, u.Status); err != nil {
			return err
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
			return fmt.Errorf("failed to deprovision or delete user: %v", err)
		}
	}
	return nil
}
