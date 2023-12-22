package modelos

import (
	"time"
)

// Las estructuras empiezan con <type> <nombre> <struct>
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (structRef *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	structRef.Id = id
	structRef.Name = name
	structRef.CreatedAt = createdAt
	structRef.Status = status
}
