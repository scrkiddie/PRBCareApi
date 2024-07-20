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

type AdminPuskesmasService struct {
	DB                       *gorm.DB
	AdminPuskesmasRepository *repository.AdminPuskesmasRepository
	PasienRepository         *repository.PasienRepository
	Validator                *validator.Validate
	Config                   *viper.Viper
}

func NewAdminPuskesmasService(db *gorm.DB,
	adminPuskesmasRepository *repository.AdminPuskesmasRepository,
	pasienRepository *repository.PasienRepository,
	validator *validator.Validate,
	config *viper.Viper) *AdminPuskesmasService {
	return &AdminPuskesmasService{db, adminPuskesmasRepository, pasienRepository, validator, config}
}

func (s *AdminPuskesmasService) List(ctx context.Context) (*[]model.AdminPuskesmasResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	adminPuskesmas := new([]entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindAll(tx, adminPuskesmas); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	var response []model.AdminPuskesmasResponse
	for _, admin := range *adminPuskesmas {
		response = append(response, model.AdminPuskesmasResponse{
			ID:            admin.ID,
			NamaPuskesmas: admin.NamaPuskesmas,
			Telepon:       admin.Telepon,
			Alamat:        admin.Alamat,
		})
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &response, nil
}

func (s *AdminPuskesmasService) Get(ctx context.Context, request *model.AdminPuskesmasGetRequest) (*model.AdminPuskesmasResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindById(tx, adminPuskesmas, request.ID); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound)
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	adminPuskesmasResponse := new(model.AdminPuskesmasResponse)
	adminPuskesmasResponse.ID = adminPuskesmas.ID
	adminPuskesmasResponse.Username = adminPuskesmas.Username
	adminPuskesmasResponse.NamaPuskesmas = adminPuskesmas.NamaPuskesmas
	adminPuskesmasResponse.Alamat = adminPuskesmas.Alamat
	adminPuskesmasResponse.Telepon = adminPuskesmas.Telepon

	return adminPuskesmasResponse, nil
}

func (s *AdminPuskesmasService) Create(ctx context.Context, request *model.AdminPuskesmasCreateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	total, err := s.AdminPuskesmasRepository.CountByUsername(tx, request.Username)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 {
		return fiber.NewError(fiber.StatusConflict, "Username sudah digunakan")
	}

	total, err = s.AdminPuskesmasRepository.CountByTelepon(tx, request.Telepon)
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

	adminPuskesmasEnity := new(entity.AdminPuskesmas)
	adminPuskesmasEnity.Username = request.Username
	adminPuskesmasEnity.NamaPuskesmas = request.NamaPuskesmas
	adminPuskesmasEnity.Alamat = request.Alamat
	adminPuskesmasEnity.Telepon = request.Telepon
	adminPuskesmasEnity.Password = string(password)

	if err := s.AdminPuskesmasRepository.Create(tx, adminPuskesmasEnity); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminPuskesmasService) Update(ctx context.Context, request *model.AdminPuskesmasUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindById(tx, adminPuskesmas, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound)
	}

	total, err := s.AdminPuskesmasRepository.CountByUsername(tx, request.Username)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 && adminPuskesmas.Username != request.Username {
		return fiber.NewError(fiber.StatusConflict, "Username sudah digunakan")
	}

	total, err = s.AdminPuskesmasRepository.CountByTelepon(tx, request.Telepon)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 && adminPuskesmas.Telepon != request.Telepon {
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

	adminPuskesmas.Username = request.Username
	adminPuskesmas.NamaPuskesmas = request.NamaPuskesmas
	adminPuskesmas.Alamat = request.Alamat
	adminPuskesmas.Telepon = request.Telepon
	if string(password) != "" {
		adminPuskesmas.Password = string(password)
	}

	if err := s.AdminPuskesmasRepository.Update(tx, adminPuskesmas); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminPuskesmasService) Delete(ctx context.Context, request *model.AdminPuskesmasDeleteRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindById(tx, adminPuskesmas, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound)
	}

	if err := s.PasienRepository.FindByIdAdminPuskesmas(tx, &entity.Pasien{}, request.ID); err == nil {
		return fiber.NewError(fiber.StatusConflict, "Admin puskesmas masih terkait dengan data pasien yang ada")
	}

	if err := s.AdminPuskesmasRepository.Delete(tx, adminPuskesmas); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminPuskesmasService) Login(ctx context.Context, request *model.AdminPuskesmasLoginRequest) (*model.AdminPuskesmasResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindByUsername(tx, adminPuskesmas, request.Username); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(adminPuskesmas.Password), []byte(request.Password)); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Username atau password salah")
	}

	key := s.Config.GetString("jwt.secret")
	exp := s.Config.GetInt("jwt.exp")
	claims := jwt.MapClaims{
		"sub":  adminPuskesmas.ID,
		"role": constant.RoleAdminPuskesmas,
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

	return &model.AdminPuskesmasResponse{Token: token}, nil
}

