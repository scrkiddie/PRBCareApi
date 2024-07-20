package model

type PengambilanObatResponse struct {
	ID                 int             `json:"id"`
	Resi               string          `json:"resi,omitempty"`
	IdPasien           int             `json:"idPasien,omitempty"`
	PasienResponse     *PasienResponse `json:"pasien,omitempty"`
	IdObat             int             `json:"idObat,omitempty"`
	Obat               *ObatResponse   `json:"obat,omitempty"`
	TanggalPengambilan int64           `json:"tanggalPengambilan"`
	Jumlah             int             `json:"jumlah"`
	Status             string          `json:"status,omitempty"`
}

type PengambilanObatSearchRequest struct {
	IdPengguna       int    `validate:"omitempty,numeric"`
	IdAdminPuskesmas int    `validate:"omitempty,numeric"`
	IdAdminApotek    int    `validate:"omitempty,numeric"`
	Status           string `json:"status" validate:"omitempty,oneof=menunggu diambil batal"`
}

type PengambilanObatGetRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}
type PengambilanObatCreateRequest struct {
	IdPasien           int   `json:"idPasien" validate:"required,numeric"`
	IdObat             int   `json:"idObat" validate:"required,numeric"`
	Jumlah             int   `json:"jumlah" validate:"required,numeric,gt=0"`
	TanggalPengambilan int64 `json:"tanggalPengambilan" validate:"required,numeric"`
	IdAdminPuskesmas   int   `validate:"omitempty,numeric"`
}
type PengambilanObatUpdateRequest struct {
	ID                 int   `json:"id" validate:"required,numeric"`
	IdPasien           int   `json:"idPasien" validate:"required,numeric"`
	IdObat             int   `json:"idObat" validate:"required,numeric"`
	Jumlah             int   `json:"jumlah" validate:"required,numeric,gt=0"`
	TanggalPengambilan int64 `json:"tanggalPengambilan" validate:"required,numeric"`
	IdAdminPuskesmas   int   `validate:"omitempty,numeric"`
}
type PengambilanObatDeleteRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}

type PengambilanObatDiambilRequest struct {
	ID            int `json:"id" validate:"required,numeric"`
	IdAdminApotek int `validate:"omitempty,numeric"`
}

type PengambilanObatBatalRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}
