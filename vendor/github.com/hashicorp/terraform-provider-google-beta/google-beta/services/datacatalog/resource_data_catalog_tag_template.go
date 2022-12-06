// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package datacatalog

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/verify"
)

// Use it to delete TagTemplate Field
func deleteTagTemplateField(d *schema.ResourceData, config *transport_tpg.Config, name, billingProject, userAgent string) error {

	url_delete, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}/fields/"+name+"?force={{force_delete}}")
	if err != nil {
		return err
	}
	var obj map[string]interface{}
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url_delete,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return fmt.Errorf("Error deleting TagTemplate Field %v: %s", name, err)
	}

	log.Printf("[DEBUG] Finished deleting TagTemplate Field %q: %#v", name, res)
	return nil
}

// Use it to create TagTemplate Field
func createTagTemplateField(d *schema.ResourceData, config *transport_tpg.Config, body map[string]interface{}, name, billingProject, userAgent string) error {

	url_create, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}/fields")
	if err != nil {
		return err
	}

	url_create, err = transport_tpg.AddQueryParams(url_create, map[string]string{"tagTemplateFieldId": name})
	if err != nil {
		return err
	}

	res_create, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url_create,
		UserAgent: userAgent,
		Body:      body,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating TagTemplate Field: %s", err)
	}

	if err != nil {
		return fmt.Errorf("Error creating TagTemplate Field %v: %s", name, err)
	} else {
		log.Printf("[DEBUG] Finished creating TagTemplate Field %v: %#v", name, res_create)
	}

	return nil
}

func ResourceDataCatalogTagTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataCatalogTagTemplateCreate,
		Read:   resourceDataCatalogTagTemplateRead,
		Update: resourceDataCatalogTagTemplateUpdate,
		Delete: resourceDataCatalogTagTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataCatalogTagTemplateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"fields": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: `Set of tag template field IDs and the settings for the field. This set is an exhaustive list of the allowed fields. This set must contain at least one field and at most 500 fields. The change of field_id will be resulting in re-creating of field. The change of primitive_type will be resulting in re-creating of field, however if the field is a required, you cannot update it.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `The type of value this tag field can contain.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enum_type": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `Represents an enum type.
 Exactly one of 'primitive_type' or 'enum_type' must be set`,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allowed_values": {
													Type:     schema.TypeSet,
													Required: true,
													Description: `The set of allowed values for this enum. The display names of the
values must be case-insensitively unique within this set. Currently,
enum values can only be added to the list of allowed values. Deletion
and renaming of enum values are not supported.
Can have up to 500 allowed values.`,
													Elem: datacatalogTagTemplateFieldsFieldsTypeEnumTypeAllowedValuesSchema(),
													// Default schema.HashSchema is used.
												},
											},
										},
									},
									"primitive_type": {
										Type:         schema.TypeString,
										Computed:     true,
										Optional:     true,
										ValidateFunc: verify.ValidateEnum([]string{"DOUBLE", "STRING", "BOOL", "TIMESTAMP", ""}),
										Description: `Represents primitive types - string, bool etc.
 Exactly one of 'primitive_type' or 'enum_type' must be set Possible values: ["DOUBLE", "STRING", "BOOL", "TIMESTAMP"]`,
									},
								},
							},
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: `A description for this field.`,
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: `The display name for this field.`,
						},
						"is_required": {
							Type:        schema.TypeBool,
							Computed:    true,
							Optional:    true,
							Description: `Whether this is a required field. Defaults to false.`,
						},
						"order": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
							Description: `The order of this field with respect to other fields in this tag template.
