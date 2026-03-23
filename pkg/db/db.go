package db

import (
	"database/sql"
	"errors"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT,
    repeat VARCHAR(128)
);

CREATE INDEX idx_date ON scheduler(date);
`

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	install := false
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		install = true
	}

	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	if install {
		_, err = DB.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetTasks(limit int, search string) ([]Task, error) {
	if DB == nil {
		return nil, errors.New("db is not initialized")
	}

	query := `SELECT id, date, title, comment, repeat FROM scheduler`
	args := make([]any, 0, 3)

	if search != "" {
		if t, err := time.Parse(`02.01.2006`, search); err == nil {
			query += ` WHERE date = ?`
			args = append(args, t.Format(`20060102`))
		} else {
			like := `%` + search + `%`
			query += ` WHERE title LIKE ? OR comment LIKE ?`
			args = append(args, like, like)
		}
	}

	query += ` ORDER BY date, id LIMIT ?`
	args = append(args, limit)

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0, limit)
	for rows.Next() {
		var t Task
		if err = rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
