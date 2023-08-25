package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Abrir la conexión a la base de datos
    db, err := sql.Open("sqlite3", "./crud.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Crear la tabla si no existe
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS items (
            id INTEGER PRIMARY KEY,
            name TEXT,
            description TEXT
        )
    `)
    if err != nil {
        log.Fatal(err)
    }

    // Insertar un registro
    _, err = db.Exec("INSERT INTO items(name, description) VALUES(?, ?)", "Ejemplo", "Descripción de ejemplo")
    if err != nil {
        log.Fatal(err)
    }

    // Leer registros
    rows, err := db.Query("SELECT id, name, description FROM items")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name, description string
        err := rows.Scan(&id, &name, &description)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Description: %s\n", id, name, description)
    }

    // Actualizar un registro
    _, err = db.Exec("UPDATE items SET description = ? WHERE id = ?", "Nueva descripción", 1)
    if err != nil {
        log.Fatal(err)
    }

    // Eliminar un registro
    _, err = db.Exec("DELETE FROM items WHERE id = ?", 1)
    if err != nil {
        log.Fatal(err)
    }
}
