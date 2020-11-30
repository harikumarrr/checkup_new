package checkup

import (
	"encoding/json"
	"fmt"

	"checkup/storage/appinsights"
	"checkup/storage/fs"
	"checkup/storage/github"
	"checkup/storage/mysql"
	"checkup/storage/postgres"
	"checkup/storage/s3"
	"checkup/storage/sql"
	"checkup/storage/sqlite3"
)

func storageDecode(typeName string, config json.RawMessage) (Storage, error) {
	switch typeName {
	case sqlite3.Type:
		return sqlite3.New(config)
	case mysql.Type:
		return mysql.New(config)
	case postgres.Type:
		return postgres.New(config)
	case s3.Type:
		return s3.New(config)
	case github.Type:
		return github.New(config)
	case fs.Type:
		return fs.New(config)
	case sql.Type:
		return sql.New(config)
	case appinsights.Type:
		return appinsights.New(config)
	default:
		return nil, fmt.Errorf(errUnknownStorageType, typeName)
	}
}
