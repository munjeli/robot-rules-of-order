package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type Transcript struct {
	Base
	Meeting    string
	Public     bool
	Contents   []Action
	Start      timestamp.Timestamp
	End        timestamp.Timestamp
	Archived   timestamp.Timestamp
	ArchiveURL string
}
