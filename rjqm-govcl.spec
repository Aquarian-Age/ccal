#
# spec file for package rjqm-govcl.tar.gz
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


Name:          rjqm-govcl
Version:       0.0.0
Release:       1
Summary:       rjqm-govcl
License:       MIT
Group:         Application
Url:           https://github.com/Aquarian-Age/rjqm/releases
Source0:       rjqm-govcl.tar.gz
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
日家奇门 天乙贵人 时辰黄黑 截路空


%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/rjqm-govcl.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}
/bin/bash ./build.sh

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/
cp -f %{_builddir}/%{name}-%{version}/rjqm-govcl %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/rjqm-govcl
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/
cp -f %{_builddir}/%{name}-%{version}/rjqm-govcl.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/rjqm-govcl.png
cp -f %{_builddir}/%{name}-%{version}/rjqm-govcl.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/rjqm-govcl.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
/usr/local/bin/rjqm-govcl
/usr/local/share/applications/rjqm-govcl.desktop
/usr/local/share/icons/rjqm-govcl.png


%changelog

 
