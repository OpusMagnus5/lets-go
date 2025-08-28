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

/*
Używając kanałów jako parametrów funkcji, można określić, czy kanał jest przeznaczony tylko do wysyłania lub odbierania wartości. 
Ta specyfika zwiększa bezpieczeństwo typu programu.
*/
func ping(pings chan<- string, msg string) { // Funkcja ping akceptuje tylko kanał do wysyłania wartości
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // Funkcja pong akceptuje jeden kanał do odbioru i drugi dla wysyłania
    msg := <-pings
    pongs <- msg
}

	
func channelDirections() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}

/*
Timeouts
Implementacja timeoutów w Go jest łatwa i elegancka dzięki channels i select.
Zwróć uwagę, że kanał jest buforowany, więc wysyłanie w goroutine jest nieblokujące.
<-time.After oczekuje na wysłanie wartości po upływie limitu czasu wynoszącego 1s.
*/

func testTimeouts() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

/*
Iterate over 2 values in the queue channel.
Ten zakres iteruje po każdym elemencie odebranym z kolejki. Ponieważ zamknęliśmy kanał powyżej, iteracja kończy się po otrzymaniu 2 elementów.
Ten przykład pokazaje również, że możliwe jest zamknięcie niepustego kanału, ale pozostałe wartości nadal będą odbierane.
*/
func rangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

/*
Często chcemy wykonać kod Go w pewnym momencie w przyszłości lub wielokrotnie w pewnych odstępach czasu. 
Wbudowane w Go funkcje timera i tickera ułatwiają oba te zadania.
*/

func testTimers() {

	/*
	Timery reprezentują pojedyncze zdarzenie w przyszłości. Mówisz timerowi, 
	jak długo chcesz czekać, a on zapewnia kanał, który zostanie powiadomiony w tym czasie. Ten timer będzie czekał 2 sekundy.
	*/
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // <-timer1.C blokuje się na kanale C timera, dopóki nie wyśle wartości wskazującej, że timer został uruchomiony.
	fmt.Println("Timer 1 fired")

	/*
	Jeśli chciałeś tylko poczekać, mogłeś użyć time.Sleep. 
	Jednym z powodów, dla których timer może być przydatny, jest możliwość anulowania timera przed jego uruchomieniem. Oto przykład.
	*/
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

/*

Timery są przeznaczone do robienia czegoś raz w przyszłości - tickery są przeznaczone do robienia czegoś wielokrotnie 
w regularnych odstępach czasu. Oto przykład tickera, który tyka cyklicznie, dopóki go nie zatrzymamy.
*/
func testTickers() {
	/*
	Tickery używają mechanizmu podobnego do timerów: kanału, do którego wysyłane są wartości. 
	W tym przypadku użyjemy wbudowanej funkcji select na kanale, aby oczekiwać na wartości przychodzące co 500 ms.
	*/
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	/*
	Tickery mogą być zatrzymywane podobnie jak timery. 
	Gdy ticker zostanie zatrzymany, nie będzie już odbierać żadnych wartości na swoim kanale. Zatrzymamy nasz po 1600 ms.
	*/
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

/*
Workery będą odbierać pracę na kanale zadań i wysyłać odpowiednie wyniki na kanale wyników.
*/
func worker2(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func testWorkerPools() {
	// Aby korzystać z naszej puli pracowników, musimy wysyłać im zadania i zbierać ich wyniki. W tym celu tworzymy 2 kanały.
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Uruchomiamy 3 pracowników, początkowo zablokowanych, ponieważ nie ma jeszcze żadnych zadań.
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	// Wysyłamy 5 zadań, a następnie zamykamy ten kanał, aby wskazać, że to wszystkie zadania, które mamy.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Na koniec zbieramy wszystkie wyniki pracy. Zapewnia to również, że goroutines worker zostały zakończone. 
	// Alternatywnym sposobem oczekiwania na wiele goroutines jest użycie WaitGroup.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}