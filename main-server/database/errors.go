package database

type UsernameAlreadyTakenError struct {
}

func (u UsernameAlreadyTakenError) Error() string {
	return "database:UsernameAlreadyTaken"
}
