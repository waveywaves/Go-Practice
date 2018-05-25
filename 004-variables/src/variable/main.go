package main;

import "fmt";
//import "strings";
import "math/big";
import "math";
import "time";

func main() {

	// Implicit and explicit declaration for variables 

	str1 := "An implicitly typed string";
	var str2 string = "An explicitly typed string";

	strlen, err := fmt.Println(str1, str2);
	if err == nil {
		fmt.Println(strlen);
	} else {
		fmt.Println(err);
	}

	//Math Operators

	var num1 float64 = 100;
	var num2 float64 = 200;
	fmt.Println("100 + 200 ",num1+num2);
	fmt.Println("100 - 200 ",num1-num2);
	fmt.Println("100 * 200 ",num1*num2);
	fmt.Println("100 / 200 ",float64(num1/num2));
	fmt.Println("100 % 200 ",(int(num1)%int(num2)));

	f1,f2,f3 := 23.5, 45.7, 12.9;
	floatSum := f1 + f2 + f3;

	fmt.Println("Float Sum : ", floatSum);

	var b1,b2,b3,bigSum big.Float;
	b1.SetFloat64(23.5);
	b2.SetFloat64(45.7);
	b3.SetFloat64(12.9);

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3);
	fmt.Printf("bigSum : %.10g \n",&bigSum);

	circleRadius := 15.5;
	circumference := circleRadius * math.Pi;
	fmt.Printf("circumference %.2f \n",circumference)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC);
	fmt.Println("Go launched at %s \n",t);

	fmt.Println(t.Month());
	fmt.Println(t.Day());

}