// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Xing Li (@lix-fortinet), Jie Xue (@JieX19)
// Yue Wang (@yuew-ftnt), Liang Liu (@Liang_FTNT)

// Description: Description: SDK for FortiSwitch Provider

package forticlient

type creatUpdateOutput struct {
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// UpdateAlertemailSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the alertemail - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateAlertemailSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/alertemail/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteAlertemailSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the alertemail - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteAlertemailSetting(mkey string) (err error) {

	//No unset API for alertemail - setting
	return
}

// ReadAlertemailSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the alertemail - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadAlertemailSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/alertemail/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateGuiConsole API operation for FortiSwitch updates the specified Console.
// Returns the index value of the Console and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the gui - console chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateGuiConsole(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/gui/console"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteGuiConsole API operation for FortiSwitch deletes the specified Console.
// Returns error for service API and SDK errors.
// See the gui - console chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteGuiConsole(mkey string) (err error) {

	//No unset API for gui - console
	return
}

// ReadGuiConsole API operation for FortiSwitch gets the Console
// with the specified index value.
// Returns the requested Console value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the gui - console chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadGuiConsole(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/gui/console"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateLogCustomField API operation for FortiSwitch creates a new Custom Field.
// Returns the index value of the Custom Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateLogCustomField(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/log/custom-field"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateLogCustomField API operation for FortiSwitch updates the specified Custom Field.
// Returns the index value of the Custom Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogCustomField(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/custom-field"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogCustomField API operation for FortiSwitch deletes the specified Custom Field.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogCustomField(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/log/custom-field"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadLogCustomField API operation for FortiSwitch gets the Custom Field
// with the specified index value.
// Returns the requested Custom Field value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - custom-field chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogCustomField(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/custom-field"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateLogEventfilter API operation for FortiSwitch updates the specified Eventfilter.
// Returns the index value of the Eventfilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - eventfilter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogEventfilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/eventfilter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogEventfilter API operation for FortiSwitch deletes the specified Eventfilter.
// Returns error for service API and SDK errors.
// See the log - eventfilter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogEventfilter(mkey string) (err error) {

	//No unset API for log - eventfilter
	return
}

// ReadLogEventfilter API operation for FortiSwitch gets the Eventfilter
// with the specified index value.
// Returns the requested Eventfilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - eventfilter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogEventfilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/eventfilter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogGui API operation for FortiSwitch updates the specified Gui.
// Returns the index value of the Gui and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - gui chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogGui(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log/gui"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogGui API operation for FortiSwitch deletes the specified Gui.
// Returns error for service API and SDK errors.
// See the log - gui chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogGui(mkey string) (err error) {

	//No unset API for log - gui
	return
}

// ReadLogGui API operation for FortiSwitch gets the Gui
// with the specified index value.
// Returns the requested Gui value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log - gui chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogGui(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log/gui"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogDiskFilter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogDiskFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.disk/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogDiskFilter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.disk - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogDiskFilter(mkey string) (err error) {

	//No unset API for log.disk - filter
	return
}

// ReadLogDiskFilter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogDiskFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.disk/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogDiskSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogDiskSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.disk/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogDiskSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.disk - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogDiskSetting(mkey string) (err error) {

	//No unset API for log.disk - setting
	return
}

// ReadLogDiskSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.disk - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogDiskSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.disk/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzerFilter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerFilter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerFilter(mkey string) (err error) {

	//No unset API for log.fortianalyzer - filter
	return
}

// ReadLogFortianalyzerFilter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzerOverrideFilter API operation for FortiSwitch updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerOverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerOverrideFilter API operation for FortiSwitch deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerOverrideFilter(mkey string) (err error) {

	//No unset API for log.fortianalyzer - override-filter
	return
}

// ReadLogFortianalyzerOverrideFilter API operation for FortiSwitch gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerOverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzerOverrideSetting API operation for FortiSwitch updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerOverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerOverrideSetting API operation for FortiSwitch deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerOverrideSetting(mkey string) (err error) {

	//No unset API for log.fortianalyzer - override-setting
	return
}

// ReadLogFortianalyzerOverrideSetting API operation for FortiSwitch gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - override-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerOverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzerSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzerSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzerSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzerSetting(mkey string) (err error) {

	//No unset API for log.fortianalyzer - setting
	return
}

// ReadLogFortianalyzerSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzerSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzer2Filter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer2Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer2/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer2Filter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer2Filter(mkey string) (err error) {

	//No unset API for log.fortianalyzer2 - filter
	return
}

// ReadLogFortianalyzer2Filter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer2Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer2/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzer2Setting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer2Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer2/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer2Setting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer2Setting(mkey string) (err error) {

	//No unset API for log.fortianalyzer2 - setting
	return
}

// ReadLogFortianalyzer2Setting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer2 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer2Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer2/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzer3Filter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer3Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer3/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer3Filter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer3Filter(mkey string) (err error) {

	//No unset API for log.fortianalyzer3 - filter
	return
}

// ReadLogFortianalyzer3Filter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer3Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer3/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortianalyzer3Setting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortianalyzer3Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortianalyzer3/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortianalyzer3Setting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortianalyzer3Setting(mkey string) (err error) {

	//No unset API for log.fortianalyzer3 - setting
	return
}

// ReadLogFortianalyzer3Setting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortianalyzer3 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortianalyzer3Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortianalyzer3/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogFortiguardSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogFortiguardSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.fortiguard/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogFortiguardSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.fortiguard - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogFortiguardSetting(mkey string) (err error) {

	//No unset API for log.fortiguard - setting
	return
}

// ReadLogFortiguardSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.fortiguard - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogFortiguardSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.fortiguard/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogMemoryFilter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogMemoryFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.memory/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogMemoryFilter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.memory - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogMemoryFilter(mkey string) (err error) {

	//No unset API for log.memory - filter
	return
}

// ReadLogMemoryFilter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogMemoryFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.memory/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogMemoryGlobalSetting API operation for FortiSwitch updates the specified Global Setting.
// Returns the index value of the Global Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - global-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogMemoryGlobalSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.memory/global-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogMemoryGlobalSetting API operation for FortiSwitch deletes the specified Global Setting.
// Returns error for service API and SDK errors.
// See the log.memory - global-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogMemoryGlobalSetting(mkey string) (err error) {

	//No unset API for log.memory - global-setting
	return
}

// ReadLogMemoryGlobalSetting API operation for FortiSwitch gets the Global Setting
// with the specified index value.
// Returns the requested Global Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - global-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogMemoryGlobalSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.memory/global-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogMemorySetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogMemorySetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.memory/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogMemorySetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.memory - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogMemorySetting(mkey string) (err error) {

	//No unset API for log.memory - setting
	return
}

// ReadLogMemorySetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.memory - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogMemorySetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.memory/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogRemoteSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.remote - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogRemoteSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.remote/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogRemoteSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.remote - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogRemoteSetting(mkey string) (err error) {

	//No unset API for log.remote - setting
	return
}

// ReadLogRemoteSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.remote - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogRemoteSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.remote/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogdFilter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdFilter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdFilter(mkey string) (err error) {

	//No unset API for log.syslogd - filter
	return
}

// ReadLogSyslogdFilter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogdOverrideFilter API operation for FortiSwitch updates the specified Override Filter.
// Returns the index value of the Override Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdOverrideFilter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/override-filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdOverrideFilter API operation for FortiSwitch deletes the specified Override Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdOverrideFilter(mkey string) (err error) {

	//No unset API for log.syslogd - override-filter
	return
}

// ReadLogSyslogdOverrideFilter API operation for FortiSwitch gets the Override Filter
// with the specified index value.
// Returns the requested Override Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdOverrideFilter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/override-filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogdOverrideSetting API operation for FortiSwitch updates the specified Override Setting.
// Returns the index value of the Override Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdOverrideSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/override-setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdOverrideSetting API operation for FortiSwitch deletes the specified Override Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdOverrideSetting(mkey string) (err error) {

	//No unset API for log.syslogd - override-setting
	return
}

// ReadLogSyslogdOverrideSetting API operation for FortiSwitch gets the Override Setting
// with the specified index value.
// Returns the requested Override Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - override-setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdOverrideSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/override-setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogdSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogdSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogdSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogdSetting(mkey string) (err error) {

	//No unset API for log.syslogd - setting
	return
}

// ReadLogSyslogdSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogdSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogd2Filter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd2Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd2/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd2Filter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd2Filter(mkey string) (err error) {

	//No unset API for log.syslogd2 - filter
	return
}

// ReadLogSyslogd2Filter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd2Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd2/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogd2Setting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd2Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd2/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd2Setting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd2Setting(mkey string) (err error) {

	//No unset API for log.syslogd2 - setting
	return
}

// ReadLogSyslogd2Setting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd2 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd2Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd2/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogd3Filter API operation for FortiSwitch updates the specified Filter.
// Returns the index value of the Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd3Filter(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd3/filter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd3Filter API operation for FortiSwitch deletes the specified Filter.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd3Filter(mkey string) (err error) {

	//No unset API for log.syslogd3 - filter
	return
}

// ReadLogSyslogd3Filter API operation for FortiSwitch gets the Filter
// with the specified index value.
// Returns the requested Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - filter chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd3Filter(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd3/filter"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateLogSyslogd3Setting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateLogSyslogd3Setting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/log.syslogd3/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteLogSyslogd3Setting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteLogSyslogd3Setting(mkey string) (err error) {

	//No unset API for log.syslogd3 - setting
	return
}

// ReadLogSyslogd3Setting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the log.syslogd3 - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadLogSyslogd3Setting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/log.syslogd3/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterAccessList API operation for FortiSwitch creates a new Access List.
// Returns the index value of the Access List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAccessList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/access-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAccessList API operation for FortiSwitch updates the specified Access List.
// Returns the index value of the Access List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAccessList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/access-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAccessList API operation for FortiSwitch deletes the specified Access List.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAccessList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/access-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAccessList API operation for FortiSwitch gets the Access List
// with the specified index value.
// Returns the requested Access List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAccessList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/access-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterAccessList6 API operation for FortiSwitch creates a new Access List6.
// Returns the index value of the Access List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAccessList6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/access-list6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAccessList6 API operation for FortiSwitch updates the specified Access List6.
// Returns the index value of the Access List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAccessList6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/access-list6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAccessList6 API operation for FortiSwitch deletes the specified Access List6.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAccessList6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/access-list6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAccessList6 API operation for FortiSwitch gets the Access List6
// with the specified index value.
// Returns the requested Access List6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - access-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAccessList6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/access-list6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterAspathList API operation for FortiSwitch creates a new Aspath List.
// Returns the index value of the Aspath List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAspathList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/aspath-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAspathList API operation for FortiSwitch updates the specified Aspath List.
// Returns the index value of the Aspath List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAspathList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/aspath-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAspathList API operation for FortiSwitch deletes the specified Aspath List.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAspathList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/aspath-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAspathList API operation for FortiSwitch gets the Aspath List
// with the specified index value.
// Returns the requested Aspath List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - aspath-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAspathList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/aspath-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterAuthPath API operation for FortiSwitch creates a new Auth Path.
// Returns the index value of the Auth Path and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterAuthPath(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/auth-path"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterAuthPath API operation for FortiSwitch updates the specified Auth Path.
// Returns the index value of the Auth Path and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterAuthPath(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/auth-path"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterAuthPath API operation for FortiSwitch deletes the specified Auth Path.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterAuthPath(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/auth-path"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterAuthPath API operation for FortiSwitch gets the Auth Path
// with the specified index value.
// Returns the requested Auth Path value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - auth-path chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterAuthPath(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/auth-path"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateRouterBgp API operation for FortiSwitch updates the specified Bgp.
// Returns the index value of the Bgp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bgp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterBgp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/bgp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterBgp API operation for FortiSwitch deletes the specified Bgp.
// Returns error for service API and SDK errors.
// See the router - bgp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterBgp(mkey string) (err error) {

	//No unset API for router - bgp
	return
}

