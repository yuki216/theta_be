package enum

type Status string

const (
	StatusDraft     Status = "DRAFT"
	StatusActive    Status = "ACTIVE"
	StatusSuspended Status = "SUSPENDED"
	StatusArchived  Status = "ARCHIVED"
)

func (s Status) String() string {
	return string(s)
}

func (s Status) IsValid() bool {
	switch s {
	case StatusDraft, StatusActive, StatusSuspended, StatusArchived:
		return true
	default:
		return false
	}
}
