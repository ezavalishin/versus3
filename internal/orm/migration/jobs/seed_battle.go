package jobs

import (
	"github.com/ezavalishin/versus3/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	firstBattle = &models.Battle{
		Title: "test battle",
	}

	units = []*models.Unit{
		&models.Unit{
			Title: "unit1",
		},
		&models.Unit{
			Title: "unit2",
		},
		&models.Unit{
			Title: "unit3",
		},
		&models.Unit{
			Title: "unit4",
		},
	}
)

// SeedUsers inserts the first users
var SeedBattles = &gormigrate.Migration{
	ID: "SEED_BATTLES",
	Migrate: func(db *gorm.DB) error {
		db.Create(&firstBattle)

		for _, u := range units {
			u.BattleID = firstBattle.ID
			db.Create(&u)
		}

		return nil
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstBattle).Error
	},
}
