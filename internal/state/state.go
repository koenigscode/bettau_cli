package state

type State int
type Language string

const (
	Learn          State = iota
	Review         State = iota
	Import         State = iota
	AddLanguage    State = iota
	SelectLanguage State = iota
	Exit           State = iota
)

type ApplicationState struct {
	CurrentState    State
	CurrentLanguage Language
	Languages       []Language
}
