package collection

type Collection struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CreateDate *string `json:"createDate"`
	Createor   *string `json:"createor"`
}
