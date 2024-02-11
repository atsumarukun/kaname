package requests

type SearchComputersQuery struct {
	Sort  *string `form:"sort"`
	Order *string `form:"order"`
}
