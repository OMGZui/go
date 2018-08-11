package main

import "fmt"

var arr [10]int
var ar = [10]int{2, 4, 1}
var rating = map[string]int{"C": 5, "Go": 4, "Python": 6, "C++": 2}

type testInt func(int) bool // 声明了一个函数类型

type person struct {
	name string
	age  int
}

/* array、slice、map */
func ShowArr() {
	arr[0] = 10
	arr[1] = 40
	FPrint("全局数组", arr[0])

	a := [...]int{2, 4, 1}
	FPrint("简短声明", a[2])

	double := [2][2]int{{1, 2}, {2, 1}}
	FPrint("二维数组", double[1][0])

	sliceA := ar[2:5:8]
	FPrint("slice", sliceA[0])

	cRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", cRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	FPrint("map", rating["Go"])
}

/* 流程控制 */
func Expression() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	for k, v := range rating {
		fmt.Printf("%s=>%d\n", k, v)
	}

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	x := 3
	y := 4
	xPLUSy, xTIMESy := SumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)

	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}

var P person

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human  // 匿名字段，struct
	Skills // 匿名字段，自定义的类型string slice
	int    // 内置类型作为匿名字段
	speciality string
}

func ShowStr() {
	P.name = "o"                                    // 赋值"o"给P的name属性.
	P.age = 25                                      // 赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s\n", P.name) // 访问P的name属性.

	p := new(person)
	p.name = "m"
	p.age = 23
	fmt.Printf("name=>%s, age=>%d\n", p.name, p.age) // 访问P的name属性.

	// 初始化学生Jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

func main() {
	ShowArr()
	Expression()
	ShowStr()
}

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func FPrint(name string, value int) {
	fmt.Printf("%s=>%d\n", name, value)
}