// ReadRouterBgp API operation for FortiSwitch gets the Bgp
// with the specified index value.
// Returns the requested Bgp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - bgp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterBgp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/bgp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterCommunityList API operation for FortiSwitch creates a new Community List.
// Returns the index value of the Community List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterCommunityList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/community-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterCommunityList API operation for FortiSwitch updates the specified Community List.
// Returns the index value of the Community List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterCommunityList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/community-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterCommunityList API operation for FortiSwitch deletes the specified Community List.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterCommunityList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/community-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterCommunityList API operation for FortiSwitch gets the Community List
// with the specified index value.
// Returns the requested Community List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - community-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterCommunityList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/community-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterGwdetect API operation for FortiSwitch creates a new Gwdetect.
// Returns the index value of the Gwdetect and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - gwdetect chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterGwdetect(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/gwdetect"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterGwdetect API operation for FortiSwitch updates the specified Gwdetect.
// Returns the index value of the Gwdetect and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - gwdetect chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterGwdetect(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/gwdetect"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterGwdetect API operation for FortiSwitch deletes the specified Gwdetect.
// Returns error for service API and SDK errors.
// See the router - gwdetect chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterGwdetect(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/gwdetect"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterGwdetect API operation for FortiSwitch gets the Gwdetect
// with the specified index value.
// Returns the requested Gwdetect value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - gwdetect chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterGwdetect(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/gwdetect"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateRouterIsis API operation for FortiSwitch updates the specified Isis.
// Returns the index value of the Isis and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - isis chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterIsis(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/isis"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterIsis API operation for FortiSwitch deletes the specified Isis.
// Returns error for service API and SDK errors.
// See the router - isis chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterIsis(mkey string) (err error) {

	//No unset API for router - isis
	return
}

// ReadRouterIsis API operation for FortiSwitch gets the Isis
// with the specified index value.
// Returns the requested Isis value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - isis chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterIsis(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/isis"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterKeyChain API operation for FortiSwitch creates a new Key Chain.
// Returns the index value of the Key Chain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterKeyChain(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/key-chain"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterKeyChain API operation for FortiSwitch updates the specified Key Chain.
// Returns the index value of the Key Chain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterKeyChain(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/key-chain"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterKeyChain API operation for FortiSwitch deletes the specified Key Chain.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterKeyChain(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/key-chain"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterKeyChain API operation for FortiSwitch gets the Key Chain
// with the specified index value.
// Returns the requested Key Chain value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - key-chain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterKeyChain(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/key-chain"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateRouterMulticast API operation for FortiSwitch updates the specified Multicast.
// Returns the index value of the Multicast and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterMulticast(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/multicast"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterMulticast API operation for FortiSwitch deletes the specified Multicast.
// Returns error for service API and SDK errors.
// See the router - multicast chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterMulticast(mkey string) (err error) {

	//No unset API for router - multicast
	return
}

// ReadRouterMulticast API operation for FortiSwitch gets the Multicast
// with the specified index value.
// Returns the requested Multicast value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterMulticast(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/multicast"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterMulticastFlow API operation for FortiSwitch creates a new Multicast Flow.
// Returns the index value of the Multicast Flow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterMulticastFlow(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/multicast-flow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterMulticastFlow API operation for FortiSwitch updates the specified Multicast Flow.
// Returns the index value of the Multicast Flow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterMulticastFlow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/multicast-flow"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterMulticastFlow API operation for FortiSwitch deletes the specified Multicast Flow.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterMulticastFlow(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/multicast-flow"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterMulticastFlow API operation for FortiSwitch gets the Multicast Flow
// with the specified index value.
// Returns the requested Multicast Flow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - multicast-flow chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterMulticastFlow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/multicast-flow"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateRouterOspf API operation for FortiSwitch updates the specified Ospf.
// Returns the index value of the Ospf and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterOspf(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterOspf API operation for FortiSwitch deletes the specified Ospf.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterOspf(mkey string) (err error) {

	//No unset API for router - ospf
	return
}

// ReadRouterOspf API operation for FortiSwitch gets the Ospf
// with the specified index value.
// Returns the requested Ospf value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterOspf(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateRouterOspf6 API operation for FortiSwitch updates the specified Ospf6.
// Returns the index value of the Ospf6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterOspf6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterOspf6 API operation for FortiSwitch deletes the specified Ospf6.
// Returns error for service API and SDK errors.
// See the router - ospf6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterOspf6(mkey string) (err error) {

	//No unset API for router - ospf6
	return
}

// ReadRouterOspf6 API operation for FortiSwitch gets the Ospf6
// with the specified index value.
// Returns the requested Ospf6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ospf6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterOspf6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf6"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateRouterPolicy API operation for FortiSwitch updates the specified Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterPolicy API operation for FortiSwitch deletes the specified Policy.
// Returns error for service API and SDK errors.
// See the router - policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterPolicy(mkey string) (err error) {

	//No unset API for router - policy
	return
}

// ReadRouterPolicy API operation for FortiSwitch gets the Policy
// with the specified index value.
// Returns the requested Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/policy"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterPrefixList API operation for FortiSwitch creates a new Prefix List.
// Returns the index value of the Prefix List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterPrefixList(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/prefix-list"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterPrefixList API operation for FortiSwitch updates the specified Prefix List.
// Returns the index value of the Prefix List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterPrefixList(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/prefix-list"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterPrefixList API operation for FortiSwitch deletes the specified Prefix List.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterPrefixList(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/prefix-list"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterPrefixList API operation for FortiSwitch gets the Prefix List
// with the specified index value.
// Returns the requested Prefix List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterPrefixList(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/prefix-list"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterPrefixList6 API operation for FortiSwitch creates a new Prefix List6.
// Returns the index value of the Prefix List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterPrefixList6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/prefix-list6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterPrefixList6 API operation for FortiSwitch updates the specified Prefix List6.
// Returns the index value of the Prefix List6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterPrefixList6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/prefix-list6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterPrefixList6 API operation for FortiSwitch deletes the specified Prefix List6.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterPrefixList6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/prefix-list6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterPrefixList6 API operation for FortiSwitch gets the Prefix List6
// with the specified index value.
// Returns the requested Prefix List6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - prefix-list6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterPrefixList6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/prefix-list6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateRouterRip API operation for FortiSwitch updates the specified Rip.
// Returns the index value of the Rip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - rip chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterRip(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/rip"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterRip API operation for FortiSwitch deletes the specified Rip.
// Returns error for service API and SDK errors.
// See the router - rip chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterRip(mkey string) (err error) {

	//No unset API for router - rip
	return
}

// ReadRouterRip API operation for FortiSwitch gets the Rip
// with the specified index value.
// Returns the requested Rip value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - rip chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterRip(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/rip"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateRouterRipng API operation for FortiSwitch updates the specified Ripng.
// Returns the index value of the Ripng and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ripng chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterRipng(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ripng"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterRipng API operation for FortiSwitch deletes the specified Ripng.
// Returns error for service API and SDK errors.
// See the router - ripng chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterRipng(mkey string) (err error) {

	//No unset API for router - ripng
	return
}

// ReadRouterRipng API operation for FortiSwitch gets the Ripng
// with the specified index value.
// Returns the requested Ripng value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - ripng chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterRipng(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ripng"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterRouteMap API operation for FortiSwitch creates a new Route Map.
// Returns the index value of the Route Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterRouteMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/route-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterRouteMap API operation for FortiSwitch updates the specified Route Map.
// Returns the index value of the Route Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterRouteMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/route-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterRouteMap API operation for FortiSwitch deletes the specified Route Map.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterRouteMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/route-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterRouteMap API operation for FortiSwitch gets the Route Map
// with the specified index value.
// Returns the requested Route Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - route-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterRouteMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/route-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateRouterSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the router - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterSetting(mkey string) (err error) {

	//No unset API for router - setting
	return
}

// ReadRouterSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateRouterStatic API operation for FortiSwitch creates a new Static.
// Returns the index value of the Static and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterStatic(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/static"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterStatic API operation for FortiSwitch updates the specified Static.
// Returns the index value of the Static and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterStatic(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/static"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterStatic API operation for FortiSwitch deletes the specified Static.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterStatic(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/static"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterStatic API operation for FortiSwitch gets the Static
// with the specified index value.
// Returns the requested Static value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterStatic(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/static"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterStatic6 API operation for FortiSwitch creates a new Static6.
// Returns the index value of the Static6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterStatic6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/static6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterStatic6 API operation for FortiSwitch updates the specified Static6.
// Returns the index value of the Static6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterStatic6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/static6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterStatic6 API operation for FortiSwitch deletes the specified Static6.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterStatic6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/static6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterStatic6 API operation for FortiSwitch gets the Static6
// with the specified index value.
// Returns the requested Static6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - static6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterStatic6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/static6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterVrf API operation for FortiSwitch creates a new Vrf.
// Returns the index value of the Vrf and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - vrf chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterVrf(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/vrf"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterVrf API operation for FortiSwitch updates the specified Vrf.
// Returns the index value of the Vrf and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - vrf chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterVrf(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/vrf"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterVrf API operation for FortiSwitch deletes the specified Vrf.
// Returns error for service API and SDK errors.
// See the router - vrf chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterVrf(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/vrf"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterVrf API operation for FortiSwitch gets the Vrf
// with the specified index value.
// Returns the requested Vrf value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router - vrf chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterVrf(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/vrf"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchAutoIslPortGroup API operation for FortiSwitch creates a new Auto Isl Port Group.
// Returns the index value of the Auto Isl Port Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - auto-isl-port-group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchAutoIslPortGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/auto-isl-port-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchAutoIslPortGroup API operation for FortiSwitch updates the specified Auto Isl Port Group.
// Returns the index value of the Auto Isl Port Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - auto-isl-port-group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAutoIslPortGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/auto-isl-port-group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAutoIslPortGroup API operation for FortiSwitch deletes the specified Auto Isl Port Group.
// Returns error for service API and SDK errors.
// See the switch - auto-isl-port-group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAutoIslPortGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/auto-isl-port-group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchAutoIslPortGroup API operation for FortiSwitch gets the Auto Isl Port Group
// with the specified index value.
// Returns the requested Auto Isl Port Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - auto-isl-port-group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAutoIslPortGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/auto-isl-port-group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchAutoNetwork API operation for FortiSwitch updates the specified Auto Network.
// Returns the index value of the Auto Network and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - auto-network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAutoNetwork(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/auto-network"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAutoNetwork API operation for FortiSwitch deletes the specified Auto Network.
// Returns error for service API and SDK errors.
// See the switch - auto-network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAutoNetwork(mkey string) (err error) {

	//No unset API for switch - auto-network
	return
}

// ReadSwitchAutoNetwork API operation for FortiSwitch gets the Auto Network
// with the specified index value.
// Returns the requested Auto Network value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - auto-network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAutoNetwork(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/auto-network"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchDomain API operation for FortiSwitch creates a new Domain.
// Returns the index value of the Domain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - domain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchDomain(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/domain"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchDomain API operation for FortiSwitch updates the specified Domain.
// Returns the index value of the Domain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - domain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchDomain(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/domain"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchDomain API operation for FortiSwitch deletes the specified Domain.
// Returns error for service API and SDK errors.
// See the switch - domain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchDomain(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/domain"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchDomain API operation for FortiSwitch gets the Domain
// with the specified index value.
// Returns the requested Domain value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - domain chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchDomain(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/domain"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchGlobal API operation for FortiSwitch updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchGlobal API operation for FortiSwitch deletes the specified Global.
// Returns error for service API and SDK errors.
// See the switch - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchGlobal(mkey string) (err error) {

	//No unset API for switch - global
	return
}

