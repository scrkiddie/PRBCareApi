package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"log"
	"prbcare_be/internal/entity"
	"prbcare_be/internal/model"
	"prbcare_be/internal/repository"
)

type PasienService struct {
	DB                       *gorm.DB
	PasienRepository         *repository.PasienRepository
	AdminPuskesmasRepository *repository.AdminPuskesmasRepository
	PenggunaRepository       *repository.PenggunaRepository
	Validator                *validator.Validate
}

func NewPasienService(
	db *gorm.DB,
	pasienRepository *repository.PasienRepository,
	adminPuskesmasRepository *repository.AdminPuskesmasRepository,
	penggunaRepository *repository.PenggunaRepository,
	validator *validator.Validate,
) *PasienService {
	return &PasienService{db, pasienRepository, adminPuskesmasRepository, penggunaRepository, validator}
}

func (s *PasienService) Search(ctx context.Context, request *model.PasienSearchRequest) (*[]model.PasienResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	pasien := new([]entity.Pasien)
	if request.IdPengguna > 0 {
		if err := s.PasienRepository.SearchAsPengguna(tx, pasien, request.IdPengguna, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else if request.IdAdminPuskesmas > 0 {
		if err := s.PasienRepository.SearchAsAdminPuskesmas(tx, pasien, request.IdAdminPuskesmas, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else {
		if err := s.PasienRepository.Search(tx, pasien, request.Status); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	}

	var response []model.PasienResponse
	for _, perPasien := range *pasien {
		response = append(response, model.PasienResponse{
			ID:             perPasien.ID,
			NoRekamMedis:   perPasien.NoRekamMedis,
			BeratBadan:     perPasien.BeratBadan,
			TinggiBadan:    perPasien.TinggiBadan,
			TekananDarah:   perPasien.TekananDarah,
			DenyutNadi:     perPasien.DenyutNadi,
			HasilLab:       perPasien.HasilLab,
			HasilEkg:       perPasien.HasilEkg,
			TanggalPeriksa: perPasien.TanggalPeriksa,
			Status:         perPasien.Status,
			Pengguna: model.PenggunaResponse{
				ID:              perPasien.Pengguna.ID,
				NamaLengkap:     perPasien.Pengguna.NamaLengkap,
				Telepon:         perPasien.Pengguna.Telepon,
				TeleponKeluarga: perPasien.Pengguna.TeleponKeluarga,
				Alamat:          perPasien.Pengguna.Alamat,
			},
			AdminPuskesmas: model.AdminPuskesmasResponse{
				ID:            perPasien.AdminPuskesmas.ID,
				NamaPuskesmas: perPasien.AdminPuskesmas.NamaPuskesmas,
				Telepon:       perPasien.AdminPuskesmas.Telepon,
				Alamat:        perPasien.AdminPuskesmas.Alamat,
			},
		})
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &response, nil
}

func (s *PasienService) Get(ctx context.Context, request *model.PasienGetRequest) (*model.PasienResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	pasien := new(entity.Pasien)
	if request.IdAdminPuskesmas > 0 {
		if err := s.PasienRepository.FindByIdAndIdAdminPuskesmas(tx, pasien, request.ID, request.IdAdminPuskesmas); err != nil {
			log.Println(err.Error())
			return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else if err := s.PasienRepository.FindById(tx, pasien, request.ID); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	pasienResponse := new(model.PasienResponse)
	pasienResponse.ID = pasien.ID
	pasienResponse.NoRekamMedis = pasien.NoRekamMedis
	pasienResponse.BeratBadan = pasien.BeratBadan
	pasienResponse.TinggiBadan = pasien.TinggiBadan
	pasienResponse.TekananDarah = pasien.TekananDarah
	pasienResponse.DenyutNadi = pasien.DenyutNadi
	pasienResponse.HasilLab = pasien.HasilLab
	pasienResponse.HasilEkg = pasien.HasilEkg
	pasienResponse.TanggalPeriksa = pasien.TanggalPeriksa
	pasienResponse.Status = pasien.Status
	pasienResponse.Pengguna = model.PenggunaResponse{
		ID:              pasien.Pengguna.ID,
		NamaLengkap:     pasien.Pengguna.NamaLengkap,
		Alamat:          pasien.Pengguna.Alamat,
		Telepon:         pasien.Pengguna.Telepon,
		TeleponKeluarga: pasien.Pengguna.TeleponKeluarga,
	}
	pasienResponse.AdminPuskesmas = model.AdminPuskesmasResponse{
		ID:            pasien.AdminPuskesmas.ID,
		NamaPuskesmas: pasien.AdminPuskesmas.NamaPuskesmas,
		Telepon:       pasien.AdminPuskesmas.Telepon,
		Alamat:        pasien.AdminPuskesmas.Alamat,
	}

	return pasienResponse, nil
}

func (s *PasienService) Create(ctx context.Context, request *model.PasienCreateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	if err := s.AdminPuskesmasRepository.FindById(tx, &entity.AdminPuskesmas{}, request.IdAdminPuskesmas); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := s.PenggunaRepository.FindById(tx, &entity.Pengguna{}, request.IdPengguna); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	pasienEntity := new(entity.Pasien)
	pasienEntity.NoRekamMedis = request.NoRekamMedis
	pasienEntity.IdPengguna = request.IdPengguna
	pasienEntity.IdAdminPuskesmas = request.IdAdminPuskesmas
	pasienEntity.BeratBadan = request.BeratBadan
	pasienEntity.TinggiBadan = request.TinggiBadan
	pasienEntity.TekananDarah = request.TekananDarah
	pasienEntity.DenyutNadi = request.DenyutNadi
	pasienEntity.HasilLab = request.HasilLab
	pasienEntity.HasilEkg = request.HasilEkg
	pasienEntity.TanggalPeriksa = request.TanggalPeriksa
	pasienEntity.Status = request.Status

	if err := s.PasienRepository.Create(tx, pasienEntity); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *PasienService) Update(ctx context.Context, request *model.PasienUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	pasien := new(entity.Pasien)
	if request.CurrentAdminPuskesmas {
		if err := s.PasienRepository.FindByIdAndIdAdminPuskesmas(tx, pasien, request.ID, request.IdAdminPuskesmas); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else if err := s.PasienRepository.FindById(tx, pasien, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := s.AdminPuskesmasRepository.FindById(tx, &entity.AdminPuskesmas{}, request.IdAdminPuskesmas); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}
	if err := s.PenggunaRepository.FindById(tx, &entity.Pengguna{}, request.IdPengguna); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}
	pasien.NoRekamMedis = request.NoRekamMedis
	pasien.IdPengguna = request.IdPengguna
	pasien.IdAdminPuskesmas = request.IdAdminPuskesmas
	pasien.BeratBadan = request.BeratBadan
	pasien.TinggiBadan = request.TinggiBadan
	pasien.TekananDarah = request.TekananDarah
	pasien.DenyutNadi = request.DenyutNadi
	pasien.HasilLab = request.HasilLab
	pasien.HasilEkg = request.HasilEkg
	pasien.TanggalPeriksa = request.TanggalPeriksa
	pasien.Status = request.Status
	pasien.Pengguna = entity.Pengguna{}
	pasien.AdminPuskesmas = entity.AdminPuskesmas{}

	if err := s.PasienRepository.Update(tx, pasien); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *PasienService) Delete(ctx context.Context, request *model.PasienDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	pasien := new(entity.Pasien)
	if request.IdAdminPuskesmas > 0 {
		if err := s.PasienRepository.FindByIdAndIdAdminPuskesmas(tx, pasien, request.ID, request.IdAdminPuskesmas); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else if err := s.PasienRepository.FindById(tx, pasien, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := s.PasienRepository.Delete(tx, pasien); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}
