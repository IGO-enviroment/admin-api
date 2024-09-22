package students

import (
	"admin-api/gen"
	"admin-api/sql"
	"fmt"
)

func (s Service) GetStudentProfile(email string) (gen.StudentProfile, error) {
	query := fmt.Sprintf(sql.GetStudentProfile, email)
	rows, err := s.pg.Query(query)
	if err != nil {
		return gen.StudentProfile{}, err
	}

	var profile gen.StudentProfile
	for rows.Next() {
		err := rows.Scan(
			&profile.Id,
			&profile.Email,
			&profile.FirstName,
			&profile.LastName,
			&profile.MiddleName,
			&profile.PhoneNumber,
			&profile.Birthdate,
			&profile.Sex,
			&profile.Description,
			&profile.Competencies,
		)

		if err != nil {
			return gen.StudentProfile{}, err
		}
	}

	return profile, nil
}
