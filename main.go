package main

import (
	"fmt" 
 "kothi/Mypack"   
)

func main() {
	var var1 int = 16;
	var var2 int = 5;
	var var3 float64 = 5.765;
    Println("Hi!!, Hope you are doing good")
    fmt.Println("Add = ",caluclator.Add(var1,var2))
	fmt.Println("Sub = ",caluclator.Sub(var1,var2))
	fmt.Println("Div = ",caluclator.Div(var1,var2))
	fmt.Println("Mul = ",caluclator.Mul(var1,var2))
	fmt.Println("Rem = ",caluclator.Rem(var1,var2))
	fmt.Printf("Sqrt = %.3f \n",caluclator.Sqrt(var3))
	fmt.Printf("Sin = %.3f \n",caluclator.Sine(var3))
	fmt.Printf("cos = %.3f",caluclator.Cos(var3))
}