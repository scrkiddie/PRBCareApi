package model

type PasienResponse struct {
	ID               int32                   `json:"id"`
	NoRekamMedis     string                  `json:"noRekamMedis"`
	Pengguna         *PenggunaResponse       `json:"pengguna,omitempty"`
	IdPengguna       int32                   `json:"idPengguna,omitempty"`
	AdminPuskesmas   *AdminPuskesmasResponse `json:"adminPuskesmas,omitempty"`
	IdAdminPuskesmas int32                   `json:"idAdminPuskesmas,omitempty"`
	BeratBadan       int32                   `json:"beratBadan"`
	TinggiBadan      int32                   `json:"tinggiBadan"`
	TekananDarah     string                  `json:"tekananDarah"`
	DenyutNadi       int32                   `json:"denyutNadi"`
	HasilLab         string                  `json:"hasilLab"`
	HasilEkg         string                  `json:"hasilEkg"`
	TanggalPeriksa   int64                   `json:"tanggalPeriksa"`
	Status           string                  `json:"status,omitempty"`
}

type PasienSearchRequest struct {
	IdPengguna       int32  `validate:"omitempty,numeric"`
	IdAdminPuskesmas int32  `validate:"omitempty,numeric"`
	Status           string `json:"status" validate:"omitempty,oneof=aktif selesai"`
}
type PasienGetRequest struct {
	ID               int32 `json:"id" validate:"required,numeric"`
	IdPengguna       int32 `validate:"omitempty,numeric"`
	IdAdminPuskesmas int32 `validate:"omitempty,numeric"`
}
type PasienCreateRequest struct {
	NoRekamMedis     string `json:"noRekamMedis" mod:"normalize_spaces" validate:"required,min=3,max=50"`
	IdPengguna       int32  `json:"idPengguna" validate:"required,numeric"`
	IdAdminPuskesmas int32  `json:"idAdminPuskesmas" validate:"required,numeric"`
	BeratBadan       int32  `json:"beratBadan" validate:"required,numeric,gt=0"`
	TinggiBadan      int32  `json:"tinggiBadan" validate:"required,numeric,gt=0"`
	TekananDarah     string `json:"tekananDarah" mod:"normalize_spaces" validate:"required,min=3,max=20"`
	DenyutNadi       int32  `json:"denyutNadi" validate:"required,numeric,gt=0"`
	HasilLab         string `json:"hasilLab" mod:"normalize_spaces"`
	HasilEkg         string `json:"hasilEkg" mod:"normalize_spaces"`
	TanggalPeriksa   int64  `json:"tanggalPeriksa" validate:"required,numeric"`
}
type PasienUpdateRequest struct {
	ID                    int32  `json:"id" validate:"required,numeric"`
	NoRekamMedis          string `json:"noRekamMedis" mod:"normalize_spaces" validate:"required,min=3,max=50"`
	IdPengguna            int32  `json:"idPengguna" validate:"required,numeric"`
	CurrentAdminPuskesmas bool   `validate:"omitempty"`
	IdAdminPuskesmas      int32  `json:"idAdminPuskesmas" validate:"required,numeric"`
	BeratBadan            int32  `json:"beratBadan" validate:"required,numeric,gt=0"`
	TinggiBadan           int32  `json:"tinggiBadan" validate:"required,numeric,gt=0"`
	TekananDarah          string `json:"tekananDarah" mod:"normalize_spaces" validate:"required,min=3,max=20"`
	DenyutNadi            int32  `json:"denyutNadi" validate:"required,numeric,gt=0"`
	HasilLab              string `json:"hasilLab" mod:"normalize_spaces"`
	HasilEkg              string `json:"hasilEkg" mod:"normalize_spaces"`
	TanggalPeriksa        int64  `json:"tanggalPeriksa" validate:"required,numeric"`
}
type PasienDeleteRequest struct {
	ID               int32 `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int32 `validate:"omitempty,numeric"`
}

type PasienSelesaiRequest struct {
	ID               int32 `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int32 `validate:"omitempty,numeric"`
}
