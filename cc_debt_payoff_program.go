// my first go program

package main
import ("fmt")

func main() {
  var debt_amount, cc_interest_rate, spending_amount, payment_amount float32

  fmt.Print("Enter your amount of credit card debt: ")
  fmt.Scan(&debt_amount)
  
  fmt.Print("Enter your credit card interest rate as a decimal: ")
  fmt.Scan(&cc_interest_rate)
  
  fmt.Print("Enter your monthly spending amount: ")
  fmt.Scan(&spending_amount)
  
  fmt.Print("Enter your monthly payment amount: ")
  fmt.Scan(&payment_amount)
  
  var monthly_interest_rate float32 = cc_interest_rate / 12
  
  var pay_off_time float32 = ((debt_amount * monthly_interest_rate) + debt_amount) / (payment_amount - spending_amount)
  
  fmt.Println("You will pay off this card in ", pay_off_time, " months.")
  
}