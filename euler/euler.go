// Package for Project Euler shortcuts and stuff

package euler

import "math"
// import "math/big"

func isPrimeNaive(n uint64) bool {
    // Returns whether n is a prime number or not, using naive brute-force.
    // Complexity: O(sqrt(n))

    var lim uint64 = uint64(math.Sqrt(float64(n)))
    var i uint64

    for i = 2; i <= lim; i++ {
        if n % i == 0 {
            return false
        }
    }

    return true
}

func GCD(a uint64, b uint64) uint64 {
    // Returns the greatest common divisors of a,b using the euclidean algorithm.
    // Complexity: O(log(n))

    r := b % a

    if r == 0 {
        return a
    } else {
        return GCD(r,a)
    }

}

func ExpNaive(a int, b int) int {
    // Returns a**b in the most naive way
    // Complexity: O(n)

    return ExpModNaive(a,b,0)
}

func ExpModNaive(a int, b int, m int) int {
    // Returns a**b mod m in the most naive way
    // Complexity: O(n)

    res := 1
    for i := 0; i < b; i++ {
        res *= a
        if m != 0 {
            res %= m
        }
    }

    return res
    
}

func Fibonacci(n int) uint64 {
    // Returns the nth fibonacci number using simple iterative method.
    // Complexity: O(n)

    return FibonacciMod(n, 0)

}

func FibonacciMod(n int, m uint64) uint64 {
    // Returns the nth fibonacci number module m using simple iterative method.
    // Complexity: O(n)

    var a uint64 = 1
    var b uint64 = 1
    var tmp uint64

    for i := 1; i < n; i++ {
        tmp = a

        if m == 0 {
            a = a + b
        } else {
            a = (a + b) % m
        }

        b = tmp
    }

    return a

}


func FibonacciFast(n int) uint64 {
    // Returns the nth fibonacci number using matrix exponantion by squaring algoithm.
    // Complexity: O(log(n))
    
    return FibonacciModFast(n, 0)

}

func FibonacciModFast(n int, m uint64) uint64 {
    // Returns the nth fibonacci number modulo m using matrix exponantion by squaring algoithm.
    // Complexity: O(log(n))

    // Initilizing values
    log2n := int(math.Log2(float64(n)))
    var x uint64 = 1
    var y uint64 = 1
    var z uint64 = 0
    var xres, yres, zres uint64
    if n % 2 == 0 {
        xres = 1
        yres = 0
        zres = 1
    } else {
        xres = 1
        yres = 1
        zres = 0
    }
    
    // Squaring loop
    for i := 1; i <= log2n; i++ {
        xnew := (x*x + y*y)
        ynew := (y*x + y*z)
        znew := (y*y + z*z)
        if m != 0 {
            xnew %= m
            ynew %= m
            znew %= m
        }
        x = xnew
        y = ynew
        z = znew

        if (n / ExpNaive(2,i)) % 2 == 1 {
            xresnew := (xres*x + yres*y)
            yresnew := (xres*y + yres*z)
            zresnew := (yres*y + zres*z)
            if m != 0 {
                xresnew %= m
                yresnew %= m
                zresnew %= m
            }
            xres = xresnew
            yres = yresnew
            zres = zresnew
        }
    }

    return xres

}
