package myInterface

type CustomerSlice []Customer

// ---Syarat untuk menggunakan sort standard library golang adalah dengan menyesuaikan kontrak di sort.Interface---
func (c CustomerSlice) Len() int { return len(c) }

func (c CustomerSlice) Less(i, j int) bool { return c[i].Id < c[j].Id }

func (c CustomerSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
	// c[i] = c[i]
	// c[j] = c[j]
}
