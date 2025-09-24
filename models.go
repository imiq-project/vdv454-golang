package main

import "time"

type AbbringerFahrtLoeschen struct {
	Protokolleintrag []Protokolleintrag  `xml:"Protokolleintrag"`
	ASBID            []ID                `xml:"ASBID"`
	AbbringerInfo    []AbbringerInfo     `xml:"AbbringerInfo"`
	FahrtIDExt       []FahrtIDExt        `xml:"FahrtIDExt"`
	UrsacheText      []TextMitSprachCode `xml:"UrsacheText"`
	Zst              time.Time           `xml:"Zst,attr"`
}
type AbbringerInfo struct {
	FahrtID                 []FahrtID                    `xml:"FahrtID"`
	HstSeqZaehler           []HstSeqZaehler              `xml:"HstSeqZaehler"`
	LinienID                []LinienID                   `xml:"LinienID"`
	LinienText              []TextMitSprachCode          `xml:"LinienText"`
	RichtungsID             []RichtungsID                `xml:"RichtungsID"`
	RichtungsText           []TextMitSprachCode          `xml:"RichtungsText"`
	VonRichtungsText        []TextMitSprachCodeMinLenght `xml:"VonRichtungsText"`
	AbfahrtszeitASBPlan     []time.Time                  `xml:"AbfahrtszeitASBPlan"`
	HaltID                  []HaltID                     `xml:"HaltID"`
	AbfahrtssteigText       []TextMitSprachCodeMinLenght `xml:"AbfahrtssteigText"`
	AbfahrtsSektorenText    []TextMitSprachCodeMinLenght `xml:"AbfahrtsSektorenText"`
	AbfahrtsLinienfahrwegID []LinienfahrwegID            `xml:"AbfahrtsLinienfahrwegID"`
	FahrtInfo               []FahrtInfo                  `xml:"FahrtInfo"`
}
type Abbringernachricht struct {
	AboID AboID `xml:"AboID,attr"`
}
type AbfahrtAusfallGroup struct {
	AbfahrtFaelltAus            []bool              `xml:"AbfahrtFaelltAus"`
	AbfahrtFaelltAusUrsacheText []TextMitSprachCode `xml:"AbfahrtFaelltAusUrsacheText"`
}
type AbfahrtAusGroup struct {
}
type AbfahrtDfiGroup struct {
}
type AbfahrtEchtzeitGroup struct {
	IstAbfahrtPrognose          []time.Time       `xml:"IstAbfahrtPrognose"`
	IstAbfahrtDisposition       []time.Time       `xml:"IstAbfahrtDisposition"`
	IstAbfahrtPrognoseUngenau   []PrognoseUngenau `xml:"IstAbfahrtPrognoseUngenau"`
	IstAbfahrtPrognoseQualitaet []ZeitQualitaet   `xml:"IstAbfahrtPrognoseQualitaet"`
}
type AbfahrtSollGroup struct {
	Abfahrtszeit            []time.Time         `xml:"Abfahrtszeit"`
	AbfahrtssteigText       []TextMitSprachCode `xml:"AbfahrtssteigText"`
	AbfahrtsSektorenText    []TextMitSprachCode `xml:"AbfahrtsSektorenText"`
	Einsteigeverbot         []bool              `xml:"Einsteigeverbot"`
	RichtungsText           []TextMitSprachCode `xml:"RichtungsText"`
	Via                     []Via               `xml:"Via"`
	AbfahrtsLinienfahrwegID []LinienfahrwegID   `xml:"AbfahrtsLinienfahrwegID"`
}
type AbfahrtSollMitAusfallGroup struct {
}
type AbfahrtStatusGroup struct {
	IstAbfahrtPrognoseStatus []PrognoseStatus   `xml:"IstAbfahrtPrognoseStatus"`
	Auslastungsstufe         []Auslastungsstufe `xml:"Auslastungsstufe"`
}
type AboAND struct {
	KanalID    []KanalID `xml:"KanalID"`
	AboID      AboID     `xml:"AboID,attr"`
	VerfallZst time.Time `xml:"VerfallZst,attr"`
}
type AboAnfrage struct {
	AboLoeschen     []AboID   `xml:"AboLoeschen"`
	AboLoeschenAlle []bool    `xml:"AboLoeschenAlle"`
	Sender          Sender    `xml:"Sender,attr"`
	Zst             time.Time `xml:"Zst,attr"`
	XSDVersionID    string    `xml:"XSDVersionID,attr"`
}
type AboAntwort struct {
	Bestaetigung []Bestaetigung `xml:"Bestaetigung"`
	XSDVersionID string         `xml:"XSDVersionID,attr"`
}
type AboASBRef struct {
	ASBID                 []ID           `xml:"ASBID"`
	LinienFilter          []LinienFilter `xml:"LinienFilter"`
	FruehesteAnkunftszeit []time.Time    `xml:"FruehesteAnkunftszeit"`
	SpaetesteAnkunftszeit []time.Time    `xml:"SpaetesteAnkunftszeit"`
	AboID                 AboID          `xml:"AboID,attr"`
	VerfallZst            time.Time      `xml:"VerfallZst,attr"`
}
type AboASB struct {
	ASBID         []ID            `xml:"ASBID"`
	Hysterese     []Hysterese     `xml:"Hysterese"`
	AbbringerInfo []AbbringerInfo `xml:"AbbringerInfo"`
	AboID         AboID           `xml:"AboID,attr"`
	VerfallZst    time.Time       `xml:"VerfallZst,attr"`
}
type AboAUSRef struct {
	Zeitfenster                  []Zeitfenster            `xml:"Zeitfenster"`
	LinienFilter                 []LinienFilter           `xml:"LinienFilter"`
	BetreiberFilter              []BetreiberFilter        `xml:"BetreiberFilter"`
	ProduktFilter                []ProduktFilter          `xml:"ProduktFilter"`
	VerkehrsmittelIDFilter       []VerkehrsmittelIDFilter `xml:"VerkehrsmittelIDFilter"`
	HaltFilter                   []HaltFilter             `xml:"HaltFilter"`
	MitGesAnschluss              []bool                   `xml:"MitGesAnschluss"`
	MitZusaetzlichenZeitfenstern []bool                   `xml:"MitZusaetzlichenZeitfenstern"`
	MitFormation                 []bool                   `xml:"MitFormation"`
	AboID                        AboID                    `xml:"AboID,attr"`
	VerfallZst                   time.Time                `xml:"VerfallZst,attr"`
}
type AboAUS struct {
	LinienFilter           []LinienFilter           `xml:"LinienFilter"`
	BetreiberFilter        []BetreiberFilter        `xml:"BetreiberFilter"`
	ProduktFilter          []ProduktFilter          `xml:"ProduktFilter"`
	VerkehrsmittelIDFilter []VerkehrsmittelIDFilter `xml:"VerkehrsmittelIDFilter"`
	HaltFilter             []HaltFilter             `xml:"HaltFilter"`
	Hysterese              []Hysterese              `xml:"Hysterese"`
	Vorschauzeit           []Vorschauzeit           `xml:"Vorschauzeit"`
	MitGesAnschluss        []bool                   `xml:"MitGesAnschluss"`
	MitRealZeiten          []bool                   `xml:"MitRealZeiten"`
	MitFormation           []bool                   `xml:"MitFormation"`
	NurAktualisierung      []bool                   `xml:"NurAktualisierung"`
	AboID                  AboID                    `xml:"AboID,attr"`
	VerfallZst             time.Time                `xml:"VerfallZst,attr"`
}
type AboAZBRef struct {
	AZBID                 []ID           `xml:"AZBID"`
	LinienFilter          []LinienFilter `xml:"LinienFilter"`
	FruehesteAbfahrtszeit []time.Time    `xml:"FruehesteAbfahrtszeit"`
	SpaetesteAbfahrtszeit []time.Time    `xml:"SpaetesteAbfahrtszeit"`
	AboID                 AboID          `xml:"AboID,attr"`
	VerfallZst            time.Time      `xml:"VerfallZst,attr"`
}
type AboAZB struct {
	AZBID             []ID            `xml:"AZBID"`
	LinienFilter      []LinienFilter  `xml:"LinienFilter"`
	Vorschauzeit      []Vorschauzeit  `xml:"Vorschauzeit"`
	MaxAnzahlFahrten  []AnzahlFahrten `xml:"MaxAnzahlFahrten"`
	Hysterese         []Hysterese     `xml:"Hysterese"`
	MaxTextLaenge     []MaxTextLaenge `xml:"MaxTextLaenge"`
	NurAktualisierung []bool          `xml:"NurAktualisierung"`
	AboID             AboID           `xml:"AboID,attr"`
	VerfallZst        time.Time       `xml:"VerfallZst,attr"`
}
type AboGroup struct {
}
type AboVIS struct {
	VISID        []ID           `xml:"VISID"`
	LinienFilter []LinienFilter `xml:"LinienFilter"`
	AboID        AboID          `xml:"AboID,attr"`
	VerfallZst   time.Time      `xml:"VerfallZst,attr"`
}
type AktiveAbos struct {
}
type ANDMeldung struct {
	Protokolleintrag []Protokolleintrag `xml:"Protokolleintrag"`
	ANDMeldungsart   []string           `xml:"ANDMeldungsart"`
	KanalID          []KanalID          `xml:"KanalID"`
	MeldungsID       []MeldungsID       `xml:"MeldungsID"`
	FormatID         []FormatID         `xml:"FormatID"`
	Zst              time.Time          `xml:"Zst,attr"`
	VerfallZst       time.Time          `xml:"VerfallZst,attr"`
}
type ANDNachricht struct {
	ANDMeldung []ANDMeldung `xml:"ANDMeldung"`
	AboID      AboID        `xml:"AboID,attr"`
}
type AnkunftAusfallGroup struct {
	AnkunftFaelltAus            []bool              `xml:"AnkunftFaelltAus"`
	AnkunftFaelltAusUrsacheText []TextMitSprachCode `xml:"AnkunftFaelltAusUrsacheText"`
}
type AnkunftAusGroup struct {
}
type AnkunftDfiGroup struct {
}
type AnkunftEchtzeitGroup struct {
	IstAnkunftPrognose          []time.Time       `xml:"IstAnkunftPrognose"`
	IstAnkunftDisposition       []time.Time       `xml:"IstAnkunftDisposition"`
	IstAnkunftPrognoseUngenau   []PrognoseUngenau `xml:"IstAnkunftPrognoseUngenau"`
	IstAnkunftPrognoseQualitaet []ZeitQualitaet   `xml:"IstAnkunftPrognoseQualitaet"`
}
type AnkunftSollGroup struct {
	AnkunftHaltID           []HaltID            `xml:"AnkunftHaltID"`
	Ankunftszeit            []time.Time         `xml:"Ankunftszeit"`
	AnkunftssteigText       []TextMitSprachCode `xml:"AnkunftssteigText"`
	AnkunftsSektorenText    []TextMitSprachCode `xml:"AnkunftsSektorenText"`
	Aussteigeverbot         []bool              `xml:"Aussteigeverbot"`
	VonRichtungsText        []TextMitSprachCode `xml:"VonRichtungsText"`
	AnkunftsLinienfahrwegID []LinienfahrwegID   `xml:"AnkunftsLinienfahrwegID"`
}
type AnkunftSollMitAusfallGroup struct {
}
type AnkunftStatusGroup struct {
	IstAnkunftPrognoseStatus []PrognoseStatus `xml:"IstAnkunftPrognoseStatus"`
}
type AnschlussPlan struct {
	Protokolleintrag      []Protokolleintrag `xml:"Protokolleintrag"`
	Zubringer             []FahrtIDAusGlobal `xml:"Zubringer"`
	HaltIDZubringer       []HaltID           `xml:"HaltIDZubringer"`
	Abbringer             []FahrtIDAusGlobal `xml:"Abbringer"`
	HaltIDAbbringer       []HaltID           `xml:"HaltIDAbbringer"`
	Umsteigewegezeit      []int              `xml:"Umsteigewegezeit"`
	MaxAutoVerzoegerung   []int              `xml:"MaxAutoVerzoegerung"`
	Prioritaet            []Prioritaet       `xml:"Prioritaet"`
	AnkunftszeitZubringer []time.Time        `xml:"AnkunftszeitZubringer"`
	AbfahrtszeitAbbringer []time.Time        `xml:"AbfahrtszeitAbbringer"`
	AnschlussID           AnschlussID        `xml:"AnschlussID,attr"`
}
type AnschlussStatus struct {
	Protokolleintrag              []Protokolleintrag `xml:"Protokolleintrag"`
	WarteInfo                     []WarteInfo        `xml:"WarteInfo"`
	AbfahrtszeitAbbringerPrognose []time.Time        `xml:"AbfahrtszeitAbbringerPrognose"`
	SicherungAufgehoben           []bool             `xml:"SicherungAufgehoben"`
	AnschlussID                   AnschlussID        `xml:"AnschlussID,attr"`
}
type ASBFahrplanlage struct {
	Protokolleintrag        []Protokolleintrag  `xml:"Protokolleintrag"`
	ASBZubringerMeldungsart []string            `xml:"ASBZubringerMeldungsart"`
	ASBID                   []ID                `xml:"ASBID"`
	FahrtID                 []FahrtID           `xml:"FahrtID"`
	HstSeqZaehler           []HstSeqZaehler     `xml:"HstSeqZaehler"`
	LinienID                []LinienID          `xml:"LinienID"`
	LinienText              []TextMitSprachCode `xml:"LinienText"`
	RichtungsID             []RichtungsID       `xml:"RichtungsID"`
	RichtungsText           []TextMitSprachCode `xml:"RichtungsText"`
	VonRichtungsText        []TextMitSprachCode `xml:"VonRichtungsText"`
	AnkunftszeitASBPlan     []time.Time         `xml:"AnkunftszeitASBPlan"`
	AnkunftszeitASBPrognose []time.Time         `xml:"AnkunftszeitASBPrognose"`
	PrognoseMoeglich        []bool              `xml:"PrognoseMoeglich"`
	UmsteigeWillige         []UmsteigeWillige   `xml:"UmsteigeWillige"`
	ZubringerHstnameLang    []TextMitSprachCode `xml:"ZubringerHstnameLang"`
	SpaetesteAbbringerInfo  []time.Time         `xml:"SpaetesteAbbringerInfo"`
	AnkunftHaltID           []HaltID            `xml:"AnkunftHaltID"`
	AnkunftssteigText       []TextMitSprachCode `xml:"AnkunftssteigText"`
	AnkunftsSektorenText    []TextMitSprachCode `xml:"AnkunftsSektorenText"`
	PrognoseUngenau         []PrognoseUngenau   `xml:"PrognoseUngenau"`
	AnkunftsLinienfahrwegID []LinienfahrwegID   `xml:"AnkunftsLinienfahrwegID"`
	FaelltAusUrsacheText    []TextMitSprachCode `xml:"FaelltAusUrsacheText"`
	FahrtInfo               []FahrtInfo         `xml:"FahrtInfo"`
	Zst                     time.Time           `xml:"Zst,attr"`
	VerfallZst              time.Time           `xml:"VerfallZst,attr"`
}
type ASBLinienFahrplan struct {
	Protokolleintrag []Protokolleintrag `xml:"Protokolleintrag"`
	ASBID            []ID               `xml:"ASBID"`
	LinienID         []LinienID         `xml:"LinienID"`
	RichtungsID      []RichtungsID      `xml:"RichtungsID"`
	BetreiberID      []Betreiber        `xml:"BetreiberID"`
	Zeitfenster      []Zeitfenster      `xml:"Zeitfenster"`
	ASBSollFahrt     []ASBSollFahrt     `xml:"ASBSollFahrt"`
}
type ASBSollFahrt struct {
	FahrtID                 []FahrtID           `xml:"FahrtID"`
	HstSeqZaehler           []HstSeqZaehler     `xml:"HstSeqZaehler"`
	LinienText              []TextMitSprachCode `xml:"LinienText"`
	RichtungsText           []TextMitSprachCode `xml:"RichtungsText"`
	VonRichtungsText        []TextMitSprachCode `xml:"VonRichtungsText"`
	AnkunftszeitASBPlan     []time.Time         `xml:"AnkunftszeitASBPlan"`
	AnkunftHaltID           []HaltID            `xml:"AnkunftHaltID"`
	AnkunftssteigText       []TextMitSprachCode `xml:"AnkunftssteigText"`
	AnkunftsSektorenText    []TextMitSprachCode `xml:"AnkunftsSektorenText"`
	AnkunftsLinienfahrwegID []LinienfahrwegID   `xml:"AnkunftsLinienfahrwegID"`
	FahrtInfo               []FahrtInfo         `xml:"FahrtInfo"`
	Zst                     time.Time           `xml:"Zst,attr"`
}
type AUSNachricht struct {
	FahrtVerband []FahrtVerband `xml:"FahrtVerband"`
	AboID        AboID          `xml:"AboID,attr"`
}
type AZBFahrplanlage struct {
	Protokolleintrag           []Protokolleintrag  `xml:"Protokolleintrag"`
	AZBID                      []ID                `xml:"AZBID"`
	FahrtID                    []FahrtID           `xml:"FahrtID"`
	HstSeqZaehler              []HstSeqZaehler     `xml:"HstSeqZaehler"`
	Traktion                   []Traktions         `xml:"Traktion"`
	BetrieblicheFahrzeugnummer []string            `xml:"BetrieblicheFahrzeugnummer"`
	LinienID                   []LinienID          `xml:"LinienID"`
	LinienText                 []TextMitSprachCode `xml:"LinienText"`
	FahrtBezeichnerText        []TextMitSprachCode `xml:"FahrtBezeichnerText"`
	RichtungsID                []RichtungsID       `xml:"RichtungsID"`
	AbmeldeID                  []AbmeldeID         `xml:"AbmeldeID"`
	ZielHstnameKurz            []TextMitSprachCode `xml:"ZielHstnameKurz"`
	AufAZB                     []bool              `xml:"AufAZB"`
	PrognoseMoeglich           []bool              `xml:"PrognoseMoeglich"`
	Fahrtspezialtext           []TextMitSprachCode `xml:"Fahrtspezialtext"`
	FaelltAusUrsacheText       []TextMitSprachCode `xml:"FaelltAusUrsacheText"`
	FahrtInfo                  []FahrtInfo         `xml:"FahrtInfo"`
	Zst                        time.Time           `xml:"Zst,attr"`
	VerfallZst                 time.Time           `xml:"VerfallZst,attr"`
}
type AZBLinienFahrplan struct {
	Protokolleintrag []Protokolleintrag `xml:"Protokolleintrag"`
	AZBID            []ID               `xml:"AZBID"`
	LinienID         []LinienID         `xml:"LinienID"`
	RichtungsID      []RichtungsID      `xml:"RichtungsID"`
	BetreiberID      []Betreiber        `xml:"BetreiberID"`
	Zeitfenster      []Zeitfenster      `xml:"Zeitfenster"`
	AZBSollFahrt     []AZBSollFahrt     `xml:"AZBSollFahrt"`
}
type AZBLinienspezialtext struct {
	Protokolleintrag             []Protokolleintrag  `xml:"Protokolleintrag"`
	LinienspezialtextMeldungsart []string            `xml:"LinienspezialtextMeldungsart"`
	AZBID                        []ID                `xml:"AZBID"`
	LinienID                     []LinienID          `xml:"LinienID"`
	LinienText                   []TextMitSprachCode `xml:"LinienText"`
	RichtungsID                  []RichtungsID       `xml:"RichtungsID"`
	Linienspezialtext            []TextMitSprachCode `xml:"Linienspezialtext"`
	Prioritaet                   []Prioritaet        `xml:"Prioritaet"`
	Zst                          time.Time           `xml:"Zst,attr"`
	VerfallZst                   time.Time           `xml:"VerfallZst,attr"`
}
type AZBNachricht struct {
	AZBLinienFahrplan    []AZBLinienFahrplan    `xml:"AZBLinienFahrplan"`
	AZBFahrplanlage      []AZBFahrplanlage      `xml:"AZBFahrplanlage"`
	AZBLinienspezialtext []AZBLinienspezialtext `xml:"AZBLinienspezialtext"`
	AZBSondertext        []AZBSondertext        `xml:"AZBSondertext"`
	AboID                AboID                  `xml:"AboID,attr"`
}
type AZBSollFahrt struct {
	FahrtID             []FahrtID           `xml:"FahrtID"`
	HstSeqZaehler       []HstSeqZaehler     `xml:"HstSeqZaehler"`
	LinienText          []TextMitSprachCode `xml:"LinienText"`
	FahrtBezeichnerText []TextMitSprachCode `xml:"FahrtBezeichnerText"`
	FahrtInfo           []FahrtInfo         `xml:"FahrtInfo"`
	Zst                 time.Time           `xml:"Zst,attr"`
}
type AZBSondertext struct {
	Protokolleintrag      []Protokolleintrag  `xml:"Protokolleintrag"`
	SondertextMeldungsart []string            `xml:"SondertextMeldungsart"`
	AZBID                 []ID                `xml:"AZBID"`
	MeldungsID            []MeldungsID        `xml:"MeldungsID"`
	Sondertext            []TextMitSprachCode `xml:"Sondertext"`
	Prioritaet            []Prioritaet        `xml:"Prioritaet"`
	Zst                   time.Time           `xml:"Zst,attr"`
	VerfallZst            time.Time           `xml:"VerfallZst,attr"`
}
type Bestaetigung struct {
	Fehlertext   []Fehlertext `xml:"Fehlertext"`
	Zst          time.Time    `xml:"Zst,attr"`
	Ergebnis     Ergebnis     `xml:"Ergebnis,attr"`
	Fehlernummer Fehlernummer `xml:"Fehlernummer,attr"`
}
type BetreiberFilter struct {
	BetreiberID []Betreiber `xml:"BetreiberID"`
}
type BetrieblicheFahrtnummern struct {
	BetrieblicheFahrtnummer []BetrieblicheFahrtnummer `xml:"BetrieblicheFahrtnummer"`
}
type BetrieblicheFahrtnummer struct {
	BetrieblicheFahrtnummerBezeichner     []BetrieblicheFahrtnummerBezeichner `xml:"BetrieblicheFahrtnummerBezeichner"`
	BetrieblicheFahrtnummerFahrtAbschnitt []FahrtStartEnde                    `xml:"BetrieblicheFahrtnummerFahrtAbschnitt"`
}
type BetrieblicheFahrtnummerBezeichner = string
type Beziehung struct {
	BeziehungsTyp    []string           `xml:"BeziehungsTyp"`
	StreckenBezug    []StreckenBezug    `xml:"StreckenBezug"`
	BeziehungZuFahrt []BeziehungZuFahrt `xml:"BeziehungZuFahrt"`
}
type BeziehungZuFahrt struct {
	FahrtBezug    []FahrtIDAusGlobal `xml:"FahrtBezug"`
	StreckenBezug []StreckenBezug    `xml:"StreckenBezug"`
}
type ClientStatusAnfrage struct {
	StartDienstZst []time.Time `xml:"StartDienstZst"`
	DatenVersionID []string    `xml:"DatenVersionID"`
	Sender         Sender      `xml:"Sender,attr"`
	Zst            time.Time   `xml:"Zst,attr"`
	MitAbos        bool        `xml:"MitAbos,attr"`
}
type ClientStatusAntwort struct {
	Status         []Status     `xml:"Status"`
	StartDienstZst []time.Time  `xml:"StartDienstZst"`
	AktiveAbos     []AktiveAbos `xml:"AktiveAbos"`
}
type DatenAbrufenAnfrage struct {
	DatensatzAlle []bool    `xml:"DatensatzAlle"`
	Sender        Sender    `xml:"Sender,attr"`
	Zst           time.Time `xml:"Zst,attr"`
}
type DatenAbrufenAntwort struct {
	Bestaetigung []Bestaetigung `xml:"Bestaetigung"`
	WeitereDaten []bool         `xml:"WeitereDaten"`
}
type DatenBereitAnfrage struct {
	Sender Sender    `xml:"Sender,attr"`
	Zst    time.Time `xml:"Zst,attr"`
}
type DatenBereitAntwort struct {
	Bestaetigung []Bestaetigung `xml:"Bestaetigung"`
}
type Direktruf struct {
	Telefonnummer []Telefonnummer `xml:"Telefonnummer"`
	IPAdresse     []IPAdresse     `xml:"IPAdresse"`
}
type FahrtFilter struct {
	FahrtID             []FahrtID       `xml:"FahrtID"`
	HstSeqZaehler       []HstSeqZaehler `xml:"HstSeqZaehler"`
	AnkunftszeitASBPlan []time.Time     `xml:"AnkunftszeitASBPlan"`
	Vorschauzeit        []Vorschauzeit  `xml:"Vorschauzeit"`
}
type FahrtIDAusGlobal struct {
	FahrtRef      []FahrtRef      `xml:"FahrtRef"`
	LinienID      []LinienID      `xml:"LinienID"`
	LeitstellenID []LeitstellenID `xml:"LeitstellenID"`
}
type FahrtIDExt struct {
	FahrtBezeichner []FahrtBezeichner `xml:"FahrtBezeichner"`
	Betriebstag     []time.Time       `xml:"Betriebstag"`
	HstSeqZaehler   []HstSeqZaehler   `xml:"HstSeqZaehler"`
}
type FahrtIDGlobal struct {
	FahrtIDExt    []FahrtIDExt    `xml:"FahrtIDExt"`
	LinienID      []LinienID      `xml:"LinienID"`
	LeitstellenID []LeitstellenID `xml:"LeitstellenID"`
}
type FahrtID struct {
	FahrtBezeichner []FahrtBezeichner `xml:"FahrtBezeichner"`
	Betriebstag     []time.Time       `xml:"Betriebstag"`
}
type FahrtInfo struct {
	FahrzeugID           []FahrzeugID                 `xml:"FahrzeugID"`
	LinienNr             []LinienNr                   `xml:"LinienNr"`
	VerkehrsmittelNummer []string                     `xml:"VerkehrsmittelNummer"`
	UmlaufNr             []UmlaufNr                   `xml:"UmlaufNr"`
	KursNr               []KursNr                     `xml:"KursNr"`
	StartHstnameLang     []TextMitSprachCodeMinLenght `xml:"StartHstnameLang"`
	StartHstnameKurz     []TextMitSprachCodeMinLenght `xml:"StartHstnameKurz"`
	ZielHstnameLang      []TextMitSprachCodeMinLenght `xml:"ZielHstnameLang"`
	ZielHstnameKurz      []TextMitSprachCodeMinLenght `xml:"ZielHstnameKurz"`
	LinienfahrwegID      []LinienfahrwegID            `xml:"LinienfahrwegID"`
	AbfahrtszeitStartHst []time.Time                  `xml:"AbfahrtszeitStartHst"`
	AnkunftszeitZielHst  []time.Time                  `xml:"AnkunftszeitZielHst"`
	ProduktID            []ProduktID                  `xml:"ProduktID"`
	BetreiberID          []Betreiber                  `xml:"BetreiberID"`
	AusfuehrenderID      []Betreiber                  `xml:"AusfuehrenderID"`
	Direktruf            []Direktruf                  `xml:"Direktruf"`
	VerkehrsmittelID     []VerkehrsmittelID           `xml:"VerkehrsmittelID"`
	ServiceAttribut      []ServiceAttribut            `xml:"ServiceAttribut"`
}
type FahrtRef struct {
	FahrtID        []FahrtID        `xml:"FahrtID"`
	FahrtStartEnde []FahrtStartEnde `xml:"FahrtStartEnde"`
}
type FahrtStartEnde struct {
	StartHaltID []HaltID    `xml:"StartHaltID"`
	Startzeit   []time.Time `xml:"Startzeit"`
	EndHaltID   []HaltID    `xml:"EndHaltID"`
	Endzeit     []time.Time `xml:"Endzeit"`
}
type GesAnschluss struct {
}
type HaltAusGroup struct {
}
type HaltGroup struct {
}
type HaltepositionsAenderung struct {
	Protokolleintrag []Protokolleintrag `xml:"Protokolleintrag"`
	ASBID            []ID               `xml:"ASBID"`
	AbbringerInfo    []AbbringerInfo    `xml:"AbbringerInfo"`
	FahrtIDExt       []FahrtIDExt       `xml:"FahrtIDExt"`
	Zst              time.Time          `xml:"Zst,attr"`
}
type HaltExtended struct {
	HaltID       []HaltID    `xml:"HaltID"`
	Ankunftszeit []time.Time `xml:"Ankunftszeit"`
	Abfahrtszeit []time.Time `xml:"Abfahrtszeit"`
}
type HaltFilter struct {
	HaltID []HaltID `xml:"HaltID"`
}
type HaltInfoDfiGroup struct {
}
type HaltInfoSollGroup struct {
}
type HaltInfoSollMitAusfallGroup struct {
}
type HaltIstGroup struct {
	Zusatzhalt []bool `xml:"Zusatzhalt"`
}
type HaltSollGroup struct {
	HaltID                []HaltID            `xml:"HaltID"`
	HaltestellenName      []TextMitSprachCode `xml:"HaltestellenName"`
	Durchfahrt            []bool              `xml:"Durchfahrt"`
	FahrtHaltspezialText  []TextMitSprachCode `xml:"FahrtHaltspezialText"`
	FahrtSteigspezialText []TextMitSprachCode `xml:"FahrtSteigspezialText"`
	ServiceAttribut       []ServiceAttribut   `xml:"ServiceAttribut"`
	StoerungsInfo         []StoerungsInfo     `xml:"StoerungsInfo"`
}
type IstFahrt struct {
	Protokolleintrag         []Protokolleintrag         `xml:"Protokolleintrag"`
	LinienID                 []LinienID                 `xml:"LinienID"`
	RichtungsID              []RichtungsID              `xml:"RichtungsID"`
	FahrtRef                 []FahrtRef                 `xml:"FahrtRef"`
	BetrieblicheFahrtnummern []BetrieblicheFahrtnummern `xml:"BetrieblicheFahrtnummern"`
	Komplettfahrt            []bool                     `xml:"Komplettfahrt"`
	UmlaufID                 []UmlaufID                 `xml:"UmlaufID"`
	KursNr                   []KursNr                   `xml:"KursNr"`
	BetreiberID              []Betreiber                `xml:"BetreiberID"`
	AusfuehrenderID          []Betreiber                `xml:"AusfuehrenderID"`
	OriginalSollFahrtverlauf []OriginalFahrtverlauf     `xml:"OriginalSollFahrtverlauf"`
	DispoID                  []string                   `xml:"DispoID"`
	IstHalt                  []IstHalt                  `xml:"IstHalt"`
	FahrtBezeichnerText      []TextMitSprachCode        `xml:"FahrtBezeichnerText"`
	VerkehrsmittelNummer     []string                   `xml:"VerkehrsmittelNummer"`
	LinienText               []TextMitSprachCode        `xml:"LinienText"`
	ProduktID                []ProduktID                `xml:"ProduktID"`
	VonRichtungsText         []TextMitSprachCode        `xml:"VonRichtungsText"`
	FahrtspezialText         []TextMitSprachCode        `xml:"FahrtspezialText"`
	LinienfahrwegID          []LinienfahrwegID          `xml:"LinienfahrwegID"`
	Zugname                  []TextMitSprachCode        `xml:"Zugname"`
	VerkehrsmittelID         []VerkehrsmittelID         `xml:"VerkehrsmittelID"`
	PrognoseMoeglich         []bool                     `xml:"PrognoseMoeglich"`
	PrognoseUngenau          []PrognoseUngenau          `xml:"PrognoseUngenau"`
	Zusatzfahrt              []bool                     `xml:"Zusatzfahrt"`
	FaelltAus                []bool                     `xml:"FaelltAus"`
	FahrtZuruecksetzen       []bool                     `xml:"FahrtZuruecksetzen"`
	StoerungsInfo            []StoerungsInfo            `xml:"StoerungsInfo"`
	FahrradMitnahme          []bool                     `xml:"FahrradMitnahme"`
	FahrzeugTypID            []string                   `xml:"FahrzeugTypID"`
	Auslastungsstufe         []Auslastungsstufe         `xml:"Auslastungsstufe"`
	ServiceAttribut          []ServiceAttribut          `xml:"ServiceAttribut"`
	IstFormation             []Formation                `xml:"IstFormation"`
	FahrtBeziehung           []Beziehung                `xml:"FahrtBeziehung"`
	Zst                      time.Time                  `xml:"Zst,attr"`
}
type IstHalt struct {
}
type LinienFahrplan struct {
	Protokolleintrag []Protokolleintrag `xml:"Protokolleintrag"`
	LinienID         []LinienID         `xml:"LinienID"`
	RichtungsID      []RichtungsID      `xml:"RichtungsID"`
	BetreiberID      []Betreiber        `xml:"BetreiberID"`
	AusfuehrenderID  []Betreiber        `xml:"AusfuehrenderID"`
}
type LinienFilter struct {
	LinienID    []LinienID    `xml:"LinienID"`
	RichtungsID []RichtungsID `xml:"RichtungsID"`
}
type OriginalFahrtverlauf struct {
	OriginalSollHalt []OriginalSollHalt `xml:"OriginalSollHalt"`
}
type OriginalSollHalt struct {
}
type ProduktFilter struct {
	ProduktID []ProduktID `xml:"ProduktID"`
}
type Protokolleintrag struct {
	Zst           time.Time `xml:"Zst,attr"`
	Systemkennung string    `xml:"Systemkennung,attr"`
	Aktion        string    `xml:"Aktion,attr"`
	Details       string    `xml:"Details,attr"`
}
type ServiceAttribut struct {
	Name []string `xml:"Name"`
	Wert []string `xml:"Wert"`
}
type SollAnschluss struct {
	FahrtID          []FahrtID `xml:"FahrtID"`
	HaltID           []HaltID  `xml:"HaltID"`
	Umsteigewegezeit []int     `xml:"Umsteigewegezeit"`
	Sitzenbleiben    []bool    `xml:"Sitzenbleiben"`
}
type SollFahrt struct {
	FahrtID                  []FahrtID                  `xml:"FahrtID"`
	OriginalSollFahrtverlauf []OriginalFahrtverlauf     `xml:"OriginalSollFahrtverlauf"`
	BetrieblicheFahrtnummern []BetrieblicheFahrtnummern `xml:"BetrieblicheFahrtnummern"`
	AusfuehrenderID          []Betreiber                `xml:"AusfuehrenderID"`
	DispoID                  []string                   `xml:"DispoID"`
	StoerungsInfo            []StoerungsInfo            `xml:"StoerungsInfo"`
	SollHalt                 []SollHalt                 `xml:"SollHalt"`
	UmlaufID                 []UmlaufID                 `xml:"UmlaufID"`
	KursNr                   []KursNr                   `xml:"KursNr"`
	FahrtBezeichnerText      []TextMitSprachCode        `xml:"FahrtBezeichnerText"`
	VerkehrsmittelNummer     []string                   `xml:"VerkehrsmittelNummer"`
	LinienText               []TextMitSprachCode        `xml:"LinienText"`
	ProduktID                []ProduktID                `xml:"ProduktID"`
	VonRichtungsText         []TextMitSprachCode        `xml:"VonRichtungsText"`
	FahrtspezialText         []TextMitSprachCode        `xml:"FahrtspezialText"`
	LinienfahrwegID          []LinienfahrwegID          `xml:"LinienfahrwegID"`
	Zugname                  []TextMitSprachCode        `xml:"Zugname"`
	VerkehrsmittelID         []VerkehrsmittelID         `xml:"VerkehrsmittelID"`
	Zusatzfahrt              []bool                     `xml:"Zusatzfahrt"`
	FaelltAus                []bool                     `xml:"FaelltAus"`
	FahrradMitnahme          []bool                     `xml:"FahrradMitnahme"`
	FahrzeugTypID            []string                   `xml:"FahrzeugTypID"`
	ServiceAttribut          []ServiceAttribut          `xml:"ServiceAttribut"`
	SollFormation            []Formation                `xml:"SollFormation"`
	FahrtBeziehung           []Beziehung                `xml:"FahrtBeziehung"`
	Zst                      time.Time                  `xml:"Zst,attr"`
}
type SollHalt struct {
	SollAnschluss []SollAnschluss `xml:"SollAnschluss"`
}
type StatusAnfrage struct {
	Sender Sender    `xml:"Sender,attr"`
	Zst    time.Time `xml:"Zst,attr"`
}
type StatusAntwort struct {
	Status         []Status    `xml:"Status"`
	DatenBereit    []bool      `xml:"DatenBereit"`
	StartDienstZst []time.Time `xml:"StartDienstZst"`
	DatenVersionID []string    `xml:"DatenVersionID"`
}
type Status struct {
	Zst      time.Time `xml:"Zst,attr"`
	Ergebnis Ergebnis  `xml:"Ergebnis,attr"`
}
type StoerungsInfo struct {
	UrsacheText []TextMitSprachCode `xml:"UrsacheText"`
}
type StreckenBezug struct {
	Halt           []HaltExtended   `xml:"Halt"`
	FahrtAbschnitt []FahrtStartEnde `xml:"FahrtAbschnitt"`
}
type Traktions struct {
	TraktionsID   []TraktionsID   `xml:"TraktionsID"`
	AnzahlFahrten []AnzahlFahrten `xml:"AnzahlFahrten"`
	Position      []Positions     `xml:"Position"`
}
type VerkehrsmittelIDFilter struct {
	VerkehrsmittelID []VerkehrsmittelID `xml:"VerkehrsmittelID"`
}
type Via struct {
	HaltestellenName       []TextMitSprachCode `xml:"HaltestellenName"`
	HaltestellenPrioritaet []int               `xml:"HaltestellenPrioritaet"`
}
type VISFahrplanlage struct {
	Protokolleintrag        []Protokolleintrag  `xml:"Protokolleintrag"`
	VISMeldungsart          []string            `xml:"VISMeldungsart"`
	VISID                   []ID                `xml:"VISID"`
	FahrtID                 []FahrtID           `xml:"FahrtID"`
	LinienID                []LinienID          `xml:"LinienID"`
	LinienText              []TextMitSprachCode `xml:"LinienText"`
	RichtungsID             []RichtungsID       `xml:"RichtungsID"`
	RichtungsText           []TextMitSprachCode `xml:"RichtungsText"`
	VonRichtungsText        []TextMitSprachCode `xml:"VonRichtungsText"`
	PrognoseMoeglich        []bool              `xml:"PrognoseMoeglich"`
	Verspaetung             []Verspaetung       `xml:"Verspaetung"`
	StartHstnameKurz        []TextMitSprachCode `xml:"StartHstnameKurz"`
	EndHstnameKurz          []TextMitSprachCode `xml:"EndHstnameKurz"`
	AktHaltID               []HaltID            `xml:"AktHaltID"`
	AktHstnameKurz          []TextMitSprachCode `xml:"AktHstnameKurz"`
	AktHstSeqZaehler        []HstSeqZaehler     `xml:"AktHstSeqZaehler"`
	AufHst                  []bool              `xml:"AufHst"`
	NachHaltID              []HaltID            `xml:"NachHaltID"`
	NachHstnameKurz         []TextMitSprachCode `xml:"NachHstnameKurz"`
	NachHstSeqZaehler       []HstSeqZaehler     `xml:"NachHstSeqZaehler"`
	Distanz                 []Distanz           `xml:"Distanz"`
	Geschwindigkeit         []int               `xml:"Geschwindigkeit"`
	Longitude               []Longitude         `xml:"Longitude"`
	Latitude                []Latitude          `xml:"Latitude"`
	PrognoseUngenau         []PrognoseUngenau   `xml:"PrognoseUngenau"`
	FahrtAbbruchUrsacheText []TextMitSprachCode `xml:"FahrtAbbruchUrsacheText"`
	Auslastungsstufe        []Auslastungsstufe  `xml:"Auslastungsstufe"`
	FahrtInfo               []FahrtInfo         `xml:"FahrtInfo"`
	Zst                     time.Time           `xml:"Zst,attr"`
	VerfallZst              time.Time           `xml:"VerfallZst,attr"`
}
type VISNachricht struct {
	VISFahrplanlage []VISFahrplanlage `xml:"VISFahrplanlage"`
	AboID           AboID             `xml:"AboID,attr"`
}
type WarteInfo struct {
}
type WartetBis struct {
	Protokolleintrag        []Protokolleintrag `xml:"Protokolleintrag"`
	ASBID                   []ID               `xml:"ASBID"`
	AbbringerInfo           []AbbringerInfo    `xml:"AbbringerInfo"`
	FahrtIDExt              []FahrtIDExt       `xml:"FahrtIDExt"`
	AbfahrtszeitASBPrognose []time.Time        `xml:"AbfahrtszeitASBPrognose"`
	Verlaesslichkeit        []Verlaesslichkeit `xml:"Verlaesslichkeit"`
	Zst                     time.Time          `xml:"Zst,attr"`
}
type Zeitfenster struct {
	GueltigVon []time.Time `xml:"GueltigVon"`
	GueltigBis []time.Time `xml:"GueltigBis"`
}
type ZeitFilter struct {
	LinienFilter          []LinienFilter `xml:"LinienFilter"`
	FruehesteAnkunftszeit []time.Time    `xml:"FruehesteAnkunftszeit"`
	SpaetesteAnkunftszeit []time.Time    `xml:"SpaetesteAnkunftszeit"`
	Vorschauzeit          []Vorschauzeit `xml:"Vorschauzeit"`
}
type ZeitQualitaet struct {
	PrognoseVerlaesslichkeit []Verlaesslichkeit `xml:"PrognoseVerlaesslichkeit"`
}
type Zubringernachricht struct {
	ASBLinienFahrplan []ASBLinienFahrplan `xml:"ASBLinienFahrplan"`
	ASBFahrplanlage   []ASBFahrplanlage   `xml:"ASBFahrplanlage"`
	AboID             AboID               `xml:"AboID,attr"`
}
type ZurueckhaltungBis struct {
}
type AbmeldeID = int
type AboID = int
type AnschlussID = string
type AnzahlFahrten = int

