package repository

import "log"

func (repo *Repository) GetSetting(key string) (value string) {
	var err = repo.dbConn.QueryRow("select `value` from `sys_site_setting` where key = ?", key).Scan(&value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
