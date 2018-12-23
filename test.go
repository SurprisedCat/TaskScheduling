/*
package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Initialize two matrices, a and b.
	a := mat.NewDense(2, 2, []float64{
		4, 0,
		0, 4,
	})
	b := mat.NewDense(2, 3, []float64{
		4, 0, 0,
		0, 0, 4,
	})

	// Take the matrix product of a and b and place the result in c.
	var c mat.Dense
	c.Mul(a, b)

	// Print the result using the formatter.
	fc := mat.Formatted(&c, mat.Prefix("    "), mat.Squeeze())
	fmt.Printf("c = %v\n", fc)
}
*/

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	a1, a2 := 1, 2

// 	rs := func(a int, b int) int { //（）括号中是接受的参数  a1,a2
// 		return a + b
// 	}(a1, a2) //() 括号中的是传入的参数，并调用匿名函数

// 	fmt.Println(rs) // 结果输出3

// 	//下面匿名函数不穿入参数，结果输出 rs2
// 	func() {
// 		fmt.Println("rs2")
// 	}() //直接执行

// 	//举例匿名函数自动执行   结果输出 3
// 	func(a int, b int) {
// 		fmt.Println(a + b)
// 	}(a1, a2)

// 	e := getFun(a1, a2) //传入参数并输出结果  6
// 	e()

// 	fn := func() {
// 		fmt.Println("hello")
// 	}
// 	fmt.Println("///******////")
// 	fmt.Printf("%T\n", fn) //打印func()
// }

// func getFun(a int, b int) func() {
// 	r1 := a
// 	r2 := b
// 	r3 := 3
// 	return func() {
// 		fmt.Println(r1 + r2 + r3)
// 	}
// }
// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"log"
// 	"strings"
// )

// func main() {
// 	in := `first_name,last_name,username
// "Rob","Pike",rob
// Ken,Thompson,ken
// "Robert","Griesemer","gri"
// `
// 	r := csv.NewReader(strings.NewReader(in))

// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(record)
// 	}

// 	in = `first_name;last_name;username
// "Rob";"Pike";rob
// # lines beginning with a # character are ignored
// Ken;Thompson;ken
// "Robert";"Griesemer";"gri"
// `
// 	r = csv.NewReader(strings.NewReader(in))
// 	r.Comma = ';'
// 	r.Comment = '#'

// 	records, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(records)

// 	records = [][]string{
// 		{"first_name", "last_name", "username"},
// 		{"Rob", "Pike", "rob"},
// 		{"Ken", "Thompson", "ken"},
// 		{"Robert", "Griesemer", "gri"},
// 	}
// 	var out strings.Builder
// 	w := csv.NewWriter(&out)

// 	for _, record := range records {
// 		fmt.Println(record)
// 		if err := w.Write(record); err != nil {
// 			log.Fatalln("error writing record to csv:", err)
// 		}
// 		w.Flush()
// 		fmt.Println(out.String())
// 		out.Reset()
// 	}

// 	// Write any buffered data to the underlying writer (standard output).
// 	w.Flush()

// 	if err := w.Error(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(out.String())
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 60; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Printf("%d\n", rand.Intn(10000))

	}
}
