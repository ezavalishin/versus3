package migration

import (
	"fmt"
	log "github.com/ezavalishin/versus3/internal/logger"
	"github.com/ezavalishin/versus3/internal/orm/migration/jobs"
	"github.com/ezavalishin/versus3/internal/orm/models"
	"github.com/ezavalishin/versus3/pkg/consts"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Battle{},
		&models.BattleUser{},
		&models.Unit{},
		&models.Round{},
		&models.Pair{},
		&models.Friend{},
	).Error

	if err != nil {
		return err
	}
	return addIndexes(db)
}

func addIndexes(db *gorm.DB) (err error) {
	// Entity names
	//db.NewScope(&models.User{}).GetModelStruct().TableName(db)
	battlesTableName := consts.Tablenames.Battles
	// FKs
	if err := db.Model(&models.Unit{}).
		AddForeignKey("battle_id", battlesTableName+"(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}
	// Indexes
	// None needed so far
	return nil
}

func ServiceAutoMigration(db *gorm.DB) error {
	fmt.Println("migrate")
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("[Migration.InitSchema] Initializing database schema")

		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// Add more jobs, etc here
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}

	fmt.Println("here")

	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedBattles,
	})

	return m.Migrate()
}
