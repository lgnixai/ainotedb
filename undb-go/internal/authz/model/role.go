
package model

type Role string

const (
	RoleOwner  Role = "owner"
	RoleAdmin  Role = "admin"
	RoleMember Role = "member"
	RoleViewer Role = "viewer"
)

type Permission struct {
	Action string `json:"action"`
	Allow  bool   `json:"allow"`
}

type RolePermissions map[Role][]Permission

var DefaultPermissions = RolePermissions{
	RoleOwner: {
		{Action: "*", Allow: true},
	},
	RoleAdmin: {
		{Action: "read:*", Allow: true},
		{Action: "write:*", Allow: true},
		{Action: "delete:*", Allow: true},
	},
	RoleMember: {
		{Action: "read:*", Allow: true},
		{Action: "write:record", Allow: true},
	},
	RoleViewer: {
		{Action: "read:*", Allow: true},
	},
}
