package Repo

import "database/sql"

type SiteSettingDBRepo struct {
	db *sql.DB
}

func NewSiteSettingDBRepo(db *sql.DB) *SiteSettingDBRepo {
	return &SiteSettingDBRepo{
		db: db,
	}
}
