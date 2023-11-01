// Declara que el archivo pertenece al paquete principal, lo que indica que este es un programa ejecutable.
package main

// OS proporciona funciones y métodos para interactuar con el sistema operativo.
import (
	"encoding/json"
	"io"
	"os"

	"github.com/LuisZentenxx/go-cli-crud/tasks"
)

// Es el punto de entrada de cualquier programa Go.
func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0666)
	// Abre o crea, si no existe, un archivo llamado "tasks.json" en modo lectura y escritura.
	// `os.O_RDWR` indica que se abrirá en modo lectura y escritura.
	// `0666` representa los permisos dados al archivo una vez abierto.

	if err != nil {
		// Verifica si hubo un error al abrir o crear el archivo.
		// Si se produce un error, el programa se detiene y se imprime el error.
		panic(err)
	}

	defer file.Close()
	// Utiliza `defer` para asegurarse de que el archivo se cierre después de que la función `main()` termine o si se produce un error.
	// `defer` pospone la ejecución de `file.Close()` hasta que `main()` termine.
	
	var tasks []tasks.Task 
	info, err := file.Stat()

	if err != nil {
		panic(err)
	}
	
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		
		if err != nil {
			panic(err)
		}

		json.Unmarshal(byte)
	}

	else {
		
	}
}
