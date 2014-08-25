package service

import (
	"veg_rest2/api"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

type VegetableService struct {
}

func (s *VegetableService) getDb(cfg Config) (gorm.DB, error) {
	//connectionString := cfg.DbUser + ":" + cfg.DbPassword + "@tcp(" + cfg.DbHost + ":3306)/" + cfg.DbName + "?charset=utf8&parseTime=True"
	//return gorm.Open("mysql", connectionString)
	return gorm.Open("sqlite3", "db/vegetables.sqlite3")
}

func (s *VegetableService) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	//db.SingularTable(true)

	db.AutoMigrate(api.Vegetable{})
	return nil
}
func (s *VegetableService) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	vegetableResource := &VegetableResource{db: db}

	r := gin.Default()

	r.GET("/vegetable", vegetableResource.GetAllVegetables)
	r.GET("/vegetable/:id", vegetableResource.GetVegetable)
	r.POST("/vegetable", vegetableResource.CreateVegetable)
	r.PUT("/vegetable/:id", vegetableResource.UpdateVegetable)
	r.PATCH("/vegetable/:id", vegetableResource.PatchVegetable)
	r.DELETE("/vegetable/:id", vegetableResource.DeleteVegetable)

	r.Run(cfg.SvcHost)

	return nil
}
