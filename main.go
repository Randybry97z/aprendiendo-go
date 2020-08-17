package main

import "fmt"

func main() {
	/*
		//SWITCH

		var calificacion int

		fmt.Println("Ingresa una calificación: ")
		fmt.Scanf("%d", &calificacion)

		/* switch {
		case calificacion == 10:
			fmt.Println("Calificacion perfecta")
		case calificacion == 8 || calificacion == 9:
			fmt.Println("Calificación excelente")
		case calificacion == 6 || calificacion == 7:
			fmt.Println("Calificación aprobatoria")
		case calificacion >= 0 && calificacion <= 5:
			fmt.Println("Reprobado")
		default:
			fmt.Println("Calificación no válida")
		} */
	/* switch calificacion{
	case calificacion 10:
		fmt.Println("Calificacion perfecta")
	case calificacion 8,9:
		fmt.Println("Calificación excelente")
	case calificacion 6,7:
		fmt.Println("Calificación aprobatoria")
	case 0,1,2,3,4,5:
		fmt.Println("Reprobado")
	default:
		fmt.Println("Calificación no válida")
	} */

	//OBTENER VALORES DE UN MAPA

	usuarios := make(map[string]string)

	usuarios["bryan"] = "bssandovalrod@gmail.com"

	//Un valor de un mapa retorna el valor y un valor booleano
	correo, ok := usuarios["bryan"]

	if ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
