# go-uname

Wrapper for syscall uname

```
import (
  uname "github.com/msaf1980/go-uname"
)

..

  u, err := uname.New()
  if err == nil {
    machine := u.Machine()
    sysname := u.Sysname()
    nodename := u.Nodename()
    kernelVersion := u.KernelVersion()
    kernelRelease := u.KernelRelease()
  }
  ```
