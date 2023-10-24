package main

type Nombre struct {
	nombre    string
	apellidos [2]string
}

func (n *Nombre) Nombre() string {
	return n.nombre
}

func (n *Nombre) Apellidos() [2]string {
	return n.apellidos
}

func NewNombre(NuevoNombre string) *Nombre {
	return &Nombre{
		nombre: NuevoNombre,
	}
}
