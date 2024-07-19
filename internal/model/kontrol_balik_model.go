package model

type KontrolBalikResponse struct {
	ID             int             `json:"id"`
	IdPasien       int             `json:"idPasien,omitempty"`
	PasienResponse *PasienResponse `json:"pasien,omitempty"`
	TanggalKontrol int64           `json:"tanggalKontrol"`
	Status         string          `json:"status,omitempty"`
}

type KontrolBalikSearchRequest struct {
	IdPengguna       int    `validate:"omitempty,numeric"`
	IdAdminPuskesmas int    `validate:"omitempty,numeric"`
	Status           string `json:"status" validate:"omitempty,oneof=menunggu selesai batal"`
}
type KontrolBalikGetRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}
type KontrolBalikCreateRequest struct {
	IdPasien         int   `json:"idPasien" validate:"required,numeric"`
	TanggalKontrol   int64 `json:"tanggalKontrol" validate:"required,numeric"`
	IdAdminPuskesmas int   `validate:"omitempty,numeric"`
}
type KontrolBalikUpdateRequest struct {
	ID               int   `json:"id" validate:"required,numeric"`
	IdPasien         int   `json:"idPasien" validate:"required,numeric"`
	TanggalKontrol   int64 `json:"tanggalKontrol" validate:"required,numeric"`
	IdAdminPuskesmas int   `validate:"omitempty,numeric"`
}
type KontrolBalikDeleteRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}

type KontrolBalikSelesaiRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}

type KontrolBalikBatalRequest struct {
	ID               int `json:"id" validate:"required,numeric"`
	IdAdminPuskesmas int `validate:"omitempty,numeric"`
}
