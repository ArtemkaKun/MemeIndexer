package Error

import "log"

const (
	InvalidLoginOrPass = "Неправильный логин или пароль!"
	MemeNotFound       = "Совпадений не найдено :("
	DefaultError       = "Кажется, на сервере возникли проблемы"
)

func HandleCommonError(err error) (errorMessage string) {
	log.Println(err)
	errorMessage = DefaultError
	return
}

func HandleDBServerError(err *string) (errorMessage string) {
	log.Println("Error from DB: " + *err)

	switch *err {
	case "mongo: no documents in result":
		errorMessage = InvalidLoginOrPass
	case "Nothing was found!":
		errorMessage = MemeNotFound
	default:
		errorMessage = DefaultError
	}
	return
}
