package gdrepo

import (
	"fmt"

	"github.com/thinhphamnls/gd/container"
)

func buildTableName(schema, table, alias string) string {
	if schema == "" {
		schema = container.DefaultSchemaGorillaDesk
	}
	return fmt.Sprintf("%s.%s AS %s", schema, table, alias)
}
