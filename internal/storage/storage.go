package storage

import (
	"pets/storage/internal/file"

	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (st *Storage) Upload(name string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(name, blob)
	if err != nil {
		return nil, err
	}

	st.files[newFile.UUID] = newFile

	return newFile, err
}

func (st *Storage) ByUUID(uuid uuid.UUID) *file.File {
	return st.files[uuid]
}
