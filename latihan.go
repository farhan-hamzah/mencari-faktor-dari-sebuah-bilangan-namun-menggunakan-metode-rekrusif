package main
import "fmt"
func faktor(i int, c int){
    if c <= i {
        if i > 0{
            if i%c == 0{
                fmt.Print(c, " ")
            }
            faktor(i, c+1)
        }
    }    
}

func main(){
    var masukan int
    fmt.Scan(&masukan)
    faktor(masukan, 1)
}