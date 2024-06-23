package seeds

import (
	"fmt"
	"mikti-depublic/helper"
)

func SeedAdmin(gormOps *helper.GormOpsImpl) {
	admin := helper.FakerAdmin()

	err := gormOps.Create(&admin)
	if err != nil {
		return
	}

	fmt.Println("Seed Admin Success")
}
