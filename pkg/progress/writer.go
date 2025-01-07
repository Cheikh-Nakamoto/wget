package progress

type progressWriter struct {
	progress chan<- int64
	done     chan<- bool
}

func (pw progressWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	pw.progress <- int64(n)
	if n == 0 {
		pw.done <- true
	}
	return n, nil
}

func (pw progressWriter) Close() error {
	close(pw.done)
	return nil
}
