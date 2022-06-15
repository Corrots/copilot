package main

import "database/sql"

type CategorySummary struct {
	Title    string  `json:"title"`
	Tasks    int     `json:"tasks"`
	AvgValue float64 `json:"avg_value"`
}

func CreateTables(db *sql.DB) {
	sql := `db.Exec("CREATE TABLE category_summary (id INTEGER PRIMARY KEY, title TEXT, value INTEGER, category TEXT)")`
	db.Exec(sql)
}

func GetCategories(db *sql.DB) ([]CategorySummary, error) {
	sql := `SELECT category, AVG(value) AS avg_value FROM category_summary GROUP BY category`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// TODO: return a list of CategorySummary
	var summaries []CategorySummary
	for rows.Next() {
		var summary CategorySummary
		err := rows.Scan(&summary.Title, &summary.Tasks, &summary.AvgValue)
		if err != nil {
			return nil, err
		}
		summaries = append(summaries, summary)
	}
	return summaries, nil
}

func CreateCategorySummary(db *sql.DB, title string, value int, category string) error {
	sql := `INSERT INTO category_summary (title, value, category) VALUES (?, ?, ?)`
	_, err := db.Exec(sql, title, value, category)
	return err
}

func DeleteCategorySummary(db *sql.DB, title string) error {
	sql := `DELETE FROM category_summary WHERE title = ?`
	_, err := db.Exec(sql, title)
	return err
}

func main() {

}
