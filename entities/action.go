package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type Action struct {
	Base
	Timestamp timestamp.Timestamp
	// privileged, subsidiary, statement, robot
	Type string
	// Citizen name or robot
	Byline string
}
