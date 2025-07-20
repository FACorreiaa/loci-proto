package utils

import (
	//g "github.com/FACorreiaa/loci-proto/protocol/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func NewConnection(serverAddress string) (*grpc.ClientConn, error) {
	tu := Transport
	if tu == nil {
		return nil, errors.New("transport utils are required")
	}

	//conn, err := g.BootstrapClient(
	//	serverAddress,
	//	tu.Logger,
	//	tu.Prometheus,
	//	tu.TraceProvider,
	//)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to connect to upstream host")
	//}

	return nil, nil
}
