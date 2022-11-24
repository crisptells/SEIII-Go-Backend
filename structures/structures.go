package structures

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	FirstName string `json:"firstName"`
	Subjects  string `json:"subjects"`
	Role	  string `json:"role"`
}

type Tutoring struct {
	Tutor 	  User `json:"user"`
	Subject   string `json:"subject"`
	Student   User `json:"student"`
	MaxStudents string `json:"maxStudents"`
	Appointments []Appointment
}

type Appointment struct {
	Date	  string `json:"date"`
	Duration  string `json:"duration"`
}
