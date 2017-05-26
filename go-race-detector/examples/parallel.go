package main // ParallelWrite writes data to file1 and file2, returns the errors.
import "os"

func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}

func main() {
	ch := ParallelWrite([]byte("Golang Meetup!!"))
	<-ch
	<-ch

}
