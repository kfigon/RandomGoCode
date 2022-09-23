## Simple take on parsing DMI decode output
https://linux.die.net/man/8/dmidecode

credits to https://xmonader.github.io/nimdays/day01_dmidecode.html

## example dmidecode outputs:

```
# dmidecode 3.1
Getting SMBIOS data from sysfs.
SMBIOS 2.6 present.

Handle 0x0001, DMI type 1, 27 bytes
System Information
        Manufacturer: LENOVO
        Product Name: 20042
        Version: Lenovo G560
        Serial Number: 2677240001087
        UUID: CB3E6A50-A77B-E011-88E9-B870F4165734
        Wake-up Type: Power Switch
        SKU Number: Calpella_CRB
        Family: Intel_Mobile
```

```
Getting SMBIOS data from sysfs.
SMBIOS 2.6 present.

Handle 0x0000, DMI type 0, 24 bytes
BIOS Information
        Vendor: LENOVO
        Version: 29CN40WW(V2.17)
        Release Date: 04/13/2011
        ROM Size: 2048 kB
        Characteristics:
                PCI is supported
                BIOS is upgradeable
                BIOS shadowing is allowed
                Boot from CD is supported
                Selectable boot is supported
                EDD is supported
                Japanese floppy for NEC 9800 1.2 MB is supported (int 13h)
                Japanese floppy for Toshiba 1.2 MB is supported (int 13h)
                5.25"/360 kB floppy services are supported (int 13h)
                5.25"/1.2 MB floppy services are supported (int 13h)
                3.5"/720 kB floppy services are supported (int 13h)
                3.5"/2.88 MB floppy services are supported (int 13h)
                8042 keyboard services are supported (int 9h)
                CGA/mono video services are supported (int 10h)
                ACPI is supported
                USB legacy is supported
                BIOS boot specification is supported
                Targeted content distribution is supported
        BIOS Revision: 1.40
```

* DMIDecode output is some meta like comments, versions and one or more sections
* Section: consists of a
    * handle line
    * title line
    * one or more indented properties
* Property: consists of
    * key
    * optional value
    * optional list of indented items