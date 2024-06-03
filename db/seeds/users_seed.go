package seeds

import (
	"fmt"
	"log"
	"mikti-depublic/helper"
)

func SeedUsers(gormOps *helper.GormOpsImpl) {
	users := helper.FakerUser()
	err := gormOps.Create(&users)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println("Seed Users Success")
}
