package helpers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopher-camp/pkg/seed"
	"gopher-camp/pkg/types"
)

func SeedDatabaseData(db *gorm.DB, logger types.Logger) error {
	databaseSeeder := seed.NewDatabaseSeeder(db)
	founder, err := databaseSeeder.CreateFounder()
	if err != nil {
		logger.LogError(err, "helpers.SeedDatabaseData", "seed")
		return err
	}
	logger.LogInfo(
		fmt.Sprintf("** Seeded a founder: %v **", founder),
		"helpers.SeedDatabaseData",
		"helpers",
	)

	company, err := databaseSeeder.CreateCompany(founder)
	if err != nil {
		logger.LogError(err, "helpers.SeedDatabaseData", "seed")
		return err
	}
	logger.LogInfo(
		fmt.Sprintf("** Seeded a company: %v **", company),
		"helpers.SeedDatabaseData",
		"helpers",
	)
	return nil
}
