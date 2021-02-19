package main
import (
    "fmt"
    "time"
    "strings"
    )
func calc(){
    var num1 float32
    var num2 float32 
    var wtdo string
    var mode string
    fmt.Println ("Pls Enter your first No.")
    fmt.Scanln (&num1)
    fmt.Println ("Pls Enter your second No.")
    fmt.Scanln (&num2)
    fmt.Println ("Now enter your mode of calculating","\n M for multiplication, D for division, S for subtraction and A for addition")
    fmt.Scanln (&mode)
    mode = strings.ToUpper(mode)
    if (mode == "M"){
        fmt.Println ("Processing...")
        time.Sleep(time.Second*1)
        fmt.Println ("Answer:",num1*num2,"\n Want a Do Over (Want to make an another calculation? {Yes/NO})")
        fmt.Scanln (&wtdo)
    } else if (mode == "D"){
            fmt.Println ("Processing...")
            time.Sleep(time.Second*1)
            fmt.Println ("Answer:",num1/num2,"\n Want a Do Over (Want to make an another calculation {Yes/NO})")
            fmt.Scanln (&wtdo)
        } else if (mode == "A"){
            fmt.Println ("Processing...")
            time.Sleep(time.Second*1)
            fmt.Println ("Answer:",num1+num2,"\n Want a Do Over (Want to make an another calculation {Yes/NO})")
            fmt.Scanln (&wtdo)
        } else if (mode == "S"){
            fmt.Println ("Processing...")
            time.Sleep(time.Second*1)
            fmt.Println ("Answer:",num1-num2,"\n Want a Do Over (Want to make an another calculation {Yes/NO})")
            fmt.Scanln (&wtdo)
        } else {
            fmt.Println ("Invalid Input ! You have Typed the wrong mode of Calculation.","\n Restarting....")
            time.Sleep(time.Second*2)
            calc()
            
        }
    wtdo = strings.ToUpper(wtdo)   
    if (wtdo == "YES"){
        fmt.Println ("Getting ready for an another calculation","\n Getting Ready......")
        time.Sleep(time.Second*1)
        calc()
    } else if (wtdo == "NO"){
      fmt.Println ("Okay. See you Next Time.", "\n Bye")
       } else {
           fmt.Println ("Huh ? I think you have given a invalid answer.","\n If you want to make an another calculation, Please restart the program.", "\n Bye")
       }
        
    }
 
func main() {
    fmt.Println ("Hi User! This is Caculator Go.","\n This is a calculator made with the Programing Language Go","\n Enjoy Calculating")
    calc()
} 

