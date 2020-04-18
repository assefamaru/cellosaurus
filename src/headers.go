package cellosaurus

import "strconv"

// Release models release specific information for the Cellosaurus.
type Release struct {
	Name        string  `json:"database"`
	Description string  `json:"description"`
	Stat        relStat `json:"release-information"`
}

type relStat struct {
	Version      string `json:"version"`
	Updated      string `json:"updated"`
	Total        int    `json:"cell-lines"`
	Human        int    `json:"human"`
	Mouse        int    `json:"mouse"`
	Rat          int    `json:"rat"`
	Species      int    `json:"species"`
	Synonyms     int    `json:"synonyms"`
	CrossRefs    int    `json:"cross-references"`
	References   int    `json:"references"`
	DistinctRefs int    `json:"distinct-references"`
	WebLinks     int    `json:"web-links"`
	CellsWithSTR int    `json:"cells-with-str-profiles"`
}

// Create returns release information for the Cellosaurus.
func (rel *Release) Create() error {
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT attribute, content FROM releaseInfo;")
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
		case "name":
			rel.Name = content
		case "description":
			rel.Description = content
		case "totalCells":
			rel.Stat.Total, _ = strconv.Atoi(content)
		case "human":
			rel.Stat.Human, _ = strconv.Atoi(content)
		case "mouse":
			rel.Stat.Mouse, _ = strconv.Atoi(content)
		case "rat":
			rel.Stat.Rat, _ = strconv.Atoi(content)
		case "species":
			rel.Stat.Species, _ = strconv.Atoi(content)
		case "synonyms":
			rel.Stat.Synonyms, _ = strconv.Atoi(content)
		case "crossReferences":
			rel.Stat.CrossRefs, _ = strconv.Atoi(content)
		case "references":
			rel.Stat.References, _ = strconv.Atoi(content)
		case "distinctRefs":
			rel.Stat.DistinctRefs, _ = strconv.Atoi(content)
		case "webLinks":
			rel.Stat.WebLinks, _ = strconv.Atoi(content)
		case "cellsWithSTR":
			rel.Stat.CellsWithSTR, _ = strconv.Atoi(content)
		case "version":
			rel.Stat.Version = content
		case "updated":
			rel.Stat.Updated = content
		}
	}

	return nil
}

// Terminology models terminology information contained in database.
type Terminology struct {
	Abbreviation string `json:"Abbreviation"`
	Name         string `json:"Name"`
	Server       string `json:"Server"`
	DbURL        string `json:"Db_URL"`
	Cat          string `json:"Cat"`
}

// Terminologies is a list of Terminology.
type Terminologies []Terminology

// List returns a list of terminologies.
func (terms *Terminologies) List() error {
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT abbreviation, name, server, db_url, cat FROM terminologies;")
	defer rows.Close()
	if err != nil {
		logSentry(err)
		return err
	}
	for rows.Next() {
		var term Terminology
		err := rows.Scan(&term.Abbreviation, &term.Name, &term.Server, &term.DbURL, &term.Cat)
		if err != nil {
			logSentry(err)
			return err
		}
		*terms = append(*terms, term)
	}

	return nil
}
