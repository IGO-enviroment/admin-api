package organization

import (
	"admin-api/gen"
	"admin-api/sql"
	"fmt"
)

func (s Service) GetOrgProfile(email string) (gen.UserOrganization, error) {
	query := fmt.Sprintf(sql.GetOrgProfile, email)
	rows, err := s.pg.Query(query)
	if err != nil {
		return gen.UserOrganization{}, err
	}

	var profile gen.UserOrganization
	for rows.Next() {
		err := rows.Scan(
			&profile.Id,
			&profile.Email,
			&profile.FirstName,
			&profile.LastName,
			&profile.MiddleName,
			&profile.PhoneNumber,
			&profile.OrgName,
			&profile.OrgDescription,
			&profile.INN,
			&profile.KPP,
			&profile.OGRN,
			&profile.Place,
		)

		if err != nil {
			return gen.UserOrganization{}, err
		}
	}

	return profile, nil
}
