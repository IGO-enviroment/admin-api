package gen

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
