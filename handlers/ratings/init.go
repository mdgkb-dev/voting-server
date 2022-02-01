package ratings

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers"
	httpHelper2 "mdgkb/mdgkb-server/helpers/httpHelperV2"
	"mdgkb/mdgkb-server/models"
)

type IHandler interface {
	//GetAll(c *gin.Context)
	//Get(c *gin.Context)
	//GetByDivisionID(c *gin.Context)
	Create(c *gin.Context)
	//Delete(c *gin.Context)
	//Update(c *gin.Context)
	//CreateComment(c *gin.Context)
	//UpdateComment(c *gin.Context)
	//RemoveComment(c *gin.Context)
	//CreateSlugs(c *gin.Context)
	//Search(c *gin.Context)
}

type IService interface {
	//setQueryFilter(*gin.Context) error

	Create(models.Ratings) error
	//GetAll() (models.Events, error)
	//Get(string) (*models.Event, error)
	//Delete(string) error
	//Update(*models.Doctor) error
}

type IRepository interface {
	//setQueryFilter(*gin.Context) error
	getDB() *bun.DB
	create(models.Ratings) error
	//getAll() (models.Events, error)
	//get(string) (*models.Event, error)
	//getByDivisionID(string) (models.Doctors, error)
	//delete(string) error
	//update(*models.Doctor) error
	//createComment(*models.DoctorComment) error
	//updateComment(*models.DoctorComment) error
	//removeComment(string) error
	//upsertMany(models.Doctors) error
	//search(string) (models.Doctors, error)
}

type IFilesService interface {
	//Upload(*gin.Context, *models.Doctor, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helpers.Helper
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db          *bun.DB
	ctx         context.Context
	helper      *helpers.Helper
	queryFilter *httpHelper2.QueryFilter
}

type FilesService struct {
	helper *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

func CreateService(db *bun.DB, helper *helpers.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helpers.Helper) *FilesService {
	return &FilesService{helper: helper}
}
