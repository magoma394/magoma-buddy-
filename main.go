package main
import(
	"time"
	"fmt"
)
func main() {
	fmt.Println("Magoma Buddy is running...")
	time.Sleep(10 * time.Second)
	fmt.Println("💧 Water check!")
	fmt.Println("Did you drink water? [y/n]")
	var ans string
	fmt.Scanln(&ans)
	if ans == "y" || ans == "yes"{
		fmt.Println("Nice, Stay Hydrated!")
	} else {
		fmt.Println("Kidding!! Go get some hydration right now!")
	}
}
