package seeds

import (
	"mikti-depublic/app"
	"mikti-depublic/helper"
)

func Run() {
	db := app.GetDB()
	gormOps := helper.NewGormOps(db)

	SeedUsers(gormOps)
	SeedAdmin(gormOps)
	SeedEvent(gormOps)
	SeedTransactions(gormOps)
}
