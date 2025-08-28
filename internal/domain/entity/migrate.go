package entity

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Products{},
		&Users{},
	)
}

func Drop(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&Products{},
		&Users{},
	)
}
