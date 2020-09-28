package main

import "fmt"

func main()  {
	var ( 
		dia uint;
		m uint;
	)

	fmt.Scan(&dia);
	fmt.Scan(&m);

	if (m == 1 && dia >= 20) || (m == 2 && dia <= 18) {
        fmt.Println("acuario");
    } else if (m == 2 && dia >= 19) || (m == 3 && dia <= 20) {
        fmt.Println("piscis");
    } else if (m == 3 && dia >= 21) || (m == 4 && dia <= 19) {
        fmt.Println("aries");
	} else if (m == 4 && dia >= 20) || (m == 5 && dia <= 20) {
        fmt.Println("tauro");
    } else if (m == 5 && dia >= 21) || (m == 6 && dia <= 20) {
        fmt.Println("geminis");
    } else if (m == 6 && dia >= 21) || (m == 7 && dia <= 22) {
        fmt.Println("cancer");
    } else if (m == 7 && dia >= 23) || (m == 8 && dia <= 22) {
        fmt.Println("leo");
    } else if (m == 8 && dia >= 23) || (m == 9 && dia <= 22) {
        fmt.Println("virgo");
    } else if (m == 9 && dia >= 23) || (m == 10 && dia <= 22) {
        fmt.Println("libra");
    } else if (m == 10 && dia >= 23) || (m == 11 && dia <= 21) {
        fmt.Println("esc||pio");
    } else if (m == 11 && dia >= 22) || (m == 12 && dia <= 21) {
        fmt.Println("sagitario");
    } else if (m == 12 && dia >= 22) || (m == 1 && dia <= 19) {
		fmt.Println("capric||nio");
	}
}