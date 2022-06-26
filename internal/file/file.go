package file

import (
	"fmt"

	"github.com/google/uuid"
)


type File struct {
	UUID uuid.UUID
	Name string 
	Data []byte
}

func NewFile(name string, blob []byte) (*File, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("cannot gen uuid for file: %w", err)
	}

	return &File{
		UUID: uuid,
		Name: name,
		Data: blob,
	}, nil
}

func (f *File)  String() string {
	return fmt.Sprintf("File(%s, %d)", f.Name, f.UUID)
}