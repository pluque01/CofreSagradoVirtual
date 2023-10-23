package main

type Telefono struct {
	telefono string
}

func (t *Telefono) SetTelefono(telefono string) {
	t.telefono = telefono
}

func (t *Telefono) Telefono() string {
	return t.telefono
}

func NewTelefono(NuevoTelefono string) *Telefono {
	return &Telefono{
		telefono: NuevoTelefono,
	}
}
