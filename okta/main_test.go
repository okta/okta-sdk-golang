package okta

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
)

var apiClient *APIClient

func init() {
	configuration := NewConfiguration(WithCache(false))
	configuration.Debug = false

	apiClient = NewAPIClient(configuration)
}

func TestMain(m *testing.M) {
	err := sweep()
	if err != nil {
		fmt.Printf("failed to clean up organization before integration tests: %v", err)
	}
	exitVal := m.Run()
	apiClient = apiClient.RefreshNext()
	err = sweep()
	if err != nil {
		fmt.Printf("failed to clean up organization after integration tests: %v", err)
	}
	os.Exit(exitVal)
}

func sweep() (err error) {
	log.Println("[INFO] sweeping test users, groups, idps")
	err = sweepUsers()
	if err != nil {
		return
	}
	err = sweepGroups()
	if err != nil {
		return
	}
	err = sweepGroupRules()
	if err != nil {
		return
	}
	// err = sweepApps()
	// if err != nil {
	// 	return
	// }
	return sweepIdps()
}

func sweepUsers() error {
	req := apiClient.UserApi.ListUsers(apiClient.cfg.Context).Limit(200)
	req = req.Q("SDK_TEST_")
	users, resp, err := req.Execute()
	if err != nil {
		return err
	}
	for _, u := range users {
		if err := cleanUpUser(u.GetId()); err != nil {
			return err
		}
	}
	for resp.HasNextPage() {
		var users []*User
		resp, err = resp.Next(&users)
		if err != nil {
			return err
		}
		for _, u := range users {
			if err := cleanUpUser(u.GetId()); err != nil {
				return err
			}
		}
	}
	return nil
}

func sweepGroups() error {
	req := apiClient.GroupApi.ListGroups(apiClient.cfg.Context).Limit(200)
	req = req.Q("SDK_TEST")
	groups, resp, err := req.Execute()
	if err != nil {
		return err
	}
	for _, g := range groups {
		if err := cleanUpGroup(g.GetId()); err != nil {
			return err
		}
	}
	for resp.HasNextPage() {
		var groups []*Group
		resp, err = resp.Next(&groups)
		if err != nil {
			return err
		}
		for _, g := range groups {
			if err := cleanUpGroup(g.GetId()); err != nil {
				return err
			}
		}
	}
	return nil
}

func sweepIdps() error {
	req := apiClient.IdentityProviderApi.ListIdentityProviders(apiClient.cfg.Context).Limit(200)
	req = req.Q("SDK_TEST")
	idps, resp, err := req.Execute()
	if err != nil {
		return err
	}
	for _, idp := range idps {
		if err := cleanUpIdp(idp.GetId()); err != nil {
			return err
		}
	}
	for resp.HasNextPage() {
		var idps []*IdentityProvider
		resp, err = resp.Next(&idps)
		if err != nil {
			return err
		}
		for _, idp := range idps {
			if err := cleanUpIdp(idp.GetId()); err != nil {
				return err
			}
		}
	}
	return nil
}

func sweepGroupRules() error {
	req := apiClient.GroupApi.ListGroupRules(apiClient.cfg.Context).Limit(200)
	req = req.Search("SDK_TEST")
	groupRules, resp, err := req.Execute()
	if err != nil {
		return err
	}
	for _, gr := range groupRules {
		if gr.GetStatus() == "ACTIVE" {
			_, err = apiClient.GroupApi.DeactivateGroupRule(apiClient.cfg.Context, gr.GetId()).Execute()
			if err != nil {
				return err
			}
		}
		if err := cleanUpGroupRule(gr.GetId()); err != nil {
			return err
		}
	}
	for resp.HasNextPage() {
		var groupRules []*GroupRule
		resp, err = resp.Next(&groupRules)
		if err != nil {
			return err
		}
		for _, gr := range groupRules {
			if gr.GetStatus() == "ACTIVE" {
				_, err = apiClient.GroupApi.DeactivateGroupRule(apiClient.cfg.Context, gr.GetId()).Execute()
				if err != nil {
					return err
				}
			}
			if err := cleanUpGroupRule(gr.GetId()); err != nil {
				return err
			}
		}
	}
	return nil
}

func sweepApps() error {
	req := apiClient.ApplicationApi.ListApplications(apiClient.cfg.Context).Limit(200)
	req = req.Q("SDK_TEST")
	apps, _, err := req.Execute()
	if err != nil {
		return err
	}
	for _, a := range apps {
		id, err := getAppId(a)
		if err != nil {
			fmt.Printf("unknown app type %v", a)
			continue
		}
		if err := cleanUpApplication(id); err != nil {
			return err
		}
	}
	return nil
}

func getAppId(app ListApplications200ResponseInner) (string, error) {
	if app.AutoLoginApplication != nil {
		return app.AutoLoginApplication.GetId(), nil
	} else if app.BasicAuthApplication != nil {
		return app.BasicAuthApplication.GetId(), nil
	} else if app.BookmarkApplication != nil {
		return app.BookmarkApplication.GetId(), nil
	} else if app.BrowserPluginApplication != nil {
		return app.BrowserPluginApplication.GetId(), nil
	} else if app.OpenIdConnectApplication != nil {
		return app.OpenIdConnectApplication.GetId(), nil
	} else if app.SamlApplication != nil {
		return app.SamlApplication.GetId(), nil
	} else if app.SecurePasswordStoreApplication != nil {
		return app.SecurePasswordStoreApplication.GetId(), nil
	} else if app.WsFederationApplication != nil {
		return app.WsFederationApplication.GetId(), nil
	}
	return "", errors.New("unknown app type")
}
