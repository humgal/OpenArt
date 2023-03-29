package collection

type Collection struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	CreateDate string  `json:"createDate"`
	Creator    *string `json:"creator"`
}
