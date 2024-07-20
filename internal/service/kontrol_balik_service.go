package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"log"
	"prbcare_be/internal/constant"
	"prbcare_be/internal/entity"
	"prbcare_be/internal/model"
	"prbcare_be/internal/repository"
)

type KontrolBalikService struct {
	DB                     *gorm.DB
	KontrolBalikRepository *repository.KontrolBalikRepository
	PasienRepository       *repository.PasienRepository
	Validator              *validator.Validate
}

func NewKontrolBalikService(
	db *gorm.DB,
	kontrolBalikRepository *repository.KontrolBalikRepository,
	pasienRepository *repository.PasienRepository,
	validator *validator.Validate,
) *KontrolBalikService {
	return &KontrolBalikService{db, kontrolBalikRepository, pasienRepository, validator}
}

func (s *KontrolBalikService) Search(ctx context.Context, request *model.KontrolBalikSearchRequest) (*[]model.KontrolBalikResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	kontrolBalik := new([]entity.KontrolBalik)
	if request.IdPengguna > 0 {
		if err := s.KontrolBalikRepository.SearchAsPengguna(tx, kontrolBalik, request.IdPengguna, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else if request.IdAdminPuskesmas > 0 {
		if err := s.KontrolBalikRepository.SearchAsAdminPuskesmas(tx, kontrolBalik, request.IdAdminPuskesmas, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else {
		if err := s.KontrolBalikRepository.Search(tx, kontrolBalik, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	}

	var response []model.KontrolBalikResponse
	for _, perKontrolBalik := range *kontrolBalik {
		response = append(response, model.KontrolBalikResponse{
			ID: perKontrolBalik.ID,
			PasienResponse: &model.PasienResponse{
				ID:           perKontrolBalik.Pasien.ID,
				NoRekamMedis: perKontrolBalik.Pasien.NoRekamMedis,
				Pengguna: &model.PenggunaResponse{
					ID:              perKontrolBalik.Pasien.Pengguna.ID,
					NamaLengkap:     perKontrolBalik.Pasien.Pengguna.NamaLengkap,
					Telepon:         perKontrolBalik.Pasien.Pengguna.Telepon,
					TeleponKeluarga: perKontrolBalik.Pasien.Pengguna.TeleponKeluarga,
					Alamat:          perKontrolBalik.Pasien.Pengguna.Alamat,
				},
				AdminPuskesmas: &model.AdminPuskesmasResponse{
					ID:            perKontrolBalik.Pasien.AdminPuskesmas.ID,
					NamaPuskesmas: perKontrolBalik.Pasien.AdminPuskesmas.NamaPuskesmas,
					Telepon:       perKontrolBalik.Pasien.AdminPuskesmas.Telepon,
					Alamat:        perKontrolBalik.Pasien.AdminPuskesmas.Alamat,
				},
				BeratBadan:     perKontrolBalik.Pasien.BeratBadan,
				TinggiBadan:    perKontrolBalik.Pasien.TinggiBadan,
				TekananDarah:   perKontrolBalik.Pasien.TekananDarah,
				DenyutNadi:     perKontrolBalik.Pasien.DenyutNadi,
				HasilLab:       perKontrolBalik.Pasien.HasilLab,
				HasilEkg:       perKontrolBalik.Pasien.HasilEkg,
				TanggalPeriksa: perKontrolBalik.Pasien.TanggalPeriksa,
				Status:         perKontrolBalik.Pasien.Status,
			},
			TanggalKontrol: perKontrolBalik.TanggalKontrol,
			Status:         perKontrolBalik.Status,
		})
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &response, nil
}

func (s *KontrolBalikService) Get(ctx context.Context, request *model.KontrolBalikGetRequest) (*model.KontrolBalikResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	kontrolBalik := new(entity.KontrolBalik)
	if request.IdAdminPuskesmas > 0 {
		if err := s.KontrolBalikRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, kontrolBalik, request.ID, request.IdAdminPuskesmas, constant.StatusKontrolBalikMenunggu); err != nil {
			log.Println(err.Error())
			return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else if err := s.KontrolBalikRepository.FindByIdAndStatus(tx, kontrolBalik, request.ID, constant.StatusKontrolBalikMenunggu); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	response := new(model.KontrolBalikResponse)
	response.ID = kontrolBalik.ID
	response.TanggalKontrol = kontrolBalik.TanggalKontrol
	response.IdPasien = kontrolBalik.IdPasien
	return response, nil
}

func (s *KontrolBalikService) Create(ctx context.Context, request *model.KontrolBalikCreateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	if request.IdAdminPuskesmas > 0 {
		if err := s.PasienRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, &entity.Pasien{}, request.IdPasien, request.IdAdminPuskesmas, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else {
		if err := s.PasienRepository.FindByIdAndStatus(tx, &entity.Pasien{}, request.IdPasien, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	}

	kontrolBalikEntity := new(entity.KontrolBalik)
	kontrolBalikEntity.IdPasien = request.IdPasien
	kontrolBalikEntity.TanggalKontrol = request.TanggalKontrol
	kontrolBalikEntity.Status = constant.StatusKontrolBalikMenunggu

	if err := s.KontrolBalikRepository.Create(tx, kontrolBalikEntity); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *KontrolBalikService) Update(ctx context.Context, request *model.KontrolBalikUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	// cek harus status menunggu untuk edit
	kontrolBalik := new(entity.KontrolBalik)
	if request.IdAdminPuskesmas > 0 {
		if err := s.KontrolBalikRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, kontrolBalik, request.ID, request.IdAdminPuskesmas, constant.StatusKontrolBalikMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else {
		if err := s.KontrolBalikRepository.FindByIdAndStatus(tx, kontrolBalik, request.ID, constant.StatusKontrolBalikMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	}
	// cek harus pasien status aktif untuk edit
	if request.IdAdminPuskesmas > 0 {
		if err := s.PasienRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, &entity.Pasien{}, request.IdPasien, request.IdAdminPuskesmas, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else {
		if err := s.PasienRepository.FindByIdAndStatus(tx, &entity.Pasien{}, request.IdPasien, constant.StatusPasienAktif); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	}

	kontrolBalik.IdPasien = request.IdPasien
	kontrolBalik.TanggalKontrol = request.TanggalKontrol

	if err := s.KontrolBalikRepository.Update(tx, kontrolBalik); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *KontrolBalikService) Delete(ctx context.Context, request *model.KontrolBalikDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	kontrolBalik := new(entity.KontrolBalik)
	if request.IdAdminPuskesmas > 0 {
		if err := s.KontrolBalikRepository.FindByIdAndIdAdminPuskesmasAndStatusOrStatus(tx, kontrolBalik, request.ID, request.IdAdminPuskesmas, constant.StatusKontrolBalikBatal, constant.StatusKontrolBalikSelesai); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else {
		if err := s.KontrolBalikRepository.FindByIdAndStatusOrStatus(tx, kontrolBalik, request.ID, constant.StatusKontrolBalikBatal, constant.StatusKontrolBalikSelesai); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	}

	if err := s.KontrolBalikRepository.Delete(tx, kontrolBalik); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *KontrolBalikService) Batal(ctx context.Context, request *model.KontrolBalikBatalRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	// cek harus status menunggu untuk batal
	kontrolBalik := new(entity.KontrolBalik)
	if request.IdAdminPuskesmas > 0 {
		if err := s.KontrolBalikRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, kontrolBalik, request.ID, request.IdAdminPuskesmas, constant.StatusKontrolBalikMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else {
		if err := s.KontrolBalikRepository.FindByIdAndStatus(tx, kontrolBalik, request.ID, constant.StatusKontrolBalikMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	}

	kontrolBalik.Status = constant.StatusKontrolBalikBatal

	if err := s.KontrolBalikRepository.Update(tx, kontrolBalik); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *KontrolBalikService) Selesai(ctx context.Context, request *model.KontrolBalikSelesaiRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}
	// cek harus status menunggu untuk selesai
	kontrolBalik := new(entity.KontrolBalik)
	if request.IdAdminPuskesmas > 0 {
		if err := s.KontrolBalikRepository.FindByIdAndIdAdminPuskesmasAndStatus(tx, kontrolBalik, request.ID, request.IdAdminPuskesmas, constant.StatusKontrolBalikMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else {
		if err := s.KontrolBalikRepository.FindByIdAndStatus(tx, kontrolBalik, request.ID, constant.StatusKontrolBalikMenunggu); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	}

	kontrolBalik.Status = constant.StatusKontrolBalikSelesai

	if err := s.KontrolBalikRepository.Update(tx, kontrolBalik); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}
