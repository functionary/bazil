package pb // import "bazil.org/bazil/pb"

import "errors"

var ErrEmptyMessage = errors.New("Unmarshaled an empty message")
