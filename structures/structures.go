package structures

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Vorname  string `json:"firstName"`
	Geld     string `json:"cash"`
}

type Tutoring struct {
	Tutoring_id string `json:"tutoring_id"`
	Tutor       string `json:"tutor"` //just the name
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type Appointment struct {
	Appointment_id int    `json:"appointment_id"`
	Date           string `json:"date"`
	Duration       string `json:"duration"`
	Tutoring_id    string `json:"tutoring_id"`
}

type Experience struct {
	User_email  string `json:"user_email"`
	Math        int    `json:"math"`
	German      int    `json:"german"`
	English     int    `json:"english"`
	Physics     int    `json:"physics"`
	Chemistry   int    `json:"chemistry"`
	Informatics int    `json:"informatics"`
}

type User_tutoring struct {
	User_email  string `json:"user_email"`
	Tutoring_id string `json:"tutoring_id"`
}
