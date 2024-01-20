package interfaceCase

type Keluarga interface {
	GetNama() string
	GetUmur() uint
}

type Anggota struct {
	Nama string
	Umur uint
}

func (a Anggota) GetNama() string { return a.Nama }

func (a Anggota) GetUmur() uint { return a.Umur }
