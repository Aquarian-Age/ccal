#
# spec file for package ctzw
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


Name:           ctzw
Version:        0.3.6
Release:        0
Summary:        ziWei
License:        MIT
URL:            https://github.com/Aquarian-Age/zwds
Source0:        ctzw.tar.gz
BuildRequires: golang(API) = 1.15
BuildRequires:  gcc 
BuildRequires:  pkg-config 
BuildRoot:      %{_tmppath}/%{name}-%{version}-build   

%description
ctzw

%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/ctzw.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}

go build -o ctzw -mod=vendor -tags tempdll -ldflags="-s -w" -trimpath .

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -m0755 %{_builddir}/%{name}-%{version}/ctzw %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/ctzw

mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}
install -m0755 %{_builddir}/%{name}-%{version}/ctzw.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/ctzw.png
install -m0755 %{_builddir}/%{name}-%{version}/ctzw.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/ctzw.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/*.go
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/go.*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/ctzw*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/ctzw.png
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
%{_bindir}/ctzw
%{_datadir}/applications/ctzw.desktop
%{_datadir}/icons/ctzw.png


%changelog