// ReadSwitchGlobal API operation for FortiSwitch gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchInterface API operation for FortiSwitch creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchInterface API operation for FortiSwitch updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchInterface API operation for FortiSwitch deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the switch - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchInterface API operation for FortiSwitch gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchIpMacBinding API operation for FortiSwitch creates a new Ip Mac Binding.
// Returns the index value of the Ip Mac Binding and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - ip-mac-binding chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchIpMacBinding(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/ip-mac-binding"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchIpMacBinding API operation for FortiSwitch updates the specified Ip Mac Binding.
// Returns the index value of the Ip Mac Binding and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - ip-mac-binding chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchIpMacBinding(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/ip-mac-binding"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchIpMacBinding API operation for FortiSwitch deletes the specified Ip Mac Binding.
// Returns error for service API and SDK errors.
// See the switch - ip-mac-binding chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchIpMacBinding(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/ip-mac-binding"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchIpMacBinding API operation for FortiSwitch gets the Ip Mac Binding
// with the specified index value.
// Returns the requested Ip Mac Binding value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - ip-mac-binding chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchIpMacBinding(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/ip-mac-binding"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchMirror API operation for FortiSwitch creates a new Mirror.
// Returns the index value of the Mirror and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - mirror chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchMirror(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/mirror"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchMirror API operation for FortiSwitch updates the specified Mirror.
// Returns the index value of the Mirror and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - mirror chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchMirror(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/mirror"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchMirror API operation for FortiSwitch deletes the specified Mirror.
// Returns error for service API and SDK errors.
// See the switch - mirror chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchMirror(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/mirror"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchMirror API operation for FortiSwitch gets the Mirror
// with the specified index value.
// Returns the requested Mirror value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - mirror chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchMirror(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/mirror"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchPhyMode API operation for FortiSwitch updates the specified Phy Mode.
// Returns the index value of the Phy Mode and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - phy-mode chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchPhyMode(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/phy-mode"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchPhyMode API operation for FortiSwitch deletes the specified Phy Mode.
// Returns error for service API and SDK errors.
// See the switch - phy-mode chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchPhyMode(mkey string) (err error) {

	//No unset API for switch - phy-mode
	return
}

// ReadSwitchPhyMode API operation for FortiSwitch gets the Phy Mode
// with the specified index value.
// Returns the requested Phy Mode value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - phy-mode chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchPhyMode(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/phy-mode"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchPhysicalPort API operation for FortiSwitch creates a new Physical Port.
// Returns the index value of the Physical Port and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - physical-port chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchPhysicalPort(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/physical-port"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchPhysicalPort API operation for FortiSwitch updates the specified Physical Port.
// Returns the index value of the Physical Port and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - physical-port chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchPhysicalPort(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/physical-port"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchPhysicalPort API operation for FortiSwitch deletes the specified Physical Port.
// Returns error for service API and SDK errors.
// See the switch - physical-port chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchPhysicalPort(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/physical-port"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchPhysicalPort API operation for FortiSwitch gets the Physical Port
// with the specified index value.
// Returns the requested Physical Port value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - physical-port chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchPhysicalPort(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/physical-port"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchQuarantine API operation for FortiSwitch creates a new Quarantine.
// Returns the index value of the Quarantine and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - quarantine chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchQuarantine(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/quarantine"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchQuarantine API operation for FortiSwitch updates the specified Quarantine.
// Returns the index value of the Quarantine and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - quarantine chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchQuarantine(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/quarantine"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchQuarantine API operation for FortiSwitch deletes the specified Quarantine.
// Returns error for service API and SDK errors.
// See the switch - quarantine chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchQuarantine(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/quarantine"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchQuarantine API operation for FortiSwitch gets the Quarantine
// with the specified index value.
// Returns the requested Quarantine value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - quarantine chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchQuarantine(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/quarantine"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchRaguardPolicy API operation for FortiSwitch creates a new Raguard Policy.
// Returns the index value of the Raguard Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - raguard-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchRaguardPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/raguard-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchRaguardPolicy API operation for FortiSwitch updates the specified Raguard Policy.
// Returns the index value of the Raguard Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - raguard-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchRaguardPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/raguard-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchRaguardPolicy API operation for FortiSwitch deletes the specified Raguard Policy.
// Returns error for service API and SDK errors.
// See the switch - raguard-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchRaguardPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/raguard-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchRaguardPolicy API operation for FortiSwitch gets the Raguard Policy
// with the specified index value.
// Returns the requested Raguard Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - raguard-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchRaguardPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/raguard-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchSecurityFeature API operation for FortiSwitch updates the specified Security Feature.
// Returns the index value of the Security Feature and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - security-feature chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchSecurityFeature(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/security-feature"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchSecurityFeature API operation for FortiSwitch deletes the specified Security Feature.
// Returns error for service API and SDK errors.
// See the switch - security-feature chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchSecurityFeature(mkey string) (err error) {

	//No unset API for switch - security-feature
	return
}

// ReadSwitchSecurityFeature API operation for FortiSwitch gets the Security Feature
// with the specified index value.
// Returns the requested Security Feature value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - security-feature chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchSecurityFeature(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/security-feature"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchStaticMac API operation for FortiSwitch creates a new Static Mac.
// Returns the index value of the Static Mac and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - static-mac chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchStaticMac(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/static-mac"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchStaticMac API operation for FortiSwitch updates the specified Static Mac.
// Returns the index value of the Static Mac and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - static-mac chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchStaticMac(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/static-mac"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchStaticMac API operation for FortiSwitch deletes the specified Static Mac.
// Returns error for service API and SDK errors.
// See the switch - static-mac chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchStaticMac(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/static-mac"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchStaticMac API operation for FortiSwitch gets the Static Mac
// with the specified index value.
// Returns the requested Static Mac value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - static-mac chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchStaticMac(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/static-mac"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchStormControl API operation for FortiSwitch updates the specified Storm Control.
// Returns the index value of the Storm Control and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - storm-control chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchStormControl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/storm-control"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchStormControl API operation for FortiSwitch deletes the specified Storm Control.
// Returns error for service API and SDK errors.
// See the switch - storm-control chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchStormControl(mkey string) (err error) {

	//No unset API for switch - storm-control
	return
}

// ReadSwitchStormControl API operation for FortiSwitch gets the Storm Control
// with the specified index value.
// Returns the requested Storm Control value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - storm-control chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchStormControl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/storm-control"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchTrunk API operation for FortiSwitch creates a new Trunk.
// Returns the index value of the Trunk and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - trunk chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchTrunk(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/trunk"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchTrunk API operation for FortiSwitch updates the specified Trunk.
// Returns the index value of the Trunk and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - trunk chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchTrunk(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/trunk"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchTrunk API operation for FortiSwitch deletes the specified Trunk.
// Returns error for service API and SDK errors.
// See the switch - trunk chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchTrunk(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/trunk"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchTrunk API operation for FortiSwitch gets the Trunk
// with the specified index value.
// Returns the requested Trunk value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - trunk chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchTrunk(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/trunk"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchVirtualWire API operation for FortiSwitch creates a new Virtual Wire.
// Returns the index value of the Virtual Wire and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - virtual-wire chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchVirtualWire(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/virtual-wire"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchVirtualWire API operation for FortiSwitch updates the specified Virtual Wire.
// Returns the index value of the Virtual Wire and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - virtual-wire chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchVirtualWire(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/virtual-wire"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchVirtualWire API operation for FortiSwitch deletes the specified Virtual Wire.
// Returns error for service API and SDK errors.
// See the switch - virtual-wire chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchVirtualWire(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/virtual-wire"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchVirtualWire API operation for FortiSwitch gets the Virtual Wire
// with the specified index value.
// Returns the requested Virtual Wire value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - virtual-wire chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchVirtualWire(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/virtual-wire"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchVlan API operation for FortiSwitch creates a new Vlan.
// Returns the index value of the Vlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - vlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchVlan(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/vlan"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchVlan API operation for FortiSwitch updates the specified Vlan.
// Returns the index value of the Vlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - vlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchVlan(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/vlan"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchVlan API operation for FortiSwitch deletes the specified Vlan.
// Returns error for service API and SDK errors.
// See the switch - vlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchVlan(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/vlan"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchVlan API operation for FortiSwitch gets the Vlan
// with the specified index value.
// Returns the requested Vlan value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - vlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchVlan(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/vlan"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchVlanTpid API operation for FortiSwitch creates a new Vlan Tpid.
// Returns the index value of the Vlan Tpid and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - vlan-tpid chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchVlanTpid(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch/vlan-tpid"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchVlanTpid API operation for FortiSwitch updates the specified Vlan Tpid.
// Returns the index value of the Vlan Tpid and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - vlan-tpid chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchVlanTpid(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch/vlan-tpid"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchVlanTpid API operation for FortiSwitch deletes the specified Vlan Tpid.
// Returns error for service API and SDK errors.
// See the switch - vlan-tpid chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchVlanTpid(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch/vlan-tpid"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchVlanTpid API operation for FortiSwitch gets the Vlan Tpid
// with the specified index value.
// Returns the requested Vlan Tpid value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch - vlan-tpid chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchVlanTpid(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch/vlan-tpid"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchControllerGlobal API operation for FortiSwitch updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchControllerGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch-controller/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchControllerGlobal API operation for FortiSwitch deletes the specified Global.
// Returns error for service API and SDK errors.
// See the switch-controller - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchControllerGlobal(mkey string) (err error) {

	//No unset API for switch-controller - global
	return
}

// ReadSwitchControllerGlobal API operation for FortiSwitch gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch-controller - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchControllerGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch-controller/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchAcl8021X API operation for FortiSwitch creates a new 802 1X.
// Returns the index value of the 802 1X and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - 802-1X chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchAcl8021X(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.acl/802-1X"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchAcl8021X API operation for FortiSwitch updates the specified 802 1X.
// Returns the index value of the 802 1X and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - 802-1X chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAcl8021X(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.acl/802-1X"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAcl8021X API operation for FortiSwitch deletes the specified 802 1X.
// Returns error for service API and SDK errors.
// See the switch.acl - 802-1X chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAcl8021X(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.acl/802-1X"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchAcl8021X API operation for FortiSwitch gets the 802 1X
// with the specified index value.
// Returns the requested 802 1X value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - 802-1X chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAcl8021X(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.acl/802-1X"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchAclEgress API operation for FortiSwitch creates a new Egress.
// Returns the index value of the Egress and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - egress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchAclEgress(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.acl/egress"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchAclEgress API operation for FortiSwitch updates the specified Egress.
// Returns the index value of the Egress and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - egress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAclEgress(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.acl/egress"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAclEgress API operation for FortiSwitch deletes the specified Egress.
// Returns error for service API and SDK errors.
// See the switch.acl - egress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAclEgress(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.acl/egress"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchAclEgress API operation for FortiSwitch gets the Egress
// with the specified index value.
// Returns the requested Egress value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - egress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAclEgress(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.acl/egress"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchAclIngress API operation for FortiSwitch creates a new Ingress.
// Returns the index value of the Ingress and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - ingress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchAclIngress(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.acl/ingress"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchAclIngress API operation for FortiSwitch updates the specified Ingress.
// Returns the index value of the Ingress and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - ingress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAclIngress(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.acl/ingress"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAclIngress API operation for FortiSwitch deletes the specified Ingress.
// Returns error for service API and SDK errors.
// See the switch.acl - ingress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAclIngress(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.acl/ingress"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchAclIngress API operation for FortiSwitch gets the Ingress
// with the specified index value.
// Returns the requested Ingress value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - ingress chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAclIngress(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.acl/ingress"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchAclPolicer API operation for FortiSwitch creates a new Policer.
// Returns the index value of the Policer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - policer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchAclPolicer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.acl/policer"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchAclPolicer API operation for FortiSwitch updates the specified Policer.
// Returns the index value of the Policer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - policer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAclPolicer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.acl/policer"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAclPolicer API operation for FortiSwitch deletes the specified Policer.
// Returns error for service API and SDK errors.
// See the switch.acl - policer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAclPolicer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.acl/policer"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchAclPolicer API operation for FortiSwitch gets the Policer
// with the specified index value.
// Returns the requested Policer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - policer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAclPolicer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.acl/policer"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchAclPrelookup API operation for FortiSwitch creates a new Prelookup.
// Returns the index value of the Prelookup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - prelookup chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchAclPrelookup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.acl/prelookup"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchAclPrelookup API operation for FortiSwitch updates the specified Prelookup.
// Returns the index value of the Prelookup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - prelookup chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAclPrelookup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.acl/prelookup"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAclPrelookup API operation for FortiSwitch deletes the specified Prelookup.
// Returns error for service API and SDK errors.
// See the switch.acl - prelookup chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAclPrelookup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.acl/prelookup"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchAclPrelookup API operation for FortiSwitch gets the Prelookup
// with the specified index value.
// Returns the requested Prelookup value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - prelookup chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAclPrelookup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.acl/prelookup"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchAclSettings API operation for FortiSwitch updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAclSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.acl/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAclSettings API operation for FortiSwitch deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the switch.acl - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAclSettings(mkey string) (err error) {

	//No unset API for switch.acl - settings
	return
}

