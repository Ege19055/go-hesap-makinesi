package main

import(
  "fmt"
)

func main(){
    var calc bool = true
    var sayi1 int
    var sayi2 int
    var islem string

    for calc {
        fmt.Println("işlem seçiniz")
        fmt.Scan(&islem)
        switch(islem){
            case "+":
            fmt.Println("1.sayıyı giriniz");
            fmt.Scan(&sayi1)
            fmt.Println("2.sayıyı giriniz");
            fmt.Scan(&sayi2)
            fmt.Println(sayi1+sayi2)
            break; 
            case "-":
				fmt.Println("1.sayıyı giriniz");
				fmt.Scan(&sayi1)
				fmt.Println("2.sayıyı giriniz");
				fmt.Scan(&sayi2)
				fmt.Println(sayi1-sayi2)
            break; 
        }
    }
}
