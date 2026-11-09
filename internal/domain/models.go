package domain

type Ski struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	Size       int     `json:"size"`
	Status     string  `json:"status"`     // available oder rented
	DailyPrice float64 `json:"dailyPrice"` // Mietpreis pro Tag in Euro
}

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Loan struct {
	ID         int     `json:"id"`
	SkiID      int     `json:"skiId"`
	PersonID   int     `json:"personId"`
	Date       string  `json:"date"`
	Days       int     `json:"days"`
	DailyPrice float64 `json:"dailyPrice"`
	TotalPrice float64 `json:"totalPrice"`
}

type Reservation struct {
	ID       int    `json:"id"`
	SkiID    int    `json:"skiId"`
	PersonID int    `json:"personId"`
	Date     string `json:"date"`
	Status   string `json:"status"`
}

type Dashboard struct {
	AvailableSkis int `json:"availableSkis"`
	RentedSkis    int `json:"rentedSkis"`
	TotalSkis     int `json:"totalSkis"`
	Persons       int `json:"persons"`
}
