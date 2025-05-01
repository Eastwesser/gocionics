package user

type User struct {
	ID          int
	Name        string
	Email       string
	Password    string
	CharacterID int // ссылка на Character
}
