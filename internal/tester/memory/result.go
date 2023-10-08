package memory

type Result struct {
	bufSize  int
	workerID int
}

func (r *Result) BufferSize() int {
	return r.bufSize
}

func (r *Result) WorkerID() int {
	return r.workerID
}
