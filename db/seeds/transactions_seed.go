package seeds

import (
	"fmt"
	"mikti-depublic/helper"
)

func SeedTransactions(gormOps *helper.GormOpsImpl) {
	transactions := helper.FakerTransaction()

	err := gormOps.Create(&transactions)
	if err != nil {
		return
	}

	fmt.Println("Seed Transactions Success")
}
