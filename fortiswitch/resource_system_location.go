// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Configure Location table.

package fortiswitch

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocationCreate,
		Read:   resourceSystemLocationRead,
		Update: resourceSystemLocationUpdate,
		Delete: resourceSystemLocationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"address_civic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country_subdivision": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"sub_branch_road": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street_name_pre_mod": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"number": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"seat": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"county": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"unit": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"city": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"additional": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"zip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"floor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"branch_road": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street_name_post_mod": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"post_office_box": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"primary_road": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"place_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"direction": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"street_suffix": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"road_section": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"number_suffix": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"building": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"room": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"language": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"additional_code": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"country": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"script": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"city_division": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"trailing_str_suffix": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"landmark": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"block": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
						"postal_community": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"elin_number": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"elin_number": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"coordinates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"latitude": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"datum": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"altitude": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"altitude_unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"longitude": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemLocationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemLocation(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemLocation resource while getting object: %v", err)
	}

	o, err := c.CreateSystemLocation(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemLocation resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemLocation")
	}

	return resourceSystemLocationRead(d, m)
}

func resourceSystemLocationUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemLocation(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocation resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemLocation(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemLocation")
	}

	return resourceSystemLocationRead(d, m)
}

func resourceSystemLocationDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemLocation(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocationRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemLocation(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocation(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocation resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocationAddressCivic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "country_subdivision"
	if _, ok := i["country-subdivision"]; ok {

		result["country_subdivision"] = flattenSystemLocationAddressCivicCountrySubdivision(i["country-subdivision"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sub_branch_road"
	if _, ok := i["sub-branch-road"]; ok {

		result["sub_branch_road"] = flattenSystemLocationAddressCivicSubBranchRoad(i["sub-branch-road"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street_name_pre_mod"
	if _, ok := i["street-name-pre-mod"]; ok {

		result["street_name_pre_mod"] = flattenSystemLocationAddressCivicStreetNamePreMod(i["street-name-pre-mod"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "number"
	if _, ok := i["number"]; ok {

		result["number"] = flattenSystemLocationAddressCivicNumber(i["number"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "seat"
	if _, ok := i["seat"]; ok {

		result["seat"] = flattenSystemLocationAddressCivicSeat(i["seat"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "county"
	if _, ok := i["county"]; ok {

		result["county"] = flattenSystemLocationAddressCivicCounty(i["county"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street"
	if _, ok := i["street"]; ok {

		result["street"] = flattenSystemLocationAddressCivicStreet(i["street"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unit"
	if _, ok := i["unit"]; ok {

		result["unit"] = flattenSystemLocationAddressCivicUnit(i["unit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "city"
	if _, ok := i["city"]; ok {

		result["city"] = flattenSystemLocationAddressCivicCity(i["city"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "additional"
	if _, ok := i["additional"]; ok {

		result["additional"] = flattenSystemLocationAddressCivicAdditional(i["additional"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "zip"
	if _, ok := i["zip"]; ok {

		result["zip"] = flattenSystemLocationAddressCivicZip(i["zip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "floor"
	if _, ok := i["floor"]; ok {

		result["floor"] = flattenSystemLocationAddressCivicFloor(i["floor"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "branch_road"
	if _, ok := i["branch-road"]; ok {

		result["branch_road"] = flattenSystemLocationAddressCivicBranchRoad(i["branch-road"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street_name_post_mod"
	if _, ok := i["street-name-post-mod"]; ok {

		result["street_name_post_mod"] = flattenSystemLocationAddressCivicStreetNamePostMod(i["street-name-post-mod"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "post_office_box"
	if _, ok := i["post-office-box"]; ok {

		result["post_office_box"] = flattenSystemLocationAddressCivicPostOfficeBox(i["post-office-box"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "primary_road"
	if _, ok := i["primary-road"]; ok {

		result["primary_road"] = flattenSystemLocationAddressCivicPrimaryRoad(i["primary-road"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "place_type"
	if _, ok := i["place-type"]; ok {

		result["place_type"] = flattenSystemLocationAddressCivicPlaceType(i["place-type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "direction"
	if _, ok := i["direction"]; ok {

		result["direction"] = flattenSystemLocationAddressCivicDirection(i["direction"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "street_suffix"
	if _, ok := i["street-suffix"]; ok {

		result["street_suffix"] = flattenSystemLocationAddressCivicStreetSuffix(i["street-suffix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "road_section"
	if _, ok := i["road-section"]; ok {

		result["road_section"] = flattenSystemLocationAddressCivicRoadSection(i["road-section"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "number_suffix"
	if _, ok := i["number-suffix"]; ok {

		result["number_suffix"] = flattenSystemLocationAddressCivicNumberSuffix(i["number-suffix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {

		result["name"] = flattenSystemLocationAddressCivicName(i["name"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "building"
	if _, ok := i["building"]; ok {

		result["building"] = flattenSystemLocationAddressCivicBuilding(i["building"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "room"
	if _, ok := i["room"]; ok {

		result["room"] = flattenSystemLocationAddressCivicRoom(i["room"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "language"
	if _, ok := i["language"]; ok {

		result["language"] = flattenSystemLocationAddressCivicLanguage(i["language"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "additional_code"
	if _, ok := i["additional-code"]; ok {

		result["additional_code"] = flattenSystemLocationAddressCivicAdditionalCode(i["additional-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "country"
	if _, ok := i["country"]; ok {

		result["country"] = flattenSystemLocationAddressCivicCountry(i["country"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "script"
	if _, ok := i["script"]; ok {

		result["script"] = flattenSystemLocationAddressCivicScript(i["script"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "city_division"
	if _, ok := i["city-division"]; ok {

		result["city_division"] = flattenSystemLocationAddressCivicCityDivision(i["city-division"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "trailing_str_suffix"
	if _, ok := i["trailing-str-suffix"]; ok {

		result["trailing_str_suffix"] = flattenSystemLocationAddressCivicTrailingStrSuffix(i["trailing-str-suffix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "landmark"
	if _, ok := i["landmark"]; ok {

		result["landmark"] = flattenSystemLocationAddressCivicLandmark(i["landmark"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "block"
	if _, ok := i["block"]; ok {

		result["block"] = flattenSystemLocationAddressCivicBlock(i["block"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "postal_community"
	if _, ok := i["postal-community"]; ok {

		result["postal_community"] = flattenSystemLocationAddressCivicPostalCommunity(i["postal-community"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLocationAddressCivicCountrySubdivision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicSubBranchRoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicStreetNamePreMod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicSeat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicCounty(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicStreet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicCity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicAdditional(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicZip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicFloor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicBranchRoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicStreetNamePostMod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicPostOfficeBox(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicPrimaryRoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicPlaceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicStreetSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicRoadSection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicNumberSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicBuilding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicRoom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicLanguage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicAdditionalCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicCountry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicCityDivision(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicTrailingStrSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicLandmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationAddressCivicPostalCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationElinNumber(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "elin_number"
	if _, ok := i["elin-number"]; ok {

		result["elin_number"] = flattenSystemLocationElinNumberElinNumber(i["elin-number"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLocationElinNumberElinNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationCoordinates(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "latitude"
	if _, ok := i["latitude"]; ok {

		result["latitude"] = flattenSystemLocationCoordinatesLatitude(i["latitude"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "datum"
	if _, ok := i["datum"]; ok {

		result["datum"] = flattenSystemLocationCoordinatesDatum(i["datum"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "altitude"
	if _, ok := i["altitude"]; ok {

		result["altitude"] = flattenSystemLocationCoordinatesAltitude(i["altitude"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "altitude_unit"
	if _, ok := i["altitude-unit"]; ok {

		result["altitude_unit"] = flattenSystemLocationCoordinatesAltitudeUnit(i["altitude-unit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "longitude"
	if _, ok := i["longitude"]; ok {

		result["longitude"] = flattenSystemLocationCoordinatesLongitude(i["longitude"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLocationCoordinatesLatitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationCoordinatesDatum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationCoordinatesAltitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationCoordinatesAltitudeUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemLocationCoordinatesLongitude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemLocation(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("address_civic", flattenSystemLocationAddressCivic(o["address-civic"], d, "address_civic", sv)); err != nil {
			if !fortiAPIPatch(o["address-civic"]) {
				return fmt.Errorf("Error reading address_civic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("address_civic"); ok {
			if err = d.Set("address_civic", flattenSystemLocationAddressCivic(o["address-civic"], d, "address_civic", sv)); err != nil {
				if !fortiAPIPatch(o["address-civic"]) {
					return fmt.Errorf("Error reading address_civic: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("elin_number", flattenSystemLocationElinNumber(o["elin-number"], d, "elin_number", sv)); err != nil {
			if !fortiAPIPatch(o["elin-number"]) {
				return fmt.Errorf("Error reading elin_number: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("elin_number"); ok {
			if err = d.Set("elin_number", flattenSystemLocationElinNumber(o["elin-number"], d, "elin_number", sv)); err != nil {
				if !fortiAPIPatch(o["elin-number"]) {
					return fmt.Errorf("Error reading elin_number: %v", err)
				}
			}
		}
	}

	if err = d.Set("name", flattenSystemLocationName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("coordinates", flattenSystemLocationCoordinates(o["coordinates"], d, "coordinates", sv)); err != nil {
			if !fortiAPIPatch(o["coordinates"]) {
				return fmt.Errorf("Error reading coordinates: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("coordinates"); ok {
			if err = d.Set("coordinates", flattenSystemLocationCoordinates(o["coordinates"], d, "coordinates", sv)); err != nil {
				if !fortiAPIPatch(o["coordinates"]) {
					return fmt.Errorf("Error reading coordinates: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemLocationFortiTestDebug(d *schema.ResourceData, fswdebugsn int, fswdebugbeg int, fswdebugend int) {
	log.Printf(strconv.Itoa(fswdebugsn))
	e := validation.IntBetween(fswdebugbeg, fswdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiSwitch Ver", " "), e)
}

func expandSystemLocationAddressCivic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "country_subdivision"
	if _, ok := d.GetOk(pre_append); ok {

		result["country-subdivision"], _ = expandSystemLocationAddressCivicCountrySubdivision(d, i["country_subdivision"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sub_branch_road"
	if _, ok := d.GetOk(pre_append); ok {

		result["sub-branch-road"], _ = expandSystemLocationAddressCivicSubBranchRoad(d, i["sub_branch_road"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street_name_pre_mod"
	if _, ok := d.GetOk(pre_append); ok {

		result["street-name-pre-mod"], _ = expandSystemLocationAddressCivicStreetNamePreMod(d, i["street_name_pre_mod"], pre_append, sv)
	}
	pre_append = pre + ".0." + "number"
	if _, ok := d.GetOk(pre_append); ok {

		result["number"], _ = expandSystemLocationAddressCivicNumber(d, i["number"], pre_append, sv)
	}
	pre_append = pre + ".0." + "seat"
	if _, ok := d.GetOk(pre_append); ok {

		result["seat"], _ = expandSystemLocationAddressCivicSeat(d, i["seat"], pre_append, sv)
	}
	pre_append = pre + ".0." + "county"
	if _, ok := d.GetOk(pre_append); ok {

		result["county"], _ = expandSystemLocationAddressCivicCounty(d, i["county"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street"
	if _, ok := d.GetOk(pre_append); ok {

		result["street"], _ = expandSystemLocationAddressCivicStreet(d, i["street"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unit"
	if _, ok := d.GetOk(pre_append); ok {

		result["unit"], _ = expandSystemLocationAddressCivicUnit(d, i["unit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "city"
	if _, ok := d.GetOk(pre_append); ok {

		result["city"], _ = expandSystemLocationAddressCivicCity(d, i["city"], pre_append, sv)
	}
	pre_append = pre + ".0." + "additional"
	if _, ok := d.GetOk(pre_append); ok {

		result["additional"], _ = expandSystemLocationAddressCivicAdditional(d, i["additional"], pre_append, sv)
	}
	pre_append = pre + ".0." + "zip"
	if _, ok := d.GetOk(pre_append); ok {

		result["zip"], _ = expandSystemLocationAddressCivicZip(d, i["zip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "floor"
	if _, ok := d.GetOk(pre_append); ok {

		result["floor"], _ = expandSystemLocationAddressCivicFloor(d, i["floor"], pre_append, sv)
	}
	pre_append = pre + ".0." + "branch_road"
	if _, ok := d.GetOk(pre_append); ok {

		result["branch-road"], _ = expandSystemLocationAddressCivicBranchRoad(d, i["branch_road"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street_name_post_mod"
	if _, ok := d.GetOk(pre_append); ok {

		result["street-name-post-mod"], _ = expandSystemLocationAddressCivicStreetNamePostMod(d, i["street_name_post_mod"], pre_append, sv)
	}
	pre_append = pre + ".0." + "post_office_box"
	if _, ok := d.GetOk(pre_append); ok {

		result["post-office-box"], _ = expandSystemLocationAddressCivicPostOfficeBox(d, i["post_office_box"], pre_append, sv)
	}
	pre_append = pre + ".0." + "primary_road"
	if _, ok := d.GetOk(pre_append); ok {

		result["primary-road"], _ = expandSystemLocationAddressCivicPrimaryRoad(d, i["primary_road"], pre_append, sv)
	}
	pre_append = pre + ".0." + "place_type"
	if _, ok := d.GetOk(pre_append); ok {

		result["place-type"], _ = expandSystemLocationAddressCivicPlaceType(d, i["place_type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "direction"
	if _, ok := d.GetOk(pre_append); ok {

		result["direction"], _ = expandSystemLocationAddressCivicDirection(d, i["direction"], pre_append, sv)
	}
	pre_append = pre + ".0." + "street_suffix"
	if _, ok := d.GetOk(pre_append); ok {

		result["street-suffix"], _ = expandSystemLocationAddressCivicStreetSuffix(d, i["street_suffix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "road_section"
	if _, ok := d.GetOk(pre_append); ok {

		result["road-section"], _ = expandSystemLocationAddressCivicRoadSection(d, i["road_section"], pre_append, sv)
	}
	pre_append = pre + ".0." + "number_suffix"
	if _, ok := d.GetOk(pre_append); ok {

		result["number-suffix"], _ = expandSystemLocationAddressCivicNumberSuffix(d, i["number_suffix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok {

		result["name"], _ = expandSystemLocationAddressCivicName(d, i["name"], pre_append, sv)
	}
	pre_append = pre + ".0." + "building"
	if _, ok := d.GetOk(pre_append); ok {

		result["building"], _ = expandSystemLocationAddressCivicBuilding(d, i["building"], pre_append, sv)
	}
	pre_append = pre + ".0." + "room"
	if _, ok := d.GetOk(pre_append); ok {

		result["room"], _ = expandSystemLocationAddressCivicRoom(d, i["room"], pre_append, sv)
	}
	pre_append = pre + ".0." + "language"
	if _, ok := d.GetOk(pre_append); ok {

		result["language"], _ = expandSystemLocationAddressCivicLanguage(d, i["language"], pre_append, sv)
	}
	pre_append = pre + ".0." + "additional_code"
	if _, ok := d.GetOk(pre_append); ok {

		result["additional-code"], _ = expandSystemLocationAddressCivicAdditionalCode(d, i["additional_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "country"
	if _, ok := d.GetOk(pre_append); ok {

		result["country"], _ = expandSystemLocationAddressCivicCountry(d, i["country"], pre_append, sv)
	}
	pre_append = pre + ".0." + "script"
	if _, ok := d.GetOk(pre_append); ok {

		result["script"], _ = expandSystemLocationAddressCivicScript(d, i["script"], pre_append, sv)
	}
	pre_append = pre + ".0." + "city_division"
	if _, ok := d.GetOk(pre_append); ok {

		result["city-division"], _ = expandSystemLocationAddressCivicCityDivision(d, i["city_division"], pre_append, sv)
	}
	pre_append = pre + ".0." + "trailing_str_suffix"
	if _, ok := d.GetOk(pre_append); ok {

		result["trailing-str-suffix"], _ = expandSystemLocationAddressCivicTrailingStrSuffix(d, i["trailing_str_suffix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "landmark"
	if _, ok := d.GetOk(pre_append); ok {

		result["landmark"], _ = expandSystemLocationAddressCivicLandmark(d, i["landmark"], pre_append, sv)
	}
	pre_append = pre + ".0." + "block"
	if _, ok := d.GetOk(pre_append); ok {

		result["block"], _ = expandSystemLocationAddressCivicBlock(d, i["block"], pre_append, sv)
	}
	pre_append = pre + ".0." + "postal_community"
	if _, ok := d.GetOk(pre_append); ok {

		result["postal-community"], _ = expandSystemLocationAddressCivicPostalCommunity(d, i["postal_community"], pre_append, sv)
	}

	return result, nil
}

func expandSystemLocationAddressCivicCountrySubdivision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicSubBranchRoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicStreetNamePreMod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicSeat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicCounty(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicStreet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicCity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicAdditional(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicZip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicFloor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicBranchRoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicStreetNamePostMod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicPostOfficeBox(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicPrimaryRoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicPlaceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicStreetSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicRoadSection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicNumberSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicBuilding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicRoom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicLanguage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicAdditionalCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicCountry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicCityDivision(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicTrailingStrSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicLandmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationAddressCivicPostalCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationElinNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "elin_number"
	if _, ok := d.GetOk(pre_append); ok {

		result["elin-number"], _ = expandSystemLocationElinNumberElinNumber(d, i["elin_number"], pre_append, sv)
	}

	return result, nil
}

func expandSystemLocationElinNumberElinNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationCoordinates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "latitude"
	if _, ok := d.GetOk(pre_append); ok {

		result["latitude"], _ = expandSystemLocationCoordinatesLatitude(d, i["latitude"], pre_append, sv)
	}
	pre_append = pre + ".0." + "datum"
	if _, ok := d.GetOk(pre_append); ok {

		result["datum"], _ = expandSystemLocationCoordinatesDatum(d, i["datum"], pre_append, sv)
	}
	pre_append = pre + ".0." + "altitude"
	if _, ok := d.GetOk(pre_append); ok {

		result["altitude"], _ = expandSystemLocationCoordinatesAltitude(d, i["altitude"], pre_append, sv)
	}
	pre_append = pre + ".0." + "altitude_unit"
	if _, ok := d.GetOk(pre_append); ok {

		result["altitude-unit"], _ = expandSystemLocationCoordinatesAltitudeUnit(d, i["altitude_unit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "longitude"
	if _, ok := d.GetOk(pre_append); ok {

		result["longitude"], _ = expandSystemLocationCoordinatesLongitude(d, i["longitude"], pre_append, sv)
	}

	return result, nil
}

func expandSystemLocationCoordinatesLatitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationCoordinatesDatum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationCoordinatesAltitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationCoordinatesAltitudeUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemLocationCoordinatesLongitude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocation(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address_civic"); ok {

		t, err := expandSystemLocationAddressCivic(d, v, "address_civic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-civic"] = t
		}
	}

	if v, ok := d.GetOk("elin_number"); ok {

		t, err := expandSystemLocationElinNumber(d, v, "elin_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["elin-number"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemLocationName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("coordinates"); ok {

		t, err := expandSystemLocationCoordinates(d, v, "coordinates", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["coordinates"] = t
		}
	}

	return &obj, nil
}
