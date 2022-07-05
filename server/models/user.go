package models

type User struct {
	ID	uint	`json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func CreateUser (fullname string, email string, password string) *User {
	salt := "12345678"; // TODO
	query, err := GetDB().Prepare("INSERT INTO users(fullname, email, password, salt) VALUES (?, ?, ?, ?)");
	if err != nil {
		panic(err.Error());
		// TODO
	}

	res, err := query.Exec(fullname, email, password, salt);

	if err != nil {
		panic(err.Error());
		// TODO
	}

	lastId, err := res.LastInsertId();
	if err != nil {
		panic(err.Error());
		// TODO
	}

	return &User{
		ID: uint(lastId),
		Fullname: fullname,
		Email: email,
		Password: password,
	}
}
