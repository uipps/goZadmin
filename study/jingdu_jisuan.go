// 精度计算问题，在PHP中有bcmath，bcadd等等
//

package main

import (
    "fmt"
    "github.com/shopspring/decimal"
)

func main() {

    m1 := 8.2
    m2 := 3.8
    fmt.Println(m1 - m2) // 期望是4.4，结果打印出了4.3999999999999995

    fmt.Println("")
    m1D := decimal.NewFromFloat(m1)
    m2D := decimal.NewFromFloat(m2)
    fmt.Println(m1D.Sub(m2D))        // 4.4
    // fmt.Println(m1D - m2D)    // 报错： invalid operation: m1D - m2D (operator - not defined on struct)

    price, err := decimal.NewFromString("136.02")
    if err != nil {
        panic(err)
    }

    quantity := decimal.NewFromInt(3)

    fee, _ := decimal.NewFromString(".035")
    taxRate, _ := decimal.NewFromString(".08875")

    subtotal := price.Mul(quantity)

    preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

    total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

    fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
    fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
    fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
    fmt.Println("Total:", total)                            // Total: 459.824961375
    fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875

    //os.Exit()
}
