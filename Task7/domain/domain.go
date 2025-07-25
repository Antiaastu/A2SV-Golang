package domain
import "time"

type User struct{
	ID string
	Username string
	Password string
	Role string
}

type Task struct{
	ID string
	Title string
	Description string
	DueDate time.Time
	Status string
	OwnerID string
}
