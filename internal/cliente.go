package main

type Cliente struct {
	telefono *Telefono
	nombre   *Nombre
}

func (c *Cliente) Telefono() string {
	return c.telefono.Telefono()
}

func (c *Cliente) Nombre() string {
	return c.nombre.Nombre()
}
func (c *Cliente) Apellidos() [2]string {
	return c.nombre.Apellidos()
}

func NewCliente(NuevoTelefono string, NuevoNombre string) *Cliente {
	return &Cliente{
		telefono: NewTelefono(NuevoTelefono),
		nombre:   NewNombre(NuevoNombre),
	}
}
