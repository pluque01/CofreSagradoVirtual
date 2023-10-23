package main

type Nombre struct {
	nombre    string
	apellido1 string
	apellido2 string
}

func (n *Nombre) Nombre() string {
	return n.nombre
}

func (n *Nombre) Apellido1() string {
	return n.apellido1
}

func (n *Nombre) Apellido2() string {
	return n.apellido2
}

func NewNombre(NuevoNombre string) *Nombre {
	return &Nombre{
		nombre: NuevoNombre,
	}
}
