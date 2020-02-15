package bst

type Stats struct {
	Count int
	Total int
	Rows  int
	Cols  int
}

func (t Tree) Stats() Stats {
	return Stats{
		Total: t.Total(),
		Count: t.Count(),
		Rows:  t.Rows(),
		Cols:  t.Cols(),
	}
}
