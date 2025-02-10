package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	_ "modernc.org/sqlite"
)

const (
	webDir = "./web" //директория веб-приложения
)




func main() {

   //получение пути до исполняемого файла приложения
   appPath, err := os.Executable()
   if err != nil {
       log.Fatal(err)
   }
   //определение наличия базы данных
   dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
   _, err = os.Stat(dbFile)
   
   if err != nil {
      log.Println("База данных не обнаружена. Создание новой.")
   file, err := os.Create("scheduler.db")
if err != nil {
    log.Fatal(err)
    return
}
log.Println("Новый файл базы успешно создан")
file.Close()

db, err := sql.Open("sqlite", "scheduler.db")
    if err != nil {
        log.Println(err)
        return
    }
    defer db.Close()

table := `CREATE TABLE scheduler (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
     date CHAR(8) NOT NULL DEFAULT "",
     title VARCHAR(128) NOT NULL DEFAULT "",
     comment VARCHAR(128) NOT NULL DEFAULT "",
     repeat VARCHAR(128) NOT NULL DEFAULT "");
     CREATE INDEX idx_date ON scheduler (date)`
     _, err = db.Exec(table)
   if err != nil {
       log.Fatal(err)
       return 
   }
   log.Println("Таблица успешно создана!")  
}
   
   

   // если install равен true, после открытия БД требуется выполнить 
   // sql-запрос с CREATE TABLE и CREATE INDEX 

   //определение порта для запуска приложения
   port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

   //
	r := chi.NewRouter()
   r.Handle("/*", http.FileServer(http.Dir(webDir)))
	//r.Get("/js/scripts.min.js", handleScripts)
	
   //запуск сервера приложения
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}
