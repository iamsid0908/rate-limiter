package service

import "my-echo-app/domain"

type HealthService struct {
	HealthDomain domain.HealthDomain
}

func (c *HealthService) Check() (string, error) {
	health, err := c.HealthDomain.GetHealth()
	if err != nil {
		return "", err
	}

	return health, nil
}
