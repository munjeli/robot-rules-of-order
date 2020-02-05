package entity

import "github.com/golang/protobuf/ptypes/timestamp"

type Meeting struct {
	Base
	Date     string
	Public   bool
	Invitees []Citizen
	Host     Citizen
	Status   Status
	// transcript self url
	Transcript string
	// invitation self url
	Invitation string
	Quorum     int
	// options are majInvitee, majAttendees, hardLimit
	QuorumType string
}

type Status struct {
	Timestamp timestamp.Timestamp
	State     string
	Latest    Action
	// Citizen names
	Attendees []string
	// Citizen on the floor
	Floor string
}

func CreateMeeting()  {}
func DeleteMeeting()  {}
func RunMeeting()     {}
func UpdateMeeting()  {}
func ArchiveMeeting() {}