// ReadSwitchAclSettings API operation for FortiSwitch gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAclSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.acl/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchAclServiceCustom API operation for FortiSwitch creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl.service - custom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchAclServiceCustom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.acl.service/custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchAclServiceCustom API operation for FortiSwitch updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl.service - custom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchAclServiceCustom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.acl.service/custom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchAclServiceCustom API operation for FortiSwitch deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the switch.acl.service - custom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchAclServiceCustom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.acl.service/custom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchAclServiceCustom API operation for FortiSwitch gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.acl.service - custom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchAclServiceCustom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.acl.service/custom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchIgmpSnoopingGlobals API operation for FortiSwitch updates the specified Globals.
// Returns the index value of the Globals and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.igmp-snooping - globals chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchIgmpSnoopingGlobals(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.igmp-snooping/globals"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchIgmpSnoopingGlobals API operation for FortiSwitch deletes the specified Globals.
// Returns error for service API and SDK errors.
// See the switch.igmp-snooping - globals chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchIgmpSnoopingGlobals(mkey string) (err error) {

	//No unset API for switch.igmp-snooping - globals
	return
}

// ReadSwitchIgmpSnoopingGlobals API operation for FortiSwitch gets the Globals
// with the specified index value.
// Returns the requested Globals value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.igmp-snooping - globals chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchIgmpSnoopingGlobals(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.igmp-snooping/globals"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchLldpProfile API operation for FortiSwitch creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.lldp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchLldpProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.lldp/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchLldpProfile API operation for FortiSwitch updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.lldp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchLldpProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.lldp/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchLldpProfile API operation for FortiSwitch deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the switch.lldp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchLldpProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.lldp/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchLldpProfile API operation for FortiSwitch gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.lldp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchLldpProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.lldp/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchLldpSettings API operation for FortiSwitch updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.lldp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchLldpSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.lldp/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchLldpSettings API operation for FortiSwitch deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the switch.lldp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchLldpSettings(mkey string) (err error) {

	//No unset API for switch.lldp - settings
	return
}

// ReadSwitchLldpSettings API operation for FortiSwitch gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.lldp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchLldpSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.lldp/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchMacsecProfile API operation for FortiSwitch creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.macsec - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchMacsecProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.macsec/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchMacsecProfile API operation for FortiSwitch updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.macsec - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchMacsecProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.macsec/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchMacsecProfile API operation for FortiSwitch deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the switch.macsec - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchMacsecProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.macsec/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchMacsecProfile API operation for FortiSwitch gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.macsec - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchMacsecProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.macsec/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchMldSnoopingGlobals API operation for FortiSwitch updates the specified Globals.
// Returns the index value of the Globals and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.mld-snooping - globals chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchMldSnoopingGlobals(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.mld-snooping/globals"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchMldSnoopingGlobals API operation for FortiSwitch deletes the specified Globals.
// Returns error for service API and SDK errors.
// See the switch.mld-snooping - globals chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchMldSnoopingGlobals(mkey string) (err error) {

	//No unset API for switch.mld-snooping - globals
	return
}

// ReadSwitchMldSnoopingGlobals API operation for FortiSwitch gets the Globals
// with the specified index value.
// Returns the requested Globals value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.mld-snooping - globals chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchMldSnoopingGlobals(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.mld-snooping/globals"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchNetworkMonitorDirected API operation for FortiSwitch creates a new Directed.
// Returns the index value of the Directed and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.network-monitor - directed chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchNetworkMonitorDirected(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.network-monitor/directed"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchNetworkMonitorDirected API operation for FortiSwitch updates the specified Directed.
// Returns the index value of the Directed and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.network-monitor - directed chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchNetworkMonitorDirected(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.network-monitor/directed"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchNetworkMonitorDirected API operation for FortiSwitch deletes the specified Directed.
// Returns error for service API and SDK errors.
// See the switch.network-monitor - directed chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchNetworkMonitorDirected(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.network-monitor/directed"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchNetworkMonitorDirected API operation for FortiSwitch gets the Directed
// with the specified index value.
// Returns the requested Directed value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.network-monitor - directed chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchNetworkMonitorDirected(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.network-monitor/directed"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchNetworkMonitorSettings API operation for FortiSwitch updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.network-monitor - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchNetworkMonitorSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.network-monitor/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchNetworkMonitorSettings API operation for FortiSwitch deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the switch.network-monitor - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchNetworkMonitorSettings(mkey string) (err error) {

	//No unset API for switch.network-monitor - settings
	return
}

// ReadSwitchNetworkMonitorSettings API operation for FortiSwitch gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.network-monitor - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchNetworkMonitorSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.network-monitor/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchPtpPolicy API operation for FortiSwitch creates a new Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.ptp - policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchPtpPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.ptp/policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchPtpPolicy API operation for FortiSwitch updates the specified Policy.
// Returns the index value of the Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.ptp - policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchPtpPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.ptp/policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchPtpPolicy API operation for FortiSwitch deletes the specified Policy.
// Returns error for service API and SDK errors.
// See the switch.ptp - policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchPtpPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.ptp/policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchPtpPolicy API operation for FortiSwitch gets the Policy
// with the specified index value.
// Returns the requested Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.ptp - policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchPtpPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.ptp/policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchPtpSettings API operation for FortiSwitch updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.ptp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchPtpSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.ptp/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchPtpSettings API operation for FortiSwitch deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the switch.ptp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchPtpSettings(mkey string) (err error) {

	//No unset API for switch.ptp - settings
	return
}

// ReadSwitchPtpSettings API operation for FortiSwitch gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.ptp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchPtpSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.ptp/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSwitchQosDot1PMap API operation for FortiSwitch creates a new Dot1P Map.
// Returns the index value of the Dot1P Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - dot1p-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchQosDot1PMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.qos/dot1p-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchQosDot1PMap API operation for FortiSwitch updates the specified Dot1P Map.
// Returns the index value of the Dot1P Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - dot1p-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchQosDot1PMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.qos/dot1p-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchQosDot1PMap API operation for FortiSwitch deletes the specified Dot1P Map.
// Returns error for service API and SDK errors.
// See the switch.qos - dot1p-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchQosDot1PMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.qos/dot1p-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchQosDot1PMap API operation for FortiSwitch gets the Dot1P Map
// with the specified index value.
// Returns the requested Dot1P Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - dot1p-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchQosDot1PMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.qos/dot1p-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchQosIpDscpMap API operation for FortiSwitch creates a new Ip Dscp Map.
// Returns the index value of the Ip Dscp Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - ip-dscp-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchQosIpDscpMap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.qos/ip-dscp-map"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchQosIpDscpMap API operation for FortiSwitch updates the specified Ip Dscp Map.
// Returns the index value of the Ip Dscp Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - ip-dscp-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchQosIpDscpMap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.qos/ip-dscp-map"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchQosIpDscpMap API operation for FortiSwitch deletes the specified Ip Dscp Map.
// Returns error for service API and SDK errors.
// See the switch.qos - ip-dscp-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchQosIpDscpMap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.qos/ip-dscp-map"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchQosIpDscpMap API operation for FortiSwitch gets the Ip Dscp Map
// with the specified index value.
// Returns the requested Ip Dscp Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - ip-dscp-map chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchQosIpDscpMap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.qos/ip-dscp-map"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchQosQosPolicy API operation for FortiSwitch creates a new Qos Policy.
// Returns the index value of the Qos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - qos-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchQosQosPolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.qos/qos-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchQosQosPolicy API operation for FortiSwitch updates the specified Qos Policy.
// Returns the index value of the Qos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - qos-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchQosQosPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.qos/qos-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchQosQosPolicy API operation for FortiSwitch deletes the specified Qos Policy.
// Returns error for service API and SDK errors.
// See the switch.qos - qos-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchQosQosPolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.qos/qos-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchQosQosPolicy API operation for FortiSwitch gets the Qos Policy
// with the specified index value.
// Returns the requested Qos Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.qos - qos-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchQosQosPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.qos/qos-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSwitchStpInstance API operation for FortiSwitch creates a new Instance.
// Returns the index value of the Instance and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.stp - instance chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSwitchStpInstance(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/switch.stp/instance"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSwitchStpInstance API operation for FortiSwitch updates the specified Instance.
// Returns the index value of the Instance and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.stp - instance chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchStpInstance(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.stp/instance"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchStpInstance API operation for FortiSwitch deletes the specified Instance.
// Returns error for service API and SDK errors.
// See the switch.stp - instance chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchStpInstance(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/switch.stp/instance"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSwitchStpInstance API operation for FortiSwitch gets the Instance
// with the specified index value.
// Returns the requested Instance value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.stp - instance chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchStpInstance(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.stp/instance"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSwitchStpSettings API operation for FortiSwitch updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.stp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSwitchStpSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/switch.stp/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSwitchStpSettings API operation for FortiSwitch deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the switch.stp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSwitchStpSettings(mkey string) (err error) {

	//No unset API for switch.stp - settings
	return
}

// ReadSwitchStpSettings API operation for FortiSwitch gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the switch.stp - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSwitchStpSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/switch.stp/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemAccprofile API operation for FortiSwitch creates a new Accprofile.
// Returns the index value of the Accprofile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAccprofile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/accprofile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAccprofile API operation for FortiSwitch updates the specified Accprofile.
// Returns the index value of the Accprofile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAccprofile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAccprofile API operation for FortiSwitch deletes the specified Accprofile.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAccprofile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAccprofile API operation for FortiSwitch gets the Accprofile
// with the specified index value.
// Returns the requested Accprofile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - accprofile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAccprofile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/accprofile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAdmin API operation for FortiSwitch creates a new Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdmin(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAdmin API operation for FortiSwitch updates the specified Admin.
// Returns the index value of the Admin and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdmin(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAdmin API operation for FortiSwitch deletes the specified Admin.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdmin(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAdmin API operation for FortiSwitch gets the Admin
// with the specified index value.
// Returns the requested Admin value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdmin(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/admin"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemAlarm API operation for FortiSwitch updates the specified Alarm.
// Returns the index value of the Alarm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alarm chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlarm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/alarm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAlarm API operation for FortiSwitch deletes the specified Alarm.
// Returns error for service API and SDK errors.
// See the system - alarm chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlarm(mkey string) (err error) {

	//No unset API for system - alarm
	return
}

