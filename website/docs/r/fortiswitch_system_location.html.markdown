---
subcategory: "FortiSwitch System"
layout: "fortiswitch"
page_title: "FortiSwitch: fortiswitch_system_location"
description: |-
  Configure Location table.
---

# fortiswitch_system_location
Configure Location table.

## Example Usage

```hcl
resource "fortiswitch_system_location" "name" {
        coordinates {
            altitude = "164"
            altitude_unit = "m"
            datum = "WGS84"
            latitude = "224"
            longitude = "563"
        }
        name = "default_name_45"
}
```

## Argument Reference

The following arguments are supported:

* `address_civic` - Configure Location Civic Address. The structure of `address_civic` block is documented below.
* `elin_number` - Configure Location ELIN Number. The structure of `elin_number` block is documented below.
* `name` - Unique Location Item Name.
* `coordinates` - Configure Location GPS Coordinates. The structure of `coordinates` block is documented below.

The `address_civic` block supports:

* `country_subdivision` - National subdivisions (state, canton, region, province, or prefecture).
* `sub_branch_road` - Sub branch road name.
* `street_name_pre_mod` - Street name pre modifier.
* `number` - House number.
* `seat` - Seat number.
* `county` - County, parish, gun (JP), or district (IN).
* `street` - Street.
* `unit` - Unit (apartment, suite).
* `city` - City, township, or shi (JP).
* `additional` - Additional location information.
* `zip` - Postal/zip code.
* `floor` - Floor.
* `branch_road` - Branch road name.
* `street_name_post_mod` - Street name post modifier.
* `post_office_box` - Post office box (P.O. box).
* `primary_road` - Primary road name.
* `place_type` - Placetype.
* `direction` - Leading street direction.
* `street_suffix` - Street suffix.
* `road_section` - Road section.
* `number_suffix` - House number suffix.
* `name` - Name (residence and office occupant).
* `building` - Building (structure).
* `room` - Room number.
* `language` - Language.
* `additional_code` - Additional code.
* `country` - The two-letter ISO 3166 country code in capital ASCII letters eg. US, CA, DK, DE.
* `script` - Script used to present the address information.
* `city_division` - City division, borough, city district, ward, or chou (JP).
* `trailing_str_suffix` - Trailing street suffix.
* `landmark` - Landmark or vanity address.
* `block` - Neighborhood or block.
* `postal_community` - Postal community name.

The `elin_number` block supports:

* `elin_number` - Configure Elin Callback Number, 10 to 20 bytes numerial string.

The `coordinates` block supports:

* `latitude` - Floating point start with ( +/- )  or end with ( N or S ) eg. +/-16.67 or 16.67N.
* `datum` - WGS84, NAD83, NAD83/MLLW . Valid values: `WGS84`, `NAD83`, `NAD83/MLLW`.
* `altitude` - +/- Floating point no. eg. 117.47.
* `altitude_unit` - m ( meters), f ( floors). Valid values: `m`, `f`.
* `longitude` - Floating point start with ( +/- )  or end with ( E or W ) eg. +/-26.789 or 26.789E.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Location can be imported using any of these accepted formats:
```
$ terraform import fortiswitch_system_location.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTISWITCH_IMPORT_TABLE"="false"
$ terraform import fortiswitch_system_location.labelname {{name}}
$ unset "FORTISWITCH_IMPORT_TABLE"
```
