## apk

### go version
go version go1.17 linux/amd64

### ndk version
- android-ndk-r23b-linux
ndk-build --version
```text
GNU Make 4.3.90
Built for x86_64-pc-linux-gnu
Copyright (C) 1988-2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
```

### ndk path
.bashrc |grep -i ndk
```bash
export PATH=/home/user/Android/android-ndk-r23b:$PATH
```

---

### build
```bash
go get fyne.io/fyne/v2
./build.sh
```