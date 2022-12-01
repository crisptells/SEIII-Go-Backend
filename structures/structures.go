package structures

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	Subjects  string `json:"subjects"`
	Role      string `json:"role"`
}

type Tutoring struct {
	Tutoring_id string `json:"tutoring_id"`
	Tutor       string `json:"user"` //just the email
	Subject     string `json:"subject"`
	Students    string `json:"students"` //just the emails
	MaxStudents int    `json:"maxStudents"`
}

type Appointment struct {
	Appointment_id int    `json:"appointment_id"`
	Date           string `json:"date"`
	Duration       string `json:"duration"`
	Tutoring_id    string `json:"tutoring_id"`
}
