package userPrivilege

type UserPrivilege int64

const (
	User UserPrivilege = iota
	Admin
)

func (e UserPrivilege) String() string {
	switch e {
	case User:
		return "user"
	case Admin:
		return "admin"
	}

	return "UNDEFINED_USER_PRIVILEGE_VALUE"
}
