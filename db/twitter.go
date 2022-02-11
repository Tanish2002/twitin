package db

import "fmt"

func AddRecentID(recentID string) {
	// Always updating id=1
	err := db.QueryRow("INSERT INTO twitter_details(id,recent_id) VALUES(1,$1) returning id;", recentID)
	if err != nil {
		fmt.Println(err)
	}
}
func UpdateRecentID(recentID string) {
	stmt, err := db.Prepare("update twitter_details set recent_id=$1 where id=1")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(recentID)
	if err != nil {
		fmt.Println(err)
	}
}
func GetRecentID() string {
	rows, err := db.Query("SELECT * FROM twitter_details")
	if err != nil {
		fmt.Println("AAAA")
		fmt.Println(err)
	}
	var recent_id string
	for rows.Next() {
		var id int
		err = rows.Scan(&id, &recent_id)
		if err != nil {
			fmt.Println(err)
		}
	}
	return recent_id
}
