package model

import "time"

// Base is information that is expected on all records
type Base struct {
	// id is the unique globally identifier for the record. E.g. Batch:lkjd3jkl1lkf42
	ID string `json:"id"`
	// createdAt is the time of creation of the record
	CreatedAt time.Time `json:"createdAt"`
	// updatedAt is the time of last update of the record
	UpdatedAt time.Time `json:"updatedAt"`

	// CreatedByID records the User ID of the user that created the record
	CreatedByID *string `json:"createdByID"`
}

// CreatedBy returns the User that created the record.
func (b *Base) CreatedBy() (*User, error) {
	return nil, nil
}

// Timed is for records that have a start, an end, and a status.
type Timed struct {
	Status    *Status    `json:"status"`
	StartedAt *time.Time `json:"startedAt"`
	EndedAt   *time.Time `json:"endedAt"`
}

// Sorted records keep a unique index within a scope.
type Sorted struct {
	Index *uint32 `json:"index"`
}

// Archivable records can be archived. If the value is nil, the record is not
// archived.
type Archivable struct {
	ArchivedAt *time.Time `json:"archivedAt"`
}

// Configured records take a configuration on init that can change the
// behavior of the object.
type Configured struct {
	// Configuration returns an open-ended configuration JSON for the record
	// (see documentation for details).
	Configuration *string `json:"configuration"`
}
