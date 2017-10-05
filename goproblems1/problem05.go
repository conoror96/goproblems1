// converting python games into golang
package main
import(
    "fmt"
    "math/rand"
    "time"
)

//this generates random number between given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := xrand(1, 25)
    guessTaken := 0
    var guess int

    fmt.Println("Guessing game between 1 and 25")
    fmt.Println("Take a guess...")
    //this is the while loop
    for guess != myrand {
        fmt.Scanf("%d", &guess)
        guessTaken++
        if guess < myrand {
            fmt.Println("Your guess is too low. Try again..")
        }
        if guess > myrand {
            fmt.Println("Your guess is too high. Try again..")
        }
        if guess == myrand {
            break
        }
    }
    if guess == myrand {
        fmt.Printf("You guessed the number %d in %d tries\n",myrand, guessTaken)
    }
}
