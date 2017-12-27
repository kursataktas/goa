package calc

import (
	"context"
	"log"

	calcsvc "goa.design/goa/examples/calc/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsvcSvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calcsvc.Service {
	return &calcsvcSvc{logger}
}

// Add implements add.
func (s *calcsvcSvc) Add(ctx context.Context, p *calcsvc.AddPayload) (int, error) {
	return p.A + p.B, nil
}
