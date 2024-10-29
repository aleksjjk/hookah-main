package models

type User struct {
    ID              int64
    Name            string
    Address         string
    Phone           string
    Language        string
		Delivery string
		SelectedTobaccos []string // Для хранения выбранных табаков
}
func (u *User) ResetOrderData() {
    u.SelectedTobaccos = []string{}
		u.Delivery = ""
}

