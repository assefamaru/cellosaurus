package cellosaurus

// Cell is a cell line model.
type Cell struct {
	ID string    `json:"identifier"`
	AC Accession `json:"accession"`
	SY []string  `json:"synonyms,omitempty"`
	CA string    `json:"category,omitempty"`
	SX string    `json:"sex,omitempty"`
}

// Cells is an array of cell lines.
type Cells []Cell

// Accession contains primary and secondary accession number(s).
type Accession struct {
	Pri string   `json:"primary"`
	Sec []string `json:"secondary,omitempty"`
}

// Count returns the total number of cell line records in database.
// func Count() (int, error) {
// 	var (
// 		count string
// 		total int
// 	)
//
// 	db, err := Database()
// 	defer db.Close()
// 	if err != nil {
// 		return total, err
// 	}
//
// 	err = db.QueryRow("SELECT content FROM stats where attribute = 'total';").Scan(&count)
// 	if err != nil {
// 		logSentry(err)
// 		return total, err
// 	}
//
// 	total, _ = strconv.Atoi(count)
// 	return total, nil
// }

// List returns a list of paginated cell lines.
// func (cells *Cells) List(page int, perPage int) error {
// 	db, err := Database()
// 	defer db.Close()
// 	if err != nil {
// 		return err
// 	}
//
// 	from := (page - 1) * perPage
// 	to := perPage
// 	query := fmt.Sprintf("SELECT identifier, accession FROM cells LIMIT %d,%d;", from, to)
//
// 	rows, err := db.Query(query)
// 	defer rows.Close()
// 	if err != nil {
// 		logSentry(err)
// 		return err
// 	}
// 	for rows.Next() {
// 		var cell Cell
// 		if err = rows.Scan(&cell.ID, &cell.AC); err != nil {
// 			logSentry(err)
// 			return err
// 		}
// 		*cells = append(*cells, cell)
// 	}
//
// 	return nil
// }
