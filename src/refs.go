package cellosaurus

import (
	"fmt"
	"strconv"
	"strings"
)

// Reference contains reference data for the Cellosaurus.
type Reference struct {
	RX []string `json:"identifier"`
	RA []string `json:"authors,omitempty"`
	RG []string `json:"group/consortium,omitempty"`
	RT string   `json:"title"`
	RL string   `json:"citation"`
}

// References is a list of Reference.
type References struct {
	Meta refMeta     `json:"meta"`
	Data []Reference `json:"data"`
}

type refMeta struct {
	Page     int `json:"page"`
	PerPage  int `json:"per-page"`
	LastPage int `json:"last-page"`
	Total    int `json:"total"`
}

// List returns a list of paginated references.
func (refs *References) List() error {
	db, err := Database()
	defer db.Close()
	if err != nil {
		return err
	}

	from := (refs.Meta.Page - 1) * refs.Meta.PerPage
	to := refs.Meta.PerPage
	query := fmt.Sprintf("SELECT rx, ra, rg, rt, rl FROM refs LIMIT %d,%d;", from, to)

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		logSentry(err)
		return err
	}
	for rows.Next() {
		var (
			rx  string
			ra  string
			rg  string
			rt  string
			ref Reference
		)
		if err = rows.Scan(&rx, &ra, &rg, &rt, &ref.RL); err != nil {
			logSentry(err)
			return err
		}
		ref.RX = strings.Split(strings.TrimSuffix(rx, ";"), "; ")
		if ra != "" {
			ref.RA = strings.Split(strings.TrimSuffix(ra, ";"), ", ")
		}
		if rg != "" {
			ref.RG = strings.Split(rg, ";")
		}
		if rt != "" {
			ref.RT = strings.TrimSuffix(strings.TrimPrefix(rt, " \""), "\";")
		}
		refs.Data = append(refs.Data, ref)
	}

	return nil
}

// totalRefs returns the total number of reference records in database.
func totalRefs() (int, error) {
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

	err = db.QueryRow("SELECT content FROM releaseInfo where attribute = 'totalReferences';").Scan(&count)
	if err != nil {
		logSentry(err)
		return total, err
	}

	total, _ = strconv.Atoi(count)
	return total, nil
}
