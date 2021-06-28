package hw0

import (
    "math"
)

//factorial takes in an integer n and output the integre of n!
func factorial(n int) int{
  facto := 1
  if n <= 0{
    panic("error")
  }else{
    for i := 1; i <= n; i++{
      facto *= // // IDEA:
    }
  }
  return facto
}

//Combination
func combination(n int, k int) int{
  ans := 1
  if n==0 || k==0{
    return ans
  }

if k>= (n/2) {
   k = n - k
}

for i := 0; i < k; i++ {
  ans *= (n - i)
  ans /= i + 1

 }
 return ans
}

//permutation takes
func permutation(n int, k int) int{
  if n == 0 || k == 0{
  return 1
}

permu := n-k+1
for i:= n-k+2; i <= n; i++{
  permu *= i
  }
  return permu
}

//factorial array takes
func FactorialArray(n int) []int{
  array := make([]int, n+1)
  array[0] = 1
  for i := 1; i <= n; i++{
    array[i] = array[i-1] * i
  }
  return array
}

//FibonacciArray
func FibonacciArray(n int) []int {
  array := make([]int, n)
  array[0] = 1
  if n == 1 {
    return array
  }
  array[1] = 1
  for i := 2; i < n; i++ {
    array[i] = array[i-1] + array[i-2]
  }
  return array
}

//minarray
func MinArray(a []int) int {
  min := a[0]
  for i, e := range a {
    if i == 0 || e < min {
      min = e
    }
  }
  return min
}

//GCDArray
func GCDArray(a []int) int{
  gcd := 1
  min := MinArray(a)
  for d := 2; d <= min; d++{
    for i := range a{
      if (a[i]%d != 0){
        break
      }else if i==(len(a)-1){
        gcd = d
      }
    }
 }
  return gcd
}

//isperfect takes
func IsPerfect(n int) bool{
  if n  == 1{
    return false
  }
  sum := 1
  upper := n
  for d := 2; d <= int(math.Sqrt(float64(upper))); d++ {
if n%d == 0 {
  sum = sum + d + n/d
  }
  }
  if sum == n {
    return true
  } else {
    return false
  }
}

// next perfectnumber
//larger than n
func NextPerfectNumber(n int) int{
  p := n+1
  for p >= 0{
    if IsPerfect() != true{
      p++
    } else {
      return p
    }
    }
    return p
}

//nect twin primes
func NextTwinPrimes(n int) (int, int){
  if n <= 0{
    n = 1
  }
  if n%2 == 0{
    n+=1
  }else{
    n+=2
  }

 for n>0{
   if IsPrime(n) == true{
     if IsPrime(n+2) == true{
       return n, n+2
       } else {
         n += 2
      }
    }
  n += 2}
}
  return 0, 0
  }
//isprime
  func IsPrime(n int) bool{
    if n == 0 || n == 1 {
    return false}
    for i := 2; i <= int(math.Sqrt(float64(n))); i++{
      if n%i == 0 {
        return false
      }
    }
    return true
}
//sieveoferatosthenes takoes integer in and returns a slice of n+1 booleans porimearray
func sieveOfEratosthenes(n int) []bool {
  primeArray := make([]bool, n+1)
  for k := 2; k <= n; k++ {
    primeArray[k] = true
  }

// range over prime array, cross off
for p:= 2; float64(p)<= math.Sqrt(float64(n)); p++ {
  primeArray = crossOfMultiples(primeArray, p)
   }
}

  return primeArray
}




//INSTRUCTIONS
//Place the "hw0" folder in your "go/src" directory.
//Type your functions for HW0 in this file.
//To test your functions, navigate in the terminal ("cd") into go/src/hw0.
//Then run the command "go test". You will see that your code is run on a collection
//of test datasets.  If you pass all of the tests, you will see "PASS" in the last line.
//If for any problem you fail any test, you will see "FAIL".
//Feedback is printed based on how your code's responses differ from the correct answer.
//Use this feedback to debug your code. You can call "go test" as many times as you like,
//but remember to save your work if you edit this file!
//Happy coding :)
