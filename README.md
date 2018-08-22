This tool inspects tar headers.

Example:
```
[kmeaw@secrethost:~]$ tar --xattrs -c /usr/bin/nvidia-docker-plugin | ./describe-tar
tar: Removing leading `/' from member names
&tar.Header{
    Typeflag: 0x30,
    Name:     "usr/bin/nvidia-docker-plugin",
    Linkname: "",
    Size:     7187912,
    Mode:     493,
    Uid:      0,
    Gid:      0,
    Uname:    "root",
    Gname:    "root",
    ModTime:  time.Time{
        wall: 0x0,
        ext:  63624102890,
        loc:  &time.Location{},
    },
    AccessTime: time.Time{
        wall: 0x34959764,
        ext:  63670545148,
        loc:  &time.Location{},
    },
    ChangeTime: time.Time{
        wall: 0x24df68ab,
        ext:  63666808001,
        loc:  &time.Location{},
    },
    Devmajor:   0,
    Devminor:   0,
    Xattrs:     {"security.capability":"\x01\x00\x00\x02\b\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
    PAXRecords: {"atime":"1534948348.882218852", "ctime":"1531211201.618621099", "SCHILY.xattr.security.capability":"\x01\x00\x00\x02\b\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"},
    Format:     4,
}
```
