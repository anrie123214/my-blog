---
title: Containers on Linux: Isolation Mechanism - Part 1
date: 2026-10-07
category: Containers
author: Nyatt
slug: namespaces-and-cgroups
---

## **Containers on Linux:**<br>**Isolation Mechanism - Part 1**

When I first started looking at containers, although I had some experience setting them up for personal use cases before, I found it hard to connect all the terminologies and concepts. Thankfully, I ended up with a somewhat good knowledge base and am writing this to share my findings. For the first one, let us answer some simple questions about Namespaces and Cgroups: 

# **What can a Linux process see?** 

Linux provides several types of **namespaces** , each isolating a different part of the system: 

- **PID** — isolates the process ID space. 

- **Network** — gives processes their own network interfaces, IP addresses, and routing tables. 

- **Mount** — provides a separate view of the filesystem and mount points. 

- **UTS** — isolates the hostname and domain name. 

- **IPC** — isolates inter-process communication resources. 

- **User** — isolates and maps user and group identities. 

- **Cgroup** — isolates a process's view of the cgroup hierarchy. 

**_Note: Cgroup namespaces are different from cgroups themselves_** _. Cgroups control resource usage; the cgroup namespace controls what part of the cgroup hierarchy a process can see._ 

A Linux process normally exists within the same overall system as other processes. It can see parts of the filesystem, processes, network interfaces, and other system resources. 

**Namespaces** allow Linux to isolate these views. 

A process can be placed into a namespace where it sees a different set of processes, network interfaces, mounts, or other system resources. 

The important idea is simple: 

### **Namespaces control what a process can see.** 

This is one of the main building blocks behind containers. 

Among them, **user namespaces** deserve a little more attention. They deal with something fundamental: identity and privileges, which empower securities on running containers on your machine. A process can have its own process IDs, network interfaces, and filesystem view, but we still need to understand what user that process actually is and what permissions it has. 

# **User namespaces: Who am I?** 

A process doesn't just have a view of the system. It also has an identity. 

Inside a container, a process may appear to run as **root** . But that identity does not necessarily have the same meaning or privileges on the host. 

Linux **user namespaces** provide a way to isolate and map user and group identities between different environments. 

The important idea here is: 

### **User namespaces control how a process's identity is represented and isolated.** 

Now we have two important pieces: 

```
Namespaces
    ↓
What can I see?
User namespaces
    ↓
Who am I?
```

But isolation isn't only about what a process can see or who it is. 

A process also needs resources to run. 

# **Cgroups: How much can I use?** 

This is where **cgroups** , short for control groups, come in. 

Cgroups allow Linux to group processes together and control how much system resources they can consume. 

They can be used to manage things such as: 

- CPU 

- Memory 

- Number of processes 

- I/O 

So we can add another piece: 

```
Namespaces
    ↓
What can I see?
```

```
User namespaces
    ↓
Who am I?
Cgroups
    ↓
How much can I use?
```

These are Linux kernel mechanisms. None of them are containers by themselves. 

But put these mechanisms together, along with filesystems, networking, and other security controls, and we have many of the pieces needed to build a container. 

# **And then we get to containers** 

So where does container engines like **Docker/Podman** fit into all of this? 

Docker isn't replacing these Linux mechanisms with some completely different container technology. It provides a way for us to **use these mechanisms without having to configure them manually** . 

When we run: 

- docker run nginx 

Docker takes the container image and the configuration we provide, then sets up the environment needed to run the application. 

Behind the scenes, that means creating the appropriate isolation, setting resource controls, preparing the container's filesystem, configuring networking, and finally starting the container's processes. 

We don't have to think about each of those steps every time we want to run an application. 

That's really the value Docker provides. 

Instead of thinking: 

- Create namespaces 

- Configure cgroups 

- Prepare a filesystem 

- Set up networking 

- Apply security restrictions 

- Start the process 

we can think: 

- docker run nginx 

**Containers are indeed magic. They enabled Linux isolation and resource-control mechanisms to be packaged into a convenient developer experience that can be interacted through a single command.** 
