package main

type Nombre struct {
	nombre    string
	apellidos [2]string
}

func (n *Nombre) Nombre() string {
	return n.nombre
}

func (n *Nombre) Apellido1() string {
	return n.apellidos[0]
}

func (n *Nombre) Apellido2() string {
	return n.apellidos[1]
}

func NewNombre(NuevoNombre string) *Nombre {
	return &Nombre{
		nombre: NuevoNombre,
	}
}
