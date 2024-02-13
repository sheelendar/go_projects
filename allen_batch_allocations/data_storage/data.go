package data_storage

import "github.com/sheelendar/src/go_projects/allen_batch_allocations/dto"

var (
	// we can make these as protacted so can not be modified by other source. due time limit not doing this.
	StudentList = make(map[string]dto.Student)
	AdminList   = make(map[string]dto.Admin)
	Batches     = []dto.Batch{}
)
