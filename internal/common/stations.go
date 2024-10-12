package common

import "math/rand"

type SubwayStation int

const (
	Bathurst SubwayStation = iota
	Bay
	Bayview
	Bessarion
	BloorYonge
	Broadview
	CastleFrank
	Chester
	Christie
	College
	Coxwell
	Davisville
	DonMills
	Donlands
	DownsviewPark
	Dufferin
	Dundas
	DundasWest
	Dupont
	Eglinton
	EglintonWest
	Finch
	FinchWest
	Glencairn
	Greenwood
	HighPark
	Highway407
	Islington
	Jane
	Keele
	Kennedy
	King
	Kipling
	Lansdowne
	Lawrence
	LawrenceWest
	Leslie
	MainStreet
	Museum
	NorthYorkCentre
	OldMill
	Osgoode
	Ossington
	Pape
	PioneerVillage
	Queen
	QueensPark
	Rosedale
	RoyalYork
	Runnymede
	SheppardWest
	SheppardYonge
	Sherbourne
	Spadina
	StAndrew
	StClair
	StClairWest
	StGeorge
	StPatrick
	Summerhill
	Union
	Vaughan
	VictoriaPark
	Warden
	Wellesley
	Wilson
	Woodbine
	YorkMills
	YorkUniversity
	Yorkdale
)

func (s SubwayStation) String() string {
	switch s {
	case Bathurst:
		return "Bathurst"
	case Bay:
		return "Bay"
	case Bayview:
		return "Bayview"
	case Bessarion:
		return "Bessarion"
	case BloorYonge:
		return "Bloor-Yonge"
	case Broadview:
		return "Broadview"
	case CastleFrank:
		return "Castle Frank"
	case Chester:
		return "Chester"
	case Christie:
		return "Christie"
	case College:
		return "College"
	case Coxwell:
		return "Coxwell"
	case Davisville:
		return "Davisville"
	case DonMills:
		return "DonMills"
	case Donlands:
		return "Donlands"
	case DownsviewPark:
		return "Downsview Park"
	case Dufferin:
		return "Dufferin"
	case Dundas:
		return "Dundas"
	case DundasWest:
		return "Dundas West"
	case Dupont:
		return "Dupont"
	case Eglinton:
		return "Eglinton"
	case EglintonWest:
		return "Eglinton West"
	case Finch:
		return "Finch"
	case FinchWest:
		return "Finch West"
	case Glencairn:
		return "Glencairn"
	case Greenwood:
		return "Greenwood"
	case HighPark:
		return "High Park"
	case Highway407:
		return "Highway 407"
	case Islington:
		return "Islington"
	case Jane:
		return "Jane"
	case Keele:
		return "Keele"
	case Kennedy:
		return "Kennedy"
	case King:
		return "King"
	case Kipling:
		return "Kipling"
	case Lansdowne:
		return "Lansdowne"
	case Lawrence:
		return "Lawrence"
	case LawrenceWest:
		return "Lawrence West"
	case Leslie:
		return "Leslie"
	case MainStreet:
		return "Main Street"
	case Museum:
		return "Museum"
	case NorthYorkCentre:
		return "North York Centre"
	case OldMill:
		return "OldMill"
	case Osgoode:
		return "Osgoode"
	case Ossington:
		return "Ossington"
	case Pape:
		return "Pape"
	case PioneerVillage:
		return "Pioneer Village"
	case Queen:
		return "Queen"
	case QueensPark:
		return "Queen's Park"
	case Rosedale:
		return "Rosedale"
	case RoyalYork:
		return "Royal York"
	case Runnymede:
		return "Runnymede"
	case SheppardWest:
		return "Sheppard West"
	case SheppardYonge:
		return "Sheppard-Yonge"
	case Sherbourne:
		return "Sherbourne"
	case Spadina:
		return "Spadina"
	case StAndrew:
		return "St. Andrew"
	case StClair:
		return "St. Clair"
	case StClairWest:
		return "St. Clair West"
	case StGeorge:
		return "St. George"
	case StPatrick:
		return "St. Patrick"
	case Summerhill:
		return "Summerhill"
	case Union:
		return "Union"
	case Vaughan:
		return "Vaughan"
	case VictoriaPark:
		return "Victoria Park"
	case Warden:
		return "Warden"
	case Wellesley:
		return "Wellesley"
	case Wilson:
		return "Wilson"
	case Woodbine:
		return "Woodbine"
	case YorkMills:
		return "York Mills"
	case YorkUniversity:
		return "York University"
	case Yorkdale:
		return "Yorkdale"
	}
	return "Unknown"
}

const numStations = 70

func RandomStation() string {
	randomIndex := rand.Intn(numStations)
	station := SubwayStation(randomIndex)

	return station.String()
}
