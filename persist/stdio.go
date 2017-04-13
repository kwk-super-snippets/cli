package persist

import (
	"github.com/kennygrant/sanitize"
	"io/ioutil"
	"strings"
	"path"
	"os"
	"bitbucket.com/sharingmachine/kwkcli/models"
	"bitbucket.com/sharingmachine/kwkcli/log"
	"bitbucket.com/sharingmachine/kwkcli/cache"
)

var ErrFileNotFound = models.ErrOneLine(models.Code_NotFound, "File not found.")
var ErrFileExpired = models.ErrOneLine(models.Code_NotFound, "File found but expired.")

func New() IO {
	return &StdIO{}
}

type StdIO struct {
}

func (s *StdIO) Write(subDirName string, suffixPath string, snippet string, incHoldingDir bool) (filePath string, err error) {
	fp := s.getFilePath(subDirName, suffixPath, incHoldingDir)
	log.Debug("WRITE: %s", fp)
	err = ioutil.WriteFile(fp, []byte(snippet), cache.StandardFilePermission)
	return fp, err
}

func (s *StdIO) Read(subDirName string, suffixPath string, incHoldingDir bool, after int64) (string, error) {
	fp := s.getFilePath(subDirName, suffixPath, incHoldingDir)
	log.Debug("READ: %s", fp)
	if fi, err := os.Stat(fp); err != nil {
		if os.IsNotExist(err) {
			// TODO: PUT IN STANDARD ERROR
			return "", ErrFileNotFound
		} else {
			return "", err
		}
	} else {
		if after == 0 || after < int64(fi.ModTime().Unix()) {
			bts, err := ioutil.ReadFile(fp)
			return string(bts), err
		} else {
			return "", ErrFileExpired
		}

	}
}

func (s *StdIO) DeleteAll() error {
	return os.RemoveAll(cache.Path())
}

func (s *StdIO) Delete(directoryName string, fileName string) error {
	dirPath, err := s.getSubDir(directoryName)
	if err != nil {
		return err
	}
	fp := path.Join(dirPath, fileName)
	log.Debug("DELETING:%s", fp)
	return os.RemoveAll(fp)
}

func (s *StdIO) upsertDirectory(dir string) error {
	if err := os.MkdirAll(dir, cache.StandardFilePermission); err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}

func (s *StdIO) getSubDir(directoryName string) (string, error) {
	dir := path.Join(cache.Path(), directoryName)
	log.Debug("DIR: %s", dir)
	err := s.upsertDirectory(dir)
	return dir, err
}

func (s *StdIO) getHoldingDirectory(subDirName string, fullName string) string {
	hd := strings.Replace(fullName, ".", "_", -1)
	dirPath := path.Join(cache.Path(), subDirName, hd)
	if e := s.upsertDirectory(dirPath); e != nil {
		log.Error("Could not create directory", e)
		return ""
	}
	return hd
}

func (s *StdIO) getFilePath(subDirName string, suffixPath string, incHoldingDir bool) string {
	sn := sanitize.Name(suffixPath)
	if incHoldingDir {
		hd := s.getHoldingDirectory(subDirName, sn)
		return path.Join(cache.Path(), subDirName, hd, sn)
	} else {
		if err := s.upsertDirectory(path.Join(cache.Path(), subDirName)); err != nil {
			panic(err)
		}
		return path.Join(cache.Path(), subDirName, sn)
	}
}
