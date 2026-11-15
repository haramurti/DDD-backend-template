package dto

type CreateFeatureBRequest struct {
	// TODO: field + tag binding, contoh:
	// Title string `json:"title" binding:"required,min=3,max=150"`
}

type UpdateFeatureBRequest struct {
	// TODO
}

type ListFeatureBQuery struct {
	Q    string `form:"q" binding:"max=100"`
	Sort string `form:"sort" binding:"omitempty,oneof=newest oldest"`
	Page int    `form:"page,default=1" binding:"min=1"`
	Size int    `form:"size,default=20" binding:"min=1,max=100"`
	// TODO: filter lain (status, kategori, dll)
}
