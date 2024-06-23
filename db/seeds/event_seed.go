package seeds

import (
	"fmt"
	"mikti-depublic/helper"
)

func SeedEvent(gormOps *helper.GormOpsImpl) {
	event := helper.FakerEvent()

	err := gormOps.Create(&event)
	if err != nil {
		return
	}

	fmt.Println("Seed Event Success")
}
