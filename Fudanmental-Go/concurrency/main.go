package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type Player struct {
	HP int
}

func NewPlayer(hp int) actor.Producer {
	return func() actor.Receiver {
		return &Player{
			HP: hp,
		}
	}
}

type takeDamage struct {
	amount int
}

func (p *Player) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("player actor is started (read state from DB)")
	case actor.Stopped:
		fmt.Println("player actor is stopper (save state to DB)")
	case takeDamage:
		fmt.Println("Player is taking damage:", msg.amount)
	}
}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	pid := e.Spawn(NewPlayer(100), "player", actor.WithID("myuserid358"))

	msg := takeDamage{amount: 999}
	for i := 0; i < 100; i++ {
		e.Send(pid, msg)
	}

	time.Sleep(time.Second * 2)
}
