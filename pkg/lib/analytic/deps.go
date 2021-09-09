package analytic

import (
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	NewStoreRedisLogger,
	wire.Struct(new(GlobalDBStore), "*"),
	wire.Struct(new(AppDBStore), "*"),
	wire.Struct(new(AuditDBStore), "*"),
	wire.Struct(new(UserWeeklyReport), "*"),
	wire.Struct(new(ProjectWeeklyReport), "*"),
	wire.Struct(new(Service), "*"),
	wire.Struct(new(StoreRedis), "*"),
	wire.Bind(new(CounterStore), new(*StoreRedis)),
)
