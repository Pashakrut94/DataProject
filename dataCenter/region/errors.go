package region

import (
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("no information in database about this region")
)