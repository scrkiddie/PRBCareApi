package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"prb_care_api/internal/constant"
	"prb_care_api/internal/entity"
	"prb_care_api/internal/model"
	"prb_care_api/internal/repository"
	"time"
)

type AdminApotekService struct {
	DB                    *gorm.DB
	AdminApotekRepository *repository.AdminApotekRepository
	ObatRepository        *repository.ObatRepository
	Validator             *validator.Validate
	Config                *viper.Viper
}

func NewAdminApotekService(db *gorm.DB,
	adminApotekRepository *repository.AdminApotekRepository,
	obatRepository *repository.ObatRepository,
	validator *validator.Validate,
	config *viper.Viper) *AdminApotekService {
	return &AdminApotekService{db, adminApotekRepository, obatRepository, validator, config}
}

func (s *AdminApotekService) List(ctx context.Context) (*[]model.AdminApotekResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	adminApotek := new([]entity.AdminApotek)
	if err := s.AdminApotekRepository.FindAll(tx, adminApotek); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	var response []model.AdminApotekResponse
	for _, admin := range *adminApotek {
		response = append(response, model.AdminApotekResponse{
			ID:         admin.ID,
			NamaApotek: admin.NamaApotek,
			Telepon:    admin.Telepon,
			Alamat:     admin.Alamat,
		})
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &response, nil
}

func (s *AdminApotekService) Get(ctx context.Context, request *model.AdminApotekGetRequest) (*model.AdminApotekResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindById(tx, adminApotek, request.ID); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	adminApotekResponse := new(model.AdminApotekResponse)
	adminApotekResponse.ID = adminApotek.ID
	adminApotekResponse.Username = adminApotek.Username
	adminApotekResponse.NamaApotek = adminApotek.NamaApotek
	adminApotekResponse.Alamat = adminApotek.Alamat
	adminApotekResponse.Telepon = adminApotek.Telepon

	return adminApotekResponse, nil
}

func (s *AdminApotekService) Create(ctx context.Context, request *model.AdminApotekCreateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	total, err := s.AdminApotekRepository.CountByUsername(tx, request.Username)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 {
		return fiber.NewError(fiber.StatusConflict, "Username sudah digunakan")
	}

	total, err = s.AdminApotekRepository.CountByTelepon(tx, request.Telepon)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 {
		return fiber.NewError(fiber.StatusConflict, "Telepon sudah digunakan")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	adminApotekEnity := new(entity.AdminApotek)
	adminApotekEnity.Username = request.Username
	adminApotekEnity.NamaApotek = request.NamaApotek
	adminApotekEnity.Alamat = request.Alamat
	adminApotekEnity.Telepon = request.Telepon
	adminApotekEnity.Password = string(password)

	if err := s.AdminApotekRepository.Create(tx, adminApotekEnity); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminApotekService) Update(ctx context.Context, request *model.AdminApotekUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindById(tx, adminApotek, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.ErrNotFound
	}

	total, err := s.AdminApotekRepository.CountByUsername(tx, request.Username)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 && adminApotek.Username != request.Username {
		return fiber.NewError(fiber.StatusConflict, "Username sudah digunakan")
	}

	total, err = s.AdminApotekRepository.CountByTelepon(tx, request.Telepon)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 && adminApotek.Telepon != request.Telepon {
		return fiber.NewError(fiber.StatusConflict, "Telepon sudah digunakan")
	}

	var password []byte
	if request.Password != "" {
		password, err = bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err.Error())
			return fiber.ErrInternalServerError
		}
	}

	adminApotek.Username = request.Username
	adminApotek.NamaApotek = request.NamaApotek
	adminApotek.Alamat = request.Alamat
	adminApotek.Telepon = request.Telepon
	if string(password) != "" {
		adminApotek.Password = string(password)
	}

	if err := s.AdminApotekRepository.Update(tx, adminApotek); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminApotekService) Delete(ctx context.Context, request *model.AdminApotekDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindById(tx, adminApotek, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.ErrNotFound
	}

	if err := s.ObatRepository.FindByIdAdminApotek(tx, &entity.Obat{}, request.ID); err == nil {
		return fiber.NewError(fiber.StatusConflict, "Admin apotek masih terkait dengan data obat yang ada")
	}

	if err := s.AdminApotekRepository.Delete(tx, adminApotek); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminApotekService) Login(ctx context.Context, request *model.AdminApotekLoginRequest) (*model.AdminApotekResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindByUsername(tx, adminApotek, request.Username); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(adminApotek.Password), []byte(request.Password)); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Username atau password salah")
	}

	key := s.Config.GetString("jwt.secret")
	exp := s.Config.GetInt("jwt.exp")
	claims := jwt.MapClaims{
		"sub":  adminApotek.ID,
		"role": constant.RoleAdminApotek,
		"exp":  time.Now().Add(time.Duration(exp) * time.Hour).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
	if err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &model.AdminApotekResponse{Token: token}, nil
}

func (s *AdminApotekService) Verify(ctx context.Context, request *model.AdminApotekVerifyRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindById(tx, adminApotek, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminApotekService) Current(ctx context.Context, request *model.AdminApotekGetRequest) (*model.AdminApotekResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindById(tx, adminApotek, request.ID); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	adminApotekResponse := new(model.AdminApotekResponse)
	adminApotekResponse.NamaApotek = adminApotek.NamaApotek
	adminApotekResponse.Alamat = adminApotek.Alamat
	adminApotekResponse.Telepon = adminApotek.Telepon

	return adminApotekResponse, nil
}

func (s *AdminApotekService) CurrentProfileUpdate(ctx context.Context, request *model.AdminApotekProfileUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindById(tx, adminApotek, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.ErrNotFound
	}

	total, err := s.AdminApotekRepository.CountByTelepon(tx, request.Telepon)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 && adminApotek.Telepon != request.Telepon {
		return fiber.NewError(fiber.StatusConflict, "Telepon sudah digunakan")
	}

	adminApotek.NamaApotek = request.NamaApotek
	adminApotek.Alamat = request.Alamat
	adminApotek.Telepon = request.Telepon

	if err := s.AdminApotekRepository.Update(tx, adminApotek); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminApotekService) CurrentPasswordUpdate(ctx context.Context, request *model.AdminApotekPasswordUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminApotek := new(entity.AdminApotek)
	if err := s.AdminApotekRepository.FindById(tx, adminApotek, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(adminApotek.Password), []byte(request.CurrentPassword)); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusUnauthorized, "Password saat ini salah")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	adminApotek.Password = string(password)

	if err := s.AdminApotekRepository.Update(tx, adminApotek); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}
