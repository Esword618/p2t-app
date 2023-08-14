package global

import (
	"context"
	"gorm.io/gorm"
)

var (
	GlobalCtx context.Context
	GlobalDb  *gorm.DB
)
