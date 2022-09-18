package helpers

import (
	"fmt"
	"gopher-camp/pkg/helpers/seed"
	"gopher-camp/pkg/models"
	"gopher-camp/pkg/types"
	"gorm.io/gorm"
)

func SeedDatabaseData(db *gorm.DB, logger types.Logger) error {
	var founders []*models.Founder
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
	founders = append(founders, founder)

	company, err := databaseSeeder.CreateCompany(founders)
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
