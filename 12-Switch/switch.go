package main

import "fmt"

func diaDaSemana(numero int) string{
	switch numero{ // avalia o numero
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terca-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default: // se nenhuma foi 
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string{
	var diaDaSemana string // outro modo de fazer
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // clausura que joga pro proximo, ent cai pra segunda
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terca-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sabado"
	}

	return diaDaSemana
}

func main()  {

	fmt.Println("Switch")

	dia := diaDaSemana(2)
	fmt.Println(dia)

	diaErr := diaDaSemana(200)
	fmt.Println(diaErr)
}
	
