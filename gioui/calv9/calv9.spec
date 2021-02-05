#
# spec file for package cal-auth
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


Name:           cal-auth
Version:        0.6.9
Release:        9
Summary:        Chinese Lunar Calendar
License:        me
URL:            https://github.com/Aquarian-Age/ccal/releases
Source0:        cal-gioui.tar.gz
BuildRoot:      %{_tmppath}-%{version}-%{release}-build
BuildRequires: gtk2-devel >= 2.0.0
%description
Chinese Lunar Calendar


%prep
rm -rf %{name}
rm -rf %{buildroot}

%build
tar xzvf ../SOURCES/cal-gioui.tar.gz -C %{buildroot}/
cd %{buildroot}
go build -o cal-auth -ldflags="-s -w"

%install
mkdir -p %{buildroot}/usr/local/bin/
cp -f %{buildroot}/cal-auth %{buildroot}/usr/local/bin/cal-auth
mkdir -p %{buildroot}/usr/local/share/applications/
mkdir -p %{buildroot}/usr/local/share/icons/
cp -f %{buildroot}/calendar.svg %{buildroot}/usr/local/share/icons/calendar.svg
cp -f %{buildroot}/cal-auth.desktop %{buildroot}/usr/local/share/applications/cal-auth.desktop

rm -rf %{buildroot}/*.go
rm -rf %{buildroot}/go.*
rm -rf %{buildroot}/cal-*
rm -rf %{buildroot}/*.svg

%files
%defattr(-,root,root)
/usr/local/bin/cal-auth
/usr/local/share/applications/cal-auth.desktop
/usr/local/share/icons/calendar.svg

%changelog
