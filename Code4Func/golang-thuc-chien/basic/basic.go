package basic

import (
	"fmt"
	"time"
)

type Student struct {
	firstName string
	lastName  string
	email     string
}

func (s *Student) Email() string {
	s.email = "duynghia@gmail.com"
	return s.email
}

func g1() {
	fmt.Println("1")
}

func DemoBasic() {
	s := Student{
		firstName: "nghia",
		lastName:  "Le",
		email:     "abc@gmail.com",
	}
	fmt.Println(s)

	valueEmail := s.Email()
	fmt.Println(valueEmail)

	fmt.Println(s.email)

	// go routine
	go g1()

	go func() {
		fmt.Println("2")
	}()

	// channel: propose of channel is interactive with go routine
	c := make(chan int)
	go func() {
		c <- 100
	}()

	go func() {
		fmt.Println(<-c)
	}()

	time.Sleep(time.Second)

}
