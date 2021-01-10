# concurrent vs parallel
* 2 w tym samym czasie na roznym hardwarze (cpu) - parallel
* concurrent "iluzja paralelizmu", jest switching i overlaping. Moze byc na 2 maszynach lub na jednym

# basics 

## proces
`proces` - instancja programu, unikalna pamiec, stack, heap, kod, rejestry, liby

OS switchuje procesy co np. 20ms. Czesto jest concurrency, a nie parallel. Zalezy czy mamy hardware

`scheduling` - OS zarzadza procesami i watkami tak, aby byla iluzja rownoleglosci. Przydziela zasoby i sprawiedliwie odpala rozne procesy
* round robin - kazdy po kolei
* priorytetowe - jakies UI jest wazniejsze niz cos w backgroundzie

`context switch` - gdy OS zmienia proces/watek - trzeba zapisac stan (context - unikalne rzeczy dla procesu np. pamiec, rejestry, kod) i przelaczyc sie na druga jednostke wykonawcza

## thread
proces sklada sie z wielu `watkow`.

Tez jest scheduling, context switch, ale watki maja mniej stanu (kontekstu), latwiejszy context switch.

Jest stack, kod i rejestry. Pamiec jest wspolna dla wszystkich watkow w procesie

## goroutine
typowe dla GO. To jest "watek", ale w GO. Jest lzejsze niz thread. Gorutyny moga byc w jednym watku w OSie. Switching jest w `Go runtime scheduluer`. Mamy tutaj concurrency

`logical processor` zmapowany do OSowego watku.

## interleaving
watki sie przeplataja, wiec w przypadku crashu ciezej dojsc do przyczyny crashu - reszta watkow ma inny stan. Ciezej deterministycznie przesledzic przebieg programu. Moze przeplecione byc na poziomie kodu albo instrukcji maszynowych

* order of execution within tasks is known
* order of execution between concurrent tasks is unknown
* interleaving of instructions between tasks is unknown. Nie wiemy:
    * moze byc instrukcja z tasku 1 przepleciona z taskiem 2
    * moze byc najpierw 3 instrukcje z task 1, potem 3 z task 2 itd.


## Race condition
problem zalezacy od interleavingu. Inny przebieg programu - niedeterministyczny. Zdazaja sie przez komunikacje (wspoldzielenie zasobow) pomiedzy taskami 

|task1    |task2    |
|:-:      |:-:      |
|  x = 1  |         |
|         | print x |
|  x++    |         |

vs

|task1    |task2    |
|:-:      |:-:      |
|  x = 1  |         |
|  x++    |         |
|         | print x |

# Goroutines
* gorutyna konczy sie gdy skonczy sie kod jej funkcji
* gdy `main` skonczy - wszystkie inne tez koncza, nawet w srodku przetwarzania

## Synchronisation
**WE CAN'T RELY ON TIMING**

* enforcing order of execution with synchronisation methods. Not allowing some interleavings. 
* reduce some concurrency and performance!

introduce a global event that is viewed by all tasks at the same time and run specific actions only when global event occured

|task1                |task2            |
|:-:                  |:-:              |
|  x = 1              |                 |
|  x++                |                 |
| global event        | if global event |
|                     |     print x     |

`synch package` - functions to sync between go routines
`Sync waitGroup` - calling goroutine wait for others to complete.
equivalent of java thread.join()

contains internal counter until all goroutines completed
* Add() - increment counter - number of threads to wait
* Done() - decrement counter
* Wait() - wait until counter == 0

### komunikacja miedzy watkami w GO - channel
bez channels trzeba uzyc pol w obiektach i wyciagac

```
c := make(chan int)
// send to channel in goroutine:
c <- 3
// receive from the channel:
x := <- c
```
* order not quaranteed
* thread will block until we take something from channel `<- c`, so we can use that to wait without group
* allows sending/receiving data and syncs threads


### blocking on channels
* sending blocks data intil it's received by consumer thread
* receiving thread will block until the data is available
* channel communication is synchronous it's like a wg.Wait()

|task1                                      |task2     |
|:-:                                        |:-:       |
|  c <- 3                                   |          | 
|  task1 will wait till task 2 receives     |          |
|                                           | x := <- 3|
|  task1 resumes                            |          |

**by default** - channel is unbuffered - it can't hold data in transit. It's used only to receive data

**buffered** - adding channel capacity with 
```
c := make(chan int, 2) - 2 ints to input
```
now:
* sending only blocks if the buffer is full
* receiving only blocks if the buffer is empty

when to use?
* sender and receiver needs to be in locked state
* producer is faster than consumer
* consuming too fast

### sharing variables - MUTEX
concurrency is at the machine code level - i++ contains 3 operations. It's not atomic, so we need to sync that - make it serializable.
**Shared variables can't be accessed at the same time**

`mutex` - mutual exclusion - uses a binary semaphore (flag is used or not)