#
# spec file for package sizhu-govcl
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


Name:           sizhu-govcl
Version:        0.0.3
Release:        0
Summary:        Constellation sizhu
License:        GPL
URL:            https://github.com/Aquarian-Age/ccal/
Source0:        sizhu-govcl.tar.gz
#BuildRequires:  
#Requires:     liblcl.so  
BuildRoot:      %{_tmppath}/%{name}-%{version}-build

%description
sizhu


%prep
mkdir -p %{_builddir}/%{name}-%{version}
cd %{_builddir}
tar xzvf %{_sourcedir}/sizhu-govcl.tar.gz -C %{_builddir}/%{name}-%{version}


%build
cd %{_builddir}/%{name}-%{version}
go mod tidy
go build -o sizhu-govcl -ldflags="-s -w" -tags tempdll .


%install
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/
cp -f %{_builddir}/%{name}-%{version}/sizhu-govcl %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/bin/sizhu-govcl
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/
mkdir -p %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/
cp -f %{_builddir}/%{name}-%{version}/sizhu-govcl.png %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/icons/sizhu-govcl.png
cp -f %{_builddir}/%{name}-%{version}/sizhu-govcl.desktop %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}/usr/local/share/applications/sizhu-govcl.desktop

%clean
rm -rf %{_buildrootdir}/%{name}-%{version}-%{release}.%{_arch}
rm -rf %{_builddir}/%{name}-%{version}

%files
%defattr(-,root,root)
/usr/local/bin/sizhu-govcl
/usr/local/share/applications/sizhu-govcl.desktop
/usr/local/share/icons/sizhu-govcl.png

%changelog
