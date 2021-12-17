# Concurrency in Go

`atomic` - indivisible/untinterruptible. One step, can't be interrupted by another thread.


Something may be atomic in one context, not in another (scope dependent - operation atomic in context of process might not be atomic in context of OS)

`critical section` - section of program that needs exclussive access to shared resources
(examples of critical sections: if, goroutine, print statement etc.)

Each critical section should be protected

`deadlock` - all concurrent processes are waiting on another. No recovery. 