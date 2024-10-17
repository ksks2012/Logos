# Heading Code

* `keep-empty-line`

```go
package mysql

import (
  "database/sql"
  "fmt"
  "context"

  metastore "github.com/semeqetjsakatayza/go-metastore-mysql"
)

```

# DistributedLockMeta (distributed-lock-meta) r.1

* `strip-spaces`

```sql
```

## Routines

### prepare fetch revision

```go
metaStoreInst := metastore.MetaStore{
  TableName: metaStoreTableName,
  Ctx:       m.ctx,
  Conn:      m.conn,
}
```

### fetch revision

```go
if ${SCHEMA_REV_VAR}, _, err = metaStoreInst.FetchRevision(${SCHEMA_REV_KEY}); nil != err {
  return nil, err
}
```

### update revision

```go
metaStoreInst := metastore.MetaStore{
  TableName: metaStoreTableName,
  Ctx:       m.ctx,
  Conn:      m.conn,
}
err = metaStoreInst.StoreRevision(key, rev)
```

<!-- TODO: create DB -->
