package task

import (
	"io/ioutil"
	"path/filepath"
)

type ChanTask struct {
	srcDir  string
	dstDir  string
	filter  Filter
	workers int
}

func NewChanTask(srcDir, dstDir string, filter Filter, workers int) *ChanTask {
	return &ChanTask{
		srcDir:  srcDir,
		dstDir:  dstDir,
		filter:  filter,
		workers: workers,
	}
}

func (t *ChanTask) Process() error {
	fileInfos, err := ioutil.ReadDir(t.srcDir)
	if err != nil {
		return err
	}

	jobs := make(chan string, t.workers)
	results := make(chan error)

	// Worker function
	worker := func() {
		for filename := range jobs {
			srcPath := filepath.Join(t.srcDir, filename)
			dstPath := filepath.Join(t.dstDir, filename)

			err := t.filter.Process(srcPath, dstPath)
			if err != nil {
				results <- err
			} else {
				results <- nil
			}
		}
	}

	// Start workers
	for i := 0; i < t.workers; i++ {
		go worker()
	}

	// Send jobs
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			jobs <- fileInfo.Name()
		}
	}
	close(jobs)

	// Wait for results
	for range fileInfos {
		err := <-results
		if err != nil {
			println("Error processing file:", err)
		}
	}

	return nil
}