A higher value indicates a more important field. The value can be negative.
Multiple fields can have the same order, and field orders within a tag do not have to be sequential.`,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The resource name of the tag template field in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}/fields/{field}`,
						},
					},
				},
			},
			"tag_template_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateRegexp(`^[a-z_][a-z0-9_]{0,63}$`),
				Description:  `The id of the tag template to create.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The display name for this template.`,
			},
			"force_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `This confirms the deletion of any possible tags using this template. Must be set to true in order to delete the tag template.`,
				Default:     false,
			},
			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Template location region.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The resource name of the tag template in URL format. Example: projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func datacatalogTagTemplateFieldsFieldsTypeEnumTypeAllowedValuesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The display name of the enum value.`,
			},
		},
	}
}

func resourceDataCatalogTagTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogTagTemplateDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	fieldsProp, err := expandDataCatalogTagTemplateFields(d.Get("fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(fieldsProp)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}projects/{{project}}/locations/{{region}}/tagTemplates?tagTemplateId={{tag_template_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new TagTemplate: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating TagTemplate: %s", err)
	}
	if err := d.Set("name", flattenDataCatalogTagTemplateName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating TagTemplate %q: %#v", d.Id(), res)

	return resourceDataCatalogTagTemplateRead(d, meta)
}

func resourceDataCatalogTagTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DataCatalogTagTemplate %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("region", region); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}

	if err := d.Set("name", flattenDataCatalogTagTemplateName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}
	if err := d.Set("display_name", flattenDataCatalogTagTemplateDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}
	if err := d.Set("fields", flattenDataCatalogTagTemplateFields(res["fields"], d, config)); err != nil {
		return fmt.Errorf("Error reading TagTemplate: %s", err)
	}

	return nil
}

func resourceDataCatalogTagTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogTagTemplateDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	fieldsProp, err := expandDataCatalogTagTemplateFields(d.Get("fields"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fieldsProp)) {
		obj["fields"] = fieldsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating TagTemplate %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	if len(updateMask) > 0 {

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})

		if err != nil {
			return fmt.Errorf("Error updating TagTemplate %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating TagTemplate %q: %#v", d.Id(), res)
		}

	}

	// since fields have a separate endpoint,
	// we need to handle it manually

	type FieldChange struct {
		Old, New map[string]interface{}
	}

	o, n := d.GetChange("fields")
	vals := make(map[string]*FieldChange)

	// this will create a dictionary with the value
	// of field_id as the key that will contain the
	// maps of old and new values
	for _, raw := range o.(*schema.Set).List() {
		obj := raw.(map[string]interface{})
		k := obj["field_id"].(string)
		vals[k] = &FieldChange{Old: obj}
	}

	for _, raw := range n.(*schema.Set).List() {
		obj := raw.(map[string]interface{})
		k := obj["field_id"].(string)
		if _, ok := vals[k]; !ok {
			// if key is not present in the vals,
			// then create an empty object to hold the new value
			vals[k] = &FieldChange{}
		}
		vals[k].New = obj
	}

	// fields schema to create schema.set below
	dataCatalogTagTemplateFieldsSchema := &schema.Resource{
		Schema: ResourceDataCatalogTagTemplate().Schema["fields"].Elem.(*schema.Resource).Schema,
	}

	for name, change := range vals {
		// A few different situations to deal with in here:
		// - change.Old is nil: create a new role
		// - change.New is nil: remove an existing role
		// - both are set: test if New is different than Old and update if so

		changeOldSet := schema.NewSet(schema.HashResource(dataCatalogTagTemplateFieldsSchema), []interface{}{})
		changeOldSet.Add(change.Old)
		var changeOldProp map[string]interface{}
		if len(change.Old) != 0 {
			changeOldProp, _ = expandDataCatalogTagTemplateFields(changeOldSet, nil, nil)
			changeOldProp = changeOldProp[name].(map[string]interface{})
		}

		changeNewSet := schema.NewSet(schema.HashResource(dataCatalogTagTemplateFieldsSchema), []interface{}{})
		changeNewSet.Add(change.New)
		var changeNewProp map[string]interface{}
		if len(change.New) != 0 {
			changeNewProp, _ = expandDataCatalogTagTemplateFields(changeNewSet, nil, nil)
			changeNewProp = changeNewProp[name].(map[string]interface{})
		}

		// if old state is empty, then we have a new field to create
		if len(change.Old) == 0 {
			err := createTagTemplateField(d, config, changeNewProp, name, billingProject, userAgent)
			if err != nil {
				return err
			}

			continue
		}

		// if new state is empty, then we need to delete the current field
		if len(change.New) == 0 {
			err := deleteTagTemplateField(d, config, name, billingProject, userAgent)
			if err != nil {
				return err
			}

			continue
		}

		// if we have old and new values, but are not equal, update with the new state
		if !reflect.DeepEqual(changeOldProp, changeNewProp) {
			url1, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}/fields/"+name)
			if err != nil {
				return err
			}

			oldType := changeOldProp["type"].(map[string]interface{})
			newType := changeNewProp["type"].(map[string]interface{})

			if oldType["primitiveType"] != newType["primitiveType"] {
				// As primitiveType can't be changed, it is considered as ForceNew which triggers the deletion of old field and recreation of a new field
				// Before that, we need to check that is_required is True for the newType or not, as we don't have support to add new required field in the existing TagTemplate,
				// So in such cases, we can simply return the error

				// Reason for checking the isRequired in changeNewProp -
				// Because this changeNewProp check should be ignored when the user wants to update the primitive type and make it optional rather than keeping it required.
				if changeNewProp["isRequired"] != nil && changeNewProp["isRequired"].(bool) {
					return fmt.Errorf("Updating the primitive type for a required field on an existing tag template is not supported as TagTemplateField %q is required", name)
				}

				// delete changeOldProp
				err_delete := deleteTagTemplateField(d, config, name, billingProject, userAgent)
				if err_delete != nil {
					return err_delete
				}

				// recreate changeNewProp
				err_create := createTagTemplateField(d, config, changeNewProp, name, billingProject, userAgent)
				if err_create != nil {
					return err_create
				}

				log.Printf("[DEBUG] Finished updating TagTemplate Field %q", name)
				return resourceDataCatalogTagTemplateRead(d, meta)
			}

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "PATCH",
				Project:   billingProject,
				RawURL:    url1,
				UserAgent: userAgent,
				Body:      changeNewProp,
				Timeout:   d.Timeout(schema.TimeoutDelete),
			})
			if err != nil {
				return fmt.Errorf("Error updating TagTemplate Field %v: %s", name, err)
			}

			log.Printf("[DEBUG] Finished updating TagTemplate Field %q: %#v", name, res)
		}
	}
	return resourceDataCatalogTagTemplateRead(d, meta)
}

func resourceDataCatalogTagTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for TagTemplate: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DataCatalogBasePath}}{{name}}?force={{force_delete}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting TagTemplate %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "TagTemplate")
	}

	log.Printf("[DEBUG] Finished deleting TagTemplate %q: %#v", d.Id(), res)
	return nil
}

func resourceDataCatalogTagTemplateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	egRegex := regexp.MustCompile("projects/(.+)/locations/(.+)/tagTemplates/(.+)")

	parts := egRegex.FindStringSubmatch(name)
	if len(parts) != 4 {
		return nil, fmt.Errorf("tag template name does not fit the format %s", egRegex)
	}
	if err := d.Set("project", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("region", parts[2]); err != nil {
		return nil, fmt.Errorf("Error setting region: %s", err)
	}
	if err := d.Set("tag_template_id", parts[3]); err != nil {
		return nil, fmt.Errorf("Error setting tag_template_id: %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func flattenDataCatalogTagTemplateName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFields(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"field_id":     k,
			"name":         flattenDataCatalogTagTemplateFieldsName(original["name"], d, config),
			"display_name": flattenDataCatalogTagTemplateFieldsDisplayName(original["displayName"], d, config),
			"description":  flattenDataCatalogTagTemplateFieldsDescription(original["description"], d, config),
			"type":         flattenDataCatalogTagTemplateFieldsType(original["type"], d, config),
			"is_required":  flattenDataCatalogTagTemplateFieldsIsRequired(original["isRequired"], d, config),
			"order":        flattenDataCatalogTagTemplateFieldsOrder(original["order"], d, config),
		})
	}
	return transformed
}
func flattenDataCatalogTagTemplateFieldsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["primitive_type"] =
		flattenDataCatalogTagTemplateFieldsTypePrimitiveType(original["primitiveType"], d, config)
	transformed["enum_type"] =
		flattenDataCatalogTagTemplateFieldsTypeEnumType(original["enumType"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogTagTemplateFieldsTypePrimitiveType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsTypeEnumType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["allowed_values"] =
		flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(original["allowedValues"], d, config)
	return []interface{}{transformed}
}
func flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(datacatalogTagTemplateFieldsFieldsTypeEnumTypeAllowedValuesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"display_name": flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(original["displayName"], d, config),
		})
	}
	return transformed
}
func flattenDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsIsRequired(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDataCatalogTagTemplateFieldsOrder(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func expandDataCatalogTagTemplateDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataCatalogTagTemplateFieldsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedDisplayName, err := expandDataCatalogTagTemplateFieldsDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		transformedDescription, err := expandDataCatalogTagTemplateFieldsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedType, err := expandDataCatalogTagTemplateFieldsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		transformedIsRequired, err := expandDataCatalogTagTemplateFieldsIsRequired(original["is_required"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIsRequired); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["isRequired"] = transformedIsRequired
		}

		transformedOrder, err := expandDataCatalogTagTemplateFieldsOrder(original["order"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOrder); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["order"] = transformedOrder
		}

		transformedFieldId, err := tpgresource.ExpandString(original["field_id"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedFieldId] = transformed
	}
	return m, nil
}

func expandDataCatalogTagTemplateFieldsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrimitiveType, err := expandDataCatalogTagTemplateFieldsTypePrimitiveType(original["primitive_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimitiveType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primitiveType"] = transformedPrimitiveType
	}

	transformedEnumType, err := expandDataCatalogTagTemplateFieldsTypeEnumType(original["enum_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnumType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enumType"] = transformedEnumType
	}

	return transformed, nil
}

func expandDataCatalogTagTemplateFieldsTypePrimitiveType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllowedValues, err := expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(original["allowed_values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedValues"] = transformedAllowedValues
	}

	return transformed, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDisplayName, err := expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(original["display_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisplayName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["displayName"] = transformedDisplayName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataCatalogTagTemplateFieldsTypeEnumTypeAllowedValuesDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsIsRequired(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogTagTemplateFieldsOrder(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}