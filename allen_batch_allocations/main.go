package main

/*
    Users' Abilities:
        Students can sign up.
        Admins can sign up.
        Admins can create groups.
        Students can join a specific group.
        A student can only join one group (like IIT or NEET) once.
        Admins can assign students to groups based on different factors.
            Gender: Female students are assigned first to available morning, then noon, then evening slots.
            Capacity: Assign to groups with the most space first.

    Input and Output:
        Inputs for actions should be done through methods.
        Method signature should contain enough info for all needs.
        You can change how inputs and outputs are formatted without changing how it works.

		Sample Examples
			● register(“Akhilesh”, Male, Student) - Student1
			● register(“Komal”, Female, Student) - Student2
			● register(“Rajnish”, Male,Student) - Student3
			● register(“Mayuri”, Female,Student) - Student4
			● enroll(Student1, Stream : IIT)
			● enroll(Student2, Stream: IIT )
			● enroll(Student3, Stream: NEET )
			● enroll(Student4, Stream: IIT)
			● register(“Kamesh”, Male , Admin) - Admin1
			● register(“M”, Male, Admin) - Admin2
			● createBatch(Admin1, Capacity=3, Stream : IIT, Timing : Morning) - B1
			○ Capacity of the batch is the max number of students which can be allocated to this batch
			● createBatch(Admin1, Capacity=2, Stream: NEET, Timing: Evening) - B2
			● createBatch(Admin1, Capacity=3, Stream : IIT, Timing : Evening) - B3

	Guidelines:
        Write a driver class for demonstration purposes.
        Output can be written to the console.
        Store all data in memory.
        Use any programming language you prefer.
        Don't create a user interface.
        Save your code with your name and email it. It will be executed on another machine, so specify any dependencies in your email.

    Expectations:
        Your code should be demonstrable and functionally correct.
        It should handle edge cases well and handle errors gracefully with appropriate exception handling.
        It should have a good object-oriented design.
        Code should be readable, modular, testable, and extensible.
        Use clear and intuitive names for variables, methods, and classes.
        Make it easy to add or remove functionality without rewriting large portions of the code.
        Avoid writing overly long or complex code.
        Don't use databases; store data in memory.

*/
import (
	"fmt"

	"github.com/sheelendar/src/go_projects/allen_batch_allocations/data_storage"
	"github.com/sheelendar/src/go_projects/allen_batch_allocations/handler/handler"
	"github.com/sheelendar/src/go_projects/allen_batch_allocations/utils"
)

func main() {
	handler.Register("Akhilesh", utils.MALE, utils.STUDENT)
	handler.Register("Komal", utils.FEMALE, utils.STUDENT)
	handler.Register("Rajnish", utils.MALE, utils.STUDENT)
	handler.Register("Mayuri", utils.FEMALE, utils.STUDENT)
	handler.Register("Ram", utils.MALE, utils.STUDENT)

	handler.Enroll("Akhilesh", "IIT")
	handler.Enroll("Komal", "IIT")
	handler.Enroll("Rajnish", "NEET")
	handler.Enroll("Mayuri", "IIT")
	handler.Enroll("Ram", "IIT")

	handler.Register("Kamesh", utils.MALE, utils.ADMIN)
	handler.Register("M", utils.MALE, utils.ADMIN)

	handler.CreateBatch("Kamesh", "IIT", utils.MORNING, 3, 1)
	handler.CreateBatch("Kamesh", "NEET", utils.EVENING, 2, 20)
	handler.CreateBatch("Kamesh", "IIT", utils.EVENING, 3, 21)

	handler.AllocateBatch("Kamesh", "Komal", utils.AllocationGenderPriority)
	handler.AllocateBatch("Kamesh", "Mayuri", utils.AllocationHigerCapa)
	handler.AllocateBatch("Kamesh", "Akhilesh", utils.AllocationHigerCapa)
	handler.AllocateBatch("Kamesh", "Rajnish", utils.AllocationHigerCapa)
	handler.AllocateBatch("Kamesh", "Ram", utils.AllocationHigerCapa)

	for name, student := range data_storage.StudentList {
		fmt.Println("Name: ", name)
		fmt.Println("stream: ", student.Stream, " Roll: ", student.Roll, " Gender: ", student.Gender, " batchID ", student.BatchID)
	}

	fmt.Println()
	fmt.Println()
	for name, admin := range data_storage.AdminList {
		fmt.Println("Name: ", name)
		fmt.Println(" Roll: ", admin.Roll, " Gender: ", admin.Gender)
	}

	fmt.Println()
	n := len(data_storage.Batches)

	for i := 0; i < n; i++ {
		fmt.Print(" Name ", data_storage.Batches[i].Name)
		fmt.Print(" BatchID ", data_storage.Batches[i].BatchID)
		fmt.Print(" Capacity ", data_storage.Batches[i].Capacity)
		fmt.Print(" CurrCapacity ", data_storage.Batches[i].CurrCapacity)
		fmt.Println()

	}
}