func (s *AdminPuskesmasService) Verify(ctx context.Context, request *model.VerifyAdminPuskesmasRequest) (*model.Auth, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	tokenString := request.Token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:", token.Header["alg"])
			return nil, fiber.ErrInternalServerError
		}
		return []byte(s.Config.GetString("jwt.secret")), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err.Error())
		return nil, fiber.ErrUnauthorized
	}

	var id int
	var role string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if subFloat64, ok := claims["sub"].(float64); ok {
			id = int(subFloat64)
		} else {
			return nil, fiber.ErrUnauthorized
		}
		if roleString, ok := claims["role"].(string); ok {
			role = roleString
		} else {
			return nil, fiber.ErrUnauthorized
		}
	} else {
		return nil, fiber.ErrUnauthorized
	}

	if role != constant.RoleAdminPuskesmas {
		return nil, fiber.ErrUnauthorized
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindById(tx, adminPuskesmas, id); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound)
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &model.Auth{ID: adminPuskesmas.ID, Role: role}, nil
}

func (s *AdminPuskesmasService) Current(ctx context.Context, request *model.AdminPuskesmasGetRequest) (*model.AdminPuskesmasResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrBadRequest
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindById(tx, adminPuskesmas, request.ID); err != nil {
		log.Println(err.Error())
		return nil, fiber.NewError(fiber.StatusNotFound)
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return nil, fiber.ErrInternalServerError
	}

	adminPuskesmasResponse := new(model.AdminPuskesmasResponse)
	adminPuskesmasResponse.NamaPuskesmas = adminPuskesmas.NamaPuskesmas
	adminPuskesmasResponse.Alamat = adminPuskesmas.Alamat
	adminPuskesmasResponse.Telepon = adminPuskesmas.Telepon

	return adminPuskesmasResponse, nil
}

func (s *AdminPuskesmasService) CurrentProfileUpdate(ctx context.Context, request *model.AdminPuskesmasProfileUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindById(tx, adminPuskesmas, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound)
	}

	total, err := s.AdminPuskesmasRepository.CountByTelepon(tx, request.Telepon)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	if total > 0 && adminPuskesmas.Telepon != request.Telepon {
		return fiber.NewError(fiber.StatusConflict, "Telepon sudah digunakan")
	}

	adminPuskesmas.NamaPuskesmas = request.NamaPuskesmas
	adminPuskesmas.Alamat = request.Alamat
	adminPuskesmas.Telepon = request.Telepon

	if err := s.AdminPuskesmasRepository.Update(tx, adminPuskesmas); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *AdminPuskesmasService) CurrentPasswordUpdate(ctx context.Context, request *model.AdminPuskesmasPasswordUpdateRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validator.Struct(request); err != nil {
		log.Println(err.Error())
		return fiber.ErrBadRequest
	}

	adminPuskesmas := new(entity.AdminPuskesmas)
	if err := s.AdminPuskesmasRepository.FindById(tx, adminPuskesmas, request.ID); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusNotFound)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(adminPuskesmas.Password), []byte(request.CurrentPassword)); err != nil {
		log.Println(err.Error())
		return fiber.NewError(fiber.StatusUnauthorized, "Password saat ini salah")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}
	adminPuskesmas.Password = string(password)

	if err := s.AdminPuskesmasRepository.Update(tx, adminPuskesmas); err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Println(err.Error())
		return fiber.ErrInternalServerError
	}

	return nil
}
