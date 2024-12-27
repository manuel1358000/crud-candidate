package clients

import (
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

func InitDB() *gorm.DB {
    // Cargar las variables de entorno desde el archivo .env (si existe)
    if err := godotenv.Load(); err != nil {
        log.Printf("Error al cargar el archivo .env: %v", err)
    }

    // Obtener la URL de la base de datos desde la variable de entorno DATABASE_URL
    databaseURL := os.Getenv("DATABASE_URL")
    
    // Si DATABASE_URL no está configurada, usar las variables locales
    if databaseURL == "" {
        dbUser := os.Getenv("DB_USER")
        dbPassword := os.Getenv("DB_PASSWORD")
        dbName := os.Getenv("DB_NAME")
        dbHost := os.Getenv("DB_HOST")
        dbPort := os.Getenv("DB_PORT")
        dbSSLMode := os.Getenv("DB_SSLMODE")

        // Construir la URL de conexión para el entorno local
        databaseURL = "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSLMode
    }

    // Usar la URL de la base de datos directamente en el DSN
    db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }

    // Obtener el objeto DB subyacente
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Error al obtener la base de datos subyacente: %v", err)
    }

    // Hacer ping para verificar la conexión
    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Error al hacer ping a la base de datos: %v", err)
    } else {
        log.Println("Conexión exitosa a la base de datos!")
    }

    return db
}
