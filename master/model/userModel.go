package model

type UserModel struct {
	IdUser       string `json:"id_user"`
	NamaUser     string `json:"nama_user"`
	TanggalLahir string `json:"tanggal_lahir"`
	NoKTP        string `json:"no_ktp"`
	Pekerjaan    string `json:"pekerjaan"`
	Pendidikan   string `json:"pendidikan"`
	Status       string `json:"status,omitempty"`
}
