package telefono

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
func (t *Telefono) GetTelefono() string {
	return t.telefono
}