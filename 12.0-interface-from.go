package main

// START OMIT
type Goodbyer interface {
	Goodbye()
}

type Forgiver interface {
	Forgive()
}

type Yura interface {
	Goodbyer
	Forgiver
}

// END OMIT