// ReadSystemAlarm API operation for FortiSwitch gets the Alarm
// with the specified index value.
// Returns the requested Alarm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alarm chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlarm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/alarm"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemAlertemail API operation for FortiSwitch updates the specified Alertemail.
// Returns the index value of the Alertemail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlertemail(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/alertemail"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAlertemail API operation for FortiSwitch deletes the specified Alertemail.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlertemail(mkey string) (err error) {

	//No unset API for system - alertemail
	return
}

// ReadSystemAlertemail API operation for FortiSwitch gets the Alertemail
// with the specified index value.
// Returns the requested Alertemail value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlertemail(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/alertemail"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemArpTable API operation for FortiSwitch creates a new Arp Table.
// Returns the index value of the Arp Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemArpTable(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/arp-table"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemArpTable API operation for FortiSwitch updates the specified Arp Table.
// Returns the index value of the Arp Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemArpTable(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemArpTable API operation for FortiSwitch deletes the specified Arp Table.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemArpTable(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemArpTable API operation for FortiSwitch gets the Arp Table
// with the specified index value.
// Returns the requested Arp Table value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - arp-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemArpTable(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/arp-table"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutoScript API operation for FortiSwitch creates a new Auto Script.
// Returns the index value of the Auto Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutoScript(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/auto-script"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutoScript API operation for FortiSwitch updates the specified Auto Script.
// Returns the index value of the Auto Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoScript(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoScript API operation for FortiSwitch deletes the specified Auto Script.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoScript(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutoScript API operation for FortiSwitch gets the Auto Script
// with the specified index value.
// Returns the requested Auto Script value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-script chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoScript(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/auto-script"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutomationAction API operation for FortiSwitch creates a new Automation Action.
// Returns the index value of the Automation Action and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationAction(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-action"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationAction API operation for FortiSwitch updates the specified Automation Action.
// Returns the index value of the Automation Action and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationAction(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-action"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationAction API operation for FortiSwitch deletes the specified Automation Action.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationAction(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-action"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationAction API operation for FortiSwitch gets the Automation Action
// with the specified index value.
// Returns the requested Automation Action value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-action chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationAction(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-action"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutomationDestination API operation for FortiSwitch creates a new Automation Destination.
// Returns the index value of the Automation Destination and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationDestination(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-destination"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationDestination API operation for FortiSwitch updates the specified Automation Destination.
// Returns the index value of the Automation Destination and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationDestination(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-destination"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationDestination API operation for FortiSwitch deletes the specified Automation Destination.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationDestination(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-destination"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationDestination API operation for FortiSwitch gets the Automation Destination
// with the specified index value.
// Returns the requested Automation Destination value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-destination chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationDestination(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-destination"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutomationStitch API operation for FortiSwitch creates a new Automation Stitch.
// Returns the index value of the Automation Stitch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationStitch(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-stitch"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationStitch API operation for FortiSwitch updates the specified Automation Stitch.
// Returns the index value of the Automation Stitch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationStitch(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-stitch"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationStitch API operation for FortiSwitch deletes the specified Automation Stitch.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationStitch(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-stitch"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationStitch API operation for FortiSwitch gets the Automation Stitch
// with the specified index value.
// Returns the requested Automation Stitch value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-stitch chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationStitch(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-stitch"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAutomationTrigger API operation for FortiSwitch creates a new Automation Trigger.
// Returns the index value of the Automation Trigger and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAutomationTrigger(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/automation-trigger"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAutomationTrigger API operation for FortiSwitch updates the specified Automation Trigger.
// Returns the index value of the Automation Trigger and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutomationTrigger(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/automation-trigger"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutomationTrigger API operation for FortiSwitch deletes the specified Automation Trigger.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutomationTrigger(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/automation-trigger"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAutomationTrigger API operation for FortiSwitch gets the Automation Trigger
// with the specified index value.
// Returns the requested Automation Trigger value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - automation-trigger chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutomationTrigger(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/automation-trigger"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemBugReport API operation for FortiSwitch updates the specified Bug Report.
// Returns the index value of the Bug Report and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - bug-report chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemBugReport(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/bug-report"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemBugReport API operation for FortiSwitch deletes the specified Bug Report.
// Returns error for service API and SDK errors.
// See the system - bug-report chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemBugReport(mkey string) (err error) {

	//No unset API for system - bug-report
	return
}

// ReadSystemBugReport API operation for FortiSwitch gets the Bug Report
// with the specified index value.
// Returns the requested Bug Report value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - bug-report chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemBugReport(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/bug-report"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemCentralManagement API operation for FortiSwitch updates the specified Central Management.
// Returns the index value of the Central Management and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCentralManagement(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/central-management"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCentralManagement API operation for FortiSwitch deletes the specified Central Management.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCentralManagement(mkey string) (err error) {

	//No unset API for system - central-management
	return
}

// ReadSystemCentralManagement API operation for FortiSwitch gets the Central Management
// with the specified index value.
// Returns the requested Central Management value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - central-management chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCentralManagement(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/central-management"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemConsole API operation for FortiSwitch updates the specified Console.
// Returns the index value of the Console and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - console chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemConsole(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/console"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemConsole API operation for FortiSwitch deletes the specified Console.
// Returns error for service API and SDK errors.
// See the system - console chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemConsole(mkey string) (err error) {

	//No unset API for system - console
	return
}

// ReadSystemConsole API operation for FortiSwitch gets the Console
// with the specified index value.
// Returns the requested Console value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - console chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemConsole(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/console"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemDns API operation for FortiSwitch updates the specified Dns.
// Returns the index value of the Dns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDns(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDns API operation for FortiSwitch deletes the specified Dns.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDns(mkey string) (err error) {

	//No unset API for system - dns
	return
}

// ReadSystemDns API operation for FortiSwitch gets the Dns
// with the specified index value.
// Returns the requested Dns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDns(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemDnsDatabase API operation for FortiSwitch creates a new Dns Database.
// Returns the index value of the Dns Database and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDnsDatabase(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/dns-database"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDnsDatabase API operation for FortiSwitch updates the specified Dns Database.
// Returns the index value of the Dns Database and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDnsDatabase(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDnsDatabase API operation for FortiSwitch deletes the specified Dns Database.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDnsDatabase(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDnsDatabase API operation for FortiSwitch gets the Dns Database
// with the specified index value.
// Returns the requested Dns Database value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-database chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDnsDatabase(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns-database"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDnsServer API operation for FortiSwitch creates a new Dns Server.
// Returns the index value of the Dns Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDnsServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/dns-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDnsServer API operation for FortiSwitch updates the specified Dns Server.
// Returns the index value of the Dns Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDnsServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDnsServer API operation for FortiSwitch deletes the specified Dns Server.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDnsServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDnsServer API operation for FortiSwitch gets the Dns Server
// with the specified index value.
// Returns the requested Dns Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns-server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDnsServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns-server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemEmailServer API operation for FortiSwitch updates the specified Email Server.
// Returns the index value of the Email Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemEmailServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/email-server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemEmailServer API operation for FortiSwitch deletes the specified Email Server.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemEmailServer(mkey string) (err error) {

	//No unset API for system - email-server
	return
}

// ReadSystemEmailServer API operation for FortiSwitch gets the Email Server
// with the specified index value.
// Returns the requested Email Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - email-server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemEmailServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/email-server"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFlanCloud API operation for FortiSwitch updates the specified Flan Cloud.
// Returns the index value of the Flan Cloud and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - flan-cloud chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFlanCloud(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/flan-cloud"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFlanCloud API operation for FortiSwitch deletes the specified Flan Cloud.
// Returns error for service API and SDK errors.
// See the system - flan-cloud chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFlanCloud(mkey string) (err error) {

	//No unset API for system - flan-cloud
	return
}

// ReadSystemFlanCloud API operation for FortiSwitch gets the Flan Cloud
// with the specified index value.
// Returns the requested Flan Cloud value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - flan-cloud chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFlanCloud(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/flan-cloud"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFlowExport API operation for FortiSwitch updates the specified Flow Export.
// Returns the index value of the Flow Export and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - flow-export chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFlowExport(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/flow-export"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFlowExport API operation for FortiSwitch deletes the specified Flow Export.
// Returns error for service API and SDK errors.
// See the system - flow-export chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFlowExport(mkey string) (err error) {

	//No unset API for system - flow-export
	return
}

// ReadSystemFlowExport API operation for FortiSwitch gets the Flow Export
// with the specified index value.
// Returns the requested Flow Export value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - flow-export chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFlowExport(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/flow-export"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFm API operation for FortiSwitch updates the specified Fm.
// Returns the index value of the Fm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFm(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fm"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFm API operation for FortiSwitch deletes the specified Fm.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFm(mkey string) (err error) {

	//No unset API for system - fm
	return
}

// ReadSystemFm API operation for FortiSwitch gets the Fm
// with the specified index value.
// Returns the requested Fm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fm chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFm(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fm"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFortianalyzer API operation for FortiSwitch updates the specified Fortianalyzer.
// Returns the index value of the Fortianalyzer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortianalyzer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortianalyzer"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortianalyzer API operation for FortiSwitch deletes the specified Fortianalyzer.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortianalyzer(mkey string) (err error) {

	//No unset API for system - fortianalyzer
	return
}

// ReadSystemFortianalyzer API operation for FortiSwitch gets the Fortianalyzer
// with the specified index value.
// Returns the requested Fortianalyzer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortianalyzer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortianalyzer"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFortianalyzer2 API operation for FortiSwitch updates the specified Fortianalyzer2.
// Returns the index value of the Fortianalyzer2 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer2 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortianalyzer2(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortianalyzer2"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortianalyzer2 API operation for FortiSwitch deletes the specified Fortianalyzer2.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer2 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortianalyzer2(mkey string) (err error) {

	//No unset API for system - fortianalyzer2
	return
}

// ReadSystemFortianalyzer2 API operation for FortiSwitch gets the Fortianalyzer2
// with the specified index value.
// Returns the requested Fortianalyzer2 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer2 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortianalyzer2(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortianalyzer2"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFortianalyzer3 API operation for FortiSwitch updates the specified Fortianalyzer3.
// Returns the index value of the Fortianalyzer3 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer3 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortianalyzer3(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortianalyzer3"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortianalyzer3 API operation for FortiSwitch deletes the specified Fortianalyzer3.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer3 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortianalyzer3(mkey string) (err error) {

	//No unset API for system - fortianalyzer3
	return
}

// ReadSystemFortianalyzer3 API operation for FortiSwitch gets the Fortianalyzer3
// with the specified index value.
// Returns the requested Fortianalyzer3 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortianalyzer3 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortianalyzer3(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortianalyzer3"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFortiguard API operation for FortiSwitch updates the specified Fortiguard.
// Returns the index value of the Fortiguard and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortiguard(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortiguard"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortiguard API operation for FortiSwitch deletes the specified Fortiguard.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortiguard(mkey string) (err error) {

	//No unset API for system - fortiguard
	return
}

// ReadSystemFortiguard API operation for FortiSwitch gets the Fortiguard
// with the specified index value.
// Returns the requested Fortiguard value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiguard chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortiguard(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortiguard"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFortimanager API operation for FortiSwitch updates the specified Fortimanager.
// Returns the index value of the Fortimanager and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortimanager(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fortimanager"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFortimanager API operation for FortiSwitch deletes the specified Fortimanager.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortimanager(mkey string) (err error) {

	//No unset API for system - fortimanager
	return
}

