package models

type User struct {
	ID       int
	Name     string
	Password string
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeleteAt  time.Time
}

type Menus struct {
	ID    int
	Name  string
	Price int
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeleteAt  time.Time
}

type Table struct {
	ID     int
	Number int
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeleteAt  time.Time
}

type Order struct {
	ID         int
	UserID     int
	TableID    int
	Quantity   int
	TotalPrice int
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeleteAt  time.Time
}
