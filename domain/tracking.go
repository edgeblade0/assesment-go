package domain

const MessageSucces = "Delivery tracking detail fetched successfully"
const MessageFailed = "Error fetching data"

type Result struct {
	Status Status
	Data   Data
}

type History struct {
	Description string    `json:"description"`
	CreateAt    string    `json:"createdAt"`
	Formatted   Formatted `json:"formatted"`
}

type Formatted struct {
	CreateAt string `json:"createdAt"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	ReceivedBy string    `json:"receivedBy"`
	Histories  []History `josn:"histories"`
}
