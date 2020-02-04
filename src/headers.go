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
	Total        int    `json:"total"`
	Human        int    `json:"human"`
	Mouse        int    `json:"mouse"`
	Rat          int    `json:"rat"`
	Publications int    `json:"publications"`
	References   int    `json:"references"`
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
		case "publications":
			rel.Stat.Publications, _ = strconv.Atoi(content)
		case "totalReferences":
			rel.Stat.References, _ = strconv.Atoi(content)
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
	Name        string `json:"name"`
	Source      string `json:"source"`
	Description string `json:"description"`
	URL         string `json:"url"`
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

	rows, err := db.Query("SELECT name, source, description, url FROM terminologies;")
	defer rows.Close()
	if err != nil {
		logSentry(err)
		return err
	}
	for rows.Next() {
		var term Terminology
		err := rows.Scan(&term.Name, &term.Source, &term.Description, &term.URL)
		if err != nil {
			logSentry(err)
			return err
		}
		*terms = append(*terms, term)
	}

	return nil
}