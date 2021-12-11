#
# spec file for package zeRi-govcl_v5
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


Name:           zeRi-govcl
Version:        5.0.1
Release:        0
Summary:        zeRi
License:        MIT
URL:           https://github.com/Aquarian-Age/ccal/releases 
Source0:        zeRi-govcl.tar.gz
#BuildRequires:  
#Requires:       
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
协纪辩方书 时家奇门 日家奇门 


%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/zeRi-govcl.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}
go mod tidy
/bin/bash ./build.sh
#mv %{_builddir}/%{name}-%{version}/%{name}/zeri-govcl %{_builddir}/%{name}-%{version}/%{name}/zeRi-govcl_v5

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/
cp -f %{_builddir}/%{name}-%{version}/zeRi-govcl_v5 %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/zeRi-govcl_v5
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/
cp -f %{_builddir}/%{name}-%{version}/zeRi-govcl_v5.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/zeRi-govcl_v5.png
cp -f %{_builddir}/%{name}-%{version}/zeRi-govcl_v5.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/zeRi-govcl_v5.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
/usr/local/bin/zeRi-govcl_v5
/usr/local/share/applications/zeRi-govcl_v5.desktop
/usr/local/share/icons/zeRi-govcl_v5.png


%changelog
