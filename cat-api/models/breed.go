package models

//Breed have all possible breed data
type Breed struct {
	Adaptability     int64  `json:"adaptability" bson:"adaptability"`
	AffectionLevel   int64  `json:"affection_level" bson:"affection_level"`
	AltNames         string `json:"alt_names" bson:"alt_names"`
	CFAURL           string `json:"cfa_url" bson:"cfa_url"`
	ChildFriendly    int64  `json:"child_friendly" bson:"child_friendly"`
	CountryCode      string `json:"country_code" bson:"country_code"`
	CountryCodes     string `json:"country_codes" bson:"country_codes"`
	Description      string `json:"description" bson:"description"`
	DogFriendly      int64  `json:"dog_friendly" bson:"dog_friendly"`
	EnergyLevel      int64  `json:"energy_level" bson:"energy_level"`
	Experimental     int64  `json:"experimental" bson:"experimental"`
	Grooming         int64  `json:"grooming" bson:"grooming"`
	Hairless         int64  `json:"hairless" bson:"hairless"`
	HealthIssues     int64  `json:"health_issues" bson:"health_issues"`
	HypoAllergenic   int64  `json:"hypoallergenic" bson:"hypoallergenic"`
	ID               string `json:"id" bson:"id"`
	Indoor           int64  `json:"indoor" bson:"indoor"`
	Intelligence     int64  `json:"intelligence" bson:"intelligence"`
	Lap              int64  `json:"lap" bson:"lap"`
	LifeSpan         string `json:"life_span" bson:"life_span"`
	Name             string `json:"name" bson:"name"`
	Natural          int64  `json:"natural" bson:"natural"`
	Origin           string `json:"origin" bson:"origin"`
	Rare             int64  `json:"rare" bson:"rare"`
	Rex              int64  `json:"rex" bson:"rex"`
	SheddingLevel    int64  `json:"shedding_level" bson:"shedding_level"`
	ShortLegs        int64  `json:"short_legs" bson:"short_legs"`
	SocialNeeds      int64  `json:"social_needs" bson:"social_needs"`
	StrangerFriendly int64  `json:"stranger_friendly" bson:"stranger_friendly"`
	SupressedTail    int64  `json:"suppressed_tail" bson:"suppressed_tail"`
	Temperament      string `json:"temperament" bson:"temperament"`
	VCAHospitalsURL  string `json:"vcahospitals_url" bson:"vcahospitals_url"`
	VetStreetURL     string `json:"vetstreet_url" bson:"vetstreet_url"`
	Vocalisation     int64  `json:"vocalisation" bson:"vocalisation"`
	Weight           Weight `json:"weight" bson:"weight"`
	WikipediaURL     string `json:"wikipedia_url" bson:"wikipedia_url"`
}

//Weight is a struct that wields the weight of a breed in different measures
type Weight struct {
	Imperial string `json:"imperial" bson:"imperial"`
	Metric   string `json:"metric" bson:"metric"`
}
