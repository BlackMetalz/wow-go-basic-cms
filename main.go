package main

import (
	"fmt"
	"github.com/wow-go-basic-cms/config"
	"github.com/wow-go-basic-cms/models"
)

func main() {
	db := config.GetMysqlConfig()

	bnetAccountModel := models.BnetAccountModel{
		Db: db,
	}

	fmt.Println("Bnet Account list")
	bnetAccounts, err2 := bnetAccountModel.GetAllBnetAccount()

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Total Bnet Accounts: ", len(bnetAccounts))
		for _, bnetaccount := range bnetAccounts {
			fmt.Println("Id: ", bnetaccount.Id)
			fmt.Println("Email: ", bnetaccount.Email)
			fmt.Println("Sha_pass_hash: ", bnetaccount.Sha_pass_hash)
		}
	}

}