// ReadSystemFortimanager API operation for FortiSwitch gets the Fortimanager
// with the specified index value.
// Returns the requested Fortimanager value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortimanager chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortimanager(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fortimanager"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemFswCloud API operation for FortiSwitch updates the specified Fsw Cloud.
// Returns the index value of the Fsw Cloud and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fsw-cloud chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFswCloud(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/fsw-cloud"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemFswCloud API operation for FortiSwitch deletes the specified Fsw Cloud.
// Returns error for service API and SDK errors.
// See the system - fsw-cloud chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFswCloud(mkey string) (err error) {

	//No unset API for system - fsw-cloud
	return
}

// ReadSystemFswCloud API operation for FortiSwitch gets the Fsw Cloud
// with the specified index value.
// Returns the requested Fsw Cloud value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fsw-cloud chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFswCloud(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/fsw-cloud"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemGlobal API operation for FortiSwitch updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGlobal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/global"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemGlobal API operation for FortiSwitch deletes the specified Global.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGlobal(mkey string) (err error) {

	//No unset API for system - global
	return
}

// ReadSystemGlobal API operation for FortiSwitch gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGlobal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/global"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemInterface API operation for FortiSwitch creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemInterface API operation for FortiSwitch updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemInterface API operation for FortiSwitch deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemInterface API operation for FortiSwitch gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemIpv6NeighborCache API operation for FortiSwitch creates a new Ipv6 Neighbor Cache.
// Returns the index value of the Ipv6 Neighbor Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemIpv6NeighborCache(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemIpv6NeighborCache API operation for FortiSwitch updates the specified Ipv6 Neighbor Cache.
// Returns the index value of the Ipv6 Neighbor Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemIpv6NeighborCache(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemIpv6NeighborCache API operation for FortiSwitch deletes the specified Ipv6 Neighbor Cache.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemIpv6NeighborCache(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemIpv6NeighborCache API operation for FortiSwitch gets the Ipv6 Neighbor Cache
// with the specified index value.
// Returns the requested Ipv6 Neighbor Cache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ipv6-neighbor-cache chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemIpv6NeighborCache(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ipv6-neighbor-cache"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemLinkMonitor API operation for FortiSwitch creates a new Link Monitor.
// Returns the index value of the Link Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLinkMonitor(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/link-monitor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemLinkMonitor API operation for FortiSwitch updates the specified Link Monitor.
// Returns the index value of the Link Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLinkMonitor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/link-monitor"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemLinkMonitor API operation for FortiSwitch deletes the specified Link Monitor.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLinkMonitor(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/link-monitor"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemLinkMonitor API operation for FortiSwitch gets the Link Monitor
// with the specified index value.
// Returns the requested Link Monitor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - link-monitor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLinkMonitor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/link-monitor"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemLocation API operation for FortiSwitch creates a new Location.
// Returns the index value of the Location and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - location chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLocation(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/location"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemLocation API operation for FortiSwitch updates the specified Location.
// Returns the index value of the Location and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - location chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocation(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/location"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemLocation API operation for FortiSwitch deletes the specified Location.
// Returns error for service API and SDK errors.
// See the system - location chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocation(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/location"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemLocation API operation for FortiSwitch gets the Location
// with the specified index value.
// Returns the requested Location value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - location chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocation(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/location"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemMacAddressTable API operation for FortiSwitch creates a new Mac Address Table.
// Returns the index value of the Mac Address Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMacAddressTable(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/mac-address-table"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemMacAddressTable API operation for FortiSwitch updates the specified Mac Address Table.
// Returns the index value of the Mac Address Table and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMacAddressTable(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/mac-address-table"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemMacAddressTable API operation for FortiSwitch deletes the specified Mac Address Table.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMacAddressTable(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/mac-address-table"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemMacAddressTable API operation for FortiSwitch gets the Mac Address Table
// with the specified index value.
// Returns the requested Mac Address Table value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mac-address-table chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMacAddressTable(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/mac-address-table"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemManagementTunnel API operation for FortiSwitch updates the specified Management Tunnel.
// Returns the index value of the Management Tunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemManagementTunnel(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/management-tunnel"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemManagementTunnel API operation for FortiSwitch deletes the specified Management Tunnel.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemManagementTunnel(mkey string) (err error) {

	//No unset API for system - management-tunnel
	return
}

// ReadSystemManagementTunnel API operation for FortiSwitch gets the Management Tunnel
// with the specified index value.
// Returns the requested Management Tunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - management-tunnel chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemManagementTunnel(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/management-tunnel"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemNtp API operation for FortiSwitch updates the specified Ntp.
// Returns the index value of the Ntp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNtp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/ntp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemNtp API operation for FortiSwitch deletes the specified Ntp.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNtp(mkey string) (err error) {

	//No unset API for system - ntp
	return
}

// ReadSystemNtp API operation for FortiSwitch gets the Ntp
// with the specified index value.
// Returns the requested Ntp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNtp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/ntp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemObjectTag API operation for FortiSwitch creates a new Object Tag.
// Returns the index value of the Object Tag and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tag chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemObjectTag(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/object-tag"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemObjectTag API operation for FortiSwitch updates the specified Object Tag.
// Returns the index value of the Object Tag and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tag chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemObjectTag(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/object-tag"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemObjectTag API operation for FortiSwitch deletes the specified Object Tag.
// Returns error for service API and SDK errors.
// See the system - object-tag chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemObjectTag(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/object-tag"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemObjectTag API operation for FortiSwitch gets the Object Tag
// with the specified index value.
// Returns the requested Object Tag value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - object-tag chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemObjectTag(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/object-tag"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemPasswordPolicy API operation for FortiSwitch updates the specified Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPasswordPolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/password-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPasswordPolicy API operation for FortiSwitch deletes the specified Password Policy.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPasswordPolicy(mkey string) (err error) {

	//No unset API for system - password-policy
	return
}

// ReadSystemPasswordPolicy API operation for FortiSwitch gets the Password Policy
// with the specified index value.
// Returns the requested Password Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPasswordPolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/password-policy"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemPortPair API operation for FortiSwitch creates a new Port Pair.
// Returns the index value of the Port Pair and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - port-pair chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemPortPair(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/port-pair"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemPortPair API operation for FortiSwitch updates the specified Port Pair.
// Returns the index value of the Port Pair and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - port-pair chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPortPair(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/port-pair"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPortPair API operation for FortiSwitch deletes the specified Port Pair.
// Returns error for service API and SDK errors.
// See the system - port-pair chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPortPair(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/port-pair"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemPortPair API operation for FortiSwitch gets the Port Pair
// with the specified index value.
// Returns the requested Port Pair value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - port-pair chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPortPair(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/port-pair"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemProxyArp API operation for FortiSwitch creates a new Proxy Arp.
// Returns the index value of the Proxy Arp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemProxyArp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/proxy-arp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemProxyArp API operation for FortiSwitch updates the specified Proxy Arp.
// Returns the index value of the Proxy Arp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemProxyArp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/proxy-arp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemProxyArp API operation for FortiSwitch deletes the specified Proxy Arp.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemProxyArp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/proxy-arp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemProxyArp API operation for FortiSwitch gets the Proxy Arp
// with the specified index value.
// Returns the requested Proxy Arp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - proxy-arp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemProxyArp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/proxy-arp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemResourceLimits API operation for FortiSwitch updates the specified Resource Limits.
// Returns the index value of the Resource Limits and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemResourceLimits(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/resource-limits"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemResourceLimits API operation for FortiSwitch deletes the specified Resource Limits.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemResourceLimits(mkey string) (err error) {

	//No unset API for system - resource-limits
	return
}

// ReadSystemResourceLimits API operation for FortiSwitch gets the Resource Limits
// with the specified index value.
// Returns the requested Resource Limits value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - resource-limits chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemResourceLimits(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/resource-limits"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemSessionTtl API operation for FortiSwitch updates the specified Session Ttl.
// Returns the index value of the Session Ttl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSessionTtl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/session-ttl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSessionTtl API operation for FortiSwitch deletes the specified Session Ttl.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSessionTtl(mkey string) (err error) {

	//No unset API for system - session-ttl
	return
}

// ReadSystemSessionTtl API operation for FortiSwitch gets the Session Ttl
// with the specified index value.
// Returns the requested Session Ttl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - session-ttl chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSessionTtl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/session-ttl"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemSettings API operation for FortiSwitch updates the specified Settings.
// Returns the index value of the Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSettings(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSettings API operation for FortiSwitch deletes the specified Settings.
// Returns error for service API and SDK errors.
// See the system - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSettings(mkey string) (err error) {

	//No unset API for system - settings
	return
}

// ReadSystemSettings API operation for FortiSwitch gets the Settings
// with the specified index value.
// Returns the requested Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - settings chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSettings(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/settings"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemSflow API operation for FortiSwitch updates the specified Sflow.
// Returns the index value of the Sflow and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSflow(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sflow"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSflow API operation for FortiSwitch deletes the specified Sflow.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSflow(mkey string) (err error) {

	//No unset API for system - sflow
	return
}

// ReadSystemSflow API operation for FortiSwitch gets the Sflow
// with the specified index value.
// Returns the requested Sflow value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sflow chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSflow(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sflow"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemSnifferProfile API operation for FortiSwitch creates a new Sniffer Profile.
// Returns the index value of the Sniffer Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer-profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnifferProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/sniffer-profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSnifferProfile API operation for FortiSwitch updates the specified Sniffer Profile.
// Returns the index value of the Sniffer Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer-profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnifferProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/sniffer-profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnifferProfile API operation for FortiSwitch deletes the specified Sniffer Profile.
// Returns error for service API and SDK errors.
// See the system - sniffer-profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnifferProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/sniffer-profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSnifferProfile API operation for FortiSwitch gets the Sniffer Profile
// with the specified index value.
// Returns the requested Sniffer Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer-profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnifferProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/sniffer-profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemTosBasedPriority API operation for FortiSwitch creates a new Tos Based Priority.
// Returns the index value of the Tos Based Priority and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemTosBasedPriority(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/tos-based-priority"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemTosBasedPriority API operation for FortiSwitch updates the specified Tos Based Priority.
// Returns the index value of the Tos Based Priority and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemTosBasedPriority(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/tos-based-priority"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemTosBasedPriority API operation for FortiSwitch deletes the specified Tos Based Priority.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemTosBasedPriority(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/tos-based-priority"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemTosBasedPriority API operation for FortiSwitch gets the Tos Based Priority
// with the specified index value.
// Returns the requested Tos Based Priority value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - tos-based-priority chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemTosBasedPriority(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/tos-based-priority"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemVdom API operation for FortiSwitch creates a new Vdom.
// Returns the index value of the Vdom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdom(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVdom API operation for FortiSwitch updates the specified Vdom.
// Returns the index value of the Vdom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdom(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdom API operation for FortiSwitch deletes the specified Vdom.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdom(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVdom API operation for FortiSwitch gets the Vdom
// with the specified index value.
// Returns the requested Vdom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdom(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemVdomDns API operation for FortiSwitch updates the specified Vdom Dns.
// Returns the index value of the Vdom Dns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-dns chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomDns(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-dns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomDns API operation for FortiSwitch deletes the specified Vdom Dns.
// Returns error for service API and SDK errors.
// See the system - vdom-dns chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomDns(mkey string) (err error) {

	//No unset API for system - vdom-dns
	return
}

