package cellosaurus

// Cell is a cell line model.
type Cell struct {
	ID string `json:"identifier"`
	AC string `json:"accession"`
}

// Cells is an array of cell lines.
type Cells []Cell
