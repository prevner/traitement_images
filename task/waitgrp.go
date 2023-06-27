package task

import (
	"io/ioutil"
	"path/filepath"
	"sync"
)

type WaitGrpTask struct {
	srcDir  string
	dstDir  string
	filter  Filter
	workers int
}

func NewWaitGrpTask(srcDir, dstDir string, filter Filter) *WaitGrpTask {
	return &WaitGrpTask{
		srcDir:  srcDir,
		dstDir:  dstDir,
		filter:  filter,
		workers: 1, // Default to 1 worker
	}
}

func (t *WaitGrpTask) Process() error {
	fileInfos, err := ioutil.ReadDir(t.srcDir)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()

				srcPath := filepath.Join(t.srcDir, filename)
				dstPath := filepath.Join(t.dstDir, filename)

				err := t.filter.Process(srcPath, dstPath)
				if err != nil {
					println("Error processing file:", filename)
				}
			}(fileInfo.Name())
		}
	}

	wg.Wait()

	return nil
}
