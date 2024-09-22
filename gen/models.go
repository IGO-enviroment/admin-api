package gen

import "time"

const (
	AuthorizationScopes = "authorization.Scopes"
	OriginScopes        = "origin.Scopes"
)

// AddStudent defines model for addStudent.
type AddStudent struct {
	Students     []Student `json:"students"`
	UniversityId int       `json:"universityId"`
}

// CreationResult defines model for creationResult.
type CreationResult struct {
	IsCreated    bool   `json:"isCreated"`
	StudentEmail string `json:"studentEmail"`
}

// Login defines model for login.
type Login struct {
	// Email User email
	Email string `json:"email"`

	// Password User password
	Password string `json:"password"`
}

// Student defines model for student.
type Student struct {
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	MiddleName  string `json:"middleName"`
	PhoneNumber string `json:"phoneNumber"`
}

// TokenReponse defines model for tokenReponse.
type TokenReponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
}

type StudentProfile struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	MiddleName   string    `json:"middleName"`
	PhoneNumber  string    `json:"phoneNumber"`
	Birthdate    time.Time `json:"birthdate"`
	Sex          string    `json:"sex"`
	Description  string    `json:"description"`
	Competencies []string  `json:"competencies"`
}
