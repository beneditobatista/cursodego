package operadora

import (
	"strconv"

	"github.com/beneditobatista/cursodego/pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora
var NomeOperadora = "Qualquer uma"

//PrefixoDaOperadora prefixo mais o nome da opeadora
var PrefixoDaOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
