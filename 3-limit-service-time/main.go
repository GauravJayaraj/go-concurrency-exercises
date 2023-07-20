//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import "time"

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	t := make(chan bool, 2)

	go func(t chan bool) {
		process()
		t <- true
	}(t)

	if !u.IsPremium {
		// per request 10sec  if user is non-premium

		go func(t chan bool) {
			time.Sleep(10 * time.Second)
			t <- false
		}(t)
	}

	return <-t
}

func main() {
	RunMockServer()
}
