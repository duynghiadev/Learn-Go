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
	ID            string
	HP            int
	inventoryPID  *actor.PID
	shopPID       *actor.PID
	weaponShopPID *actor.PID
	guildPID      *actor.PID
	Power         int
}

type Shop struct {
	stock int
}

type WeaponShop struct {
	weapons int
}

type Guild struct {
	Name  string
	Bonus int
}

type GlobalEventManager struct{}

type drinkBottle struct {
	amount int
}

type buyBottle struct{}

type buyWeapon struct{}

type attack struct {
	target *actor.PID
}

type MyEvent struct {
	foo string
}

func (i *Inventory) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("[Inventory] Started - Loading from DB...")
		c.Engine().Subscribe(c.PID())
	case actor.Stopped:
		fmt.Println("[Inventory] Stopped - Saving to DB...")
	case drinkBottle:
		if i.Bottles > 0 {
			i.Bottles -= msg.amount
			fmt.Println("[Inventory] Drank a bottle. Remaining:", i.Bottles)
		} else {
			fmt.Println("[Inventory] No bottles left!")
		}
	case MyEvent:
		fmt.Println("[Inventory] Received global event:", msg.foo)
	}
}

func NewInventory(bottles int) actor.Producer {
	return func() actor.Receiver {
		return &Inventory{
			Bottles: bottles,
		}
	}
}

func (p *Player) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Printf("[Player %s] Started - HP: %d\n", p.ID, p.HP)
		p.inventoryPID = c.SpawnChild(NewInventory(2), "inventory_"+p.ID)
		c.Engine().Subscribe(c.PID())
	case actor.Stopped:
		fmt.Printf("[Player %s] Stopped - Saving data...\n", p.ID)
	case drinkBottle:
		c.Forward(p.inventoryPID)
	case buyBottle:
		if p.shopPID != nil {
			c.Request(p.shopPID, buyBottle{}, time.Second*5)
		}
	case buyWeapon:
		if p.weaponShopPID != nil {
			c.Request(p.weaponShopPID, buyWeapon{}, time.Second*5)
		}
	case attack:
		fmt.Printf("[Player %s] is attacking!\n", p.ID)
		c.Send(msg.target, attack{target: c.PID()})
	case MyEvent:
		fmt.Printf("[Player %s] Received global event: %s\n", p.ID, msg.foo)
	}
}

func NewPlayer(id string, hp int, shopPID, weaponShopPID, guildPID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &Player{
			ID:            id,
			HP:            hp,
			shopPID:       shopPID,
			weaponShopPID: weaponShopPID,
			guildPID:      guildPID,
			Power:         10}
	}
}

func (s *Shop) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("[Shop] Opened - Stock:", s.stock)
		c.Engine().Subscribe(c.PID())
	case buyBottle:
		if s.stock > 0 {
			s.stock--
			fmt.Println("[Shop] Bottle sold! Remaining stock:", s.stock)
			c.Respond(drinkBottle{amount: 1})
		} else {
			fmt.Println("[Shop] Out of stock!")
		}
	case MyEvent:
		fmt.Println("[Shop] Received global event:", msg.foo)
	}
}

func NewShop(stock int) actor.Producer {
	return func() actor.Receiver {
		return &Shop{
			stock: stock,
		}
	}
}

func (w *WeaponShop) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
		fmt.Println("[Weapon Shop] Opened - Weapons:", w.weapons)
	case buyWeapon:
		if w.weapons > 0 {
			w.weapons--
			fmt.Println("[Weapon Shop] Weapon sold! Remaining:", w.weapons)
		}
	}
}

func NewWeaponShop(weapons int) actor.Producer {
	return func() actor.Receiver {
		return &WeaponShop{weapons: weapons}
	}
}

func (g *Guild) Receive(c *actor.Context) {
	fmt.Println("[Guild]", g.Name, "is active with bonus:", g.Bonus)
}

func NewGuild(name string, bonus int) actor.Producer {
	return func() actor.Receiver {
		return &Guild{
			Name:  name,
			Bonus: bonus,
		}
	}
}

func (g *GlobalEventManager) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
		fmt.Println("[GlobalEventManager] Started")
	case MyEvent:
		fmt.Println("[GlobalEventManager] Broadcasting event...")
		c.Engine().BroadcastEvent(MyEvent{foo: "Server Maintenance Incoming!"})
	}
}

func NewGlobalEventManager() actor.Producer {
	return func() actor.Receiver {
		return &GlobalEventManager{}
	}
}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	shopPID := e.Spawn(NewShop(5), "shop")
	weaponShopPID := e.Spawn(NewWeaponShop(3), "weapon_shop")
	guildPID := e.Spawn(NewGuild("Warriors", 5), "guild_warriors")
	eventManager := e.Spawn(NewGlobalEventManager(), "global_events")

	player1 := e.Spawn(NewPlayer("Alice", 100, shopPID, weaponShopPID, guildPID), "player_Alice")
	player2 := e.Spawn(NewPlayer("Bob", 120, shopPID, weaponShopPID, guildPID), "player_Bob")

	e.Send(player1, drinkBottle{amount: 1})
	e.Send(player2, buyBottle{})
	e.Send(eventManager, MyEvent{})
	time.Sleep(time.Second * 2)
}
