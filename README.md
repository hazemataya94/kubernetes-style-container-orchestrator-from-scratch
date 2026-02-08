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

---
## Phase 3 — etcd fundamentals

### Goal
understand etcd as a strongly-consistent key-value store with watch semantics.

### Topics
- Raft and leader election
- MVCC and revisions
- watches and event streams

---
## Phase 4 — containerd + controlling containers (VM strongly recommended)

### Goal
understand the container runtime boundary and control containers programmatically.

### Topics
- OCI images vs containers
- containerd architecture
- CRI concepts

--- 
## Phase 5 — Mini Kubernetes-style orchestrator

### Goal:
Build a minimal orchestrator using VMs -Virtual Machies- that feels like Kubernetes
- API server
- persistent state
- reconciliation
- node agents
- control plane concepts

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
