
- [chineseLunar](https://github.com/Aquarian-Age/ccal/releases/tag/chineseLunar)

- openSuSE
```bash
sudo zypper ar -f http://download.opensuse.org/repositories/home:/liangzi/openSUSE_Tumbleweed/home:liangzi.repo
sudo zypper ref -r home_liangzi
sudo zypper in -y chineseLunar
```

- [奇门](https://github.com/Aquarian-Age/ccal/releases/tag/qm-govcl)


- [紫微斗数](https://github.com/Aquarian-Age/ccal/releases/tag/zwds-govcl)


- [协纪辩方书](https://github.com/Aquarian-Age/ccal/releases/tag/xjbfs-v6)


- [28宿日历](https://github.com/Aquarian-Age/ccal/releases/tag/28%E5%AE%BF%E6%97%A5%E5%8E%86)

- [四柱简排](https://github.com/Aquarian-Age/ccal/releases/tag/sizhu)

---

- rpm ded打包使用了[nfpm](https://github.com/goreleaser/nfpm)

- exe打包使用了[Inno Setup](https://jrsoftware.org/isinfo.php)

- rpm安装

```bash
sudo rpm -ivh xjbfs-6.2.0-a.x86_64.rpm --force --nodeps
```

- deb安装
```bash
sudo dpkg --install chineseLunar_1.0.6t-c_amd64.deb
```
