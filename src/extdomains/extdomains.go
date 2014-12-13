package extdomains

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"domains"
	"log"
)

func AddToDB(domain domains.Domaincsv) {

	db, err := sql.Open("sqlite3", "/home/juno/git/goFastCgi/goFastCgi/singo.db")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	sqlstr := "insert into extdomains(Locale,Themes,Domain,Ip) values(?,?,?,?)"

	stmt, err := tx.Prepare(sqlstr)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	if _, err = stmt.Exec(domain.Locale, domain.Themes, domain.Name, domain.Ip); err != nil {
		log.Fatal(err)

	}
	stmt.Close()
	tx.Commit()
}
