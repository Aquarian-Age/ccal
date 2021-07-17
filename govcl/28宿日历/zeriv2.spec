#
# spec file for package zeri
#
# Copyright (c) 2021 SUSE LLC
#
# All modifications and additions to the file contributed by third parties
# remain the property of their copyright owners, unless otherwise agreed
# upon. The license for this file, and modifications and additions to the
# file, is the same license as for the pristine package itself (unless the
# license for the pristine package is not an Open Source License, in which
# case the license is the MIT License). An "Open Source License" is a
# license that conforms to the Open Source Definition (Version 1.9)
# published by the Open Source Initiative.

# Please submit bugfixes or comments via https://bugs.opensuse.org/
#


Name:           zeri
Version:        0.9.9
Release:        3
Summary:        zeri
License:        MIT
URL:           https://github.com/Aquarian-Age/ccal/releases 
Source0:        zeri.tar.gz
#BuildRequires:  
#Requires:       
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
农历择日

%prep

mkdir -p /home/xuan/rpm/BUILD/%{name}-%{version}
cd /home/xuan/rpm/BUILD/%{name}-%{version}

%build

tar xzvf ../SOURCES/zeri.tar.gz -C /home/xuan/rpm/BUILD/%{name}-%{version}
cd /home/xuan/rpm/BUILD/%{name}-%{version}
go build -o zeri -i -tags tempdll -ldflags="-s -w" .


%install

mkdir -p %{buildroot}/usr/local/bin/
cp -f /home/xuan/rpm/BUILD/%{name}-%{version}/zeri %{buildroot}/usr/local/bin/zeri
mkdir -p %{buildroot}/usr/local/share/applications/
mkdir -p %{buildroot}/usr/local/share/icons/
cp -f /home/xuan/rpm/BUILD/%{name}-%{version}/zeri.png %{buildroot}/usr/local/share/icons/zeri.png
cp -f /home/xuan/rpm/BUILD/%{name}-%{version}/zeri.desktop %{buildroot}/usr/local/share/applications/zeri.desktop

rm -rf %{buildroot}/*.go
rm -rf %{buildroot}/go.*
rm -rf %{buildroot}/zeri*
rm -rf %{buildroot}/zeri.png


%files
%defattr(-,root,root)
/usr/local/bin/zeri
/usr/local/share/applications/zeri.desktop
/usr/local/share/icons/zeri.png


%changelog
