package health_check

import (
	health "github.com/AppsFlyer/go-sundheit"
	"logur.dev/logur"
)

type HealthListener struct {
	name   string
	logger logur.Logger
}

func NewHealthListener(logger logur.Logger, serviceName string) HealthListener {
	return HealthListener{
		name:   serviceName,
		logger: logger,
	}
}

func (c HealthListener) OnResultsUpdated(results map[string]health.Result) {
	for _, result := range results {
		if result.Error != nil {
			c.logger.Info("Health check failed", map[string]interface{}{"service": c.name, "error": result.Error.Error()})
			return
		}

		c.logger.Info("Health check completed", map[string]interface{}{"service": c.name})
	}
}
