package captcha

import (
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	NewCloudflareClient,
)
