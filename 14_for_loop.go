package main

import "fmt"

func for_loop() {
// counter := 0
// for counter <=10 {
// 	fmt.Println("Perulangan ke",counter)
// 	counter++
// }

for counter :=1 ; counter <=10 ; counter++ {
	fmt.Println("Perulangan ke",counter)

}

names := []string{"Paijo","Siti","Devi"}
for i := 0 ; i < len(names) ; i++ {
	fmt.Println(names[i])
}

for index,name := range names {
	fmt.Println("data ke-",index,"=",name)

}

//gunakan _ untuk memberitahu variable tidak di butuhkan
for _,name := range names {
	fmt.Println("nama =",name)

}

person:=make(map[string]string)
person["name"]="Babeng"
person["age"]="22"
person["address"]="Bengkulu"
for key,value:= range person{
	fmt.Println(key,":",value)

}
}
