package clients

import (
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

func InitDB() *gorm.DB {
    if err := godotenv.Load(); err != nil {
        log.Printf("Error al cargar el archivo .env: %v", err)
    }

    databaseURL := os.Getenv("DATABASE_URL")
    
    if databaseURL == "" {
        dbUser := os.Getenv("DB_USER")
        dbPassword := os.Getenv("DB_PASSWORD")
        dbName := os.Getenv("DB_NAME")
        dbHost := os.Getenv("DB_HOST")
        dbPort := os.Getenv("DB_PORT")
        dbSSLMode := os.Getenv("DB_SSLMODE")

        databaseURL = "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSLMode
    }

    db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Error al obtener la base de datos subyacente: %v", err)
    }

    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Error al hacer ping a la base de datos: %v", err)
    } else {
        log.Println("Conexión exitosa a la base de datos!")
    }

    return db
}