const ASBZubringerMeldungsartFahrplanlage = "ASBZubringerMeldungsartFahrplanlage"
const ASBZubringerMeldungsartAusfall = "ASBZubringerMeldungsartAusfall"
const ASBZubringerMeldungsartBereichErreicht = "ASBZubringerMeldungsartBereichErreicht"
const AZBMeldungsartFahrplanlage = "AZBMeldungsartFahrplanlage"
const AZBMeldungsartAusfall = "AZBMeldungsartAusfall"
const AZBMeldungsartBereichVerlassen = "AZBMeldungsartBereichVerlassen"

type Auslastungsstufe = string
type Betreiber = string
type Distanz = int
type ElementFuerUmwandlungInClassicHaltID = string
type Ergebnis = string
type FahrtBezeichner = string
type FahrzeugID = string

const NachrichtenMeldungsartNachricht = "NachrichtenMeldungsartNachricht"
const NachrichtenMeldungsartLoeschen = "NachrichtenMeldungsartLoeschen"

type Formation struct {
}
type FoFahrzeug struct {
	FoFahrzeugNummer []FoFahrzeugNummer `xml:"FoFahrzeugNummer"`
	FoOrdnungsNummer []FoOrdnungsNummer `xml:"FoOrdnungsNummer"`
	FoFahrzeugID     string             `xml:"FoFahrzeugID,attr"`
}
type FoFahrzeugTyp = FoString
type FoFahrzeugNummer = FoString
type FoOrdnungsNummer = int
type FoFahrzeugAusstattung struct {
	FoFahrzeugAusstattungsCode []string            `xml:"FoFahrzeugAusstattungsCode"`
	FoAnzahl                   []FoAnzahl          `xml:"FoAnzahl"`
	FoBezeichnung              []TextMitSprachCode `xml:"FoBezeichnung"`
	FoFahrzeugAusstattungID    string              `xml:"FoFahrzeugAusstattungID,attr"`
}
type FoAnzahl = int
type FoTechnischesAttribut struct {
	FoTechnischesAttributCode []string            `xml:"FoTechnischesAttributCode"`
	FoWert                    []FoWert            `xml:"FoWert"`
	FoBezeichnung             []TextMitSprachCode `xml:"FoBezeichnung"`
}
type FoWert = FoString
type FoFahrzeugGruppe struct {
	FoVerkehrlicheNummer       []FoVerkehrlicheNummer           `xml:"FoVerkehrlicheNummer"`
	FoFahrzeugGruppenZielText  []FoFahrzeugGruppenStartZielText `xml:"FoFahrzeugGruppenZielText"`
	FoFahrzeugGruppenStartText []FoFahrzeugGruppenStartZielText `xml:"FoFahrzeugGruppenStartText"`
	FoFahrzeugGruppeID         string                           `xml:"FoFahrzeugGruppeID,attr"`
}
type FoVerkehrlicheNummer = FoString
type FoFahrzeugGruppenStartZielText = FoString
type FoFahrzeugPosition struct {
	FoFahrzeugIDREF []FoFahrzeugIDREF `xml:"FoFahrzeugIDREF"`
	FoPosition      []FoPosition      `xml:"FoPosition"`
	FoOrientierung  []string          `xml:"FoOrientierung"`
}
type FoFahrzeugGruppenFahrtAbschnitt struct {
	FoAbschnitt     []FahrtStartEnde `xml:"FoAbschnitt"`
	FoFahrtrichtung []string         `xml:"FoFahrtrichtung"`
}
type FoFahrtAbschnittFahrzeugGruppe struct {
	FoFahrzeugGruppeIDREF []FoFahrzeugGruppeIDREF `xml:"FoFahrzeugGruppeIDREF"`
	FoPosition            []FoPosition            `xml:"FoPosition"`
}
type FoDurchgang struct {
	FoFahrzeugGruppeIDREF []FoFahrzeugGruppeIDREF `xml:"FoFahrzeugGruppeIDREF"`
	FoDurchgangMoeglich   []bool                  `xml:"FoDurchgangMoeglich"`
}
type FoFormationAenderung struct {
	FoAenderungsCode  []string            `xml:"FoAenderungsCode"`
	FoAenderungsTexte []FoAenderungsTexte `xml:"FoAenderungsTexte"`
}
type FoFahrzeugAusstattungFahrtAbschnitt struct {
	FoAbschnitt []FahrtStartEnde `xml:"FoAbschnitt"`
}
type FoFahrzeugAusstattungZustand struct {
	FoFahrzeugAusstattungIDREF []FoFahrzeugAusstattungIDREF `xml:"FoFahrzeugAusstattungIDREF"`
	FoZustand                  []FoZustand                  `xml:"FoZustand"`
}
type FoFahrzeugZustandFahrtAbschnitt struct {
	FoAbschnitt []FahrtStartEnde `xml:"FoAbschnitt"`
}
type FoFahrzeugZustand struct {
	FoFahrzeugIDREF []FoFahrzeugIDREF `xml:"FoFahrzeugIDREF"`
	FoZustand       []FoZustand       `xml:"FoZustand"`
}
type FoAusstattungsBelegung struct {
	FoFahrzeugAusstattungsCode []string           `xml:"FoFahrzeugAusstattungsCode"`
	FoAnzahl                   []FoAnzahl         `xml:"FoAnzahl"`
	FoBelegung                 []FoAnzahl         `xml:"FoBelegung"`
	FoAuslastungsstufe         []Auslastungsstufe `xml:"FoAuslastungsstufe"`
}
type FoFahrzeugBelegungFahrtAbschnitt struct {
	FoAbschnitt               []FahrtStartEnde            `xml:"FoAbschnitt"`
	FoFahrzeugBelegungen      []FoFahrzeugBelegungenTyp   `xml:"FoFahrzeugBelegungen"`
	FoReisegruppenProFahrzeug []FoReisegruppenProFahrzeug `xml:"FoReisegruppenProFahrzeug"`
}
type FoReisegruppenProFahrzeug struct {
	FoFahrzeugIDREF        []FoFahrzeugIDREF        `xml:"FoFahrzeugIDREF"`
	FoReisegruppeVorhanden []FoReisegruppeVorhanden `xml:"FoReisegruppeVorhanden"`
	FoReisegruppenName     []FoReisegruppenName     `xml:"FoReisegruppenName"`
}
type FoReisegruppeVorhanden = bool
type FoBelegungProzentual = int
type FoReisegruppenName = FoString
type FoFahrzeugBelegungenTyp struct {
}
type FoZaehldatenGroup struct {
	FoBelegungReal         []FoAnzahl `xml:"FoBelegungReal"`
	FoEinsteigerBeiAbfahrt []FoAnzahl `xml:"FoEinsteigerBeiAbfahrt"`
	FoAussteigerBeiAnkunft []FoAnzahl `xml:"FoAussteigerBeiAnkunft"`
}
type FoHalt struct {
	HaltID                 []HaltID           `xml:"HaltID"`
	Abfahrtszeit           []time.Time        `xml:"Abfahrtszeit"`
	Ankunftszeit           []time.Time        `xml:"Ankunftszeit"`
	StoerungsInfo          []StoerungsInfo    `xml:"StoerungsInfo"`
	FoAnkunft              []FoAnkunftAbfahrt `xml:"FoAnkunft"`
	FoAbfahrt              []FoAnkunftAbfahrt `xml:"FoAbfahrt"`
	FoAnkunftGleichAbfahrt []bool             `xml:"FoAnkunftGleichAbfahrt"`
}
type FoAnkunftAbfahrt struct {
	FoFahrtrichtungAmHalt []string `xml:"FoFahrtrichtungAmHalt"`
}
type FoFahrzeugAmHalt struct {
	FoFahrzeugIDREF []FoFahrzeugIDREF `xml:"FoFahrzeugIDREF"`
	FoHaltPosition  []FoHaltPosition  `xml:"FoHaltPosition"`
	FoZustand       []FoZustand       `xml:"FoZustand"`
	FoErweiterung   []Erweiterung     `xml:"FoErweiterung"`
}
type FoHaltPosition struct {
	FoHaltPositionStart []int `xml:"FoHaltPositionStart"`
	FoHaltPositionEnde  []int `xml:"FoHaltPositionEnde"`
}
type FoSektorHaltPosition struct {
	FoSektorBezeichnung          []FoString       `xml:"FoSektorBezeichnung"`
	FoSektorBezeichnungPosition  []int            `xml:"FoSektorBezeichnungPosition"`
	FoSektorBezeichnungVorhanden []bool           `xml:"FoSektorBezeichnungVorhanden"`
	FoHaltPosition               []FoHaltPosition `xml:"FoHaltPosition"`
}
type FoAenderungAmHalt struct {
	FoAenderungsCodeAmHalt []string            `xml:"FoAenderungsCodeAmHalt"`
	FoAenderungsTexte      []FoAenderungsTexte `xml:"FoAenderungsTexte"`
}
type FoZustand struct {
	FoZustandsCode       []string            `xml:"FoZustandsCode"`
	FoZustandsKurzform   []TextMitSprachCode `xml:"FoZustandsKurzform"`
	FoZustandsText       []TextMitSprachCode `xml:"FoZustandsText"`
	FoZustandsEmpfehlung []TextMitSprachCode `xml:"FoZustandsEmpfehlung"`
}
type FoAenderungsTexte struct {
	FoAenderungsKurzform   []TextMitSprachCode `xml:"FoAenderungsKurzform"`
	FoAenderungsText       []TextMitSprachCode `xml:"FoAenderungsText"`
	FoAenderungsEmpfehlung []TextMitSprachCode `xml:"FoAenderungsEmpfehlung"`
}
type FoPosition = int
type FoBarrierefrei = bool
type FoString = string
type TextMitSprachCode struct {
}
type TextMitSprachCodeMinLenght struct {
}
type Erweiterung struct {
}
type FoFahrzeugIDREF = string
type FoStatusIDREF = string
type FoFahrzeugAusstattungIDREF = string
type FoFahrzeugGruppeIDREF = string
type FoFahrzeugBereichIDREF = string

