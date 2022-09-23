# my_app

## go serve
```bash
GOOS=linux GOARCH=arm64 go build -o serve-arm64

adb devices
adb push ./serve-arm64 /data/local/tmp

adb shell
cd /data/local/tmp/
./serve-arm64 &
```

## flutter client

```bash
adb install my_app.apk
```
