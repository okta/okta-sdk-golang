## Migrating from 2.x to 3.x

In releases prior to version 3 we use an Open API v2 specification, and an Okta custom client generator to partially generate our SDK. A new version of the Open API specification (V3) has been released, and new well-known generators are now available and well received by the community. Planning the future of this SDK, we consider this a good opportunity to modernize by aligning with established standards for API client generation. 

### Enums

`StringEnums` are now supported. However, format might slightly change depending on the OpenAPI specification and codegen.

### Features parity

The following features have been ported to 3.x:

* Inline configuration, configuration via environment variables, appsettings.json or YAML files
* Manual pagination for collections
* Default retry strategy for 429 HTTP responses and ability to provide your own strategy
* Web proxy 
* OAuth for Okta

### New APIs

In order to provide better structuring, some endpoints have been moved from an existing client/API to their own API client:

_Before:_

```go
ctx, client, err := tests.NewClient(context.TODO())
role := okta.AssignRoleRequest{
    Type: "ORG_ADMIN",
}
createdRole, response, err := client.Group.AssignRoleToGroup(ctx, group.Id, role, nil)
roles, response, err := client.Group.ListGroupAssignedRoles(ctx, group.Id, nil)
```

_Now:_

```go
apiClient = NewAPIClient()
assignedRoleSA := NewAssignRoleRequest()
assignedRoleSA.SetType("SUPER_ADMIN")
createdRole, response, err := apiClient.RoleAssignmentApi.AssignRoleToGroup(context.TODO(), group.GetId()).AssignRoleRequest(*assignedRoleSA).Execute()
roles, response, err := apiClient.RoleAssignmentApi.ListGroupAssignedRoles(context.TODO(), group.GetId()).Execute()
```

For more details about other APIs, please check out [here](API_README.md#documentation-for-API-Endpoints).

