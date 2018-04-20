# cloudimageeditor

## About
A lot of OS provide cloud images as qcow2 format. Before launching the image, we usually want to inject something (config file, binary or script) into it.
Package cloudimageeditor aims at easy-use editor for cloud image.

## Dependence
Before using cloudimageeditor, you should pre-install some commands :
```
    mount
    umount
    cp
    sync
    rmmod
    modprobe
    partprobe
    qemu-img
    qemu-nbd
```

## Module list
| Module | Description |
| -: | -l |
| folder | copy full folder from host to another folder into image |
| hostname | set hostname |
| interface | set interface config include interface name, addresss, netmask, gateway |