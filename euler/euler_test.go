package euler

import "testing"

func TestisPrimeNaive(t *testing.T) {
    primes := [5]uint64{2, 3, 5, 7, 11}
    notprimes := [5]uint64{4, 6, 8, 9 ,10}

    for _,n := range primes {
        if isPrimeNaive(n) != true {
            t.Errorf("%d is a prime.", n)
        }
    }

    for _,n := range notprimes {
        if isPrimeNaive(n) != false {
            t.Errorf("%d is not a prime.", n)
        }
    }
}

func TestGCD(t *testing.T) {
    tuples := [6][3]uint64{
        {3,5,1}, {5,3,1},
        {10,5,5}, {5,10,5},
        {14,21,7}, {21,14,7}}

    for _,tuple := range tuples {
        res := GCD(tuple[0],tuple[1])
        if res != tuple[2] {
            t.Errorf("gcd of %d,%d should be %d, not %d",
                tuple[0], tuple[1], tuple[2], res)
        }
    }

}

func TestFibonacci(t *testing.T) {

    results := [10]uint64 { 1, 1, 2, 3, 5, 8, 13, 21, 34, 55 }
    
    for n,val := range results {
        fib := Fibonacci(n)
        if fib != val {
            t.Errorf("the %dth fibonacci number is %d, not %d",
                n, val, fib)
        }
    }

}

func TestFibonacciFast(t *testing.T) {

    results := [10]uint64 { 1, 1, 2, 3, 5, 8, 13, 21, 34, 55 }
    
    for n,val := range results {
        fib := FibonacciFast(n)
        if fib != val {
            t.Errorf("the %dth fibonacci number is %d, not %d",
                n, val, fib)
        }
    }

}

func TestFibonacciMod(t *testing.T) {

    results := [10]uint64 { 1, 1, 2, 3, 5, 8, 13, 21, 34, 55 }
    modules := [3]uint64 { 2, 5, 10 }
    
    for n,val := range results {
        for _,m := range modules {
            fib := FibonacciMod(n, m)
            if fib != val % m {
                t.Errorf("the %dth fibonacci number modulo %d is %d, not %d",
                    n, m, val % m , fib)
            }
        }
    }

}

func TestFibonacciModFast(t *testing.T) {

    results := [10]uint64 { 1, 1, 2, 3, 5, 8, 13, 21, 34, 55 }
    modules := [3]uint64 { 2, 5, 10 }
    
    for n,val := range results {
        for _,m := range modules {
            fib := FibonacciModFast(n, m)
            if fib != val % m {
                t.Errorf("the %dth fibonacci number modulo %d is %d, not %d",
                    n, m, val % m , fib)
            }
        }
    }

}

func TestIntegerSqrt(t *testing.T) {

    tuples := [5][2]uint64 {
        {2,1}, {10, 3}, {100, 10}, {9, 3}, {1,1} }

    for _,tuple := range tuples {
        sqrt := IntegerSqrt(tuple[0])
        if sqrt != tuple[1] {
            t.Errorf("the square root of %d is %d, not %d",
                tuple[0], tuple[1], sqrt)
        }
    }
}
