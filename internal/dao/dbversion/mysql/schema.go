package mysql

import (
	bsstorages "github.com/common_structure/internal/dao/dbversion"
)

func (m *schemaManager) UpgradeSchema(currentRevs *schemaRevision) (schemaChanged bool, err error) {
	status := bsstorages.SchemaUpgradeStatus{
		Changed:   false,
		LastError: nil,
	}
	return status.Changed, status.LastError
}
