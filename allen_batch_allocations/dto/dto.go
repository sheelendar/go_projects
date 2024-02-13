package dto

/*
register(“Akhilesh”, Male, Student) - Student1
register(“Komal”, Female, Student) - Student2

enroll(Student1, Stream : IIT)
enroll(Student2, Stream: IIT )

register(“Kamesh”, Male , Admin) - Admin1
register(“M”, Male, Admin) - Admin2

createBatch(Admin1, Capacity=3, Stream : IIT, Timing : Morning) - B1
Capacity of the batch is the max number of students which can be allocated to this batch
createBatch(Admin1, Capacity=2, Stream: NEET, Timing: Evening) - B2
createBatch(Admin1, Capacity=3, Stream : IIT, Timing : Evening) - B3
*/

type Student struct {
	Name    string // komal
	Gender  string // Female
	Roll    string // Student
	Stream  string // IIT/NEET
	BatchID int
	/*
		roll int
		dateOfB
		Add
		mobile
		email
	*/
}

type Admin struct {
	Name   string // admin1
	Gender string // Male
	Roll   string // Admin
	/*
		mobile
		email
	*/
}

type Batch struct {
	BatchID      int    // 1,2,3
	Name         string //B1
	CreatedBy    string // amdin1
	Capacity     int    // 5
	CurrCapacity int
	Stream       string //IIT
	Timing       string // Morning
	Time         int
}
