package container

import (
	customer "github.com/FACorreiaa/loci-proto/modules/customer/generated"

	"github.com/FACorreiaa/loci-proto/utils"
)

// Brokers is the encapsulation for all grpc clients. We use it this way so
// that we can:
//
// 1. Mock remotes in tests
// 2. Have a single point of change when contracts change
// 3. Use a common logic route for onboarding
//
// Ensure that you're using the interface type here and not the implementation
type Brokers struct {
	Customer customer.CustomerClient
}

// NewBrokers creates a common container instance for use with all cluster sevices
func NewBrokers(transportUtils *utils.TransportUtils) *Brokers {
	if transportUtils == nil {
		return nil
	}

	brokers := new(Brokers)
	brokers.TransportUtils = transportUtils
	utils.Transport = transportUtils

	return brokers
}
