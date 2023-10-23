package nombre

// Defininción de la estructura Nombre
type Nombre struct {
	nombre string
}

// Definición de los métodos de la estructura Nombre
// Función que permite asignar un nombre a la estructura Nombre
func (n *Nombre) SetNombre(nombre string) {
	n.nombre = nombre
}

// Función que permite obtener el nombre de la estructura Nombre
func (n *Nombre) Nombre() string {
	return n.nombre
}

// Función que permite crear una nueva estructura Nombre
func NewNombre(NuevoNombre string) *Nombre {
	return &Nombre{
		nombre: NuevoNombre,
	}
}
