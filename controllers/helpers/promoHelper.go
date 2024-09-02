package helpers

import (
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/repositories"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"

	"github.com/google/uuid"
)

func GetAllPromo(GetAllPromoRequestDTO requestsDTO.GetAllPromosRequestDTO) (int, interface{}){
	db := configs.GetDB()
	var promos []database.Promos

	if GetAllPromoRequestDTO.Limit == 0 || GetAllPromoRequestDTO.Limit > 100{
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit cannot be 0",
		}
		return 400, output
	}

	offset := (GetAllPromoRequestDTO.Page - 1) * GetAllPromoRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllPromoRequestDTO.OrderBy, GetAllPromoRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllPromoRequestDTO.Limit).Order(order).Find(&promos).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	if len(promos) == 0{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Promos Not Found",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int 
	db.Model(&promos).Count(&totalData)

	if totalData % int64(GetAllPromoRequestDTO.Limit) == 0{
		totalPage = int(totalData) / GetAllPromoRequestDTO.Limit
	}else{
		totalPage = int(totalData) / GetAllPromoRequestDTO.Limit + 1
	}

	output := outputs.GetAllPromoOutput{}
	output.Page = GetAllPromoRequestDTO.Page
	output.Limit = GetAllPromoRequestDTO.Limit
	output.OrderBy = GetAllPromoRequestDTO.OrderBy
	output.OrderType = GetAllPromoRequestDTO.OrderType
	output.Code = 200
	output.Message = "Success"
	output.TotalData = int(totalData)
	output.TotalTake = len(promos)
	output.TotalPage = totalPage
	
	for _, promo := range promos{
		promoResponse := responsesDTO.PromoResponseDTO{
			ID: promo.ID,
			ProductID: promo.ProductID,
			Name: promo.Name,
			Description: promo.Description,
			PromoDiscount: promo.PromoDiscount,
			IsActive: promo.IsActive,
			CreatedBy: promo.CreatedBy,
			UpdatedBy: promo.UpdatedBy,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		}
		output.Data = append(output.Data, promoResponse)
	}
	return 200, output
}

func GetAllPromoByProductID(productID string) (int, interface{}) {
	var promos []responsesDTO.PromoResponseDTO

	promos, err := repositories.GetAllPromoByProductID(productID)

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	if len(promos) == 0{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Promos Not Found",
		}
		return 404, output
	}

	output := outputs.GetAllPromoByProductID{}
	output.Code = 200
	output.Message = "Success"

	for _, promo := range promos{
		promoResponse := responsesDTO.PromoResponseDTO{
			ID: promo.ID,
			ProductID: promo.ProductID,
			Name: promo.Name,
			Description: promo.Description,
			PromoDiscount: promo.PromoDiscount,
			IsActive: promo.IsActive,
			CreatedBy: promo.CreatedBy,
			UpdatedBy: promo.UpdatedBy,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		}
		output.Data = append(output.Data, promoResponse)
	}
	return 200, output
}

func GetPromo(promoID string) (int, interface{}){
	var promo database.Promos
	promo, err := repositories.GetPromoByID(promoID)

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	if promo.ID == uuid.Nil{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Promo Not Found",
		}
		return 404, output
	}

	output := outputs.GetPromoOutput{
		BaseOutput: outputs.BaseOutput{
			Code: 200,
			Message: "Success",
		},
		Data: responsesDTO.PromoResponseDTO{
			ID: promo.ID,
			ProductID: promo.ProductID,
			Name: promo.Name,
			Description: promo.Description,
			PromoDiscount: promo.PromoDiscount,
			IsActive: promo.IsActive,
			CreatedBy: promo.CreatedBy,
			UpdatedBy: promo.UpdatedBy,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		},
	}
	return 200, output
}

func CreatePromo(CreatePromoRequestDTO requestsDTO.CreatePromosRequestDTO) (int, interface{}){
	db := configs.GetDB()
	promo := database.Promos{
		ProductID: CreatePromoRequestDTO.ProductID,
		Name: CreatePromoRequestDTO.Name,
		Description: CreatePromoRequestDTO.Description,
		PromoDiscount: CreatePromoRequestDTO.PromoDiscount,
		CreatedBy: CreatePromoRequestDTO.CreatedBy,
		IsActive: CreatePromoRequestDTO.IsActive,
	}
	
	err := db.Create(&promo).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{ 
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.CreatePromoOutput{
		BaseOutput: outputs.BaseOutput{
			Code: 200,
			Message: "Success",
		},
		Data: responsesDTO.PromoResponseDTO{
			ID: promo.ID,
			ProductID: promo.ProductID,
			Name: promo.Name,
			Description: promo.Description,
			PromoDiscount: promo.PromoDiscount,
			IsActive: promo.IsActive,
			CreatedBy: promo.CreatedBy,
			UpdatedBy: promo.UpdatedBy,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		},
	}
	return 200, output
}

func UpdatePromo(UpdatePromoRequestDTO requestsDTO.UpdatePromosRequestDTO) (int, interface{}){
	db := configs.GetDB()
	promo, err := repositories.GetPromoByID(UpdatePromoRequestDTO.ID)

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	if promo.ID == uuid.Nil{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Promo Not Found",
		}
		return 404, output
	}

	// Not Null Update Constraint -> ini tidak boleh null, kalo user tidak mengisi maka akan diisi oleh sistem

	promo.ProductID = UpdatePromoRequestDTO.ProductID

	if promo.Name != ""{
		promo.Name = UpdatePromoRequestDTO.Name
	} 

	if promo.Description != ""{
		promo.Description = UpdatePromoRequestDTO.Description
	} 

	promo.PromoDiscount = UpdatePromoRequestDTO.PromoDiscount
	
	if promo.UpdatedBy == ""{
		promo.UpdatedBy = "user"
	}else{
		promo.UpdatedBy = UpdatePromoRequestDTO.UpdatedBy
	}

	// Boolean Update Constraint
	promo.IsActive = UpdatePromoRequestDTO.IsActive

	err = db.Save(&promo).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.UpdatePromoOutput{
		BaseOutput: outputs.BaseOutput{
			Code: 200,
			Message: "Success Update",
		},
		Data: responsesDTO.PromoResponseDTO{
			ID: promo.ID,
			ProductID: promo.ProductID,
			Name: promo.Name,
			Description: promo.Description,
			PromoDiscount: promo.PromoDiscount,
			IsActive: promo.IsActive,
			CreatedBy: promo.CreatedBy,
			UpdatedBy: promo.UpdatedBy,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		},
	}
	return 200, output
}

func DeletePromo(DeletePromoRequestDTO requestsDTO.DeletePromosRequestDTO) (int, interface{}){
	db := configs.GetDB()
	promo, err := repositories.GetPromoByID(DeletePromoRequestDTO.ID)

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	if promo.ID == uuid.Nil{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Promo Not Found",
		}
		return 404, output
	}

	err = db.Delete(&promo).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.DeletePromoOutput{
		BaseOutput: outputs.BaseOutput{
			Code: 200,
			Message: "Success",
		},
		Data: responsesDTO.PromoResponseDTO{
			ID: promo.ID,
			ProductID: promo.ProductID,
			Name: promo.Name,
			Description: promo.Description,
			PromoDiscount: promo.PromoDiscount,
			IsActive: promo.IsActive,
			CreatedBy: promo.CreatedBy,
			UpdatedBy: promo.UpdatedBy,
			CreatedAt: promo.CreatedAt,
			UpdatedAt: promo.UpdatedAt,
		},
	}
	return 200, output
}
