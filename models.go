package main

import (
	"time"
)

// AboAnfrage

type AboAUS struct {
	AboID      string    `xml:"AboID"`
	VerfallZst time.Time `xml:"VerfallZst"`
}

type AboAnfrage struct {
	Sender string    `xml:"Sender,attr"`
	Zst    time.Time `xml:"Zst,attr"`
	AboAUS []AboAUS  `xml:"AboAUS"`
}

// AboBestaetigung

type Bestaetigung struct {
	Zst          time.Time `xml:"Zst,attr"`
	Ergebnis     string    `xml:"Ergebnis"`
	Fehlernummer string    `xml:"Fehlernummer"`
}

type BestaetigungMitAboID struct {
	Bestaetigung []Bestaetigung `xml:"Bestaetigung"`
	AboID        string         `xml:"AboID"`
}

type AboAntwort struct {
	Bestaetigung         []Bestaetigung         `xml:"Bestaetigung"`
	BestaetigungMitAboID []BestaetigungMitAboID `xml:"BestaetigungMitAboID"`
}

// ClientStatusAnfrage

type ClientStatusAnfrage struct {
	Sender  string    `xml:"Sender,attr"`
	Zst     time.Time `xml:"Zst,attr"`
	MitAbos bool      `xml:"MitAbos,attr"`
}

// ClientStatusAntwort

type Status struct {
	Zst      time.Time `xml:"Zst,attr"`
	Ergebnis string    `xml:"Ergebnis"`
}

type ClientStatusAntwort struct {
	Status         string    `xml:"Status"`
	StartDienstZst time.Time `xml:"StartDienstZst,attr"`
}
