package models

type Resume struct {
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Phone      string   `json:"phone"`
	Summary    string   `json:"summary"`
	Education  []Degree `json:"education"`
	Experience []Job    `json:"experience"`
	Skills     []string `json:"skills"`
}

type Degree struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
}

type Job struct {
	Company     string `json:"company"`
	Position    string `json:"position"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
	Description string `json:"description"`
}
