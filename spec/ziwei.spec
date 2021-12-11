#
# spec file for package ziwei
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


Name:           ziwei
Version:        0.3.6
Release:        0
Summary:        ziWeiDouShu
License:        MIT
URL:            https://github.com/Aquarian-Age/zwds
Source0:        ziwei.tar.gz
BuildRequires: golang(API) = 1.15
BuildRequires:  gcc 
BuildRequires:  pkg-config 
BuildRoot:      %{_tmppath}/%{name}-%{version}-build   

%description
zi wei dou shu

%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/ziwei.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}

go build -o ziwei -mod=vendor -tags tempdll -ldflags="-s -w" -trimpath .

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -m0755 %{_builddir}/%{name}-%{version}/ziwei %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/ziwei

mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}
install -m0755 %{_builddir}/%{name}-%{version}/ziwei.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/ziwei.png
install -m0755 %{_builddir}/%{name}-%{version}/ziwei.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/ziwei.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/*.go
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/go.*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/ziwei*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/ziwei.png
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
%{_bindir}/ziwei
%{_datadir}/applications/ziwei.desktop
%{_datadir}/icons/ziwei.png


%changelog