const FoOrientierungvorwaerts = "FoOrientierungvorwaerts"
const FoOrientierungrueckwaerts = "FoOrientierungrueckwaerts"
const FoAenderungsCodeFehlendeFamilienwagen = "FoAenderungsCodeFehlendeFamilienwagen"
const FoAenderungsCodeFehlendeFzgGruppe = "FoAenderungsCodeFehlendeFzgGruppe"
const FoAenderungsCodeFehlendeKurswagen = "FoAenderungsCodeFehlendeKurswagen"
const FoAenderungsCodeFehlendeRestaurantwagen = "FoAenderungsCodeFehlendeRestaurantwagen"
const FoAenderungsCodeFehlendeWagen = "FoAenderungsCodeFehlendeWagen"
const FoAenderungsCodeFehlendeRollstuhlplaetze = "FoAenderungsCodeFehlendeRollstuhlplaetze"
const FoAenderungsCodeFehlendeNiederflurwagen = "FoAenderungsCodeFehlendeNiederflurwagen"
const FoAenderungsCodeGeaenderteWagenreihung = "FoAenderungsCodeGeaenderteWagenreihung"
const FoAenderungsCodeUmgekehrteWagenreihung = "FoAenderungsCodeUmgekehrteWagenreihung"
const FoAenderungsCodeZusaetzlicheFzgGruppe = "FoAenderungsCodeZusaetzlicheFzgGruppe"
const FoAenderungsCodeZusaetzlicheKurswagen = "FoAenderungsCodeZusaetzlicheKurswagen"
const FoAenderungsCodeZusaetzlicheWagen = "FoAenderungsCodeZusaetzlicheWagen"
const FoAenderungsCodeAmHaltFehlendeFamilienwagen = "FoAenderungsCodeAmHaltFehlendeFamilienwagen"
const FoAenderungsCodeAmHaltFehlendeFzgGruppe = "FoAenderungsCodeAmHaltFehlendeFzgGruppe"
const FoAenderungsCodeAmHaltFehlendeKurswagen = "FoAenderungsCodeAmHaltFehlendeKurswagen"
const FoAenderungsCodeAmHaltFehlendeRestaurantwagen = "FoAenderungsCodeAmHaltFehlendeRestaurantwagen"
const FoAenderungsCodeAmHaltFehlendeWagen = "FoAenderungsCodeAmHaltFehlendeWagen"
const FoAenderungsCodeAmHaltGeaenderteWagenreihung = "FoAenderungsCodeAmHaltGeaenderteWagenreihung"
const FoAenderungsCodeAmHaltUmgekehrteWagenreihung = "FoAenderungsCodeAmHaltUmgekehrteWagenreihung"
const FoAenderungsCodeAmHaltZusaetzlicheFzgGruppe = "FoAenderungsCodeAmHaltZusaetzlicheFzgGruppe"
const FoAenderungsCodeAmHaltZusaetzlicheKurswagen = "FoAenderungsCodeAmHaltZusaetzlicheKurswagen"
const FoAenderungsCodeAmHaltZusaetzlicheWagen = "FoAenderungsCodeAmHaltZusaetzlicheWagen"
const FoZustandsCodedefekt = "FoZustandsCodedefekt"
const FoZustandsCodegeschlossen = "FoZustandsCodegeschlossen"
const FoZustandsCodenicht_verfuegbar = "FoZustandsCodenicht_verfuegbar"
const FoZustandsCodeoffen = "FoZustandsCodeoffen"
const FoZustandsCodereserviert = "FoZustandsCodereserviert"
const FoZustandsCodenicht_bedient = "FoZustandsCodenicht_bedient"
const FoZustandsCodeverfuegbar = "FoZustandsCodeverfuegbar"
const FoBeziehungsTypFortfuehrungDurchFahrt = "FoBeziehungsTypFortfuehrungDurchFahrt"
const FoBeziehungsTypFortfuehrungVonFahrt = "FoBeziehungsTypFortfuehrungVonFahrt"
const FoBeziehungsTypTrennungVonFahrtIn = "FoBeziehungsTypTrennungVonFahrtIn"
const FoBeziehungsTypFortfuehrungVonGetrennterFahrt = "FoBeziehungsTypFortfuehrungVonGetrennterFahrt"
const FoBeziehungsTypZusammenfuehrungVonFahrt = "FoBeziehungsTypZusammenfuehrungVonFahrt"
const FoBeziehungsTypFortfuehrungDurchZusammengefuehrteFahrt = "FoBeziehungsTypFortfuehrungDurchZusammengefuehrteFahrt"
const FoBeziehungsTypErsatzVonFahrt = "FoBeziehungsTypErsatzVonFahrt"
const FoBeziehungsTypErsatzDurchFahrt = "FoBeziehungsTypErsatzDurchFahrt"
const FoBeziehungsTypEntlastungVonFahrt = "FoBeziehungsTypEntlastungVonFahrt"
const FoBeziehungsTypEntlastungDurchFahrt = "FoBeziehungsTypEntlastungDurchFahrt"
const FoBeziehungsTypWendeAufFahrt = "FoBeziehungsTypWendeAufFahrt"
const FoBeziehungsTypWendeVonFahrt = "FoBeziehungsTypWendeVonFahrt"
const FoFahrtrichtungAmHaltgegen0 = "FoFahrtrichtungAmHaltgegen0"
const FoFahrtrichtungAmHaltgegen100 = "FoFahrtrichtungAmHaltgegen100"
const FoFahrzeugAusstattungsCodeAbteilBistro = "FoFahrzeugAusstattungsCodeAbteilBistro"
const FoFahrzeugAusstattungsCodeAbteilBusiness = "FoFahrzeugAusstattungsCodeAbteilBusiness"
const FoFahrzeugAusstattungsCodeAbteilCC = "FoFahrzeugAusstattungsCodeAbteilCC"
const FoFahrzeugAusstattungsCodeAbteilFahrrad = "FoFahrzeugAusstattungsCodeAbteilFahrrad"
const FoFahrzeugAusstattungsCodeAbteilFahrradResPflicht = "FoFahrzeugAusstattungsCodeAbteilFahrradResPflicht"
const FoFahrzeugAusstattungsCodeAbteilFamilien = "FoFahrzeugAusstattungsCodeAbteilFamilien"
const FoFahrzeugAusstattungsCodeAbteilGepaeck = "FoFahrzeugAusstattungsCodeAbteilGepaeck"
const FoFahrzeugAusstattungsCodeAbteilKinderwagen = "FoFahrzeugAusstattungsCodeAbteilKinderwagen"
const FoFahrzeugAusstattungsCodeAbteilKlasse1 = "FoFahrzeugAusstattungsCodeAbteilKlasse1"
const FoFahrzeugAusstattungsCodeAbteilKlasse2 = "FoFahrzeugAusstattungsCodeAbteilKlasse2"
const FoFahrzeugAusstattungsCodeAbteilKlasse1Plus = "FoFahrzeugAusstattungsCodeAbteilKlasse1Plus"
const FoFahrzeugAusstattungsCodeAbteilKlasse2Plus = "FoFahrzeugAusstattungsCodeAbteilKlasse2Plus"
const FoFahrzeugAusstattungsCodeAbteilKleinkind = "FoFahrzeugAusstattungsCodeAbteilKleinkind"
const FoFahrzeugAusstattungsCodeAbteilMehrzweck = "FoFahrzeugAusstattungsCodeAbteilMehrzweck"
const FoFahrzeugAusstattungsCodeAbteilRestaurant = "FoFahrzeugAusstattungsCodeAbteilRestaurant"
const FoFahrzeugAusstattungsCodeAbteilRollstuhl = "FoFahrzeugAusstattungsCodeAbteilRollstuhl"
const FoFahrzeugAusstattungsCodeAbteilRuhe = "FoFahrzeugAusstattungsCodeAbteilRuhe "
const FoFahrzeugAusstattungsCodeAbteilWL = "FoFahrzeugAusstattungsCodeAbteilWL"
const FoFahrzeugAusstattungsCodeAutoreisezugwagen = "FoFahrzeugAusstattungsCodeAutoreisezugwagen"
const FoFahrzeugAusstattungsCodeKlima = "FoFahrzeugAusstattungsCodeKlima"
const FoFahrzeugAusstattungsCodePlaetze1 = "FoFahrzeugAusstattungsCodePlaetze1"
const FoFahrzeugAusstattungsCodePlaetze2 = "FoFahrzeugAusstattungsCodePlaetze2"
const FoFahrzeugAusstattungsCodePlaetze1Plus = "FoFahrzeugAusstattungsCodePlaetze1Plus"
const FoFahrzeugAusstattungsCodePlaetze2Plus = "FoFahrzeugAusstattungsCodePlaetze2Plus"
const FoFahrzeugAusstattungsCodePlaetzeCC = "FoFahrzeugAusstattungsCodePlaetzeCC"
const FoFahrzeugAusstattungsCodePlaetzeFahrrad = "FoFahrzeugAusstattungsCodePlaetzeFahrrad"
const FoFahrzeugAusstattungsCodePlaetzeFahrradResPflicht = "FoFahrzeugAusstattungsCodePlaetzeFahrradResPflicht"
const FoFahrzeugAusstattungsCodePlaetzeGesamt = "FoFahrzeugAusstattungsCodePlaetzeGesamt"
const FoFahrzeugAusstattungsCodePlaetzeResMoeglich = "FoFahrzeugAusstattungsCodePlaetzeResMoeglich"
const FoFahrzeugAusstattungsCodePlaetzeResPflicht = "FoFahrzeugAusstattungsCodePlaetzeResPflicht"
const FoFahrzeugAusstattungsCodePlaetzeRollstuhl = "FoFahrzeugAusstattungsCodePlaetzeRollstuhl"
const FoFahrzeugAusstattungsCodePlaetzeSteh = "FoFahrzeugAusstattungsCodePlaetzeSteh"
const FoFahrzeugAusstattungsCodePlaetzeSteh1 = "FoFahrzeugAusstattungsCodePlaetzeSteh1"
const FoFahrzeugAusstattungsCodePlaetzeSteh2 = "FoFahrzeugAusstattungsCodePlaetzeSteh2"
const FoFahrzeugAusstattungsCodePlaetzeSteh1Plus = "FoFahrzeugAusstattungsCodePlaetzeSteh1Plus"
const FoFahrzeugAusstattungsCodePlaetzeSteh2Plus = "FoFahrzeugAusstattungsCodePlaetzeSteh2Plus"
const FoFahrzeugAusstattungsCodePlaetzeWL = "FoFahrzeugAusstattungsCodePlaetzeWL"
const FoFahrzeugAusstattungsCodePlaetzeWR = "FoFahrzeugAusstattungsCodePlaetzeWR"
const FoFahrzeugAusstattungsCodePlaetzeKinderwagen = "FoFahrzeugAusstattungsCodePlaetzeKinderwagen"
const FoFahrzeugAusstattungsCodePlaetzeFamilien = "FoFahrzeugAusstattungsCodePlaetzeFamilien"
const FoFahrzeugAusstattungsCodeRollstuhlToilette = "FoFahrzeugAusstattungsCodeRollstuhlToilette"
const FoFahrzeugAusstattungsCodeSteckdosen = "FoFahrzeugAusstattungsCodeSteckdosen"
const FoFahrzeugAusstattungsCodeToilette = "FoFahrzeugAusstattungsCodeToilette"
const FoFahrzeugAusstattungsCodeUSBLadebuchse = "FoFahrzeugAusstattungsCodeUSBLadebuchse"
const FoFahrzeugAusstattungsCodeWLAN = "FoFahrzeugAusstattungsCodeWLAN"
const FoFahrzeugAusstattungsCodePlaetzePRM = "FoFahrzeugAusstattungsCodePlaetzePRM"
const FoTechnischesAttributCodeAnzAchsen = "FoTechnischesAttributCodeAnzAchsen"
const FoTechnischesAttributCodeFzGesamtLaenge = "FoTechnischesAttributCodeFzGesamtLaenge"
const FoTechnischesAttributCodeFzUICAustauschregime = "FoTechnischesAttributCodeFzUICAustauschregime"
const FoTechnischesAttributCodeFzUICNummer = "FoTechnischesAttributCodeFzUICNummer"
const FoTechnischesAttributCodeFzUICOrdnungsnummer = "FoTechnischesAttributCodeFzUICOrdnungsnummer"
const FoTechnischesAttributCodeFzUICPruefziffer = "FoTechnischesAttributCodeFzUICPruefziffer"
const FoTechnischesAttributCodeFzUICTyp = "FoTechnischesAttributCodeFzUICTyp"
const FoTechnischesAttributCodeFzUICVerwaltung = "FoTechnischesAttributCodeFzUICVerwaltung"
const FoTechnischesAttributCodeGewichtLeer = "FoTechnischesAttributCodeGewichtLeer"
const FoTechnischesAttributCodeHoeheEinstieg = "FoTechnischesAttributCodeHoeheEinstieg"
const FoTechnischesAttributCodeStufenfrei = "FoTechnischesAttributCodeStufenfrei"
const FoTechnischesAttributCodeNiederflurEinstieg = "FoTechnischesAttributCodeNiederflurEinstieg"

