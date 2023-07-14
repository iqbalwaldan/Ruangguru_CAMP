package repository

import (
	"a21hc3NpZ25tZW50/model"
	"log"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	arrStudent := make([]model.Student, 0)
	// student := &model.Student{}
	data := s.db.Find(&arrStudent)
	if data.Error != nil {
		return nil, data.Error
	}
	rows, err := data.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student model.Student
		err := s.db.ScanRows(rows, &student)
		if err != nil {
			log.Fatalln(err)
		}
		arrStudent = append(arrStudent, student)
	}

	return arrStudent, nil // TODO: replace this
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	data := s.db.Create(&student)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	students := &model.Student{}
	data := s.db.Debug().Model(students).Where("id = ?", students.ID).UpdateColumns(
		map[string]interface{}{
			"name":     student.Name,
			"address":  student.Address,
			"class_id": student.ClassId,
		},
	)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *studentRepoImpl) Delete(id int) error {
	data := s.db.Delete(&model.Student{}, id)
	if data.Error != nil {
		return data.Error
	}
	return nil
	// student := &model.Student{}
	// data := s.db.Debug().Model(student).Where("id = ?", id).Delete(student)
	// if data.Error != nil {
	// 	return data.Error
	// }
	// return nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	student := &model.Student{}
	data := s.db.First(&student, id)
	if data.Error != nil {
		return &model.Student{}, data.Error
	}
	return student, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	arrClass := make([]model.StudentClass, 0)
	student := &model.Student{}
	data := s.db.First(&arrClass, student.ClassId)
	if data.Error != nil {
		return nil, data.Error
	}
	rows, err := data.Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var clas model.StudentClass
		err := s.db.ScanRows(rows, &clas)
		if err != nil {
			log.Fatalln(err)
		}
		arrClass = append(arrClass, clas)
	}

	return &arrClass, nil // TODO: replace this
}
