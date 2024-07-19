package model

type PasienResponse struct {
	ID               int                     `json:"id"`
	NoRekamMedis     string                  `json:"noRekamMedis"`
	Pengguna         *PenggunaResponse       `json:"pengguna,omitempty"`
	IdPengguna       int                     `json:"idPengguna,omitempty"`
	AdminPuskesmas   *AdminPuskesmasResponse `json:"adminPuskesmas,omitempty"`
	IdAdminPuskesmas int                     `json:"idAdminPuskesmas,omitempty"`
	BeratBadan       float64                 `json:"beratBadan"`
	TinggiBadan      float64                 `json:"tinggiBadan"`
	TekananDarah     string                  `json:"tekananDarah"`
	DenyutNadi       int                     `json:"denyutNadi"`
	HasilLab         string                  `json:"hasilLab"`
	HasilEkg         string                  `json:"hasilEkg"`
	TanggalPeriksa   int64                   `json:"tanggalPeriksa"`
	Status           string                  `json:"status,omitempty"`
}

type PasienSearchRequest struct {
	IdPengguna       int    `validate:"omitempty,numeric"`
	IdAdminPuskesmas int    `validate:"omitempty,numeric"`
	Status           string `json:"status" validate:"omitempty,oneof=aktif selesai"`
}
type PasienGetRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdPengguna       int `validate:"omitempty,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}
type PasienCreateRequest struct {
	NoRekamMedis     string  `json:"noRekamMedis" mod:"normalize_spaces" validate:"required,min=3,max=50"`
	IdPengguna       int     `json:"idPengguna" validate:"required,numeric"`
	IdAdminPuskesmas int     `json:"idAdminPuskesmas" validate:"required,numeric"`
	BeratBadan       float64 `json:"beratBadan" validate:"required,numeric"`
	TinggiBadan      float64 `json:"tinggiBadan" validate:"required,numeric"`
	TekananDarah     string  `json:"tekananDarah" mod:"normalize_spaces" validate:"required,min=3,max=20"`
	DenyutNadi       int     `json:"denyutNadi" validate:"required,numeric"`
	HasilLab         string  `json:"hasilLab" mod:"normalize_spaces"`
	HasilEkg         string  `json:"hasilEkg" mod:"normalize_spaces"`
	TanggalPeriksa   int64   `json:"tanggalPeriksa" validate:"required,numeric"`
}
type PasienUpdateRequest struct {
	ID                    int     `json:"id" validate:"required,numeric"`
	NoRekamMedis          string  `json:"noRekamMedis" mod:"normalize_spaces" validate:"required,min=3,max=50"`
	IdPengguna            int     `json:"idPengguna" validate:"required,numeric"`
	CurrentAdminPuskesmas bool    `validate:"omitempty"`
	IdAdminPuskesmas      int     `json:"idAdminPuskesmas" validate:"required,numeric"`
	BeratBadan            float64 `json:"beratBadan" validate:"required,numeric"`
	TinggiBadan           float64 `json:"tinggiBadan" validate:"required,numeric"`
	TekananDarah          string  `json:"tekananDarah" mod:"normalize_spaces" validate:"required,min=3,max=20"`
	DenyutNadi            int     `json:"denyutNadi" validate:"required,numeric"`
	HasilLab              string  `json:"hasilLab" mod:"normalize_spaces"`
	HasilEkg              string  `json:"hasilEkg" mod:"normalize_spaces"`
	TanggalPeriksa        int64   `json:"tanggalPeriksa" validate:"required,numeric"`
}
type PasienDeleteRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}

type PasienSelesaiRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}
