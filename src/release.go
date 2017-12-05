package cellosaurus

// Release models release specific information for the Cellosaurus.
type Release struct {
	Name        string      `json:"database"`
	Description string      `json:"description"`
	Stat        releaseStat `json:"release-information"`
}

type releaseStat struct {
	Version      string `json:"version"`
	Updated      string `json:"updated"`
	Total        string `json:"total"`
	Human        string `json:"human"`
	Mouse        string `json:"mouse"`
	Rat          string `json:"rat"`
	Publications string `json:"publications"`
}

// Create returns release information for the Cellosaurus.
func (rel *Release) Create() error {
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT attribute, content FROM rel_info;")
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
		case "total":
			rel.Stat.Total = content
		case "human":
			rel.Stat.Human = content
		case "mouse":
			rel.Stat.Mouse = content
		case "rat":
			rel.Stat.Rat = content
		case "publications":
			rel.Stat.Publications = content
		case "version":
			rel.Stat.Version = content
		case "updated":
			rel.Stat.Updated = content
		}
	}

	return nil
}
