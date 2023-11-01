# CLI CRUD GOLANG

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
    cd cli-crud-go

3. Compila el código para generar el ejecutable:
   ```bash
   go build -o cli-crud-go

4. Ejecuta la aplicación:
   ```bash
   ./cli-crud-go

## Uso

La aplicación admite los siguientes comandos:

- **add**: Agrega una nueva tarea.
- **list**: Muestra todas las tareas.
- **update**: Actualiza una tarea existente.
- **delete**: Elimina una tarea.
- **help**: Muestra la lista de comandos disponibles y su uso.

Ejemplo:
```bash
    # Agregar una tarea
./cli-crud-go add "Completar el informe"

# Mostrar todas las tareas
./cli-crud-go list

# Actualizar una tarea
./cli-crud-go update 1 "Completar el informe para el cliente"

# Eliminar una tarea
./cli-crud-go delete 1

