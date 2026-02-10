# Overview

## Table of contents

- [What the Kernel Is](#what-the-kernel-is)
- [User Space vs Kernel Space](#user-space-vs-kernel-space)
- [Major Kernel Components (Conceptual)](#major-kernel-components-conceptual)
- [What cgroups Are](#what-cgroups-are)
- [Kernel Directories](#kernel-directories)
- [Key Takeaways](#key-takeaways)

---

## What the Kernel Is

The **kernel** is the privileged core of the operating system.

It:
- Runs in **kernel space** (full privileges)
- Manages hardware directly (CPU, memory, disks, network)
- Exposes controlled entry points to user programs via **system calls**

Everything else (shells, Go programs, Docker, Kubernetes) runs in **user space** and must ask the kernel to do real work.

---

## User Space vs Kernel Space

- **User space**
  - Normal programs
  - No direct hardware access
  - Crashes are isolated

- **Kernel space**
  - Scheduler, memory manager, drivers, networking
  - Enforces isolation and security
  - Bugs can crash the entire system

Boundary crossing happens through **syscalls** (e.g. `fork`, `execve`, `read`, `write`, `kill`).

---


## Major Kernel Components (Conceptual)

These are *conceptual groupings*, not strict directories:

1. **Process management** – creation, scheduling, signals, exit
2. **Memory management** – virtual memory, paging, OOM
3. **Filesystems (VFS)** – files, mounts, permissions
4. **Device drivers** – disks, NICs, GPUs, input devices
5. **Networking stack** – TCP/IP, routing, sockets
6. **IPC** – pipes, shared memory, signals
7. **Security & isolation** – namespaces, capabilities, LSMs

---


## What cgroups Are

A **cgroup (control group)** is a **kernel mechanism to apply resource rules to a group of processes**.

cgroups:
- Limit and account **CPU, memory, IO, PIDs**
- Do **not** create isolation by themselves
- Are enforced entirely by the kernel

> Containers are **processes + cgroups + namespaces**.

cgroups are **not a separate kernel component**.

They are a **cross-cutting control framework** that plugs into existing subsystems:

- CPU cgroups → scheduler
- Memory cgroups → memory manager
- IO cgroups → block IO layer
- PIDs cgroups → process creation path

**Important principle**:

> cgroups do not manage resources — they tell resource managers how to behave.

cgroups v1 vs v2:
- **v1**: multiple hierarchies, complex, legacy
- **v2**: single unified hierarchy, consistent semantics

Modern systems (and Kubernetes) use **cgroup v2**.

Live kernel interface:
```
/sys/fs/cgroup/
```

---


## Kernel Implementation Reality

- Kernel is written mostly in **C**, with **assembly** for low-level paths
- Compiled into a single kernel image (`vmlinuz`)
- Runs without libc or standard runtime

Kernel-exposed interfaces as directories in the OS (most important)

- **`/proc`** — process and kernel state (virtual filesystem)
  - Examples: `/proc/<pid>/status`, `/proc/meminfo`, `/proc/schedstat`
  - Primary place to observe **process management**.

- **`/sys`** — kernel object model (virtual filesystem)
  - Examples: `/sys/devices`, `/sys/class`, `/sys/fs/cgroup`
  - Primary place to observe **devices** and **cgroups**.

- **`/dev`** — device nodes
  - Examples: `/dev/sda`, `/dev/null`, `/dev/tty`, `/dev/net/tun`
  - How user programs interact with **device drivers**.

- **`/boot`** — boot artifacts
  - Examples: `/boot/vmlinuz-*`, `/boot/initrd.img-*`
  - The **kernel image on disk** (the kernel runs from memory after boot).

- **`/run`** — runtime state (usually tmpfs)
  - Used for PID files, sockets, runtime coordination (e.g. systemd, containerd).

Kernel source code directories:
```
kernel/   # scheduler, fork, exit, signals
mm/       # memory management
fs/       # filesystems
net/      # networking
drivers/  # hardware drivers
arch/     # architecture-specific code
```

Mostly user-space directories (not kernel internals)

- `/(bin|sbin)` essential tools
- `/usr` most installed software
- `/etc` configuration
- `/var` logs and variable data
- `/tmp` temporary files
- `/home`, `/root` user homes

Bridge idea:

> **Kernel source directories explain _why_ things work; runtime directories show _what is happening now_.**

---
## Key Takeaways

- The kernel is the enforcer — orchestration is policy
- Processes are the fundamental unit
- cgroups constrain, schedulers decide
- Containers are *not* special, they are processes
