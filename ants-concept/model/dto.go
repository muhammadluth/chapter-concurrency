package model

import "time"

type Config struct {
	PoolSize         int           `json:"pool_size"`
	MinPoolSize      int           `json:"min_pool_size"`
	MaxPoolSize      int           `json:"max_pool_size"`
	PoolIncrement    int           `json:"pool_increment"`
	AutoTuneDuration time.Duration `json:"auto_tune_duration"`
}
