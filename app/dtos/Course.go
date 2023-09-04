package dtos

type Course struct {
	Id                  uint
	Code                int
	Name                string
	Family              string
	MaxNumberOfStudents int
	NumberOfStudents    int
	StartDate           string
	EndDate             string
	Availability        bool
}
