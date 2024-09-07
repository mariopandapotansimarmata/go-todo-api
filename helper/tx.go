package helper

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfErr(errorRollback)
		panic((err))
	} else {
		tx.Commit()
	}
}
