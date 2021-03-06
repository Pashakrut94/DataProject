package region

import (
	"database/sql"

	"github.com/pkg/errors"
)

func validateCreateRegion(repo RegionRepo, region Region) error {
	_, err := repo.Get(region.ISOCode)
	if errors.Is(err, sql.ErrNoRows) {
		return sql.ErrNoRows
	}
	if err != nil {
		return errors.Wrap(err, "error getting region from DB")
	}
	return nil
}

func HandleCreateRegion(repo RegionRepo, region Region) (Region, error) {
	err := validateCreateRegion(repo, region)
	switch errors.Cause(err) {
	case sql.ErrNoRows:
		if err = repo.Create(&region); err != nil {
			return Region{}, err
		}
		return region, nil
	case nil:
		if err = repo.Update(&region); err != nil {
			return Region{}, errors.Wrap(err, "error updating region")
		}
		return region, nil
	default:
		return Region{}, errors.Wrap(err, "error getting region from DB")
	}
}

func HandleGetRegion(repo RegionRepo, isoCode string) (Region, error) {
	region, err := repo.Get(isoCode)
	if errors.Is(err, sql.ErrNoRows) {
		return Region{}, ErrNotFound
	}
	if err != nil {
		return Region{}, err
	}
	return *region, nil
}

func HandleGetTotal(repo RegionRepo) (Total, error) {
	total, err := repo.Total()
	if errors.Is(err, sql.ErrNoRows) {
		return Total{}, ErrEmptyDB
	}
	if err != nil {
		return Total{}, err
	}
	return *total, nil
}
