package structures

type RESTResponse struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Records interface{} `json:"records,omitempty"`
}

type User struct {
	UserID               int                 `json:"userId"`
	CountryID            int                 `json:"countryId"`
	Email                string              `json:"email"`
	Password             string              `json:"password,omitempty"`
	InputNumberPhone     int                 `json:",omitempty"`
	InputPhoto           string              `json:",omitempty"`
	InputExtension       string              `json:",omitempty"`
	InputIsProfileLessor int                 `json:",omitempty"`
	InputNameCompany     string              `json:",omitempty"`
	InputNumberNit       string              `json:",omitempty"`
	Name                 string              `json:"name"`
	LastName             string              `json:"lastName"`
	Address              string              `json:"address"`
	CountryInfo          Country             `json:"country,omitempty"`
	PhoneInfo            Phone               `json:"phones,omitempty"`
	PhotoInfo            Photo               `json:"photo,omitempty"`
	SellerProfile        SellerProfileType   `json:"sellerProfile,omitempty"`
	ConsumerProfile      ConsumerProfileType `json:"consumerProfile,omitempty"`
}

type ConsumerProfileType struct {
	ID int `json:"consumerProfileID,omitempty"`
}

type SellerProfileType struct {
	ID            int    `json:"sellerProfileID,omitempty"`
	BussinessName string `json:"bussinesName,omitempty"`
	NitNumber     string `json:"nitNumber,omitempty"`
}

type Photo struct {
	PhotoID   int
	Photo     string `json:"photo"`
	Extension string `json:"extension"`
}

type Phone struct {
	PhoneID      int `json:"phoneId"`
	PhoneNumber  int `json:"phoneNumber"`
	CountryID    int
	CountryPhone Country `json:"country"`
}

type Country struct {
	ID        int
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

type UserPhone struct {
	UserID  int
	PhoneID int
}

type Adds struct {
	ID          int
	IDProfile   int
	IDProduct   int
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Product     ProductType `json:"product"`
	Categories  []Category  `json:"categories"`
}

type ProductType struct {
	ID              int
	SellerProfileID int
	ColorID         int
	SizeID          int
	TypeArticleID   int
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	PriceXHour      float32 `json:"priceXHour"`
	ImageURL        string  `json:"imageURL"`
	ColorInfo       Color   `json:"color"`
	SizeInfo        Size    `json:"size"`
}

type PurchaseOrderType struct {
	ID                int                       `json:"id"`
	ConsumerProfileID int                       `json:"consumerProfileId"`
	CreatedAt         string                    `json:"createdAt"`
	CreatePayment     string                    `json:"createPayment"`
	DateEndOperation  string                    `json:"dateEndOperation"`
	TotalOrder        float32                   `json:"totalOrder"`
	Status            int                       `json:"status"`
	Detail            []DetailPurchaseOrderType `json:"detail"`
}

type DetailPurchaseOrderType struct {
	ID                   int
	PurchaseOrderID      int     `json:"purchaseOrderId"`
	ProductID            int     `json:"productId"`
	DateTimeToDevolution string  `json:"dateTimeToDevolution"`
	CantHours            float32 `json:"cantHours"`
	SubTotal             float32 `json:"subTotal"`
}

type TypeArticle struct {
	ID          int
	Description string `json:"description"`
}

type Color struct {
	ID       int
	Name     string `json:"name"`
	ColorHex string `json:"colorHex"`
}

type Size struct {
	ID   int
	Size string `json:"size"`
}

type Category struct {
	ID          int
	Description string `json:"description"`
}

var MessagesCatalog = struct {
	RespondGenericMessage500Error string
}{
	RespondGenericMessage500Error: "Ocurrio un error generico",
}

var Response = RESTResponse{
	Result:  true,
	Message: "Transaccion exitosa",
}
