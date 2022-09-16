package entity

type Person struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Result []struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Username  string `json:"username"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		UUID      string `json:"uuid"`
	} `json:"result"`
}

type PersonOrdersJoined struct {
	Person
	Orders
	Items AllItems
}

type People []*Person
type PeopleOrdersJoined []*PersonOrdersJoined
