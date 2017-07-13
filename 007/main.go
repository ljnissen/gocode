package main

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (c *Car) GoToWorkIn (
	// get in car
	c.Start
	// drive to work
	c.Stop
	// get out of car
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				// ...
			case paint.Event:
				log.Print("Call OpenGL here.")
				a.Publish()
			}
		}
	})
}