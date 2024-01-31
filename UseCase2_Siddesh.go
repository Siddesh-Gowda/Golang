package main

import (
	"fmt"
	"strings"
)

type NSPEvent struct {
	EventDate      string
	EventName      string
	PrimaryContact string
	Organizingteam OrganizingTeam
}

type OrganizingTeam struct {
	TeamMembers    []string
	PrimaryContact string
}

func main() {

	var eventdate string
	var eventname string
	var contactname string
	var temaMember string
	var temaMembers []string
	var primaryContact string
	var AddEvent []NSPEvent

	fmt.Println("Enter the NSP Event to be organised")
	fmt.Println("Enter the NSP Event date")
	fmt.Scanln(&eventdate)
	fmt.Println("Enter the NSP Event Name")
	fmt.Scanln(&eventname)
	fmt.Println("Enter the primary contact if anyone from outside the organization")
	fmt.Scanln(&contactname)
	fmt.Println("Enter the organising team members name")
	fmt.Scanln(&temaMember)
	fmt.Println("Enter the Primary contact from organising team")
	fmt.Scanln(&primaryContact)

	temaMembers = strings.Split(temaMember, ",")
	var NSPEvent1 NSPEvent
	NSPEvent1.EventDate = eventdate
	NSPEvent1.EventName = eventname
	NSPEvent1.PrimaryContact = contactname
	NSPEvent1.Organizingteam.TeamMembers = temaMembers
	NSPEvent1.Organizingteam.PrimaryContact = primaryContact

	AddEvent = append(AddEvent, NSPEvent1)

	fmt.Println(AddEvent)

	// Team1 := OrganizingTeam{TeamMembers: []string{"soundarya", "abhishek"}, PrimaryContact: "soundarya"}
	// Team2 := OrganizingTeam{TeamMembers: []string{"Kishore", "achyut"}, PrimaryContact: "Kishore"}

	// //Events iniatilization
	// Event1 := NSPEvent{"jan-23", "NSP-sankranti", Team1.PrimaryContact, Team1}
	// Event2 := NSPEvent{"jan-1", "NSP-new year party", "John", Team2}

	// fmt.Println("Event 1 details are :", Event1.EventDate, Event1.EventName, Event1.PrimaryContact, Event1.Organizingteam.TeamMembers)
	// fmt.Println("Event 2 details are :", Event2.EventDate, Event2.EventName, Event2.PrimaryContact, Event2.Organizingteam.TeamMembers)

}
