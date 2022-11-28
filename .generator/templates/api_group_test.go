package okta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGroup(name string) (*Group, *APIResponse, error) {
	req := apiClient.GroupApi.CreateGroup(apiClient.cfg.Context)
	gp := NewGroupProfile()
	gp.SetName(name)
	payload := Group{Profile: gp}
	req = req.Group(payload)
	return req.Execute()
}

func cleanUpGroup(groupId string) (err error) {
	_, err = apiClient.GroupApi.DeleteGroup(apiClient.cfg.Context, groupId).Execute()
	return
}

func cleanUpGroupRule(groupRuleId string) (err error) {
	_, err = apiClient.GroupApi.DeleteGroupRule(apiClient.cfg.Context, groupRuleId).Execute()
	return
}

func Test_Get_Group(t *testing.T) {
	t.Parallel()
	group, _, err := setupGroup(randomTestString())
	require.NoError(t, err, "Creating a new group should not error")
	t.Run("get group by id", func(t *testing.T) {
		gid, _, err := apiClient.GroupApi.GetGroup(apiClient.cfg.Context, group.GetId()).Execute()
		require.NoError(t, err, "Could not get group by ID")
		assert.Equal(t, group.GetId(), gid.GetId())
	})
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
}

func Test_Get_List_Group(t *testing.T) {
	t.Parallel()
	group, _, err := setupGroup(randomTestString())
	require.NoError(t, err, "Creating a new group should not error")
	t.Run("get list group", func(t *testing.T) {
		gs, _, err := apiClient.GroupApi.ListGroups(apiClient.cfg.Context).Execute()
		require.NoError(t, err, "Could not get list group")
		var createdGroupInList bool
		for _, g := range gs {
			if group.GetId() == g.GetId() {
				createdGroupInList = true
			}
		}
		assert.True(t, createdGroupInList, "Could not find group from list")
	})
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
}

func Test_Search_Group(t *testing.T) {
	t.Parallel()
	groupName := randomTestString()
	group, _, err := setupGroup(groupName)
	require.NoError(t, err, "Creating a new group should not error")
	t.Run("search group", func(t *testing.T) {
		req := apiClient.GroupApi.ListGroups(apiClient.cfg.Context)
		req = req.Q(groupName)
		gs, _, err := req.Execute()
		require.NoError(t, err, "Could not get result from search keyword")
		var createdGroupInList bool
		for _, g := range gs {
			if group.GetId() == g.GetId() {
				createdGroupInList = true
			}
		}
		assert.True(t, createdGroupInList, "Could not find group from list")
	})
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
}

func Test_Update_Group(t *testing.T) {
	t.Parallel()
	oldGroupName := randomTestString()
	group, _, err := setupGroup(oldGroupName)
	require.NoError(t, err, "Creating a new group should not error")
	t.Run("update group", func(t *testing.T) {
		newGroupName := randomTestString()
		ngp := GroupProfile{}
		ngp.SetName(newGroupName)
		ng := Group{}
		ng.SetProfile(ngp)
		req := apiClient.GroupApi.ReplaceGroup(apiClient.cfg.Context, group.GetId())
		req = req.Group(ng)
		g, _, err := req.Execute()
		require.NoError(t, err, "Could not update group")
		assert.NotNil(t, g.Profile, "Group profile is nil")
		assert.Equal(t, g.Profile.GetName(), newGroupName)
	})
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
}

