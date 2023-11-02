// Declara que el archivo pertenece al paquete principal, lo que indica que este es un programa ejecutable.
package main

// OS proporciona funciones y métodos para interactuar con el sistema operativo.
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/LuisZentenxx/go-cli-crud/tasks"
)

// Es el punto de entrada de cualquier programa Go.
func main() {

	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
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

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		//
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)

	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Which is yout task?")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTask(file, tasks)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Must provide an ID to delete or use 'all' to delete all tasks")
			return
		}

		if os.Args[2] == "all" {
			// Eliminar todas las tareas
			tasks = task.DeleteAllTasks(tasks)

		} else {
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("ID may be a number")
				return
			}

			tasks = task.DeleteTask(tasks, id)
		}
		task.SaveTask(file, tasks)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("May get one ID to delete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("ID may be a number")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTask(file, tasks)

	}
}

func printUsage() {
	fmt.Println("Uso : go-cli-crud [list|add|complete|delete]")
}
