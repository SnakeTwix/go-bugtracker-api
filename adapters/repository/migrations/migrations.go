package migrations

import (
	"gorm.io/gorm"
	domain2 "server/core/domain"
)

func RunMigrations(db *gorm.DB) error {
	err := db.SetupJoinTable(&domain2.User{}, "Projects", &domain2.UserProject{})

	if err != nil {
		return err
	}

	err = db.SetupJoinTable(&domain2.Project{}, "Members", &domain2.UserProject{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&domain2.User{},
		&domain2.ProjectGroup{},
		&domain2.Project{},
		&domain2.TicketTag{},
		&domain2.Ticket{},
		&domain2.TicketComment{},
		&domain2.UserProjectRole{},
	)

	return err
}
