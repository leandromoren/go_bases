package users

import (
	"fmt"
	"time"

	"github.com/leandromoren/go_bases.git/modelos"
)

func AltaUsuario() {
	// Objeto
	u := new(modelos.User)
	u.AddUser(10, "Leandro", time.Now(), true)
	fmt.Println(u)
}
