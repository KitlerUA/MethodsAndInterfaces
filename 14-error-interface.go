package main

// START OMIT
type error interface {
	Error() string
}

type UploadError struct {
	Status  int
	Code    int
	Message string
}

func (e UploadError) Error() string {
	return e.Message
}

// END OMIT
func main() {

}
