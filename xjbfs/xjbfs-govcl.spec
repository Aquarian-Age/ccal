#
# spec file for package xjbfs-govcl
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


Name:           xjbfs-govcl
Version:        6.0.0
Release:        0
Summary:        xjbfs
License:        MIT
URL:           https://github.com/Aquarian-Age/ccal/releases 
Source0:        xjbfs-govcl.tar.gz
#BuildRequires:  
#Requires:       
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
协纪辩方书 


%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/xjbfs-govcl.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}
go mod tidy
/bin/bash ./build.sh


%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/
cp -f %{_builddir}/%{name}-%{version}/xjbfs-govcl %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/xjbfs-govcl
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/
cp -f %{_builddir}/%{name}-%{version}/xjbfs-govcl.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/xjbfs-govcl.png
cp -f %{_builddir}/%{name}-%{version}/xjbfs-govcl.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/xjbfs-govcl.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
/usr/local/bin/xjbfs-govcl
/usr/local/share/applications/xjbfs-govcl.desktop
/usr/local/share/icons/xjbfs-govcl.png


%changelog
