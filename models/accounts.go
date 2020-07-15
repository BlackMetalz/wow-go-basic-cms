package models

import (
	"database/sql"
	"github.com/wow-go-basic-cms/entities"
)

type BnetAccountModel struct {
	Db *sql.DB
}

func (bnetAccount BnetAccountModel) GetAllBnetAccount() ([]entities.Bnet_accounts, error) {
	rows, err := bnetAccount.Db.Query("select id,email,sha_pass_hash from bfa801_battlenet_accounts")

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
