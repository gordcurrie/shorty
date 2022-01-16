package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upShortlyUrls, downShortlyUrls)
}

func upShortlyUrls(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE urls (
			id big(20) unsigned NOT NULL AUTO_INCREMENT,
			hash varchar(20) NOT NULL,
			url varchar(500) NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downShortlyUrls(tx *sql.Tx) error {
	_, err := tx.Exec(`DROP TABLE urls;`)
	if err != nil {
		return err
	}
	return nil
}
