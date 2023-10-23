package apellidos

// Defininción de la estructura Telefono
type Telefono struct {
	telefono string
}

// Definición de los métodos de la estructura Telefono
// Función que permite asignar un telefono a la estructura Telefono
func (t *Telefono) SetTelefono(telefono string) {
	t.telefono = telefono
}

// Función que permite obtener el telefono de la estructura Telefono
func (t *Telefono) Telefono() string {
	return t.telefono
}

// Función que permite crear una nueva estructura Telefono
func NewTelefono(NuevoTelefono string) *Telefono {
	return &Telefono{
		telefono: NuevoTelefono,
	}
}
