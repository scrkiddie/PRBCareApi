package model

type ObatResponse struct {
	ID            int                  `json:"id"`
	IdAdminApotek int                  `json:"idAdminApotek,omitempty"`
	AdminApotek   *AdminApotekResponse `json:"adminApotek,omitempty"`
	NamaObat      string               `json:"namaObat"`
	Jumlah        int                  `json:"jumlah"`
}

type ObatListRequest struct {
	IdAdminApotek int `validate:"omitempty,numeric"`
}
type ObatGetRequest struct {
	ID            int `json:"id" validate:"required,numeric"`
	IdAdminApotek int `validate:"omitempty,numeric"`
}
type ObatCreateRequest struct {
	NamaObat      string `json:"namaObat" mod:"normalize_spaces" validate:"required,min=3,max=50"`
	Jumlah        int    `json:"jumlah" validate:"required,numeric,gt=0"`
	IdAdminApotek int    `validate:"required,numeric"`
}
type ObatUpdateRequest struct {
	ID                 int    `json:"id" validate:"required,numeric"`
	NamaObat           string `json:"namaObat" mod:"normalize_spaces" validate:"required,min=3,max=50"`
	Jumlah             int    `json:"jumlah" validate:"required,numeric,gt=0"`
	CurrentAdminApotek bool   `validate:"omitempty"`
	IdAdminApotek      int    `json:"idAdminApotek" validate:"required,numeric"`
}
type ObatDeleteRequest struct {
	ID            int `json:"id" validate:"required,numeric"`
	IdAdminApotek int `validate:"omitempty,numeric"`
}
