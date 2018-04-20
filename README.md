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
| -: | :-: |
| folder | copy full folder from host to another folder into image |
| hostname | set hostname |
| interface | set interface config include interface name, addresss, netmask, gateway |

## Contributing
Bug fixes and other improvements to the cloudimageeditor library are welcome at any time. The preferred submission method is to use pull request by github, or use git send-email to submit patches to cloudimageeditor@126.com. 
