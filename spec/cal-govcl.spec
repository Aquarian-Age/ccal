#
# spec file for package cal-govcl
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


Name:           cal-govcl
Version:        0.0.6
Release:        0
Summary:        Constellation calendar
License:        GPL
URL:            https://github.com/Aquarian-Age/ccal/tree/amrta/govcl
Source0:        cal-govcl.tar.gz
#BuildRequires:  
#Requires:     liblcl.so  
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
Constellation calendar


%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/cal-govcl.tar.gz -C %{_builddir}/%{name}-%{version}


%build
cd %{_builddir}/%{name}-%{version}
go mod tidy
go build -o cal-govcl -ldflags="-s -w" -tags tempdll .


%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/
cp -f %{_builddir}/%{name}-%{version}/cal-govcl %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/cal-govcl
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/
cp -f %{_builddir}/%{name}-%{version}/cal-govcl.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/cal-govcl.png
cp -f %{_builddir}/%{name}-%{version}/cal-govcl.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/cal-govcl.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
/usr/local/bin/cal-govcl
/usr/local/share/applications/cal-govcl.desktop
/usr/local/share/icons/cal-govcl.png

%changelog
