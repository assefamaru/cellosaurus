package cellosaurus

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

// Count returns the total number of records in table.
func Count(table string) (int, error) {
	var count int
	db, err := Database()
	defer db.Close()
	if err != nil {
		return count, err
	}
	err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s;", table)).Scan(&count)
	if err != nil {
		LogSentry(err)
	}
	return count, err
}

// List returns a list of all cell lines, without pagination.
func (cells *Cells) List() error {
	var cell Cell
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT identifier, accession, `as` FROM cellosaurus;")
	defer rows.Close()
	if err != nil {
		LogSentry(err)
		return err
	}
	for rows.Next() {
		err = rows.Scan(&cell.ID, &cell.AC.Primary, &cell.AC.Secondary)
		if err != nil {
			LogSentry(err)
			return err
		}
		*cells = append(*cells, cell)
	}
	return nil
}

// ListPaginated returns a list of cell lines with pagination.
func (cells *Cells) ListPaginated(page int, limit int) error {
	var cell Cell
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}
	s := (page - 1) * limit
	rows, err := db.Query(fmt.Sprintf("SELECT identifier, accession, `as` FROM cellosaurus LIMIT %d,%d;", s, limit))
	defer rows.Close()
	if err != nil {
		LogSentry(err)
		return err
	}
	for rows.Next() {
		err = rows.Scan(&cell.ID, &cell.AC.Primary, &cell.AC.Secondary)
		if err != nil {
			LogSentry(err)
			return err
		}
		*cells = append(*cells, cell)
	}
	return nil
}

// Find returns a single cell line.
func (cell *Cell) Find(id string, typ string) error {
	var (
		query string

		sy sql.NullString // synonyms
		dr sql.NullString // cross-references
		rx sql.NullString // reference-identifiers
		ww sql.NullString // web-pages
		cc sql.NullString // comments
		di sql.NullString // diseases
		hi sql.NullString // hierarchy
		ox sql.NullString // species-of-origin
		oi sql.NullString // same-origin-as
		st sql.NullString // str-profile-data
	)
	strData := &ST{}

	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}

	SQL := "SELECT identifier, accession, `as`, ca, sx, sy, dr, rx, ww, cc, di, hi, ox, oi, st FROM cellosaurus"
	if typ == "accession" || typ == "identifier" {
		query = fmt.Sprintf("%s WHERE %s = '%s';", SQL, typ, id)
	} else if typ == "synonym" {
		query = fmt.Sprintf("%s WHERE sy LIKE '%s';", SQL, "%"+id+"%")
	} else {
		return errors.New("Unknown type")
	}
	row := db.QueryRow(query)
	err = row.Scan(&cell.ID, &cell.AC.Primary, &cell.AC.Secondary, &cell.CA, &cell.SX, &sy, &dr, &rx, &ww, &cc, &di, &hi, &ox, &oi, &st)
	if err != nil {
		if err != sql.ErrNoRows {
			LogSentry(err)
		}
		return err
	}

	if sy.Valid {
		cell.SY = strings.Split(sy.String, "; ")
	}

	if dr.Valid {
		for _, ref := range strings.Split(dr.String, " | ") {
			da := strings.Split(ref, "; ")
			var newdr DR
			newdr.Database = da[0]
			newdr.Accession = da[1]
			cell.DR = append(cell.DR, newdr)
		}
	}

	if rx.Valid {
		for _, r := range strings.Split(rx.String, " | ") {
			refiden := strings.TrimRight(r, ";")
			cell.RX = append(cell.RX, refiden)
		}
	}

	if ww.Valid {
		cell.WW = strings.Split(ww.String, " | ")
	}

	if cc.Valid {
		for _, comment := range strings.Split(cc.String, " | ") {
			com := strings.Split(comment, ": ")
			var newcc CC
			newcc.Category = com[0]
			newcc.Comment = com[1]
			cell.CC = append(cell.CC, newcc)
		}
	}

	if di.Valid {
		for _, d := range strings.Split(di.String, " | ") {
			disease := strings.Split(d, "; ")
			var newdi DI
			newdi.Terminology = disease[0]
			newdi.Accession = disease[1]
			newdi.Disease = disease[2]
			cell.DI = append(cell.DI, newdi)
		}
	}

	if hi.Valid {
		for _, hierarchy := range strings.Split(hi.String, " | ") {
			h := strings.Split(hierarchy, " ! ")
			var newhi HI
			newhi.Terminology = "Cellosaurus"
			newhi.Accession = h[0]
			newhi.DF = h[1]
			cell.HI = append(cell.HI, newhi)
		}
	}

	if ox.Valid {
		for _, o := range strings.Split(ox.String, " | ") {
			f := strings.Split(o, " ! ")
			s := strings.Split(strings.TrimRight(f[0], ";"), "=")
			var newox OX
			newox.Terminology = "NCBI-Taxonomy"
			newox.Accession = s[1]
			newox.Species = f[1]
			cell.OX = append(cell.OX, newox)
		}
	}

	if oi.Valid {
		for _, j := range strings.Split(oi.String, " | ") {
			f := strings.Split(j, " ! ")
			var newoi OI
			newoi.Terminology = "Cellosaurus"
			newoi.Accession = f[0]
			newoi.Identifier = f[1]
			cell.OI = append(cell.OI, newoi)
		}
	}

	if st.Valid {
		var sources []string
		for i, str := range strings.Split(st.String, " | ") {
			if i == 0 {
				sources = strings.Split(strings.TrimLeft(str, "Source(s): "), "; ")
				continue
			}
			f := strings.Split(str, ": ")
			var newmarker Marker
			newmarker.ID = f[0]
			newmarker.Alleles = f[1]
			strData.Markers = append(strData.Markers, newmarker)
		}
		strData.Sources = sources
		cell.ST = strData
	}

	return nil
}
