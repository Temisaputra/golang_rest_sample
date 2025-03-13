package auth

type GetCurrentUserResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       User   `json:"data"`
}

type User struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	Avatar      string   `json:"avatar"`
	PhoneNumber int      `json:"phone_number"`
	Role        Role     `json:"role"`
	KodeCabang  []string `json:"kode_cabang"`
}

type Role struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	Deskripsi string      `json:"deskripsi"`
	Dept      Departement `json:"departement"`
}

type Departement struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
