package domain
import "time"
import "github.com/google/uuid"

type Usuario struct {
	ID       uuid.UUID
	TenantID uuid.UUID
	Email    string
	Password string
	Name     string
	Role     string
	Active   bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsuarioFull struct {
	Usuario
	BirthDate time.Time
	Gender    string
	Address   string
	Phone     string
	Nationality string
	MaritalStatus string
	Occupation string
}