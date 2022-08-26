## Bing Wallpaper

每天开机启动 自动下载当日壁纸到指定目录
源代码修改自[bing-pallpaper](https://github.com/TimothyYe/bing-wallpaper)

## systemd
```
mkdir -p ~/.local/share/systemd/user

vim .local/share/systemd/user/bing-wallpaper.service 

```

- bing-wallpaper.service
```text
[Unit]
Description=bing-wallpaper@home.service
After=network.target

[Service]
Type=simple
ExecStart=/home/user/bin/bing-wallpaper

[Install]
WantedBy=multi-user.target
```
- 设置开机启动
systemctl --user start bing-wallpaper.service
systemctl --user enable bing-wallpaper.service 
