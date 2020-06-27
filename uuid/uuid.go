package uuid

import (
	"github.com/Celtcoste/orc-graphql/src/postgresql"
	"github.com/google/uuid"
)

func GenerateUID(tableName string) (*string, error) {
	uid := uuid.New()
	var i = 0
	err := postgresql.WithTransaction(func(tx postgresql.Transaction) error {
		row := tx.QueryRow("SELECT COUNT(uid) FROM "+tableName+
			" WHERE uid = $1",
			uid)
		err := row.Scan(&i)
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	if i == 1 {
		return GenerateUID(tableName)
	}
	resUID := uid.String()
	return &resUID, nil
}
