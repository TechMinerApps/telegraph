package telegraph

import "errors"

const (
	// FieldShortName used as GetAccountInfo argument for getting account name.
	FieldShortName string = "short_name"

	// FieldAuthorName used as GetAccountInfo argument for getting author name.
	FieldAuthorName string = "author_name"

	// FieldAuthorURL used as GetAccountInfo argument for getting profile link.
	FieldAuthorURL string = "author_url"

	// FieldAuthURL used as GetAccountInfo argument for getting URL to authorize a browser on telegra.ph.
	FieldAuthURL string = "auth_url"

	// FieldPageCount used as GetAccountInfo argument for getting number of pages belonging to the Telegraph
	// account.
	FieldPageCount string = "page_count"
)

var (
	// ErrInvalidDataType is returned when ContentFormat function are passed a data argument of invalid type.
	ErrInvalidDataType = errors.New("invalid data type")

	// ErrNoInputData is returned when any method get nil argument.
	ErrNoInputData = errors.New("no input data")

	// ErrFloodWait is returned when recieving a flood error
	ErrFloodWait = errors.New("reaching telegraph flood limit")
)
