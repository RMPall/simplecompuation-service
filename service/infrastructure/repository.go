package infrastructure

import (
	"errors"
	"io/ioutil"
	"os"
)

const (
	fileDir = "/tmp/storeDat"
)

type Repo struct{}

func (r *Repo) Save(total int) error {
	f, err := os.Create(fileDir)
	if err != nil {
		return err
	}
	defer f.Close()

	totalByte := byte(total)
	_, err = f.Write([]byte{totalByte})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Get() (int, error) {
	dat, err := ioutil.ReadFile(fileDir)
	if err != nil {
		return 0, errors.New("failed to read file")
	}
	return int(dat[0]), nil
}
