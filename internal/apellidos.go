package apellidos

// Defininción de la estructura Apellidos
type Apellidos struct {
	apellidos string
}

// Definición de los métodos de la estructura Apellidos
// Función que permite asignar un apellido a la estructura Apellidos
func (a *Apellidos) SetApellidos(apellidos string) {
	a.apellidos = apellidos
}

// Función que permite obtener el apellido de la estructura Apellidos
func (a *Apellidos) Apellidos() string {
	return a.apellidos
}

// Función que permite crear una nueva estructura Apellidos
func NewApellidos(NuevosApellidos string) *Apellidos {
	return &Apellidos{
		apellidos: NuevosApellidos,
	}
}
