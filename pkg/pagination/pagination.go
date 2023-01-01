package pagination

import "github.com/google/uuid"

type TotalEntries int64

type PageOffset struct {
	Offset int
	Limit  int
}

type PageCursor struct {
	Cursor uuid.UUID
	Limit  int
}
