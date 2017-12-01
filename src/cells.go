package cellosaurus

import "fmt"

// Cell is a cell line model.
type Cell struct {
	ID string `json:"identifier"`
	AC string `json:"accession"`
}

// Cells is an array of cell lines.
type Cells []Cell

// Count returns the total number of cell line records in database.
func Count() (int, error) {
	var count int

	db, err := Database()
	defer db.Close()
	if err != nil {
		return count, err
	}

	err = db.QueryRow("SELECT COUNT(*) FROM cells;").Scan(&count)
	if err != nil {
		logSentry(err)
	}

	return count, err
}

// List returns a list of paginated cell lines.
func (cells *Cells) List(page int, perPage int) error {
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}

	from := (page - 1) * perPage
	to := perPage
	query := fmt.Sprintf("SELECT identifier, accession FROM cells LIMIT %d,%d;", from, to)

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		logSentry(err)
		return err
	}
	for rows.Next() {
		var cell Cell
		if err = rows.Scan(&cell.ID, &cell.AC); err != nil {
			logSentry(err)
			return err
		}
		*cells = append(*cells, cell)
	}

	return nil
}
