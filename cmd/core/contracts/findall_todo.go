package contracts

type FindAllTodoRequest struct {
	Sort     string `form:"sort" required:"false"`
	Page     int    `form:"page" required:"false"`
	PageSize int    `form:"pageSize" required:"false"`
}
