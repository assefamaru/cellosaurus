package cellosaurus

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

// Cell is a cell line model.
type Cell struct {
	ID string    `json:"identifier"`
	AC Accession `json:"accession"`
	SY []string  `json:"synonyms,omitempty"`
	CA string    `json:"category,omitempty"`
	SX string    `json:"sex,omitempty"`
	DR []DR      `json:"cross-references,omitempty"`
	RX []string  `json:"reference-identifiers,omitempty"`
	WW []string  `json:"web-pages,omitempty"`
	CC []CC      `json:"comments,omitempty"`
	ST ST        `json:"str-profile-data,omitempty"`
	DI []DI      `json:"diseases,omitempty"`
	OX []OX      `json:"species-of-origin,omitempty"`
	HI []HI      `json:"hierarchy,omitempty"`
	OI []OI      `json:"same-origin-as,omitempty"`
	DT []DT      `json:"entry-date,omitempty"`
}

// Cells is a list of cell lines.
type Cells struct {
	Meta cellMeta `json:"meta"`
	Data []Cell   `json:"data"`
}

// cellMeta contains pagination information for returned data.
type cellMeta struct {
	Page     int `json:"page"`
	PerPage  int `json:"per-page"`
	LastPage int `json:"last-page"`
	Total    int `json:"total"`
}

// Accession contains primary and secondary accession number(s).
type Accession struct {
	Pri string   `json:"primary"`
	Sec []string `json:"secondary,omitempty"`
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

// DT is a date (entry history) data.
type DT struct {
	Created     string `json:"created,omitempty"`
	LastUpdated string `json:"last-updated,omitempty"`
	Version     string `json:"version"`
}

// List returns a list of paginated cell lines.
func (cells *Cells) List() error {
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}

	from := (cells.Meta.Page - 1) * cells.Meta.PerPage
	to := cells.Meta.PerPage
	query := fmt.Sprintf("SELECT acp, id, acs, sy, sx, ca FROM cells LIMIT %d,%d;", from, to)

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		logSentry(err)
		return err
	}
	for rows.Next() {
		var (
			acs  string
			sy   string
			cell Cell
		)
		if err = rows.Scan(&cell.AC.Pri, &cell.ID, &acs, &sy, &cell.SX, &cell.CA); err != nil {
			logSentry(err)
			return err
		}
		if acs != "" {
			cell.AC.Sec = strings.Split(acs, "; ")
		}
		if sy != "" {
			cell.SY = strings.Split(sy, "; ")
		}
		cells.Data = append(cells.Data, cell)
	}

	return nil
}

// Find finds and updates a cell line with all its attributes.
func (cell *Cell) Find() error {
	var (
		query string
		acs   string
		sy    string
	)

	db, err := Database()
	defer db.Close()
	if err != nil {
		logSentry(err)
		return err
	}

	id := cell.ID

	query = fmt.Sprintf("SELECT acp, id, acs, sy, sx, ca FROM cells WHERE acp = '%s';", id)
	row := db.QueryRow(query)
	err = row.Scan(&cell.AC.Pri, &cell.ID, &acs, &sy, &cell.SX, &cell.CA)
	if err != nil {
		if err == sql.ErrNoRows {
			query = fmt.Sprintf("SELECT acp, id, acs, sy, sx, ca FROM cells WHERE id = '%s';", id)
			row := db.QueryRow(query)
			err = row.Scan(&cell.AC.Pri, &cell.ID, &acs, &sy, &cell.SX, &cell.CA)
			if err != nil {
				if err == sql.ErrNoRows {
					var acc string
					query = fmt.Sprintf("SELECT accession FROM attributes WHERE (attribute='SY' AND content='%s') LIMIT 1;", id)
					row := db.QueryRow(query)
					err = row.Scan(&acc)
					if err != nil {
						if err != sql.ErrNoRows {
							logSentry(err)
						}
						return err
					}
					query = fmt.Sprintf("SELECT acp, id, acs, sy, sx, ca FROM cells WHERE (acp='%s');", acc)
					row = db.QueryRow(query)
					err = row.Scan(&cell.AC.Pri, &cell.ID, &acs, &sy, &cell.SX, &cell.CA)
					if err != nil {
						if err != sql.ErrNoRows {
							logSentry(err)
						}
						return err
					}
				} else {
					logSentry(err)
					return err
				}
			}
		} else {
			logSentry(err)
			return err
		}
	}

	if acs != "" {
		cell.AC.Sec = strings.Split(acs, "; ")
	}
	if sy != "" {
		cell.SY = strings.Split(sy, "; ")
	}

	query = fmt.Sprintf("SELECT attribute, content FROM attributes WHERE accession = '%s';", cell.AC.Pri)
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		logSentry(err)
		return err
	}
	for rows.Next() {
		var (
			attribute string
			content   string
		)
		if err = rows.Scan(&attribute, &content); err != nil {
			logSentry(err)
			return err
		}

		switch attribute {
		case "DR":
			dr := strings.Split(content, "; ")
			cell.DR = append(cell.DR, DR{dr[0], dr[1]})
		case "RX":
			cell.RX = append(cell.RX, strings.TrimRight(content, ";"))
		case "WW":
			cell.WW = append(cell.WW, content)
		case "CC":
			cc := strings.Split(content, ": ")
			cell.CC = append(cell.CC, CC{cc[0], cc[1]})
		case "ST":
			st := strings.Split(content, ": ")
			if st[0] == "Source(s)" {
				cell.ST.Sources = strings.Split(st[1], "; ")
			} else {
				cell.ST.Markers = append(cell.ST.Markers, Marker{st[0], st[1]})
			}
		case "DI":
			di := strings.Split(content, "; ")
			cell.DI = append(cell.DI, DI{di[0], di[1], di[2]})
		case "OX":
			f := strings.Split(content, " ! ")
			s := strings.Split(strings.TrimRight(f[0], ";"), "=")
			cell.OX = append(cell.OX, OX{"NCBI-Taxonomy", s[1], f[1]})
		case "HI":
			hi := strings.Split(content, " ! ")
			cell.HI = append(cell.HI, HI{"Cellosaurus", hi[0], hi[1]})
		case "OI":
			oi := strings.Split(content, " ! ")
			cell.OI = append(cell.OI, OI{"Cellosaurus", oi[0], oi[1]})
		case "DT":
			dt := strings.Split(content, "; ")
			created := strings.Split(dt[0], ": ")[1]
			updated := strings.Split(dt[1], ": ")[1]
			version := strings.Split(dt[2], ": ")[1]
			cell.DT = append(cell.DT, DT{created, updated, version})
		}
	}

	return nil
}

// totalCells returns the total number of cell line records in database.
func totalCells() (int, error) {
	var (
		count string
		total int
	)

	db, err := Database()
	defer db.Close()
	if err != nil {
		logSentry(err)
		return total, err
	}

	err = db.QueryRow("SELECT content FROM releaseInfo where attribute = 'totalCells';").Scan(&count)
	if err != nil {
		logSentry(err)
		return total, err
	}

	total, _ = strconv.Atoi(count)
	return total, nil
}
