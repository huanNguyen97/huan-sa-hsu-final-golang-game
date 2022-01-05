package DTO

type Game struct {
	ID 				string 	`json:"id" binding:"required"`			
	Name 			string 	`json:"name" binding:"required"`
	Category 		string 	`json:"category" binding:"required"`
	Brand 			string 	`json:"brand" binding:"required"`
	Year_released 	int64 	`json:"year_released" binding:"required"`
	Price 			float64 `json:"price" binding:"required"`
	Url 			string 	`json:"url" binding:"required"`
	Id_User			int64	`json:"id_user" binding:"required"`
}