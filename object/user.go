package object

// User object
type User struct {
	ID int64
	First, Last string
}

// Name returns the full name of the user
func (u *User) Name() string {
	return u.First + " " + u.Last
}