type FahrtVerband struct {
	BetrieblicheFahrtnummer []string                 `xml:"BetrieblicheFahrtnummer"`
	FahrtVerbandsAbschnitt  []FahrtVerbandsAbschnitt `xml:"FahrtVerbandsAbschnitt"`
}
type FahrtVerbandsAbschnitt struct {
	FahrtAbschnitt   []FahrtStartEnde   `xml:"FahrtAbschnitt"`
	FahrtInAbschnitt []FahrtInAbschnitt `xml:"FahrtInAbschnitt"`
}
type FahrtInAbschnitt struct {
	FahrtId  []FahrtID `xml:"FahrtId"`
	Position []int     `xml:"Position"`
}
type Fehlernummer = int
type Fehlertext = string
type FormatID = string
type HaltIDFilter struct {
	HaltestellenID []string `xml:"HaltestellenID"`
	BereichsID     []string `xml:"BereichsID"`
	SteigID        []string `xml:"SteigID"`
	SektorenID     []string `xml:"SektorenID"`
}
type HaltID struct {
	HaltestellenID                       []string                               `xml:"HaltestellenID"`
	BereichsID                           []string                               `xml:"BereichsID"`
	SteigID                              []string                               `xml:"SteigID"`
	ElementFuerUmwandlungInClassicHaltID []ElementFuerUmwandlungInClassicHaltID `xml:"ElementFuerUmwandlungInClassicHaltID"`
	AbgeleitetAusClassicHaltID           []bool                                 `xml:"AbgeleitetAusClassicHaltID"`
}
type HstSeqZaehler = int
type Hysterese = int
type ID = string
type IPAdresse = string
type KanalID = string
type KursNr = int
type Latitude = int
type LeitstellenID = string
type LinienID = string
type LinienfahrwegID = string
type LinienNr = int
type Longitude = int
type MaxTextLaenge = int
type MeldungsID = string
type Positions = int
type Prioritaet = int
type ProduktID = string
type PrognoseStatus = string
type PrognoseUngenau = string

