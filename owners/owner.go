package owners

type Owner struct {
	ID     int
	Name   string
	Email  string
	About  string
	UserID *int // связь с пользователем (пользователь может не быть владельцем)
}
