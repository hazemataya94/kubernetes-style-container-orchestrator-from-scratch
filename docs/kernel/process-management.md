## Process Management — What the Kernel Actually Does

Process management covers:

- Creating processes
- Scheduling CPU time
- Tracking process states
- Delivering signals
- Cleaning up after exit

Kernel tracks each process using an internal structure (`task_struct`).

---


## What a Process Is (Kernel View)

A process is **not just a program file**.

It is:
- CPU execution context
- Memory mappings
- Open file descriptors
- Credentials
- Signal handlers
- Scheduling metadata

The kernel schedules **threads**, not applications.

---


## Core Process Management Topics to Understand

### Process identity
- PID, PPID
- Process tree
- PID 1 (init)

### Process creation
- `fork()`
- `exec()`
- `clone()`
- Copy-on-write (COW)

### Scheduling
- Run queues
- Time slices
- Priorities
- Fairness
- Preemption

### Process states
- Running
- Sleeping (interruptible / uninterruptible)
- Stopped
- Zombie

### Signals
- `SIGTERM`, `SIGKILL`, `SIGINT`
- Default vs handled
- Graceful vs forceful shutdown

### Exit & cleanup
- `exit()`
- `wait()` / `waitpid()`
- Zombie processes
- Orphans