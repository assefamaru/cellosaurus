package cellosaurus

import (
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
}

// Cells is a list of cell lines.
type Cells struct {
	Meta Meta   `json:"meta"`
	Data []Cell `json:"data"`
}

// Meta contains pagination information for returned data.
type Meta struct {
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
		cell.AC.Sec = strings.Split(acs, "; ")
		cell.SY = strings.Split(sy, "; ")
		cells.Data = append(cells.Data, cell)
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
		return total, err
	}

	err = db.QueryRow("SELECT content FROM releaseInfo where attribute = 'total';").Scan(&count)
	if err != nil {
		logSentry(err)
		return total, err
	}

	total, _ = strconv.Atoi(count)
	return total, nil
}
