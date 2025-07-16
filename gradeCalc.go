// Task 1
package main
import(
	"fmt"
)
func gradeCalc(){
	var name string
	var num_grade int
	fmt.Print("Enter your name : ")
	fmt.Scanln(&name)
	fmt.Print("Number of subjects : ")
	fmt.Scanln(&num_grade)
	grades := make([]float64, num_grade)
	var ans float64
	for i := 0; i < num_grade; i++ {
		fmt.Printf("grade %v : ", i + 1)
		fmt.Scanln(&grades[i])
		ans += grades[i]
	}
	fmt.Printf("Average grade for %v is : %v", name, ans / float64(num_grade))
}

func main(){
	gradeCalc()
}
