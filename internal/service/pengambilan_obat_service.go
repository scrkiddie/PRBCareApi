package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"log"
	"prb_care_api/internal/constant"
	"prb_care_api/internal/entity"
	"prb_care_api/internal/model"
	"prb_care_api/internal/repository"
)

type PengambilanObatService struct {
	DB                        *gorm.DB
	PengambilanObatRepository *repository.PengambilanObatRepository
	PasienRepository          *repository.PasienRepository
	ObatRepository            *repository.ObatRepository
	Validator                 *validator.Validate
}

func NewPengambilanObatService(
	db *gorm.DB,
	pengambilanObatRepository *repository.PengambilanObatRepository,
	pasienRepository *repository.PasienRepository,
	obatRepository *repository.ObatRepository,
	validator *validator.Validate,
) *PengambilanObatService {
	return &PengambilanObatService{db, pengambilanObatRepository, pasienRepository, obatRepository, validator}
}

func (s *PengambilanObatService) Search(ctx context.Context, request *model.PengambilanObatSearchRequest) (*[]model.PengambilanObatResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	pengambilanObat := new([]entity.PengambilanObat)
	if request.IdPengguna > 0 {
		if err := s.PengambilanObatRepository.SearchAsPengguna(tx, pengambilanObat, request.IdPengguna, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else if request.IdAdminPuskesmas > 0 {
		if err := s.PengambilanObatRepository.SearchAsAdminPuskesmas(tx, pengambilanObat, request.IdAdminPuskesmas, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else if request.IdAdminApotek > 0 {
		if err := s.PengambilanObatRepository.SearchAsAdminApotek(tx, pengambilanObat, request.IdAdminApotek, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else {
		if err := s.PengambilanObatRepository.Search(tx, pengambilanObat, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	}

	var response []model.PengambilanObatResponse
	for _, perPengambilanObat := range *pengambilanObat {
		response = append(response, model.PengambilanObatResponse{
			ID:   perPengambilanObat.ID,
			Resi: perPengambilanObat.Resi,
			PasienResponse: &model.PasienResponse{
				ID:           perPengambilanObat.Pasien.ID,
				NoRekamMedis: perPengambilanObat.Pasien.NoRekamMedis,
				Pengguna: &model.PenggunaResponse{
					ID:              perPengambilanObat.Pasien.Pengguna.ID,
					NamaLengkap:     perPengambilanObat.Pasien.Pengguna.NamaLengkap,
					Telepon:         perPengambilanObat.Pasien.Pengguna.Telepon,
					TeleponKeluarga: perPengambilanObat.Pasien.Pengguna.TeleponKeluarga,
					Alamat:          perPengambilanObat.Pasien.Pengguna.Alamat,
				},
				AdminPuskesmas: &model.AdminPuskesmasResponse{
					ID:            perPengambilanObat.Pasien.AdminPuskesmas.ID,
					NamaPuskesmas: perPengambilanObat.Pasien.AdminPuskesmas.NamaPuskesmas,
					Telepon:       perPengambilanObat.Pasien.AdminPuskesmas.Telepon,
					Alamat:        perPengambilanObat.Pasien.AdminPuskesmas.Alamat,
				},
				BeratBadan:     perPengambilanObat.Pasien.BeratBadan,
				TinggiBadan:    perPengambilanObat.Pasien.TinggiBadan,
				TekananDarah:   perPengambilanObat.Pasien.TekananDarah,
				DenyutNadi:     perPengambilanObat.Pasien.DenyutNadi,
				HasilLab:       perPengambilanObat.Pasien.HasilLab,
				HasilEkg:       perPengambilanObat.Pasien.HasilEkg,
				TanggalPeriksa: perPengambilanObat.Pasien.TanggalPeriksa,
				Status:         perPengambilanObat.Pasien.Status,
			},
			Obat: &model.ObatResponse{
				ID:       perPengambilanObat.Obat.ID,
				NamaObat: perPengambilanObat.Obat.NamaObat,
				Jumlah:   perPengambilanObat.Obat.Jumlah,
				AdminApotek: &model.AdminApotekResponse{
					ID:         perPengambilanObat.Obat.AdminApotek.ID,
					NamaApotek: perPengambilanObat.Obat.AdminApotek.NamaApotek,
					Telepon:    perPengambilanObat.Obat.AdminApotek.Telepon,
					Alamat:     perPengambilanObat.Obat.AdminApotek.Alamat,
				},
			},
			Jumlah:             perPengambilanObat.Jumlah,
			TanggalPengambilan: perPengambilanObat.TanggalPengambilan,
			Status:             perPengambilanObat.Status,
		})
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &response, nil
}

func (s *PengambilanObatService) Get(ctx context.Context, request *model.PengambilanObatGetRequest) (*model.PengambilanObatResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	pengambilanObat := new(entity.PengambilanObat)
	if request.IdAdminPuskesmas > 0 {
		if err := s.PengambilanObatRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, pengambilanObat, request.ID, request.IdAdminPuskesmas, constant.StatusPengambilanObatMenunggu); err != nil {
			log.Println(err.Error())
			return nil, fiber.NewError(fiber.StatusNotFound)
		}
	} else if err := s.PengambilanObatRepository.FindByIdAndStatus(tx, pengambilanObat, request.ID, constant.StatusPengambilanObatMenunggu); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound)
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	response := new(model.PengambilanObatResponse)
	response.ID = pengambilanObat.ID
	response.IdObat = pengambilanObat.IdObat
	response.IdPasien = pengambilanObat.IdPasien
	response.Jumlah = pengambilanObat.Jumlah
	response.TanggalPengambilan = pengambilanObat.TanggalPengambilan
	return response, nil
}

func (s *PengambilanObatService) Create(ctx context.Context, request *model.PengambilanObatCreateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	if request.IdAdminPuskesmas > 0 {
		if err := s.PasienRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, &entity.Pasien{}, request.IdPasien, request.IdAdminPuskesmas, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	} else {
		if err := s.PasienRepository.FindByIdAndStatus(tx, &entity.Pasien{}, request.IdPasien, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	}

	obat := new(entity.Obat)
	if err := s.ObatRepository.FindById(tx, obat, request.IdObat); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound)
	}
	obat.Jumlah -= request.Jumlah
	if obat.Jumlah < 0 {
		return fiber.NewError(fiber.StatusConflict, "Jumlah obat melebihi persediaan apotek")
	}

	pengambilanObatEntity := new(entity.PengambilanObat)
	pengambilanObatEntity.Resi = fmt.Sprintf("%d%d", request.IdPasien, request.TanggalPengambilan)
	pengambilanObatEntity.IdPasien = request.IdPasien
	pengambilanObatEntity.IdObat = request.IdObat
	pengambilanObatEntity.Jumlah = request.Jumlah
	pengambilanObatEntity.TanggalPengambilan = request.TanggalPengambilan
	pengambilanObatEntity.Status = constant.StatusPengambilanObatMenunggu

	if err := s.ObatRepository.Update(tx, obat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := s.PengambilanObatRepository.Create(tx, pengambilanObatEntity); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *PengambilanObatService) Update(ctx context.Context, request *model.PengambilanObatUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	// harus status menunggu
	pengambilanObat := new(entity.PengambilanObat)
	if request.IdAdminPuskesmas > 0 {
		if err := s.PengambilanObatRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, pengambilanObat, request.ID, request.IdAdminPuskesmas, constant.StatusPengambilanObatMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	} else {
		if err := s.PengambilanObatRepository.FindByIdAndStatus(tx, pengambilanObat, request.ID, constant.StatusPengambilanObatMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	}
	// harus pasien status aktif
	if request.IdAdminPuskesmas > 0 {
		if err := s.PasienRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, &entity.Pasien{}, request.IdPasien, request.IdAdminPuskesmas, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	} else {
		if err := s.PasienRepository.FindByIdAndStatus(tx, &entity.Pasien{}, request.IdPasien, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	}

	obatNew := new(entity.Obat)
	obatOld := new(entity.Obat)

	if err := s.ObatRepository.FindById(tx, obatNew, request.IdObat); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound)
	}

	if pengambilanObat.IdObat == request.IdObat {
		obatNew.Jumlah = (obatNew.Jumlah + pengambilanObat.Jumlah) - request.Jumlah
		if obatNew.Jumlah < 0 {
			return fiber.NewError(fiber.StatusConflict, "Jumlah obat melebihi persediaan apotek")
		}
	} else {
		if err := s.ObatRepository.FindById(tx, obatOld, pengambilanObat.IdObat); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
		obatOld.Jumlah += pengambilanObat.Jumlah
		obatNew.Jumlah -= request.Jumlah
		if obatNew.Jumlah < 0 {
			return fiber.NewError(fiber.StatusConflict, "Jumlah obat melebihi persediaan apotek")
		}
		if err := s.ObatRepository.Update(tx, obatOld); err != nil {
			log.Println(err.Error())
			return fiber.ErrInternalServerError
		}
	}

	if err := s.ObatRepository.Update(tx, obatNew); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	pengambilanObat.Resi = fmt.Sprintf("%d%d", request.IdPasien, request.TanggalPengambilan)
	pengambilanObat.IdPasien = request.IdPasien
	pengambilanObat.IdObat = request.IdObat
	pengambilanObat.Jumlah = request.Jumlah
	pengambilanObat.TanggalPengambilan = request.TanggalPengambilan

	if err := s.PengambilanObatRepository.Update(tx, pengambilanObat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *PengambilanObatService) Delete(ctx context.Context, request *model.PengambilanObatDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	// harus status batal atau diambil
	pengambilanObat := new(entity.PengambilanObat)
	if request.IdAdminPuskesmas > 0 {
		if err := s.PengambilanObatRepository.FindByIdAndIdAdminPuskesmasAndStatusOrStatus(tx, pengambilanObat, request.ID, request.IdAdminPuskesmas, constant.StatusPengambilanObatBatal, constant.StatusPengambilanObatDiambil); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	} else {
		if err := s.PengambilanObatRepository.FindByIdAndStatusOrStatus(tx, pengambilanObat, request.ID, constant.StatusPengambilanObatBatal, constant.StatusPengambilanObatDiambil); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	}

	if err := s.PengambilanObatRepository.Delete(tx, pengambilanObat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *PengambilanObatService) Batal(ctx context.Context, request *model.PengambilanObatBatalRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	// harus status menunggu
	pengambilanObat := new(entity.PengambilanObat)
	if request.IdAdminPuskesmas > 0 {
		if err := s.PengambilanObatRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, pengambilanObat, request.ID, request.IdAdminPuskesmas, constant.StatusPengambilanObatMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	} else {
		if err := s.PengambilanObatRepository.FindByIdAndStatus(tx, pengambilanObat, request.ID, constant.StatusPengambilanObatMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	}
	obat := new(entity.Obat)
	if err := s.ObatRepository.FindById(tx, obat, pengambilanObat.IdObat); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound)
	}
	obat.Jumlah += pengambilanObat.Jumlah

	pengambilanObat.Status = constant.StatusPengambilanObatBatal

	if err := s.ObatRepository.Update(tx, obat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := s.PengambilanObatRepository.Update(tx, pengambilanObat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *PengambilanObatService) Diambil(ctx context.Context, request *model.PengambilanObatDiambilRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	// harus status menunggu
	pengambilanObat := new(entity.PengambilanObat)
	if request.IdAdminApotek > 0 {
		if err := s.PengambilanObatRepository.FindByIdAndIdAdminApotekAndStatus(tx, pengambilanObat, request.ID, request.IdAdminApotek, constant.StatusPengambilanObatMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	} else {
		if err := s.PengambilanObatRepository.FindByIdAndStatus(tx, pengambilanObat, request.ID, constant.StatusPengambilanObatMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound)
		}
	}

	pengambilanObat.Status = constant.StatusPengambilanObatDiambil

	if err := s.PengambilanObatRepository.Update(tx, pengambilanObat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}
