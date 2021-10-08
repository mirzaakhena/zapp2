package generateendtoend

import (
	"context"
)

// Inport of GenerateEndToEnd
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase GenerateEndToEnd
type InportRequest struct {
	YamlFile string
}

// InportResponse is response payload after running the usecase GenerateEndToEnd
type InportResponse struct {
}
