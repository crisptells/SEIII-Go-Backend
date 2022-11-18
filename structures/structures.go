package structures

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	Cash      string `json:"cash"`
}

type Game struct {
	Id     string `json:"id"`
	Land1  string `json:"land1"`
	Land2  string `json:"land2"`
	Date   string `json:"date"`
	Result string `json:"result"`
	Group  string `json:"group"`
}

type Bet struct {
	User string `json:"email"`
	Game string `json:"id"`
	Result string `json:"result"`
}

type Group struct {
	Name string `json:"name"`
	AdminUser string `json:"admin"`
	Users []User
}
