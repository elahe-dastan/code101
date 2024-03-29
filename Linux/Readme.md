# Cpu throttle
If you're in a `cgroup` (cgroups (abbreviated from control groups) is a Linux kernel feature that limits, accounts for, 
and isolates the resource usage (CPU, memory, disk I/O, network, etc.) of a collection of processes.) then you can use 
the command `cat /sys/fs/cgroup/cpu/cpu.stat` it will give you a result like:

```
nr_periods 0
nr_throttled 0
throttled_time 0
```
If nr_throttled increases every time you run the command it means your process is throttling.

# SSH local port forwarding
![ssh_port_forward](ssh_port_forward.png)

`ssh -L {local_port}:{remote_internal_ip}:{remote_port} {ssh_ip}`

# venv
## Create a new virtual ENV
`python -m venv `
## Use the virtual ENV
`source environment1 /bin/activate`

# SCP
Using SCP we can transfer file over TCP so it's a fast approach to consider. We need to have SSH connection permission to use SCP.

## Large files
SCP is a fast approach itself but if your file is so huge even SCP can take so long. To resolve the issue you can **enable compression**, it will be much faster but will also increase CPU usage.

# Subscribe From Nats
The command
```shell
kubectl run nats-box --image=natsio/nats-box:latest --rm -it --restart=Never --command --
```
gives you a terminal then run
```shell
nats sub -s {your_server} {your_topic}
```