// ReadSystemVdomDns API operation for FortiSwitch gets the Vdom Dns
// with the specified index value.
// Returns the requested Vdom Dns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-dns chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomDns(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-dns"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemVdomProperty API operation for FortiSwitch creates a new Vdom Property.
// Returns the index value of the Vdom Property and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdomProperty(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom-property"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVdomProperty API operation for FortiSwitch updates the specified Vdom Property.
// Returns the index value of the Vdom Property and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomProperty(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom-property"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVdomProperty API operation for FortiSwitch deletes the specified Vdom Property.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomProperty(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom-property"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVdomProperty API operation for FortiSwitch gets the Vdom Property
// with the specified index value.
// Returns the requested Vdom Property value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom-property chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomProperty(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom-property"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemVxlan API operation for FortiSwitch creates a new Vxlan.
// Returns the index value of the Vxlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVxlan(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vxlan"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemVxlan API operation for FortiSwitch updates the specified Vxlan.
// Returns the index value of the Vxlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVxlan(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemVxlan API operation for FortiSwitch deletes the specified Vxlan.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVxlan(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemVxlan API operation for FortiSwitch gets the Vxlan
// with the specified index value.
// Returns the requested Vxlan value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vxlan chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVxlan(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vxlan"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemWeb API operation for FortiSwitch updates the specified Web.
// Returns the index value of the Web and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - web chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemWeb(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/web"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemWeb API operation for FortiSwitch deletes the specified Web.
// Returns error for service API and SDK errors.
// See the system - web chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemWeb(mkey string) (err error) {

	//No unset API for system - web
	return
}

// ReadSystemWeb API operation for FortiSwitch gets the Web
// with the specified index value.
// Returns the requested Web value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - web chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemWeb(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/web"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemZone API operation for FortiSwitch creates a new Zone.
// Returns the index value of the Zone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemZone(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/zone"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemZone API operation for FortiSwitch updates the specified Zone.
// Returns the index value of the Zone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemZone(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemZone API operation for FortiSwitch deletes the specified Zone.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemZone(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemZone API operation for FortiSwitch gets the Zone
// with the specified index value.
// Returns the requested Zone value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - zone chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemZone(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/zone"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAliasCommand API operation for FortiSwitch creates a new Command.
// Returns the index value of the Command and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.alias - command chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAliasCommand(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.alias/command"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAliasCommand API operation for FortiSwitch updates the specified Command.
// Returns the index value of the Command and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.alias - command chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAliasCommand(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.alias/command"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAliasCommand API operation for FortiSwitch deletes the specified Command.
// Returns error for service API and SDK errors.
// See the system.alias - command chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAliasCommand(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.alias/command"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAliasCommand API operation for FortiSwitch gets the Command
// with the specified index value.
// Returns the requested Command value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.alias - command chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAliasCommand(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.alias/command"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemAliasGroup API operation for FortiSwitch creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.alias - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAliasGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.alias/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemAliasGroup API operation for FortiSwitch updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.alias - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAliasGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.alias/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAliasGroup API operation for FortiSwitch deletes the specified Group.
// Returns error for service API and SDK errors.
// See the system.alias - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAliasGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.alias/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemAliasGroup API operation for FortiSwitch gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.alias - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAliasGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.alias/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemAutoupdateClientoverride API operation for FortiSwitch updates the specified Clientoverride.
// Returns the index value of the Clientoverride and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - clientoverride chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateClientoverride(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/clientoverride"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateClientoverride API operation for FortiSwitch deletes the specified Clientoverride.
// Returns error for service API and SDK errors.
// See the system.autoupdate - clientoverride chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateClientoverride(mkey string) (err error) {

	//No unset API for system.autoupdate - clientoverride
	return
}

// ReadSystemAutoupdateClientoverride API operation for FortiSwitch gets the Clientoverride
// with the specified index value.
// Returns the requested Clientoverride value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - clientoverride chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateClientoverride(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/clientoverride"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemAutoupdateOverride API operation for FortiSwitch updates the specified Override.
// Returns the index value of the Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - override chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateOverride(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/override"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateOverride API operation for FortiSwitch deletes the specified Override.
// Returns error for service API and SDK errors.
// See the system.autoupdate - override chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateOverride(mkey string) (err error) {

	//No unset API for system.autoupdate - override
	return
}

// ReadSystemAutoupdateOverride API operation for FortiSwitch gets the Override
// with the specified index value.
// Returns the requested Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - override chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateOverride(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/override"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemAutoupdatePushUpdate API operation for FortiSwitch updates the specified Push Update.
// Returns the index value of the Push Update and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdatePushUpdate(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/push-update"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdatePushUpdate API operation for FortiSwitch deletes the specified Push Update.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdatePushUpdate(mkey string) (err error) {

	//No unset API for system.autoupdate - push-update
	return
}

// ReadSystemAutoupdatePushUpdate API operation for FortiSwitch gets the Push Update
// with the specified index value.
// Returns the requested Push Update value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - push-update chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdatePushUpdate(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/push-update"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemAutoupdateSchedule API operation for FortiSwitch updates the specified Schedule.
// Returns the index value of the Schedule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateSchedule(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/schedule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateSchedule API operation for FortiSwitch deletes the specified Schedule.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateSchedule(mkey string) (err error) {

	//No unset API for system.autoupdate - schedule
	return
}

// ReadSystemAutoupdateSchedule API operation for FortiSwitch gets the Schedule
// with the specified index value.
// Returns the requested Schedule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - schedule chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateSchedule(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/schedule"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// UpdateSystemAutoupdateTunneling API operation for FortiSwitch updates the specified Tunneling.
// Returns the index value of the Tunneling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoupdateTunneling(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.autoupdate/tunneling"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemAutoupdateTunneling API operation for FortiSwitch deletes the specified Tunneling.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoupdateTunneling(mkey string) (err error) {

	//No unset API for system.autoupdate - tunneling
	return
}

// ReadSystemAutoupdateTunneling API operation for FortiSwitch gets the Tunneling
// with the specified index value.
// Returns the requested Tunneling value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.autoupdate - tunneling chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoupdateTunneling(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.autoupdate/tunneling"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemCertificateCa API operation for FortiSwitch creates a new Ca.
// Returns the index value of the Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - ca chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateCa(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.certificate/ca"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemCertificateCa API operation for FortiSwitch updates the specified Ca.
// Returns the index value of the Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - ca chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateCa(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.certificate/ca"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCertificateCa API operation for FortiSwitch deletes the specified Ca.
// Returns error for service API and SDK errors.
// See the system.certificate - ca chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateCa(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.certificate/ca"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemCertificateCa API operation for FortiSwitch gets the Ca
// with the specified index value.
// Returns the requested Ca value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - ca chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateCa(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.certificate/ca"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemCertificateCrl API operation for FortiSwitch creates a new Crl.
// Returns the index value of the Crl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - crl chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateCrl(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.certificate/crl"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemCertificateCrl API operation for FortiSwitch updates the specified Crl.
// Returns the index value of the Crl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - crl chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateCrl(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.certificate/crl"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCertificateCrl API operation for FortiSwitch deletes the specified Crl.
// Returns error for service API and SDK errors.
// See the system.certificate - crl chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateCrl(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.certificate/crl"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemCertificateCrl API operation for FortiSwitch gets the Crl
// with the specified index value.
// Returns the requested Crl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - crl chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateCrl(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.certificate/crl"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemCertificateLocal API operation for FortiSwitch creates a new Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateLocal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.certificate/local"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemCertificateLocal API operation for FortiSwitch updates the specified Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateLocal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.certificate/local"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCertificateLocal API operation for FortiSwitch deletes the specified Local.
// Returns error for service API and SDK errors.
// See the system.certificate - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateLocal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.certificate/local"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemCertificateLocal API operation for FortiSwitch gets the Local
// with the specified index value.
// Returns the requested Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateLocal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.certificate/local"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemCertificateOcsp API operation for FortiSwitch updates the specified Ocsp.
// Returns the index value of the Ocsp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - ocsp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateOcsp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.certificate/ocsp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCertificateOcsp API operation for FortiSwitch deletes the specified Ocsp.
// Returns error for service API and SDK errors.
// See the system.certificate - ocsp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateOcsp(mkey string) (err error) {

	//No unset API for system.certificate - ocsp
	return
}

// ReadSystemCertificateOcsp API operation for FortiSwitch gets the Ocsp
// with the specified index value.
// Returns the requested Ocsp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - ocsp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateOcsp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.certificate/ocsp"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemCertificateRemote API operation for FortiSwitch creates a new Remote.
// Returns the index value of the Remote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - remote chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateRemote(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.certificate/remote"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemCertificateRemote API operation for FortiSwitch updates the specified Remote.
// Returns the index value of the Remote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - remote chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateRemote(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.certificate/remote"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemCertificateRemote API operation for FortiSwitch deletes the specified Remote.
// Returns error for service API and SDK errors.
// See the system.certificate - remote chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateRemote(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.certificate/remote"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemCertificateRemote API operation for FortiSwitch gets the Remote
// with the specified index value.
// Returns the requested Remote value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.certificate - remote chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateRemote(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.certificate/remote"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemDhcpServer API operation for FortiSwitch creates a new Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemDhcpServer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.dhcp/server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemDhcpServer API operation for FortiSwitch updates the specified Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDhcpServer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemDhcpServer API operation for FortiSwitch deletes the specified Server.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDhcpServer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemDhcpServer API operation for FortiSwitch gets the Server
// with the specified index value.
// Returns the requested Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.dhcp - server chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDhcpServer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.dhcp/server"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemPtpInterfacePolicy API operation for FortiSwitch creates a new Interface Policy.
// Returns the index value of the Interface Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.ptp - interface-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemPtpInterfacePolicy(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.ptp/interface-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemPtpInterfacePolicy API operation for FortiSwitch updates the specified Interface Policy.
// Returns the index value of the Interface Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.ptp - interface-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPtpInterfacePolicy(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.ptp/interface-policy"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPtpInterfacePolicy API operation for FortiSwitch deletes the specified Interface Policy.
// Returns error for service API and SDK errors.
// See the system.ptp - interface-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPtpInterfacePolicy(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.ptp/interface-policy"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemPtpInterfacePolicy API operation for FortiSwitch gets the Interface Policy
// with the specified index value.
// Returns the requested Interface Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.ptp - interface-policy chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPtpInterfacePolicy(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.ptp/interface-policy"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemPtpProfile API operation for FortiSwitch creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.ptp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemPtpProfile(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.ptp/profile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemPtpProfile API operation for FortiSwitch updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.ptp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPtpProfile(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.ptp/profile"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemPtpProfile API operation for FortiSwitch deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the system.ptp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPtpProfile(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.ptp/profile"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemPtpProfile API operation for FortiSwitch gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.ptp - profile chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPtpProfile(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.ptp/profile"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemScheduleGroup API operation for FortiSwitch creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemScheduleGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.schedule/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemScheduleGroup API operation for FortiSwitch updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemScheduleGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.schedule/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemScheduleGroup API operation for FortiSwitch deletes the specified Group.
// Returns error for service API and SDK errors.
// See the system.schedule - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemScheduleGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.schedule/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemScheduleGroup API operation for FortiSwitch gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemScheduleGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.schedule/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemScheduleOnetime API operation for FortiSwitch creates a new Onetime.
// Returns the index value of the Onetime and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - onetime chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemScheduleOnetime(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.schedule/onetime"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemScheduleOnetime API operation for FortiSwitch updates the specified Onetime.
// Returns the index value of the Onetime and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - onetime chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemScheduleOnetime(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.schedule/onetime"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemScheduleOnetime API operation for FortiSwitch deletes the specified Onetime.
// Returns error for service API and SDK errors.
// See the system.schedule - onetime chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemScheduleOnetime(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.schedule/onetime"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemScheduleOnetime API operation for FortiSwitch gets the Onetime
// with the specified index value.
// Returns the requested Onetime value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - onetime chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemScheduleOnetime(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.schedule/onetime"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemScheduleRecurring API operation for FortiSwitch creates a new Recurring.
// Returns the index value of the Recurring and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - recurring chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemScheduleRecurring(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.schedule/recurring"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemScheduleRecurring API operation for FortiSwitch updates the specified Recurring.
// Returns the index value of the Recurring and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - recurring chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemScheduleRecurring(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.schedule/recurring"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemScheduleRecurring API operation for FortiSwitch deletes the specified Recurring.
// Returns error for service API and SDK errors.
// See the system.schedule - recurring chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemScheduleRecurring(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.schedule/recurring"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemScheduleRecurring API operation for FortiSwitch gets the Recurring
// with the specified index value.
// Returns the requested Recurring value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.schedule - recurring chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemScheduleRecurring(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.schedule/recurring"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateSystemSnmpCommunity API operation for FortiSwitch creates a new Community.
// Returns the index value of the Community and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpCommunity(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.snmp/community"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSnmpCommunity API operation for FortiSwitch updates the specified Community.
// Returns the index value of the Community and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpCommunity(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpCommunity API operation for FortiSwitch deletes the specified Community.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpCommunity(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSnmpCommunity API operation for FortiSwitch gets the Community
// with the specified index value.
// Returns the requested Community value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - community chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpCommunity(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/community"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateSystemSnmpSysinfo API operation for FortiSwitch updates the specified Sysinfo.
// Returns the index value of the Sysinfo and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpSysinfo(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/sysinfo"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpSysinfo API operation for FortiSwitch deletes the specified Sysinfo.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpSysinfo(mkey string) (err error) {

	//No unset API for system.snmp - sysinfo
	return
}

