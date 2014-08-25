package service

import (
	"log"
	"strconv"
	"time"

	"veg_rest2/api"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type VegetableResource struct {
	db gorm.DB
}

func (tr *VegetableResource) CreateVegetable(c *gin.Context) {
	var vegetable api.Vegetable

	if !c.Bind(&vegetable) {
		c.JSON(400, api.NewError("problem decoding body"))
		return
	}
	vegetable.Created_at = int32(time.Now().Unix())

	tr.db.Save(&vegetable)

	c.JSON(201, vegetable)
}

func (tr *VegetableResource) GetAllVegetables(c *gin.Context) {
	var vegetable []api.Vegetable

	tr.db.Order("created_at desc").Find(&vegetable)

	c.JSON(200, vegetable)
}

func (tr *VegetableResource) GetVegetable(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, api.NewError("problem decoding id sent"))
		return
	}

	var vegetable api.Vegetable

	if tr.db.First(&vegetable, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, vegetable)
	}
}

func (tr *VegetableResource) UpdateVegetable(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, api.NewError("problem decoding id sent"))
		return
	}

	var vegetable api.Vegetable

	if !c.Bind(&vegetable) {
		c.JSON(400, api.NewError("problem decoding body"))
		return
	}
	vegetable.Id = int32(id)

	var existing api.Vegetable

	if tr.db.First(&existing, id).RecordNotFound() {
		c.JSON(404, api.NewError("not found"))
	} else {
		tr.db.Save(&vegetable)
		c.JSON(200, vegetable)
	}

}

func (tr *VegetableResource) PatchVegetable(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, api.NewError("problem decoding id sent"))
		return
	}

	// this is a hack because Gin falsely claims my unmarshalled obj is invalid.
	// recovering from the panic and using my object that already has the json body bound to it.
	var json []api.Patch
	defer func() {
		if r := recover(); r != nil {
			if json[0].Op != "replace" && json[0].Path != "/status" {
				c.JSON(400, api.NewError("PATCH support is limited and can only replace the /status path"))
				return
			}

			var vegetable api.Vegetable

			if tr.db.First(&vegetable, id).RecordNotFound() {
				c.JSON(404, api.NewError("not found"))
			} else {
				//vegetable.Status = json[0].Value

				tr.db.Save(&vegetable)
				c.JSON(200, vegetable)
			}
		}
	}()
	c.Bind(&json)
}

func (tr *VegetableResource) DeleteVegetable(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, api.NewError("problem decoding id sent"))
		return
	}

	var vegetable api.Vegetable

	if tr.db.First(&vegetable, id).RecordNotFound() {
		c.JSON(404, api.NewError("not found"))
	} else {
		tr.db.Delete(&vegetable)
		c.Data(204, "application/json", make([]byte, 0))
	}
}

func (tr *VegetableResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
