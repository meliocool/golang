# Concurrency vs Parallel Programming

## Multi-Core Processors
- Modern processors, even in smartphones, are multi-core.

## Parallel Programming
- Divides a problem into smaller chunks and executes them **at the same time**.
- Example 1: 10 exam questions, solved by 10 people simultaneously → parallel execution.
- Example 2: Running multiple background applications at once (Spotify, Office, Browser, GoLand, etc.).

## Processes vs Threads
- **Process**
    - Represents a program execution.
    - Consumes large memory.
    - Isolated from other processes.
    - Slower to start/stop.

- **Thread**
    - A smaller unit within a process.
    - Consumes minimal memory.
    - Threads in the same process can communicate easily.
    - Faster to start/stop.
    - Example: Chrome browser is a process, and each tab runs on its own thread.

## Concurrency vs Parallelism
- **Concurrency**
    - Executes tasks by **taking turns** (time-slicing).
    - Requires fewer threads.
    - Example: In a café, you eat, talk, and drink in turns (but not all at once).

- **Parallelism**
    - Executes tasks **simultaneously**.
    - Requires many threads.

## CPU-Bound vs I/O-Bound
- **CPU-Bound**
    - Algorithms depend heavily on CPU performance.
    - Do not benefit much from concurrency.
    - Benefit more from parallel programming (e.g., Java).

- **I/O-Bound**
    - Dependent on input/output device speed (e.g., file read/write, database calls, API requests).
    - Benefit more from concurrency.
    - Example: If one thread waits 1 second for a database response, concurrency allows other tasks to use that waiting time effectively.
