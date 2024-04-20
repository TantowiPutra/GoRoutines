# GoRoutines
Repository Made to Learn GoRoutines for Concurrency &amp; Parallel Programming

What is Parallel Programming?
- Nowadays, single core processor is rarely used. Multicore processor is much more utilized nowadays to create parallel process in application
- Parallel Prorgamming is a concept of dividing a problem into smaller parts, and could be executed on the same time (parallel)
- Example: Opening several application at once (ms. office, code editor, browser, etc.)
- Real Life Example: Chefs are preparing food at a restaurant, where each of them prepares individual dishes.
                     At the bank, each teller serves their respective customers

Process Vs Thread
================================================================
Process:
- The execution of a program
- Process required more memory
- Process is isolated from another process
- Process requires more time to be stoped

Thread:
- Thread is a segment of process
- Threads require less memory
- Thread could be related to one another if it's in the same process
- Threads require less time to be stoped

Parallel vs Concurrency
================================================================
- The key difference is, paralel (execute several process at the same time), while concurrent (execute several process alternately)

CPU BOUND
================================================================
- Depends on CPU to run. For example, machine learning.

I/O Bound
================================================================
- Can implemenent parallel programming, but will benefit more using concurreny programming

What is Goroutine?
================================================================
- Goroutine is a lightweight thread managed by goroutine. Gouroutine has a size of 2kb, meanwhile thread has a size of 1mb or 1000kb
- Goroutine does not work parallel, instead it's happening concurrent

How Goroutine workss?
================================================================
- Goroutine is ran by Go scheduler, in which each thread is determined by GOMAXPROCS (most of the time equal to the amunt of CPU cores)
- It can be said that Goroutine is a replacement for thread, as it's run on thread, using Goroutines we don't have to manage thread manually as it's managed automatically by Go scheduler

Several Terminology In Go Scheduler
================================================================
- G: Goroutine
- M: Thread (Machine)
- P: Processor

How to create Goroutine?
================================================================
- It's quite simple to create Goroutine
- We just need to add keyword "go" before refering to a function which we're going to run in the goroutine

Channel
================================================================
- Channel provides a way for two or goroutines to communicate with one another
- Channel allows the communication between goroutines happens synchronoysly
- While using channel, goroutine will be blocked until there's a receiver for the data returned
- Channel allowes async await mechanism, just like in several other programming language
- In short, it works as a mediator

Channel Characteristics
================================================================
- Channel can only store one data, if another data is about to be added, the data must be taken from the channel before adding another data
- Channel must be closed or else it could lead to MEMORY LEAK
- To create map, we can use function map(), we need to define the data type of the channel

