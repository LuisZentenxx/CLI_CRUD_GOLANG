# CLI CRUD GOLANG
![](https://www.softwebsolutions.com/wp-content/uploads/2020/10/golang-Programing.jpg)

## Descripción
CLI CRUD GOLANG es una aplicación de línea de comandos escrita en Golang que permite a los usuarios gestionar una lista de tareas directamente desde la terminal. Ofrece funcionalidades básicas de creación, lectura, actualización y eliminación (CRUD) de tareas, lo que facilita la organización y gestión de tareas diarias o pendientes.

La aplicación está diseñada para ser intuitiva y fácil de usar, permitiendo a los usuarios interactuar con su lista de tareas a través de simples comandos de terminal.

## Características

- **CRUD**: Operaciones básicas de Crear, Leer, Actualizar y Eliminar tareas.
- **Interfaz de línea de comandos**: Fácil interacción a través de comandos de terminal.
- **Instalable localmente**: Posibilidad de instalar la aplicación en tu propio computador.
- **Compatibilidad multiplataforma**: Compatible con Windows, macOS y Linux.
- **Compartición del ejecutable final**: Capacidad para compartir el ejecutable final con otros usuarios.

## Requisitos
- Golang instalado en el sistema
- Conexión a internet para la instalación de dependencias, si es necesario.

## Instalación
1. Clona este repositorio en tu computador:
   ```bash
   git clone https://github.com/LuisZentenxx/CLI_CRUD_GOLANG
2. Navega hacia el directorio del proyecto:
    ```bash
    cd CLI_CRUD_GOLANG
3. Compila el código para generar el ejecutable:
   ```bash
   go build .
4. Ejecuta la aplicación:
   ```bash
   ./go-cli-crud.exe
## Uso

La aplicación admite los siguientes comandos:

- **add**: Agrega una nueva tarea.
   ```bash
   go run main.go add
- **list**: Muestra todas las tareas.
- ```bash
   go run main.go list
- **complete**: Actualiza el estado de una tarea existente usando su respectivo ID.
  ```bash
   go run main.go complete <ID>
- **delete**: Elimina una tarea indicando su respectivo ID o usando "all" para eliminar todas las tareas.
  ```bash
   go run main.go delete <ID or all>
Ejemplo:
```bash
# Agregar una tarea
   > go run main.go add
      -> My new task

# Mostrar todas las tareas
   > go run main.go list

   Resultado:
      [ ] 1 Review the Go code
      [✓] 2 Create a folder
      [ ] 3 My new task

# Actualizar estado de una tarea
   > go run main.go complete 3

   Resultado:
      [ ] 1 Review the Go code
      [✓] 2 Create a folder
      [✓] 3 My new task

# Eliminar una tarea
   > go run main.go delete 1

   Resultado:
      [✓] 2 Create a folder
      [✓] 3 My new task

# Eliminar todas las tareas
   > go run main.go delete all

   Resultado:
      [ ]

