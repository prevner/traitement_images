package task

type Filter interface {
	Process(srcPath, dstPath string) error
}

type Tasker interface {
	Process() error
}
