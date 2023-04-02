package models

// type User struct {

// 	//this represents the customer
// 	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	FirstName  string    `gorm:"type:varchar(255);not null" column:"first_name"`
// 	SecondName string    `gorm:"type:varchar(255);not null" column:"second_name"`
// 	Email      string    `gorm:"uniqueIndex" column:"email"`
// 	Order      []Order
// 	TableID    uuid.UUID `gorm:"column:table_id,omitempty"`
// 	CreatedAt  time.Time `gorm:"column:created_at,omitempty"`
// 	UpdatedAt  time.Time `gorm:"column:updated_at,omitempty"`
// }

// type Menus struct {
// 	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	Name         string    `gorm:"column:name"`
// 	OrderDetails []OrderDetails
// 	Price        float64   `gorm:"column:price,omitempty"`
// 	CreatedAt    time.Time `gorm:"column:created_at,omitempty"`
// 	UpdatedAt    time.Time `gorm:"column:updated_at,omitempty"`
// }

// type Table struct {
// 	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	User        []User
// 	TableNumber int       `gorm:"size:255"`
// 	CreatedAt   time.Time `gorm:"column:created_at"`
// 	UpdatedAt   time.Time `gorm:"column:updated_at"`
// }

// type Order struct {
// 	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	UserID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	OrderDetails []OrderDetails
// 	TableID      uuid.UUID `gorm:"column:table_id,omitempty"`
// 	Quantity     int
// 	TotalPrice   float64
// 	CreatedAt    time.Time `gorm:"column:created_at,omitempty"`
// 	UpdatedAt    time.Time `gorm:"column:updated_at,omitempty"`
// }

// type OrderDetails struct {
// 	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
// 	OrderID    int       `gorm:"column:order_id"`
// 	MenuID     uuid.UUID `gorm:"column:menu_id"`
// 	Quantity   int
// 	TotalPrice int       `gorm:"column:total_price"`
// 	CreatedAt  time.Time `gorm:"column:created_at"`
// 	UpdatedAt  time.Time `gorm:"column:updated_at"`
// }
