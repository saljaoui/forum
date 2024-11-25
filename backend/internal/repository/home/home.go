package home

type PostResponde struct {
	ID           int      `json:"ID"`
	UserID       int      `json:"UserID"`
	FirstName    string   `json:"FirstName"`
	LastName     string   `json:"LastName"`
	Title        string   `json:"Title"`
	Content      string   `json:"Content"`
	CategoryName []string `json:"CategoryName"`
	CreatedAt    string   `json:"CreatedAt"`
}