// ReadSystemSnmpSysinfo API operation for FortiSwitch gets the Sysinfo
// with the specified index value.
// Returns the requested Sysinfo value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - sysinfo chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpSysinfo(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/sysinfo"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateSystemSnmpUser API operation for FortiSwitch creates a new User.
// Returns the index value of the User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpUser(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system.snmp/user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateSystemSnmpUser API operation for FortiSwitch updates the specified User.
// Returns the index value of the User and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpUser(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteSystemSnmpUser API operation for FortiSwitch deletes the specified User.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpUser(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadSystemSnmpUser API operation for FortiSwitch gets the User
// with the specified index value.
// Returns the requested User value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system.snmp - user chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpUser(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system.snmp/user"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserGroup API operation for FortiSwitch creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserGroup(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserGroup API operation for FortiSwitch updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserGroup(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/group"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserGroup API operation for FortiSwitch deletes the specified Group.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserGroup(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/group"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserGroup API operation for FortiSwitch gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - group chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserGroup(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/group"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserLdap API operation for FortiSwitch creates a new Ldap.
// Returns the index value of the Ldap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserLdap(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/ldap"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserLdap API operation for FortiSwitch updates the specified Ldap.
// Returns the index value of the Ldap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserLdap(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/ldap"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserLdap API operation for FortiSwitch deletes the specified Ldap.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserLdap(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/ldap"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserLdap API operation for FortiSwitch gets the Ldap
// with the specified index value.
// Returns the requested Ldap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - ldap chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserLdap(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/ldap"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserLocal API operation for FortiSwitch creates a new Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserLocal(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/local"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserLocal API operation for FortiSwitch updates the specified Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserLocal(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/local"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserLocal API operation for FortiSwitch deletes the specified Local.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserLocal(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/local"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserLocal API operation for FortiSwitch gets the Local
// with the specified index value.
// Returns the requested Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - local chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserLocal(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/local"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserPeer API operation for FortiSwitch creates a new Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserPeer(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/peer"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserPeer API operation for FortiSwitch updates the specified Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserPeer(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/peer"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserPeer API operation for FortiSwitch deletes the specified Peer.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserPeer(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/peer"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserPeer API operation for FortiSwitch gets the Peer
// with the specified index value.
// Returns the requested Peer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peer chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserPeer(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/peer"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserPeergrp API operation for FortiSwitch creates a new Peergrp.
// Returns the index value of the Peergrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserPeergrp(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/peergrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserPeergrp API operation for FortiSwitch updates the specified Peergrp.
// Returns the index value of the Peergrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserPeergrp(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/peergrp"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserPeergrp API operation for FortiSwitch deletes the specified Peergrp.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserPeergrp(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/peergrp"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserPeergrp API operation for FortiSwitch gets the Peergrp
// with the specified index value.
// Returns the requested Peergrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - peergrp chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserPeergrp(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/peergrp"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateUserRadius API operation for FortiSwitch creates a new Radius.
// Returns the index value of the Radius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserRadius(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/radius"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserRadius API operation for FortiSwitch updates the specified Radius.
// Returns the index value of the Radius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserRadius(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/radius"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserRadius API operation for FortiSwitch deletes the specified Radius.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserRadius(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/radius"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserRadius API operation for FortiSwitch gets the Radius
// with the specified index value.
// Returns the requested Radius value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - radius chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserRadius(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/radius"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// UpdateUserSetting API operation for FortiSwitch updates the specified Setting.
// Returns the index value of the Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserSetting(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/setting"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserSetting API operation for FortiSwitch deletes the specified Setting.
// Returns error for service API and SDK errors.
// See the user - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserSetting(mkey string) (err error) {

	//No unset API for user - setting
	return
}

// ReadUserSetting API operation for FortiSwitch gets the Setting
// with the specified index value.
// Returns the requested Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - setting chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserSetting(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/setting"

	mapTmp, err = read(c, HTTPMethod, path, true)
	return
}

// CreateUserTacacs API operation for FortiSwitch creates a new Tacacs+.
// Returns the index value of the Tacacs+ and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateUserTacacs(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/user/tacacs+"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateUserTacacs API operation for FortiSwitch updates the specified Tacacs+.
// Returns the index value of the Tacacs+ and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateUserTacacs(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/user/tacacs+"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteUserTacacs API operation for FortiSwitch deletes the specified Tacacs+.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteUserTacacs(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/user/tacacs+"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadUserTacacs API operation for FortiSwitch gets the Tacacs+
// with the specified index value.
// Returns the requested Tacacs+ value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the user - tacacs+ chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadUserTacacs(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/user/tacacs+"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterbgpNeighbor API operation for FortiSwitch creates a new Neighbor.
// Returns the index value of the Neighbor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - neighbor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterbgpNeighbor(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/bgp/neighbor"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterbgpNeighbor API operation for FortiSwitch updates the specified Neighbor.
// Returns the index value of the Neighbor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - neighbor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterbgpNeighbor(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/bgp/neighbor"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterbgpNeighbor API operation for FortiSwitch deletes the specified Neighbor.
// Returns error for service API and SDK errors.
// See the router/bgp - neighbor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterbgpNeighbor(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/bgp/neighbor"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterbgpNeighbor API operation for FortiSwitch gets the Neighbor
// with the specified index value.
// Returns the requested Neighbor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - neighbor chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterbgpNeighbor(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/bgp/neighbor"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterbgpNetwork API operation for FortiSwitch creates a new Network.
// Returns the index value of the Network and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterbgpNetwork(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/bgp/network"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterbgpNetwork API operation for FortiSwitch updates the specified Network.
// Returns the index value of the Network and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterbgpNetwork(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/bgp/network"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterbgpNetwork API operation for FortiSwitch deletes the specified Network.
// Returns error for service API and SDK errors.
// See the router/bgp - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterbgpNetwork(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/bgp/network"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterbgpNetwork API operation for FortiSwitch gets the Network
// with the specified index value.
// Returns the requested Network value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterbgpNetwork(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/bgp/network"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterbgpNetwork6 API operation for FortiSwitch creates a new Network6.
// Returns the index value of the Network6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - network6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterbgpNetwork6(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/bgp/network6"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterbgpNetwork6 API operation for FortiSwitch updates the specified Network6.
// Returns the index value of the Network6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - network6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterbgpNetwork6(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/bgp/network6"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterbgpNetwork6 API operation for FortiSwitch deletes the specified Network6.
// Returns error for service API and SDK errors.
// See the router/bgp - network6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterbgpNetwork6(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/bgp/network6"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterbgpNetwork6 API operation for FortiSwitch gets the Network6
// with the specified index value.
// Returns the requested Network6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/bgp - network6 chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterbgpNetwork6(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/bgp/network6"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterospfInterface API operation for FortiSwitch creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterospfInterface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/ospf/interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterospfInterface API operation for FortiSwitch updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterospfInterface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf/interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterospfInterface API operation for FortiSwitch deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the router/ospf - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterospfInterface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/ospf/interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterospfInterface API operation for FortiSwitch gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterospfInterface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf/interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterospfNetwork API operation for FortiSwitch creates a new Network.
// Returns the index value of the Network and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterospfNetwork(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/ospf/network"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterospfNetwork API operation for FortiSwitch updates the specified Network.
// Returns the index value of the Network and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterospfNetwork(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf/network"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterospfNetwork API operation for FortiSwitch deletes the specified Network.
// Returns error for service API and SDK errors.
// See the router/ospf - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterospfNetwork(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/ospf/network"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterospfNetwork API operation for FortiSwitch gets the Network
// with the specified index value.
// Returns the requested Network value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - network chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterospfNetwork(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf/network"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterospfArea API operation for FortiSwitch creates a new Area.
// Returns the index value of the Area and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterospfArea(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/ospf/area"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterospfArea API operation for FortiSwitch updates the specified Area.
// Returns the index value of the Area and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterospfArea(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf/area"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterospfArea API operation for FortiSwitch deletes the specified Area.
// Returns error for service API and SDK errors.
// See the router/ospf - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterospfArea(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/ospf/area"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterospfArea API operation for FortiSwitch gets the Area
// with the specified index value.
// Returns the requested Area value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterospfArea(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf/area"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterospf6Interface API operation for FortiSwitch creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterospf6Interface(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/ospf6/interface"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterospf6Interface API operation for FortiSwitch updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterospf6Interface(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf6/interface"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterospf6Interface API operation for FortiSwitch deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the router/ospf6 - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterospf6Interface(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/ospf6/interface"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterospf6Interface API operation for FortiSwitch gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - interface chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterospf6Interface(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf6/interface"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterospf6Area API operation for FortiSwitch creates a new Area.
// Returns the index value of the Area and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterospf6Area(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/ospf6/area"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterospf6Area API operation for FortiSwitch updates the specified Area.
// Returns the index value of the Area and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterospf6Area(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf6/area"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterospf6Area API operation for FortiSwitch deletes the specified Area.
// Returns error for service API and SDK errors.
// See the router/ospf6 - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterospf6Area(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/ospf6/area"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterospf6Area API operation for FortiSwitch gets the Area
// with the specified index value.
// Returns the requested Area value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - area chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterospf6Area(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf6/area"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}

// CreateRouterospf6Redistribute API operation for FortiSwitch creates a new Redistribute.
// Returns the index value of the Redistribute and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - redistribute chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) CreateRouterospf6Redistribute(params *map[string]interface{}) (output map[string]interface{}, err error) {

	HTTPMethod := "POST"
	path := "/api/v2/cmdb/router/ospf6/redistribute"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// UpdateRouterospf6Redistribute API operation for FortiSwitch updates the specified Redistribute.
// Returns the index value of the Redistribute and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - redistribute chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateRouterospf6Redistribute(params *map[string]interface{}, mkey string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/router/ospf6/redistribute"
	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output)
	return
}

// DeleteRouterospf6Redistribute API operation for FortiSwitch deletes the specified Redistribute.
// Returns error for service API and SDK errors.
// See the router/ospf6 - redistribute chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteRouterospf6Redistribute(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/router/ospf6/redistribute"
	path += "/" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path)
	return
}

// ReadRouterospf6Redistribute API operation for FortiSwitch gets the Redistribute
// with the specified index value.
// Returns the requested Redistribute value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the router/ospf6 - redistribute chapter in the FortiSwitch Handbook - CLI Reference.
func (c *FortiSDKClient) ReadRouterospf6Redistribute(mkey string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/router/ospf6/redistribute"
	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false)
	return
}
