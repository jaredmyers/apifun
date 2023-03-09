package storage

type UserService struct {
	Store UserServicer
}

func NewUserService(store *MySqlStore) *UserService {
	return &UserService{
		Store: store,
	}
}
