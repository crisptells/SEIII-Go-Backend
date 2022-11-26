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
	Tutor        string `json:"user"` //just the email
	Subject      string `json:"subject"`
	Students     string `json:"students"` //just the emails
	MaxStudents  int    `json:"maxStudents"`
	Appointments []Appointment
}

type Appointment struct {
	Date     string `json:"date"`
	Duration string `json:"duration"`
}
