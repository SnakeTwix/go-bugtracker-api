package migrations

import (
	"gorm.io/gorm"
	"server/core/domain"
)

func RunMigrations(db *gorm.DB) error {
	err := db.SetupJoinTable(&domain.User{}, "Projects", &domain.UserProject{})

	if err != nil {
		return err
	}

	err = db.SetupJoinTable(&domain.Project{}, "Members", &domain.UserProject{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&domain.User{},
		&domain.ProjectGroup{},
		&domain.Project{},
		&domain.TicketTag{},
		&domain.Ticket{},
		&domain.TicketComment{},
		&domain.UserProjectRole{},
		&domain.Session{},
	)

	return err
}
