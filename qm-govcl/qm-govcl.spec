#
# spec file for package qm-govcl.tar.gz
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


Name:          qm-govcl
Version:       2.0.9
Release:       0
Summary:       qm-govcl
License:       MIT
Group:         Application
Url:           https://github.com/Aquarian-Age/ccal/releases
Source0:       qm-govcl.tar.gz
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
日家奇门 时家奇门 紫白飞宫


%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/qm-govcl.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}/
/bin/bash ./build.sh

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/
cp -f %{_builddir}/%{name}-%{version}/qm-govcl %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/qm-govcl
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/
cp -f %{_builddir}/%{name}-%{version}/qm-govcl.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/qm-govcl.png
cp -f %{_builddir}/%{name}-%{version}/qm-govcl.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/qm-govcl.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
/usr/local/bin/qm-govcl
/usr/local/share/applications/qm-govcl.desktop
/usr/local/share/icons/qm-govcl.png


%changelog

