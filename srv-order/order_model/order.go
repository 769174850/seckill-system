package order_model

type Order struct {
	ID        string `gorm:"primarykey"`
	UserID    string
	ProductID string
	Status    string `gorm:"default:0;not null"` //2为未支付，1为已支付，0为取消
}
