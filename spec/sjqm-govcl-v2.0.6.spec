#
# spec file for package sjqm-govcl-v2.0.6-v2.0.6.tar.gz
#
# Copyright (c) 2018 SUSE LINUX GmbH, Nuernberg, Germany.
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via http://bugs.opensuse.org/
#


Name:          sjqm-govcl
Version:       2.0.6
Release:       0
Summary:       sjqm-govcl
License:       MIT
Group:         Application
Url:           https://github.com/Aquarian-Age/sjqm/releases
Source0:       sjqm-govcl.tar.gz
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
紫白飞宫 天三门 地四户 地私门 太冲天马 孤虚 五符 五不遇
三奇入墓 时干入墓 入墓 三奇受制 天网四张 地网遮蔽
九遁 时干克应 八门克应(动应加静应) 三乍五假
荧入太白 太白入荧 六仪击刑 青龙反首 飞鸟跌穴
伏吟 反吟 门迫
年格(岁格) 月格 日格(伏干格) 时格 飞干格


%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/sjqm-govcl.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}
/bin/bash ./build.sh

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/
cp -f %{_builddir}/%{name}-%{version}/sjqm-govcl %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/sjqm-govcl
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/
cp -f %{_builddir}/%{name}-%{version}/sjqm-govcl.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/sjqm-govcl.png
cp -f %{_builddir}/%{name}-%{version}/sjqm-govcl.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/sjqm-govcl.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
/usr/local/bin/sjqm-govcl
/usr/local/share/applications/sjqm-govcl.desktop
/usr/local/share/icons/sjqm-govcl.png


%changelog