const ProtokollAktionEINGANG = "ProtokollAktionEINGANG"
const ProtokollAktionAUSGANG = "ProtokollAktionAUSGANG"
const ProtokollAktionAUSLOESUNG = "ProtokollAktionAUSLOESUNG"
const ProtokollAktionINTERN = "ProtokollAktionINTERN"
const ProtokollAktionSONSTIGES = "ProtokollAktionSONSTIGES"

type RichtungsID = string
type Sender = string
type Telefonnummer = string
type TraktionsID = string
type UmlaufID = string
type UmlaufNr = int
type UmsteigeWillige = int
type VerkehrsmittelID = string
type Verlaesslichkeit = int
type Verspaetung = int

const VISMeldungsartFahrplanlage = "VISMeldungsartFahrplanlage"
const VISMeldungsartFahrtabbruch = "VISMeldungsartFahrtabbruch"
const VISMeldungsartRegulaeresFahrtende = "VISMeldungsartRegulaeresFahrtende"
const VISMeldungsartBereichVerlassen = "VISMeldungsartBereichVerlassen"

type Vorschauzeit = int
type FoFahrzeugBereich struct {
	FoFahrzeugBereichName []string `xml:"FoFahrzeugBereichName"`
	FoFahrzeugBereichID   string   `xml:"FoFahrzeugBereichID,attr"`
}
