package models

import (
	"database/sql"
	"github.com/spf13/viper"
	"github.com/wow-go-basic-cms/entities"
)

type BnetAccountModel struct {
	Db *sql.DB
}

func (bnetAccount BnetAccountModel) GetAllBnetAccount() ([]entities.Bnet_accounts, error) {
	// battlenet_accounts
	bnetAcc := viper.GetString("auth_database.bnet_accounts_table")
	query_get_bnet_acc := "select id,email,sha_pass_hash from " + bnetAcc
	rows, err := bnetAccount.Db.Query(query_get_bnet_acc)

	if err != nil {
		return nil, err
	} else {
		bnetAccounts := []entities.Bnet_accounts{}
		for rows.Next() {
			var id int
			var email string
			var sha_pass_hash string

			err2 := rows.Scan(&id, &email, &sha_pass_hash)
			if err2 != nil {
				return nil, err2
			} else {
				bnetAccount := entities.Bnet_accounts{id, email, sha_pass_hash}
				bnetAccounts = append(bnetAccounts, bnetAccount)
			}
		}
		return bnetAccounts, nil
	}
}
