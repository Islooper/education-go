package dao

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Form   string `json:"form"`
	Path   string `json:"path"`
	IsOn   int32  `json:"is_on"`
	UserId int64  `json:"user_id"`
}

func (*File) TableName() string {
	return "file"
}

type FileDao struct {
}

func (f *FileDao) List(pageNo, pageSize int) ([]*File, error) {
	limit := pageSize
	offset := pageSize * (pageNo - 1)

	fileDos := make([]*File, 0)

	err := Db.Model(&File{}).Limit(limit).Offset(offset).Find(&fileDos).Error
	if err != nil {
		return nil, err
	}

	return fileDos, nil
}

func (f *FileDao) Create(file *File) error {
	return Db.Create(file).Error
}
