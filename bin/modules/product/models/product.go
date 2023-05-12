package models

type Product struct {
	Id                  string `json:"id,omitempty" bson:"_id,omitempty"`
	ProductName         string `json:"productname" bson:"productname" validate:"required"`
	Price               string `json:"price" bson:"price" validate:"required"`
	OptionPaymentMethod string `json:"payment" bson:"payment" validate:"required"`
	DeliveryMethod      string `json:"delivery" bson:"delivery" validate:"required"`
	StockDescription    string `json:"stock" bson:"stock" validate:"required"`
	Image               string `json:"image" bson:"image" validate:"required"`
}

type UpsertProduct struct {
	ProductName         string `json:"productname,omitempty" bson:"productname,omitempty"`
	Price               string `json:"price,omitempty" bson:"price,omitempty"`
	OptionPaymentMethod string `json:"payment,omitempty" bson:"payment,omitempty"`
	DeliveryMethod      string `json:"delivery,omitempty" bson:"delivery,omitempty"`
	StockDescription    string `json:"stock,omitempty" bson:"stock,omitempty"`
	Image               string `json:"image,omitempty" bson:"image,omitempty"`
}

func (u Product) UpsertProduct() UpsertProduct {
	return UpsertProduct{
		ProductName:         u.ProductName,
		Price:               u.Price,
		OptionPaymentMethod: u.OptionPaymentMethod,
		DeliveryMethod:      u.DeliveryMethod,
		StockDescription:    u.StockDescription,
		Image:               u.Image,
	}
}

/*
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Id          string `json:"id"`
	AccessToken string `json:"accessToken"`
}
*/

type GetProductResponse struct {
	Id                  string `json:"id,omitempty"`
	Productname         string `json:"productname"`
	Price               string `json:"price"`
	OptionPaymentMethod string `json:"optionpaymentmethod"`
	DeliveryMethod      string `json:"deliverymethod"`
	StockDescription    string `json:"stockdescription,omitempty"`
	Image               string `json:"image"`
}
