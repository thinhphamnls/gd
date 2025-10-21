package gdentity

type UserRoleType int32

const (
	UserRoleTypeSuperAdmin UserRoleType = 1
	UserRoleTypeManager    UserRoleType = 2
	UserRoleTypeSalesRep   UserRoleType = 3
	UserRoleTypeTechnician UserRoleType = 4
)

func ParseUserRoleType(s string) UserRoleType {
	switch s {
	case "super_admin":
		return UserRoleTypeSuperAdmin
	case "sales_manager":
		return UserRoleTypeManager
	case "sales_rep":
		return UserRoleTypeSalesRep
	case "technician":
		return UserRoleTypeTechnician
	default:
		return 0
	}
}
