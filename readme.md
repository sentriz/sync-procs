<h3 align=center><b>sync-procs</b></h3>
<p align=center><i>synchronise processes in different shells so that they start at the same time</i></p>

---

### installation

```shell
    go install go.senan.xyz/sync-procs@latest
```

### usage

```shell
    sync-procs <num procs> # server
    sync-procs             # client
```

### example

_session A_

```shell
    $ sync-procs 3 && date # server
    # hanging until we have all 3
    Sat 11 Jun 2022 16:28:04 IST
```

_session B_

```shell
    $ sync-procs && date
    # hanging until we have all 3
    Sat 11 Jun 2022 16:28:04 IST
```

_session C_

```shell
    $ sync-procs && date
    # hanging until we have all 3
    Sat 11 Jun 2022 16:28:04 IST
```

---
