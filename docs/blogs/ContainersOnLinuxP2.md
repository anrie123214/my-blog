---
title: Containers on Linux: Isolation Mechanism - Part 2
date: 2026-10-08
category: Containers
author: Nyatt
slug: container-isolation-p2
---
## **Containers on Linux: Isolation Mechanism - Part 2**

# **What sits behind Docker Engine?**

In **Part 1** , we learned that containers are built from Linux kernel features such as namespaces and cgroups. While these mechanisms provide isolation and resource control, interacting with them directly would be complex and inconvenient. 

Docker simplifies this process by providing a platform for building, running, and managing containers. 

At the heart of Docker is a background service called the **Docker daemon** , commonly referred to as **dockerd** . 

```
1     User
2     ↓
3     Docker CLI
4     ↓
5     dockerd
6     ↓
7     Linux Kernel
```

When a user executes a command such as: 

```text
1     docker run nginx
```

the Docker CLI sends the request to **dockerd** , which is responsible for creating the container, configuring **namespaces** and **cgroups** , mounting filesystems, and starting the application process. 

In other words, **users** interact with **Docker** commands, while **dockerd** performs the lowlevel work using Linux kernel features. 

# **Rootful Docker**

Traditionally, **Docker** operates in **rootful mode** . 

```text
1     Host
      └── 2 dockerd (root)
            └── 3 Container
```

In this setup: 

- **dockerd** runs with **root privileges** on the host. 

- Container creation requires privileged operations. 

- **Docker** can fully manage networking, storage, and resource isolation. 

The advantage is simplicity and compatibility. 

The disadvantage is that the Docker daemon is highly privileged. 

Because dockerd runs as root, a vulnerability affecting the daemon or container runtime could potentially have significant consequences for the host system. 

# **Rootless Docker**

To reduce this risk, Docker introduced **rootless mode** . 

```text
1     Host User
      └── 2 dockerd
            └── 3 Container
```

In rootless Docker: 

- Docker runs as a regular user. 

- The daemon does not require host root privileges. 

- Containers execute without a privileged daemon. 

This provides a stronger security boundary because compromising a container no longer directly implies access to a root-level Docker daemon. 

However, running containers without host root privileges creates a new challenge: 

How can a process inside a container appear to be root while remaining unprivileged on the host 

# **User Namespace Mapping**

The answer lies in **user namespaces** . 

A user namespace allows Linux to map user IDs inside a container to different user IDs on the host. 

For example: 

```text
1     Container UID 0 (root)
2     ↓
3     Host UID 100000
```

Inside the container:

```text
1     whoami
2     root
```

On the host:

```text
1     ps aux
```

might show the process running as UID 100000, which is an ordinary unprivileged user. 

This means: 

```text
1     Container root ≠! Host root
```

A process can possess root privileges within the container while having no administrative privileges outside of it. 

# **Why Namespace Mapping Matters**

Without user namespace mapping: 

```text
1     Container root
2     ↓
3     Host root
```

With user namespace mapping: 

```text
1     Container root
2     ↓
3     Host user 100000
```

The second model greatly reduces the impact of a container escape or misconfiguration. 

Instead of gaining root access to the host, an attacker would only obtain the permissions of an unprivileged user. 

# **Key Takeaway**

Containers are built using namespaces and cgroups, but Docker adds another layer of management through dockerd. 

Docker can operate in two modes: 

- **Rootful Docker** : dockerd runs as root. 

- **Rootless Docker** : dockerd runs as an unprivileged user. 

Rootless Docker relies on **user namespace mapping** , which allows a process to appear as root inside a container while remaining a non-root user on the host. 

**This is one of the most important security mechanisms in modern container environments and demonstrates how Linux namespaces are used not only for isolation but also for privilege separation.** 
