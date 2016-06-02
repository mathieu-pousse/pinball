package hardware

type InputEvent struct {
	Direction string
	InputId string
}

type Input struct {
	Name      string
	listeners []EventHandler
}

func (input *Input) Configure() {

}

func (input *Input) AddEventHandler(eh EventHandler) {
	input.listeners = append(input.listeners, eh)
}

func (input *Input) Initialize() {

}

func (input *Input) OnEvent(event InputEvent) {
	for _, eh := range input.listeners {
		eh.Handle(event)
	}
}
