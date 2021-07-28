#
# spec file for package calendar-gioui.tar.gz
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


Name:          calendar-gioui
Version:       0.9.3
Release:       0
Summary:       Chinese Lunar Calendar
License:       MIT
Group:         Application
Url:           https://github.com/Aquarian-Age/ccal/releases
Source0:       calendar-gioui.tar.gz

Requires:       Mesa-libGLESv3-devel

BuildRequires: golang(API) = 1.16
BuildRequires:  gcc 
BuildRequires:  pkg-config 
BuildRequires:  wayland-devel 
BuildRequires:  libX11-devel 
BuildRequires:  libxkbcommon-x11-devel
BuildRequires:  Mesa-libEGL-devel
BuildRequires:  libXcursor-devel
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
Chinese Astrology Lunar Calendar

%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/calendar-gioui.tar.gz -C %{_builddir}/%{name}-%{version}

%build
cd %{_builddir}/%{name}-%{version}

go build -o calendar-gioui -mod=vendor -ldflags="-s -w" -trimpath .

%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/
install -m0755 %{_builddir}/%{name}-%{version}/calendar-gioui %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_bindir}/calendar-gioui

mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/
install -d %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}
install -m0755 %{_builddir}/%{name}-%{version}/calendar-gioui.svg %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/icons/calendar-gioui.svg
install -m0755 %{_builddir}/%{name}-%{version}/calendar-gioui.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/%{_datadir}/applications/calendar-gioui.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/*.go
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/go.*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/calendar-gioui*
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/calendar-gioui.svg
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
%{_bindir}/calendar-gioui
%{_datadir}/applications/calendar-gioui.desktop
%{_datadir}/icons/calendar-gioui.svg


%changelog