func Test_Group_User_Operation(t *testing.T) {
	t.Parallel()
	group, _, err := setupGroup(randomTestString())
	require.NoError(t, err, "Creating a new group should not error")
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	t.Run("add user to group", func(t *testing.T) {
		_, err := apiClient.GroupApi.AssignUserToGroup(apiClient.cfg.Context, group.GetId(), user.GetId()).Execute()
		require.NoError(t, err, "Could not add user to group")
		users, _, err := apiClient.GroupApi.ListGroupUsers(apiClient.cfg.Context, group.GetId()).Execute()
		require.NoError(t, err)
		found := false
		for _, u := range users {
			if u.GetId() == user.GetId() {
				found = true
			}
		}
		assert.True(t, found, "Could not find user in group")
	})
	t.Run("remove user from group", func(t *testing.T) {
		_, err := apiClient.GroupApi.UnassignUserFromGroup(apiClient.cfg.Context, group.GetId(), user.GetId()).Execute()
		require.NoError(t, err, "Could not remove user from group")
		users, _, err := apiClient.GroupApi.ListGroupUsers(apiClient.cfg.Context, group.GetId()).Execute()
		require.NoError(t, err)
		found := false
		for _, u := range users {
			if u.GetId() == user.GetId() {
				found = true
			}
		}
		assert.False(t, found, "Found user in group")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
}

func Test_Group_Rule_Operation(t *testing.T) {
	t.Parallel()
	group, _, err := setupGroup(randomTestString())
	require.NoError(t, err, "Creating a new group should not error")
	user, _, _, err := setupUser(true)
	require.NoError(t, err, "Creating a new user should not error")
	grce := NewGroupRuleExpression()
	grce.SetType("urn:okta:expression:1.0")
	grce.SetValue("user.lastName==\"" + user.Profile.GetLastName() + "\"")
	grc := NewGroupRuleConditions()
	grc.SetExpression(*grce)
	grga := NewGroupRuleGroupAssignment()
	grga.SetGroupIds([]string{group.GetId()})
	gra := NewGroupRuleAction()
	gra.SetAssignUserToGroups(*grga)
	gr := NewGroupRule()
	gr.SetActions(*gra)
	gr.SetConditions(*grc)
	gr.SetType("group_rule")
	gr.SetName(randomTestString())
	req := apiClient.GroupApi.CreateGroupRule(apiClient.cfg.Context)
	req = req.GroupRule(*gr)
	groupRule, _, err := req.Execute()
	require.NoError(t, err, "Creating a new group rule should not error")
	t.Run("activate group rule", func(t *testing.T) {
		_, err = apiClient.GroupApi.ActivateGroupRule(apiClient.cfg.Context, groupRule.GetId()).Execute()
		require.NoError(t, err, "Should not error when activating rule")
		groupRules, _, err := apiClient.GroupApi.ListGroupRules(apiClient.cfg.Context).Execute()
		require.NoError(t, err, "Should not error when listing group rule")
		found := false
		for _, grs := range groupRules {
			if groupRule.GetId() == grs.GetId() {
				found = true
			}
		}
		assert.True(t, found, "Found group rule in list")
	})
	t.Run("deactivate group rule", func(t *testing.T) {
		_, err = apiClient.GroupApi.DeactivateGroupRule(apiClient.cfg.Context, groupRule.GetId()).Execute()
		require.NoError(t, err, "Should not error when deactivating rule")
	})
	t.Run("update group rule", func(t *testing.T) {
		grce := NewGroupRuleExpression()
		grce.SetType("urn:okta:expression:1.0")
		grce.SetValue("user.lastName==\"Incorrect\"")
		grc := NewGroupRuleConditions()
		grc.SetExpression(*grce)
		grga := NewGroupRuleGroupAssignment()
		grga.SetGroupIds([]string{group.GetId()})
		gra := NewGroupRuleAction()
		gra.SetAssignUserToGroups(*grga)
		gr := NewGroupRule()
		gr.SetActions(*gra)
		gr.SetConditions(*grc)
		gr.SetType("group_rule")
		gr.SetName(randomTestString())
		req := apiClient.GroupApi.ReplaceGroupRule(apiClient.cfg.Context, groupRule.GetId())
		req = req.GroupRule(*gr)
		newGroupRule, _, err := req.Execute()
		require.NoError(t, err, "Should not error when updating rule")
		_, err = apiClient.GroupApi.ActivateGroupRule(apiClient.cfg.Context, newGroupRule.GetId()).Execute()
		require.NoError(t, err, "Should not error when activating rule")
		_, err = apiClient.GroupApi.DeactivateGroupRule(apiClient.cfg.Context, groupRule.GetId()).Execute()
		require.NoError(t, err, "Should not error when deactivating rule")
	})
	err = cleanUpUser(user.GetId())
	require.NoError(t, err, "Clean up user should not error")
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
	err = cleanUpGroupRule(groupRule.GetId())
	require.NoError(t, err, "Clean up group rule should not error")
}

func Test_List_Assigned_Applications_For_Group(t *testing.T) {
	t.Parallel()
	group, _, err := setupGroup(randomTestString())
	require.NoError(t, err, "Creating a new group should not error")
	var createdApp *ListApplications200ResponseInner
	t.Run("get list assigned application for group", func(t *testing.T) {
		apps, _, err := apiClient.GroupApi.ListAssignedApplicationsForGroup(apiClient.cfg.Context, group.GetId()).Execute()
		require.NoError(t, err, "Should not error when listing assigned apps for group")
		assert.Equal(t, 0, len(apps), "there shouldn't be any apps assigned to group")
		createdApp, _, err = setupBookmarkApplication(randomTestString())
		require.NoError(t, err, "Creating an application should not error")
		aareq := apiClient.ApplicationApi.AssignGroupToApplication(apiClient.cfg.Context, createdApp.BookmarkApplication.GetId(), group.GetId())
		aareq.applicationGroupAssignment = NewApplicationGroupAssignment()
		_, _, err = aareq.Execute()
		require.NoError(t, err, "Assigning application to group should not error")
		apps, _, err = apiClient.GroupApi.ListAssignedApplicationsForGroup(apiClient.cfg.Context, group.GetId()).Execute()
		require.NoError(t, err, "Should not error when listing assigned apps for group")
		assert.Equal(t, 1, len(apps), "there shouldn't be any apps assigned to group")
	})
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
	err = cleanUpApplication(createdApp.BookmarkApplication.GetId())
	require.NoError(t, err, "Clean up group should not error")
}

func Test_Assigned_Role_To_Group_Operation(t *testing.T) {
	t.Parallel()
	group, _, err := setupGroup(randomTestString())
	require.NoError(t, err, "Creating a new group should not error")
	var createdRole *Role
	t.Run("assigned role to group", func(t *testing.T) {
		req := apiClient.RoleAssignmentApi.AssignRoleToGroup(apiClient.cfg.Context, group.GetId())
		assignedRoleSA := NewAssignRoleRequest()
		assignedRoleSA.SetType("SUPER_ADMIN")
		req = req.AssignRoleRequest(*assignedRoleSA)
		createdRole, _, err = req.Execute()
		require.NoError(t, err, "Assigned role to group should not error")
		roles, _, err := apiClient.RoleAssignmentApi.ListGroupAssignedRoles(apiClient.cfg.Context, group.GetId()).Execute()
		require.NoError(t, err, "Listing group assigned role should not error")
		var found bool
		for _, r := range roles {
			if r.GetId() == createdRole.GetId() {
				found = true
			}
		}
		assert.True(t, found)
	})
	t.Run("unassigned role to group", func(t *testing.T) {
		_, err = apiClient.RoleAssignmentApi.UnassignRoleFromGroup(apiClient.cfg.Context, group.GetId(), createdRole.GetId()).Execute()
		require.NoError(t, err, "Unassigned role to group should not error")
		roles, _, err := apiClient.RoleAssignmentApi.ListGroupAssignedRoles(apiClient.cfg.Context, group.GetId()).Execute()
		require.NoError(t, err, "Listing group assigned role should not error")
		assert.Empty(t, roles)
	})
	err = cleanUpGroup(group.GetId())
	require.NoError(t, err, "Clean up group should not error")
}
