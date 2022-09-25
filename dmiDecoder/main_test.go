package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const input1 = `# dmidecode 3.1
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
	Family: Intel_Mobile`

const input2 = `Getting SMBIOS data from sysfs.
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
	BIOS Revision: 1.40`

const input3 = `Getting SMBIOS data from sysfs.
SMBIOS 2.6 present.

Handle 0x0000, DMI type 0, 24 bytes
BIOS Information
	Vendor: LENOVO
	ROM Size: 2048 kB
	Characteristics:
		PCI is supported
		foobar is supported
	BIOS Revision: 1.40
		
Handle 0x0001, DMI type 1, 27 bytes
System Information
	Manufacturer: LENOVO
	Product Name: 20042
	SomeProps:
		Foo is supported`

func TestParse1(t *testing.T) {
	dmi := parse(input1)

	t.Run("basic structure", func(t *testing.T) {
		assertSection(t, dmi, "System Information", "Handle 0x0001, DMI type 1, 27 bytes", 8)
		assert.Len(t, dmi, 1)
	})
	
	t.Run("no items in props", func(t *testing.T) {
		props := dmi["System Information"].props
		for _,v := range props {
			assert.Len(t, v.items, 0)
		}
	})

	t.Run("properties", func(t *testing.T) {
		assert.Equal(t, dmi["System Information"].props["Manufacturer"].value, "LENOVO")
		assert.Equal(t, dmi["System Information"].props["Manufacturer"].name, "Manufacturer")
        assert.Equal(t, dmi["System Information"].props["Product Name"].value, "20042")
        assert.Equal(t, dmi["System Information"].props["Version"].value, "Lenovo G560")
        assert.Equal(t, dmi["System Information"].props["Serial Number"].value, "2677240001087")
        assert.Equal(t, dmi["System Information"].props["UUID"].value, "CB3E6A50-A77B-E011-88E9-B870F4165734")
        assert.Equal(t, dmi["System Information"].props["UUID"].name, "UUID")
        assert.Equal(t, dmi["System Information"].props["Wake-up Type"].value, "Power Switch")
        assert.Equal(t, dmi["System Information"].props["SKU Number"].value, "Calpella_CRB")
        assert.Equal(t, dmi["System Information"].props["Family"].value, "Intel_Mobile")
        assert.Equal(t, dmi["System Information"].props["Family"].name, "Family")
	})
}

func assertSection(t *testing.T, dmi dmiOutput, sectionName string, expectedHandle string, expectedProps int) {
	assert.Equal(t, expectedHandle, dmi[sectionName].handle)
	assert.Equal(t, sectionName, dmi[sectionName].title)
	assert.Len(t, dmi[sectionName].props, expectedProps)
}

func TestParse2(t *testing.T) {
	dmi := parse(input2)

	t.Run("basic structure", func(t *testing.T) {
		assertSection(t, dmi, "BIOS Information", "Handle 0x0000, DMI type 0, 24 bytes", 6)
		assert.Len(t, dmi, 1)
	})
	
	t.Run("props", func(t *testing.T) {
		props := dmi["BIOS Information"].props
		
		assert.Equal(t, props["Vendor"].name, "Vendor")
		assert.Equal(t, props["Vendor"].value, "LENOVO")
		assert.Equal(t, props["ROM Size"].name, "ROM Size")
		assert.Equal(t, props["ROM Size"].value, "2048 kB")
		assert.Equal(t, props["BIOS Revision"].name, "BIOS Revision")
		assert.Equal(t, props["BIOS Revision"].value, "1.40")

		assert.Len(t, props["Characteristics"].items, 18)
		assert.Equal(t, props["Characteristics"].name, "Characteristics")
		assert.Equal(t, props["Characteristics"].value, "")
	})
}

func TestParse3(t *testing.T) {
	dmi := parse(input3)

	t.Run("basic structure", func(t *testing.T) {
		assertSection(t, dmi, "BIOS Information", "Handle 0x0000, DMI type 0, 24 bytes", 4)
		assertSection(t, dmi, "System Information", "Handle 0x0001, DMI type 1, 27 bytes", 3)
		assert.Len(t, dmi, 2)
	})
	
	t.Run("props bios", func(t *testing.T) {
		props := dmi["BIOS Information"].props
		
		assert.Equal(t, props["Vendor"].name, "Vendor")
		assert.Equal(t, props["Vendor"].value, "LENOVO")
		assert.Equal(t, props["ROM Size"].name, "ROM Size")
		assert.Equal(t, props["ROM Size"].value, "2048 kB")
		assert.Equal(t, props["BIOS Revision"].name, "BIOS Revision")
		assert.Equal(t, props["BIOS Revision"].value, "1.40")

		assert.Len(t, props["Characteristics"].items, 2)
		assert.Equal(t, props["Characteristics"].name, "Characteristics")
		assert.Equal(t, props["Characteristics"].value, "")
	})

	t.Run("props system", func(t *testing.T) {
		props := dmi["System Information"].props
		
		assert.Equal(t, props["Manufacturer"].name, "Manufacturer")
		assert.Equal(t, props["Manufacturer"].value, "LENOVO")
		assert.Equal(t, props["Product Name"].name, "Product Name")
		assert.Equal(t, props["Product Name"].value, "20042")

		assert.Len(t, props["SomeProps"].items, 1)
		assert.Equal(t, props["SomeProps"].name, "SomeProps")
		assert.Equal(t, props["SomeProps"].value, "")
	})
}