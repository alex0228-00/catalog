package extractor

import (
	"context"
	"fmt"
)

type Extractor interface {
	Browse(ctx context.Context) error
}

type ExtractionError struct {
	artifactId string
	err        error
}

func NewExtractionError(artifactId string, err error) *ExtractionError {
	return &ExtractionError{
		artifactId: artifactId,
		err:        err,
	}
}

func (e *ExtractionError) Error() string {
	return fmt.Sprintf("Failed to extract artifact %s, error=%s", e.artifactId, e.err.Error())
}
