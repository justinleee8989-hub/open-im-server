package group

import "github.com/openimsdk/protocol/constant"

// Role assignment is reserved for owners; callers retain existing profile-edit checks.
func canUpdateMemberRole(operatorRole int32, changingRole bool) bool {
	return !changingRole || operatorRole == constant.GroupOwner
}
