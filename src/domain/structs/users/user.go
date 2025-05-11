package users

type UserColumns []string

var UserColumnsValues = UserColumns{
	"id", "name", "email", "password", "ein", "phone", "postalCode",
	"address", "addressNumber", "complement", "city", "country",
	"external_id", "created_at", "updated_at",
}

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Ein           string `json:"ein"`
	Phone         string `json:"phone"`
	PostalCode    string `json:"postalCode"`
	Address       string `json:"address"`
	AddressNumber string `json:"addressNumber"`
	Complement    string `json:"complement"`
	City          string `json:"city"`
	Country       string `json:"country"`
	ExternalId    string `json:"external_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
