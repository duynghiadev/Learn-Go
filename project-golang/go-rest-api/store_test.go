package main

// Mocks

type MockStore struct{}

func (s *MockStore) CreateProject(p *Project) error {
	return nil
}

func (s *MockStore) GetProject(id string) (*Project, error) {
	return &Project{
		ID:   1,
		Name: "Super cool project",
	}, nil
}

func (s *MockStore) GetAllProjects() ([]*Project, error) {
	return []*Project{
		{
			ID:   1,
			Name: "Test Project 1",
		},
		{
			ID:   2,
			Name: "Test Project 2",
		},
	}, nil
}

func (s *MockStore) DeleteProject(id string) error {
	return nil
}

func (s *MockStore) CreateUser(u *User) (*User, error) {
	return &User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}, nil
}

func (s *MockStore) GetUserByID(id string) (*User, error) {
	return &User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
	}, nil
}

func (s *MockStore) GetUserByEmail(email string) (*User, error) {
	hashedPassword, _ := HashPassword("123456")
	return &User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "Test",
		LastName:  "User",
		Password:  hashedPassword,
	}, nil
}

func (s *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{
		ID:           1,
		Name:         t.Name,
		Status:       "todo",
		ProjectID:    t.ProjectID,
		AssignedToID: t.AssignedToID,
	}, nil
}

func (s *MockStore) GetTask(id string) (*Task, error) {
	return &Task{
		ID:           1,
		Name:         "Demo Task",
		Status:       "in progress",
		ProjectID:    1,
		AssignedToID: 42,
	}, nil
}

func (s *MockStore) GetAllTasks() ([]*Task, error) {
	return []*Task{
		{
			ID:           1,
			Name:         "Task One",
			Status:       "todo",
			ProjectID:    1,
			AssignedToID: 42,
		},
		{
			ID:           2,
			Name:         "Task Two",
			Status:       "in progress",
			ProjectID:    2,
			AssignedToID: 24,
		},
	}, nil
}
