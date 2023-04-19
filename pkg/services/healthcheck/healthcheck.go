package healthcheck

import "github.com/vucchaid/service-template/pkg/model"

func GetHealthStatus() *model.Health {
	health := new(model.Health)
	health.Status = "UP"
	return health
}
