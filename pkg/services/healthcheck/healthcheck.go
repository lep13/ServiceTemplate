package healthcheck

import "service-template/pkg/model"

func GetHealthStatus() *model.Health {
	health := new(model.Health)
	health.Status = "UP"
	return health
}
