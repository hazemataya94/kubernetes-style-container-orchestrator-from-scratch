# Kubernetes-Style Container Orchestrator From Scratch

> From Linux primitives to Kubernetes-style control planes

This is a **learning-by-building** project to understand how Kubernetes works internally
by implementing a minimal, Kubernetes-style container orchestrator in Go.

The goal is **not** to re-implement Kubernetes, but to learn the core ideas behind it, including:
- declarative APIs
- desired vs actual state
- reconciliation loops
- watches
- control planes
- container runtimes

---

## Scope & non-goals

### In scope
- Go for systems programming
- Linux kernel primitives (processes, namespaces, cgroups)
- etcd and watch-based state propagation
- containerd and container lifecycle management
- Kubernetes-style APIs and control loops

### Explicit non-goals
- Production readiness
- Feature parity with Kubernetes
- Performance optimization
- Full networking or CNI implementation (at least initially)

This is a **learning project**, not a product.

---

## Project roadmap

The project is structured as a series of learning phases. Each phase builds on the previous one.

## Phase 1 — Golang fundamentals for systems (mostly macOS)

### Goal
be comfortable writing Go programs.

### Topics
- concurrency (goroutines, channels)
- networking (HTTP, gRPC)
- APIs, CLIs, testing
- controller-style loops

### Mini-project ideas
- Build a small REST API server with CRUD + watch-like long polling 
- Build a gRPC server that exposes CRUD for resources (Apply/Get/List/Delete) 
- Build a controller loop: “desired state” file → reconcile into “actual state” folder 
- Build a tiny kubectl-like CLI that talks to your API server 
- Controller + Work Queue: gRPC server stores desired state, controller subscribes via stream, reconciles, and reports status back via gRPC 

---
## Phase 2 — Linux kernel interfaces

### Goal
understanding kernel, how to interact with it, and what containers actually rely on.

### Topics
- processes, signals, `/proc`
- namespaces and cgroups
- filesystem mounts
- observability tools (`strace`, `lsof`, `ip`, `ss`)
- Go programs that interact with kernel concepts

### Mini-project ideas
- Build a process inspector: list PIDs and read /proc/<pid>/status, /proc/<pid>/cmdline
- Build a signal playground: start child processes, handle SIGINT/SIGTERM, implement graceful shutdown
- Build a namespace viewer: print a process namespaces from /proc/<pid>/ns/* and compare two processes
- Build a cgroup v2 limiter: run a process and apply CPU/memory limits (in VM), then observe behavior
- Strace a simple Go server and write a short “syscalls used” summary (open/read/write/socket/epoll)

---
## Phase 3 — etcd fundamentals

### Goal
understand etcd as a strongly-consistent key-value store with watch semantics.

### Topics
- Raft and leader election
- MVCC and revisions
- watches and event streams

Mini-project ideas

- Build a Go etcd client: put/get/list + watch a prefix and print events
- Build a controller reacting to etcd: watch desired/ keys → write status/ keys
- Build leader election: 3 controllers compete via etcd lease; only leader reconciles
- Build a “config rollout” toy: update a key, watchers apply it to a local file and report status
- Build a failure demo: kill etcd leader (3-node cluster) and observe re-election + watch continuity

---
## Phase 4 — containerd + controlling containers (VM strongly recommended)

### Goal
understand the container runtime boundary and control containers programmatically.

### Topics
- OCI images vs containers
- containerd architecture
- CRI concepts

### Mini-project ideas
- Use ctr to pull/run/stop a container and document the lifecycle you observed
- Use crictl to run a pod sandbox + container (CRI flow) and inspect what gets created
- Build a Go program using containerd client: pull image → create → start → stop → delete
- Build a “run workload” CLI: run image cmd, ps, logs, stop
- Build a simple health watcher: periodically check running tasks and restart if exited

--- 
## Phase 5 — Mini Kubernetes-style orchestrator

### Goal:
Build a minimal orchestrator using VMs -Virtual Machies- that feels like Kubernetes

### Topics
- API server
- persistent state
- reconciliation
- node agents
- control plane concepts

Mini-project ideas
- Build an API server for Workload objects: Apply/Get/List/Delete + Status field
- Build a reconciler: desired workloads → running containers via containerd + update status
- Build a node-agent daemon: control plane assigns workloads → agent runs them locally
- Build a tiny scheduler: pick a node based on free CPU/memory (simple static numbers is fine)
- Build watch streams: clients watch workloads and see updates (long-poll → SSE later)

---

## How to use this repository

This repo is meant to be explored, not consumed passively.

- Read the code
- Read the notes
- Run the experiments
- Break things

---

## Inspiration

This project is inspired by:
- Kubernetes architecture
- containerd and OCI runtimes
- distributed systems principles
- learning from first principles

---

## Status

This project is actively evolving.
Structure and implementations will change as understanding deepens.
