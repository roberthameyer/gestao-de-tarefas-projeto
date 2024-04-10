package model

type Tarefa struct {
	Titulo       string
	Descricao    string
	StatusTarefa int
	Prazo        string
	Categoria    string
}

// Cria o tipo enum para o status da tarefa
const (
	ToDo  = 0
	Doing = 1
	Done  = 2
)

// Imprime os status correspondentes do enum
func (t Tarefa) Status() string {
	switch t.StatusTarefa {
	case ToDo:
		return "To Do"
	case Doing:
		return "Doing"
	case Done:
		return "Done"
	}
	return "Unknown"
}
