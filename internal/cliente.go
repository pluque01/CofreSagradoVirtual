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
func (c *Cliente) Apellido1() string {
	return c.nombre.Apellido1()
}

func (c *Cliente) Apellido2() string {
	return c.nombre.Apellido2()
}

func NewCliente(NuevoTelefono string, NuevoNombre string) *Cliente {
	return &Cliente{
		telefono: NewTelefono(NuevoTelefono),
		nombre:   NewNombre(NuevoNombre),
	}
}
