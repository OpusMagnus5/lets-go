package basics

import (
	"fmt"
	"time"
	"sync"
)

func TestConcurrency() {
	fmt.Println("--concurrency----------------------------------------------------------------------------------------")
	concurrency()
}

func concurrency() {
	/*
	W Go gorutyna jest wydajnym wątkiem którym zarządza biblioteka runtime.
	Ewaluacja parametrów odbywa się w obecnej gorutynie, a wykonanie funkcji już nowej.
	Gorutyny działają w tej samej przestrzeni adresowów, tak więc dostęp do współdzielonej pamięci musi być zsynchronizowany.
	*/
	go say("world")
	say("hello")

	/*
	Kanały są posiadającymi typ „kablami” przez które możesz wysyłać i odbierać wartości używając operatora <-.
	ch <- v    // Wysyła v do kanału ch.
	v := <-ch  // Odbiera wartość z kanału ch i przypisuje ją do v.

	Domyślnie, wysyłanie oraz odbieranie są zablokowane dopóki druga strona nie jest gotowa. 
	To pozwala gorutynom na synchronizację bez jawnej blokady lub zmiennych warunkowych.
	*/
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // Tworzy nowy kanał typu int
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // odbiera z c

	fmt.Println(x, y, x+y)

	/*
	Kanały buforowane (ang. buffered channels)
	Kanały mogą być buforowane. Aby zaincjalizować kanał buforowany podaj wielkość bufora jako drugi argument funkcji make
	Wysyłanie do kanału buforowanego wprowadza blokadę tylko gdy bufor jest pełny. Odbieranie wprowadza blokadę tylko gdy bufor jest pusty.
	 */
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	/*
	Zakres i zamykanie kanału (ang. range, closing of channel)
	Wysyłający może zamknąć kanał by zasygnalizować, iż do danego kanału nie zostanie już przesłana żadna wartość. 
	Odbierając może sprawdzić czy kanał został zamknięty poprzez wprowadzenie drugiego parametru do wyrażenia odbierającego wartości z kanału

	v, ok := <-ch
	Jeśli ok przyjmie wartość boolowską false to znaczy, że nie ma więcej wartości które można odebrać i kanał jest zamknięty.
	Uwaga: Tylko wysyłający powinien zamykać kanał, nigdy odbierający. Próba wysłania czegoś przez zamknięty kanał wywoła panikę.
	Kanały nie są jak pliki; zwykle nie musisz ich zamykać. Zamykanie kanałów jest niezbędne tylko wtedy, 
	gdy odbiorca musi wiedzieć, że więcej wartość nie zostanie już nim przesłanych, na przykład po to by mógł zakończyć działanie wyrażenia range w pętli.
	*/
	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}

	/*
	Instrukcja select pozwala gorutynom oczekiwać w tym samym czasie na sygnał z kilku różnych źródeł.
	Instrukcja select blokuje działanie gorutyny dopóki jedna z jej warunków nie uruchomi się, wówczas wykonuje odpowiadający 
	temu warunkowi przypadek (ang. case). Jeśli kilka warunków zostaje uruchomionych jednocześnie, select wybiera losowo 
	jeden dostępnych przypadków i wykonuje go.
	*/
	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		quit <- 0
	}()
	fibonacci2(c3, quit)


	/*
	Przypadek default w select jest uruchamiany w momencie gdy żaden z warunków dostępnych przypadków nie został uruchomiony w danej chwili.
	Użyj przypadku default by wysłać lub otrzymać wartość bez blokowania
	*/
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // wysyła sum do c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("Wyjście")
				return
		}
	}
}

	/*
	Co jeśli nie potrzebujemy komunikacji? Co zrobić w sytuacji gdy chcemy by tylko jedna gorutyna mogła mieć w danym momencie 
	dostęp do danej zmiennej, by unikać konfliktów między gorutynami?
	Ta koncepcja nosi nazwę _wzajemnego wykluczania_ (ang. mutual exclusion), a strukturę danych która pozwala 
	nam z niej korzystać zwyczajowo nazywamy mutex.
	Biblioteka standardowa Go implementuje mechanizm wzajemnego wykluczania za pomocą typu sync.Mutex oraz jego dwóch metod:

	Lock
	Unlock
	Możemy napisać blok kodu który będzie wykonywany w trybie wzajemnego wykluczania poprzez otoczenie go wywołaniami metod Lock i Unlock
	Możemy też użyć instrukcji defer by upewnić się, że mutex zostanie otwarty (ang. unlocked)
	*/
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock użyty by tylko jedna gorutyna mogła mieć dostęp do mapy c.v. w danym czasie
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock użyty by tylko jedna gorutyna mogła mieć dostęp do mapy c.v. w danym czasie
	defer c.mu.Unlock()
	return c.v[key]
}

/*
Możemy użyć kanałów do synchronizacji wykonywania między goroutines.
W przypadku oczekiwania na zakończenie wielu goroutines, lepiej jest użyć WaitGroup.
*/
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
		
    done <- true
}

func worketTest() {

	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
}