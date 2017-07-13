package main

// Cell is a cell line data type.
type Cell struct {
	ID         int       `json:"id"`
	Identifier string    `json:"identifier"`
	Accession  Accession `json:"accession"`
	CA         string    `json:"category,omitempty"`
	SX         *string   `json:"sex,omitempty"`
	SY         []string  `json:"synonyms,omitempty"`
	OX         []OX      `json:"species-of-origin,omitempty"`
	DR         []DR      `json:"cross-references,omitempty"`
	RX         []string  `json:"reference-identifiers,omitempty"`
	WW         []string  `json:"web-pages,omitempty"`
	CC         []CC      `json:"comments,omitempty"`
	ST         *ST       `json:"str-profile-data,omitempty"`
	DI         []DI      `json:"diseases,omitempty"`
	HI         []HI      `json:"hierarchy,omitempty"`
	OI         []OI      `json:"same-origin-as,omitempty"`
}

// Accession is a unique accession id.
type Accession struct {
	Primary   string  `json:"primary"`
	Secondary *string `json:"secondary"`
}

// OX is a species of origin data.
type OX struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	Species     string `json:"species,omitempty"`
}

// DR is a cross references data.
type DR struct {
	Database  string `json:"database,omitempty"`
	Accession string `json:"accession,omitempty"`
}

// CC is a comments data.
type CC struct {
	Category string `json:"category,omitempty"`
	Comment  string `json:"comment,omitemtpy"`
}

// Marker is an str profile data marker.
type Marker struct {
	ID      string `json:"id,omitempty"`
	Alleles string `json:"alleles,omitempty"`
}

// ST is an str profile data.
type ST struct {
	Sources []string `json:"sources,omitempty"`
	Markers []Marker `json:"markers,omitempty"`
}

// DI is a disease data.
type DI struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	Disease     string `json:"disease,omitempty"`
}

// HI is a hierarchy data.
type HI struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	DF          string `json:"derived-from,omitempty"`
}

// OI is a same-origin-as data.
type OI struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}
