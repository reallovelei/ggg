package id

import (
	"github.com/reallovelei/ggg/framework"
	"github.com/reallovelei/ggg/framework/contract"
)

type GGGIDProvider struct {
}

// Register registe a new function for make a service instance
func (provider *GGGIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewGGGIDService
}

// Boot will called when the service instantiate
func (provider *GGGIDProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (provider *GGGIDProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *GGGIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

/// Name define the name for this service
func (provider *GGGIDProvider) Name() string {
	return contract.IDKey
}
