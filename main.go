package main

import (
	"fmt"
	"tienda-celulares/catalogo"
)

func main() {
	fmt.Println("Sistema de Gestion de Venta de Celulares")
	fmt.Println("Marcas disponibles:")
	fmt.Println(catalogo.MostrarMarcas())
}
