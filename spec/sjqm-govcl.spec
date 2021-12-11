#
# spec file for package sjqm-govcl.tar.gz
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
Version:       0.9.9 
Release:       2
Summary:       QiMen
License:       MIT
Group:         Application
Url:           https://github.com/Aquarian-Age/sjqm/releases
Source0:       sjqm-govcl.tar.gz
BuildRequires: golang(API) = 1.15
BuildRequires:  gcc 
BuildRequires:  pkg-config 
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
QiMen

%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/sjqm-govcl.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}

go build -o sjqm-govcl -mod=vendor -ldflags="-s -w" -trimpath -tags tempdll .

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -m0755 %{_builddir}/%{name}-%{version}/sjqm-govcl %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/sjqm-govcl

mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}
install -m0755 %{_builddir}/%{name}-%{version}/sjqm-govcl.svg %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/sjqm-govcl.svg
install -m0755 %{_builddir}/%{name}-%{version}/sjqm-govcl.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/sjqm-govcl.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/*.go
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/go.*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/sjqm-govcl*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/sjqm-govcl.svg
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
%{_bindir}/sjqm-govcl
%{_datadir}/applications/sjqm-govcl.desktop
%{_datadir}/icons/sjqm-govcl.svg


%changelog
