package region

import (
	"database/sql"
)

type RegionRepo struct {
	db *sql.DB
}

func NewRegionRepo(db *sql.DB) *RegionRepo {
	return &RegionRepo{db: db}
}

func (repo *RegionRepo) Create(region *Region) error {
	q := "insert into regions (isocode, region, infected, recovered, deceased, country) values ($1,$2,$3,$4,$5,$6)"
	_, err := repo.db.Exec(q, region.ISOCode, region.Region, region.Infected, region.Recovered, region.Deceased, region.Country)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RegionRepo) Get(ISOCode string) (*Region, error) {
	row := repo.db.QueryRow("select isocode, region, infected, recovered, deceased, country from regions where isocode = $1", ISOCode)
	var region Region
	err := row.Scan(&region.ISOCode, &region.Region, &region.Infected, &region.Recovered, &region.Deceased, &region.Country)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &region, nil
}

func (repo *RegionRepo) Update(region *Region) error {
	q := "update regions set infected = $1, recovered = $2, deceased = $3 where isocode = $4"
	_, err := repo.db.Exec(q, region.Infected, region.Recovered, region.Deceased, region.ISOCode)
	if err != nil {
		return err
	}
	return nil
}

func (repo *RegionRepo) Total() (*Total, error) {
	row := repo.db.QueryRow("select sum (infected), sum (recovered), sum (deceased) from regions")
	var total Total
	total.Country = "Russia" // ?Корректное задание?: сделай агрегацию по стране(where country = "Russia")? или нужен другой запрос в базу?
	err := row.Scan(&total.Infected, &total.Recovered, &total.Deceased)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &total, nil
}
