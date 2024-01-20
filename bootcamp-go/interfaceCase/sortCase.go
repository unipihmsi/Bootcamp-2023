package interfaceCase

type Family []Keluarga

func (c Family) Len() int { return len(c) }

func (c Family) Less(i, j int) bool { return c[i].GetUmur() < c[j].GetUmur() }

func (c Family) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
