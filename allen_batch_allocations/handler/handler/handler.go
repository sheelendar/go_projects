package handler

import (
	"fmt"
	"sort"

	"github.com/sheelendar/go_projects/allen_batch_allocations/data_storage"
	"github.com/sheelendar/go_projects/allen_batch_allocations/dto"
	"github.com/sheelendar/go_projects/allen_batch_allocations/utils"
)

func Register(name, gender, roll string) {
	if roll == utils.ADMIN {
		data_storage.AdminList[name] = dto.Admin{
			Name:   name,
			Gender: gender,
			Roll:   roll,
		}
		fmt.Println("success: registration done ")
	} else if roll == utils.STUDENT {
		data_storage.StudentList[name] = dto.Student{
			Name:   name,
			Gender: gender,
			Roll:   roll,
		}
		fmt.Println("success: registration done ")
	} else {
		fmt.Println("err: no roll found into system")
	}
}

// createBatch(Admin1, Capacity=3, Stream : IIT, Timing : Morning) - B1
func CreateBatch(createdBy, stream, timing string, cap, time int) {
	_, ok := data_storage.AdminList[createdBy]
	if !ok {
		fmt.Print("err: No admin found into system", createdBy)
		return
	}
	batch := dto.Batch{BatchID: 0, CreatedBy: createdBy, Stream: stream, Capacity: cap, Timing: timing, Time: time}
	id := len(data_storage.Batches)
	batch.BatchID = id
	data_storage.Batches = append(data_storage.Batches, batch)
	fmt.Print("success: created batch successfully", createdBy)
}

// enroll(Student1, Stream : IIT)
func Enroll(studentName string, stream string) bool {
	student, ok := data_storage.StudentList[studentName]
	if !ok {
		fmt.Print("err: student not found with given name", studentName)
		return false
	}
	student.Stream = stream
	data_storage.StudentList[studentName] = student
	fmt.Print("success: enroll for given name", studentName)
	return true
}

//allocateBatch(Admin2, Student2, “Gender Based”)  Note : Student2 is a female

func AllocateBatch(admin, student, priority string) {
	_, ok := data_storage.AdminList[admin]
	if !ok {
		fmt.Print("err: No admin found into system", admin)
		return
	}
	if priority == utils.AllocationGenderPriority {
		allocateBatchBasedOnGender(admin, student)
		return
	}
	allocateBatchBasedOnCapacity(admin, student)

}

// we will refactor both function  allocateBatchBasedOnGender and allocateBatchBasedOnCapacity reduce number of lines.
func allocateBatchBasedOnCapacity(admin string, studentName string) {
	n := len(data_storage.Batches)
	if n == 0 {
		fmt.Println("there is no batch created")
		return
	}

	student, ok := data_storage.StudentList[studentName]
	if !ok {
		fmt.Print("err: student not found with given name", studentName)
		return
	}
	if n == 1 {
		// will move this common part into another function
		if data_storage.Batches[0].CurrCapacity < data_storage.Batches[0].Capacity && student.Stream == data_storage.Batches[0].Stream {
			data_storage.Batches[0].CurrCapacity++
			if assignBatchID(studentName, data_storage.Batches[0].BatchID) {
				fmt.Println(admin, " allocated bath to student ", student)
				return
			}
			fmt.Print("there is no capacity into any batch for student")
		}
		return
	}
	sort.Slice(data_storage.Batches, func(i, j int) bool {
		return data_storage.Batches[i].CurrCapacity > data_storage.Batches[j].CurrCapacity
	})
	for i := 0; i < n; i++ {
		// will move this common part into another function
		if data_storage.Batches[i].CurrCapacity < data_storage.Batches[i].Capacity && student.Stream == data_storage.Batches[i].Stream {
			data_storage.Batches[i].CurrCapacity++
			if assignBatchID(studentName, data_storage.Batches[i].BatchID) {
				fmt.Println(admin, " allocated bath to student ", student)
				return
			}
		}
	}
	fmt.Print("there is no capacity into any batch for student")
}

func allocateBatchBasedOnGender(admin string, studentName string) {
	n := len(data_storage.Batches)
	if n == 0 {
		fmt.Println("there is no batch created")
		return
	}
	student, ok := data_storage.StudentList[studentName]
	if !ok {
		fmt.Print("err: student not found with given name", studentName)
		return
	}

	if n == 1 {
		// will move this common part into another function
		if data_storage.Batches[0].CurrCapacity < data_storage.Batches[0].Capacity && student.Stream == data_storage.Batches[0].Stream {
			data_storage.Batches[0].CurrCapacity++
			if assignBatchID(studentName, data_storage.Batches[0].BatchID) {
				fmt.Println(admin, " allocated bath to student ", student)
				return
			}
			fmt.Print("there is no capacity into any batch for student")
		}
		return
	}
	sort.Slice(data_storage.Batches, func(i, j int) bool {
		if data_storage.Batches[i].Time == data_storage.Batches[j].Time {
			return data_storage.Batches[i].CurrCapacity > data_storage.Batches[j].CurrCapacity
		}
		return data_storage.Batches[i].Time < data_storage.Batches[j].Time
	})
	for i := 0; i < n; i++ {
		// will move this common part into another function
		if data_storage.Batches[i].CurrCapacity < data_storage.Batches[i].Capacity && student.Stream == data_storage.Batches[i].Stream {
			data_storage.Batches[i].CurrCapacity++
			if assignBatchID(studentName, data_storage.Batches[i].BatchID) {
				fmt.Println(admin, " allocated bath to student ", student)
				return
			}
		}
	}
	fmt.Print("there is no capacity into any batch for student")
}

func assignBatchID(studentName string, batchID int) bool {

	student, ok := data_storage.StudentList[studentName]
	if !ok {
		fmt.Print("err: student not found with given name", studentName)
		return false
	}
	student.BatchID = batchID
	data_storage.StudentList[studentName] = student
	return true
}
