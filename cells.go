package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// IndexCell returns a list of all cell lines in database (paginated by default).
func IndexCell(c *gin.Context) {
	var (
		cell  Cell
		cells []Cell
	)

	db, err := initDB()
	defer db.Close()
	if err != nil {
		handleError(c, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	shouldIndent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true"))

	if isTrue, _ := strconv.ParseBool(c.DefaultQuery("all", "false")); isTrue {
		rows, er := db.Query("SELECT id, identifier, accession, `as` FROM cells;")
		defer rows.Close()
		if er != nil {
			handleError(c, er, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		for rows.Next() {
			err = rows.Scan(&cell.ID, &cell.Identifier, &cell.Accession.Primary, &cell.Accession.Secondary)
			if err != nil {
				handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
				return
			}
			cells = append(cells, cell)
		}
		if shouldIndent {
			c.IndentedJSON(http.StatusOK, gin.H{
				"data":        cells,
				"total":       len(cells),
				"description": "List of all cell lines in Cellosaurus",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":        cells,
				"total":       len(cells),
				"description": "List of all cell lines in Cellosaurus",
			})
		}
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("per_page", "30"))

	s := (page - 1) * limit
	SQL := "SELECT id, identifier, accession, `as` FROM cells"
	query := fmt.Sprintf("%s limit %d,%d", SQL, s, limit)
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	for rows.Next() {
		err = rows.Scan(&cell.ID, &cell.Identifier, &cell.Accession.Primary, &cell.Accession.Secondary)
		if err != nil {
			handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		cells = append(cells, cell)
	}
	row := db.QueryRow("SELECT COUNT(*) FROM cells;")
	var total int
	err = row.Scan(&total)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// Write pagination links in response header.
	writeHeaderLinks(c, "/cell_lines", page, total, limit)

	if shouldIndent {
		c.IndentedJSON(http.StatusOK, gin.H{
			"data":        cells,
			"total":       total,
			"page":        page,
			"per_page":    limit,
			"description": "List of all cell lines in Cellosaurus",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":        cells,
			"total":       total,
			"page":        page,
			"per_page":    limit,
			"description": "List of all cell lines in Cellosaurus",
		})
	}
}

// ShowCell returns a cell line using ID, Identifier or Accession.
func ShowCell(c *gin.Context) {
	var (
		cell Cell

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

	db, err := initDB()
	defer db.Close()
	if err != nil {
		handleError(c, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	id := c.Param("id")
	searchType := c.DefaultQuery("type", "identifier")

	SQL := "SELECT id, identifier, accession, `as`, ca, sx, sy, dr, rx, ww, cc, di, hi, ox, oi, st FROM cells"
	query := fmt.Sprintf("%s WHERE %s = ?;", SQL, searchType)
	row := db.QueryRow(query, id)
	err = row.Scan(&cell.ID, &cell.Identifier, &cell.Accession.Primary, &cell.Accession.Secondary, &cell.CA, &cell.SX, &sy, &dr, &rx, &ww, &cc, &di, &hi, &ox, &oi, &st)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(c, nil, http.StatusNotFound, "Cell line not found in database")
		} else {
			handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
		}
		return
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

	if shouldIndent, _ := strconv.ParseBool(c.DefaultQuery("indent", "true")); shouldIndent {
		c.IndentedJSON(http.StatusOK, cell)
	} else {
		c.JSON(http.StatusOK, cell)
	}
}
