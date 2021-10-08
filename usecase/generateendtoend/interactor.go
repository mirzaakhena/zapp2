package generateendtoend

import (
	"context"
	"github.com/mirzaakhena/zapp2/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type generateEndToEndInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenerateEndToEnd
func NewUsecase(outputPort Outport) Inport {
	return &generateEndToEndInteractor{
		outport: outputPort,
	}
}

// Execute the usecase GenerateEndToEnd
func (r *generateEndToEndInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	service.RunProcess(req.YamlFile)

	return res, nil
}
