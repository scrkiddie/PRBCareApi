package service

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"prbcare_be/internal/entity"
	"prbcare_be/internal/model"
	"prbcare_be/internal/repository"
)

type ObatService struct {
	DB                    *gorm.DB
	ObatRepository        *repository.ObatRepository
	AdminApotekRepository *repository.AdminApotekRepository
	Validator             *validator.Validate
	Config                *viper.Viper
}

func NewObatService(db *gorm.DB,
	obatRepository *repository.ObatRepository,
	validator *validator.Validate, adminApotekRepository *repository.AdminApotekRepository,
	config *viper.Viper) *ObatService {
	return &ObatService{db, obatRepository, adminApotekRepository, validator, config}
}

func (s *ObatService) List(ctx context.Context, request *model.ObatListRequest) (*[]model.ObatResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	obat := new([]entity.Obat)

	if request.IdAdminApotek > 0 {
		if err := s.ObatRepository.FindAllByIdAdminApotek(tx, obat, request.IdAdminApotek); err != nil {
			log.Println(err.Error())
			return nil, fiber.ErrInternalServerError
		}
	} else if err := s.ObatRepository.FindAll(tx, obat); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	var response []model.ObatResponse
	for _, perObat := range *obat {
		response = append(response, model.ObatResponse{
			ID:       perObat.ID,
			NamaObat: perObat.NamaObat,
			Jumlah:   perObat.Jumlah,
			AdminApotek: model.AdminApotekResponse{
				ID:         perObat.AdminApotek.ID,
				NamaApotek: perObat.AdminApotek.NamaApotek,
				Telepon:    perObat.AdminApotek.Telepon,
				Alamat:     perObat.AdminApotek.Alamat,
			},
		})
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &response, nil
}

func (s *ObatService) Get(ctx context.Context, request *model.ObatGetRequest) (*model.ObatResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	obat := new(entity.Obat)
	if request.IdAdminApotek > 0 {
		if err := s.ObatRepository.FindByIdAndIdAdminApotek(tx, obat, request.ID, request.IdAdminApotek); err != nil {
			log.Println(err.Error())
			return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else if err := s.ObatRepository.FindById(tx, obat, request.ID); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	obatResponse := new(model.ObatResponse)
	obatResponse.ID = obat.ID
	obatResponse.NamaObat = obat.NamaObat
	obatResponse.Jumlah = obat.Jumlah
	obatResponse.AdminApotek = model.AdminApotekResponse{
		ID:         obat.AdminApotek.ID,
		NamaApotek: obat.AdminApotek.NamaApotek,
		Telepon:    obat.AdminApotek.Telepon,
		Alamat:     obat.AdminApotek.Alamat,
	}

	return obatResponse, nil
}

func (s *ObatService) Create(ctx context.Context, request *model.ObatCreateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	if err := s.AdminApotekRepository.FindById(tx, &entity.AdminApotek{}, request.IdAdminApotek); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	obatEnity := new(entity.Obat)
	obatEnity.NamaObat = request.NamaObat
	obatEnity.IdAdminApotek = request.IdAdminApotek
	obatEnity.Jumlah = request.Jumlah

	if err := s.ObatRepository.Create(tx, obatEnity); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *ObatService) Update(ctx context.Context, request *model.ObatUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	obat := new(entity.Obat)
	fmt.Println(request.CurrentAdminApotek)
	if request.CurrentAdminApotek {
		if err := s.ObatRepository.FindByIdAndIdAdminApotek(tx, obat, request.ID, request.IdAdminApotek); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else if err := s.ObatRepository.FindById(tx, obat, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := s.AdminApotekRepository.FindById(tx, &entity.AdminApotek{}, request.IdAdminApotek); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	obat.IdAdminApotek = request.IdAdminApotek
	obat.NamaObat = request.NamaObat
	obat.Jumlah = request.Jumlah
	obat.AdminApotek = entity.AdminApotek{}
	fmt.Println(request.IdAdminApotek)
	if err := s.ObatRepository.Update(tx, obat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *ObatService) Delete(ctx context.Context, request *model.ObatDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	obat := new(entity.Obat)
	if request.IdAdminApotek > 0 {
		if err := s.ObatRepository.FindByIdAndIdAdminApotek(tx, obat, request.ID, request.IdAdminApotek); err != nil {
			log.Println(err.Error())
			return fiber.NewError(fiber.StatusNotFound, "Not found")
		}
	} else if err := s.ObatRepository.FindById(tx, obat, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound, "Not found")
	}

	if err := s.ObatRepository.Delete(tx, obat); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}
