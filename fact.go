package fact

import (
	// "flag"
	// "fmt"
	"math/big"
	// "time"
)

func FactNative(n int64) (ret *big.Int) {
	var N = big.NewInt(n)
	ret = big.NewInt(1)
	var one = big.NewInt(1)
	for i := big.NewInt(2); N.Cmp(i) == 1 || N.Cmp(i) == 0; i.Add(i, one) {
		ret.Mul(ret, i)
	}

	return
}

func prodTree(l, r int64) (ret *big.Int) {
	if l > r {
		return big.NewInt(1)
	}
	if l == r {
		return big.NewInt(l)
	}
	if r-l == 1 {
		return big.NewInt(l * r)
	}

	m := (l + r) / 2

	ret = big.NewInt(1)
	return ret.Mul(prodTree(l, m), prodTree(m+1, r))
}

func FactTree(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	if n == 1 || n == 2 {
		return big.NewInt(n)
	}
	return prodTree(2, n)
}

// func main() {

// 	numbPtr := flag.Int("n", 42, "an int")
// 	flag.Parse()

// 	var N = int64(*numbPtr)

// 	t := time.Now()

// 	mybig1 := FactTree(N)
// 	fmt.Println("Время поиска: ", time.Since(t))
// 	fmt.Println(mybig1.String())
// }
