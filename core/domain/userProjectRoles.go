package domain

type UserProjectRole struct {
	ProjectID uint64
	Project   *Project

	UserID uint64
	User   *User

	RoleID uint64
	Role   ProjectRole `gorm:"foreignKey:RoleID"`
}
