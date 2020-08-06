package data

//Cartoon struct
type Cartoon struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Ratings float64 `json:"ratings"`
}

// var Cartoons []Cartoon

// LoadData to load the data
func LoadData() []Cartoon {
	Cartoons := []Cartoon{
		Cartoon{ID: 1, Title: "The FLintstones", Ratings: 8},
		Cartoon{ID: 2, Title: "The Jetsons", Ratings: 8.5},
		Cartoon{ID: 3, Title: "Samurai Jack", Ratings: 9.2},
		Cartoon{ID: 4, Title: "Tom And Jerry", Ratings: 10},
		Cartoon{ID: 5, Title: "Ben 10", Ratings: 9.5},
	}

	return Cartoons
}
