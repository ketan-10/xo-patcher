package loaders

import (
	_ "github.com/go-sql-driver/mysql" // empty import, to load drivers
	"github.com/ketan-10/xo-patcher/xo/internal"
	"github.com/ketan-10/xo-patcher/xo/loaders/models"
)

// init is a special function like main, which will be auto-called on start
func init() {

	// Register mysql loader into the system.
	// though in xo-patcher we only have mysql loader
	internal.AllLoaders[internal.MYSQL] = &internal.LoaderImp{
		EnumList: models.MysqlEnums,
		DatabaseName: models.MysqlDatabaseName,
		EnumValueList: models.MysqlEnumValueList,
	}
}
