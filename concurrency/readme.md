# Concurrency in Go

`atomic` - indivisible/untinterruptible. One step, can't be interrupted by another thread.


Something may be atomic in one context, not in another (scope dependent - operation atomic in context of process might not be atomic in context of OS)

`critical section` - section of program that needs exclussive access to shared resources
(examples of critical sections: if, goroutine, print statement etc.)

Each critical section should be protected

# deadlock
`deadlock` - all concurrent processes are waiting on another. No recovery. 

## Coffman Conditions - for deadlock to occur
* mutual exclussion - concurrent process holds exclusive right to a resource at any one time
* wait for condition - process holds resource and wait for additional resource
* no preemption - resource held by a process can be released only by that process
* circular wait - P1 waits on P2 and P2 waits for P1

**if at least one law is not true - NO DEADLOCK!**

# livelock
programs that actively perform concurrent operations, but those operations do nothing to move the state of the program forward

13
