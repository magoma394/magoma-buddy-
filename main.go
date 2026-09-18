package main
import(
	"time"
	"fmt"
)
type Reminder struct {
	Name string
	Question string
}

func (r Reminder) Execute(){
	fmt.Printf("Magoma Buddy is running for [%s]...\n", r.Name)
	time.Sleep(2*time.Second)

	fmt.Printf("\n💧 %s check!\n", r.Name)
	fmt.Println(r.Question)

	var ans string
	fmt.Scanln(&ans)

	if ans == "y" || ans == "yes" {
		fmt.Println("Nice, stay hydrated!")
	} else {
		fmt.Println("Kidding!! Go get some hydration right now!")
	}
}

func main() {
	waterReminder := Reminder{
		Name: "Water",
		Question: "Did you drink water? [y,n]",
	}
	waterReminder.Execute()
}
