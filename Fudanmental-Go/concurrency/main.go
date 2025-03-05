package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

// 1. spawning children
// 2. poisoning actors (stopping actors)
// 3. Event stream

type Inventory struct {
	Bottles int
}

type Player struct {
	HP           int
	inventoryPID *actor.PID
}

func (i *Inventory) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		_ = msg
		fmt.Println("inventory actor is started (read state from DB)")
		c.Engine().Subscribe(c.PID())
	case actor.Stopped:
		fmt.Println("inventory actor is stopped (save state to DB)")
	case drinkBottle:
		fmt.Println("inventory received drink bottle message")
		i.Bottles -= msg.amount
	case MyEvent:
		fmt.Println("just got a string from the event stream in the inventory BTW:", msg.foo)
	}
}

func NewInventory(bottles int) actor.Producer {
	return func() actor.Receiver {
		return &Inventory{
			Bottles: bottles,
		}
	}
}

func NewPlayer(hp int) actor.Producer {
	return func() actor.Receiver {
		return &Player{
			HP: hp,
		}
	}
}

type drinkBottle struct {
	amount int
}

func (p *Player) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("player actor is started (read state from DB)")
		p.inventoryPID = c.SpawnChild(NewInventory(1), "inventory")
		c.Engine().Subscribe(c.PID())
	case actor.Stopped:
		fmt.Println("player actor is stopper (save state to DB)")
	case drinkBottle:
		c.Forward(p.inventoryPID)
	case MyEvent:
		fmt.Println("just got a string from the event stream:", msg.foo)
	}
}

type MyEvent struct {
	foo string
}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	pid := e.Spawn(NewPlayer(100), "player", actor.WithID("myuserid358"))

	e.Send(pid, drinkBottle{amount: 1})
	time.Sleep(time.Second * 2)
	e.BroadcastEvent(MyEvent{foo: "yeah buddy"})
	time.Sleep(time.Second * 2)
}
