package main

import (
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// Get the target directory from the command line argument
	if len(os.Args) < 2 {
		println("Usage: go run script.go <directory>")
		return
	}
	dir := os.Args[1]

	for _, rule := range replaceRules {
		// Get a list of all files with the given extension in the directory
		files, err := filepath.Glob(filepath.Join(dir, rule.fileSelectPattern))
		if err != nil {
			panic(err)
		}

		// Iterate over the files and apply the regex replacement rules
		for _, file := range files {
			// Read the file contents
			content, err := os.ReadFile(filepath.Clean(file))
			if err != nil {
				panic(err)
			}

			// Apply the regex replacement rule
			re := regexp.MustCompile(rule.pattern)
			content = re.ReplaceAll(content, []byte(rule.repl))

			// Write the updated file contents
			err = os.WriteFile(file, content, os.ModeAppend)
			if err != nil {
				panic(err)
			}
		}
	}
}

var (
	// Define the full list of regex replacement rules
	replaceRules = []struct {
		fileSelectPattern string
		pattern           string
		repl              string
	}{
		{
			fileSelectPattern: "*Api.md",
			pattern:           `PACKAGENAME`,
			repl:              `authorize`,
		},
		// EntityArrayEmbeddedPermissionsInner model
		{
			fileSelectPattern: "model_entity_array__embedded_permissions_inner.go",
			pattern:           `(func \(dst \*EntityArrayEmbeddedPermissionsInner\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *EntityArrayEmbeddedPermissionsInner) UnmarshalJSON(data []byte) error {

	var err error
	// try to unmarshal JSON data into ApplicationRolePermission
	err = json.Unmarshal(data, &dst.ApplicationRolePermission)
	if err == nil {
		jsonApplicationRolePermission, _ := json.Marshal(dst.ApplicationRolePermission)
		if string(jsonApplicationRolePermission) == "{}" { // empty struct
			dst.ApplicationRolePermission = nil
		} else {
			if dst.ApplicationRolePermission.Key != nil {
				return nil // data stored in dst.ApplicationRolePermission, return on the first match
			} else {
				dst.ApplicationRolePermission = nil
			}
		}
	} else {
		dst.ApplicationRolePermission = nil
	}

	// try to unmarshal JSON data into ApplicationResourcePermission
	err = json.Unmarshal(data, &dst.ApplicationResourcePermission)
	if err == nil {
		jsonApplicationResourcePermission, _ := json.Marshal(dst.ApplicationResourcePermission)
		if string(jsonApplicationResourcePermission) == "{}" { // empty struct
			dst.ApplicationResourcePermission = nil
		} else {
			if dst.ApplicationResourcePermission.Action != "" { // we expect an resource for this data type
				dst.ApplicationResourcePermission = nil
			} else {
				return nil // data stored in dst.ApplicationResourcePermission, return on the first match
			}
		}
	} else {
		dst.ApplicationResourcePermission = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EntityArrayEmbeddedPermissionsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataProcessorDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_processor_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataProcessorDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataProcessorDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataProcessorDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataProcessorsChainProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsCollectionTransformProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsJsonPathProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsReferenceProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsSpelProcessorDTO = nil
	dst.AuthorizeEditorDataProcessorsXPathProcessorDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_CHAIN:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsChainProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_COLLECTION_FILTER:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_COLLECTION_TRANSFORM:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsCollectionTransformProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_JSON_PATH:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsJsonPathProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_REFERENCE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsReferenceProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_SPEL:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsSpelProcessorDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAPROCESSORDTOTYPE_XPATH:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataProcessorsXPathProcessorDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataProcessorDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataConditionDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_condition_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataConditionDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataConditionDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataConditionDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataConditionsAndConditionDTO = nil
	dst.AuthorizeEditorDataConditionsComparisonConditionDTO = nil
	dst.AuthorizeEditorDataConditionsEmptyConditionDTO = nil
	dst.AuthorizeEditorDataConditionsNotConditionDTO = nil
	dst.AuthorizeEditorDataConditionsOrConditionDTO = nil
	dst.AuthorizeEditorDataConditionsReferenceConditionDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_AND:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsAndConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_COMPARISON:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsComparisonConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_EMPTY:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsEmptyConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_NOT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsNotConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_OR:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsOrConditionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONDTOTYPE_REFERENCE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsReferenceConditionDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataConditionDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataDefinitionsServiceDefinitionDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_definitions_service_definition_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataDefinitionsServiceDefinitionDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataDefinitionsServiceDefinitionDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataDefinitionsServiceDefinitionDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO = nil
	dst.AuthorizeEditorDataServicesHttpServiceDefinitionDTO = nil
	dst.AuthorizeEditorDataServicesNoneServiceDefinitionDTO = nil

	switch common.GetServiceType() {
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_CONNECTOR:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesConnectorServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_HTTP:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesHttpServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOSERVICETYPE_NONE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataServicesNoneServiceDefinitionDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataDefinitionsServiceDefinitionDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataResolverDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_resolver_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataResolverDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataResolverDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataResolverDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataAttributeResolversAttributeResolverDTO = nil
	dst.AuthorizeEditorDataAttributeResolversConstantResolverDTO = nil
	dst.AuthorizeEditorDataAttributeResolversCurrentRepetitionValueResolverDTO = nil
	dst.AuthorizeEditorDataAttributeResolversCurrentUserIDResolverDTO = nil
	dst.AuthorizeEditorDataAttributeResolversRequestResolverDTO = nil
	dst.AuthorizeEditorDataAttributeResolversServiceResolverDTO = nil
	dst.AuthorizeEditorDataAttributeResolversSystemResolverDTO = nil
	dst.AuthorizeEditorDataAttributeResolversUserResolverDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_ATTRIBUTE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversAttributeResolverDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_CONSTANT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversConstantResolverDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_CURRENT_REPETITION_VALUE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversCurrentRepetitionValueResolverDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_CURRENT_USER_ID:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversCurrentUserIDResolverDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_REQUEST:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversRequestResolverDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_SERVICE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversServiceResolverDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_SYSTEM:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversSystemResolverDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_USER:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversUserResolverDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataResolverDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataRulesEffectSettingsDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_rules_effect_settings_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataRulesEffectSettingsDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataRulesEffectSettingsDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataRulesEffectSettingsDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataRulesEffectSettingsConditionalDenyElsePermitDTO = nil
	dst.AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO = nil
	dst.AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO = nil
	dst.AuthorizeEditorDataRulesEffectSettingsUnconditionalPermitDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATARULESEFFECTSETTINGSDTOTYPE_CONDITIONAL_DENY_ELSE_PERMIT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataRulesEffectSettingsConditionalDenyElsePermitDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARULESEFFECTSETTINGSDTOTYPE_CONDITIONAL_PERMIT_ELSE_DENY:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataRulesEffectSettingsConditionalPermitElseDenyDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARULESEFFECTSETTINGSDTOTYPE_UNCONDITIONAL_DENY:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataRulesEffectSettingsUnconditionalDenyDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATARULESEFFECTSETTINGSDTOTYPE_UNCONDITIONAL_PERMIT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataRulesEffectSettingsUnconditionalPermitDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataRulesEffectSettingsDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataAttributeResolversUserQueryDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_attribute_resolvers_user_query_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataAttributeResolversUserQueryDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataAttributeResolversUserQueryDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataAttributeResolversUserQueryDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATAATTRIBUTERESOLVERSUSERQUERYDTOTYPE_USER_ID:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataAttributeResolversUserQueryDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataAuthenticationDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_authentication_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataAuthenticationDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataAuthenticationDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataAuthenticationDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataAuthenticationsBasicAuthenticationDTO = nil
	dst.AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO = nil
	dst.AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO = nil
	dst.AuthorizeEditorDataAuthenticationsTokenAuthenticationDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATAAUTHENTICATIONDTOTYPE_BASIC:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAuthenticationsBasicAuthenticationDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAAUTHENTICATIONDTOTYPE_CLIENT_CREDENTIALS:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAAUTHENTICATIONDTOTYPE_NONE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAAUTHENTICATIONDTOTYPE_TOKEN:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataAuthenticationsTokenAuthenticationDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataAuthenticationDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataConditionsComparandDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_conditions_comparand_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataConditionsComparandDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataConditionsComparandDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataConditionsComparandDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO = nil
	dst.AuthorizeEditorDataConditionsComparandsConstantComparandDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATACONDITIONSCOMPARANDDTOTYPE_ATTRIBUTE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsComparandsAttributeComparandDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATACONDITIONSCOMPARANDDTOTYPE_CONSTANT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataConditionsComparandsConstantComparandDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataConditionsComparandDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataInputDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_input_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataInputDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataInputDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataInputDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataInputsAttributeInputDTO = nil
	dst.AuthorizeEditorDataInputsConstantInputDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATAINPUTDTOTYPE_ATTRIBUTE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataInputsAttributeInputDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAINPUTDTOTYPE_CONSTANT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataInputsConstantInputDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataInputDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},

		// AuthorizeEditorDataInputMappingDTO model
		{
			fileSelectPattern: "model_authorize_editor_data_input_dto.go",
			pattern:           `(func \(dst \*AuthorizeEditorDataInputMappingDTO\) UnmarshalJSON\(data \[\]byte\) error \{\n)((.*)\n)*\}\n\n\/\/ Marshal data from the first non-nil pointers in the struct to JSON`,
			repl: `func (dst *AuthorizeEditorDataInputMappingDTO) UnmarshalJSON(data []byte) error {

	var common AuthorizeEditorDataInputMappingDTOCommon

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.AuthorizeEditorDataInputMappingsAttributeInputMappingDTO = nil
	dst.AuthorizeEditorDataInputMappingsInputInputMappingDTO = nil

	switch common.GetType() {
	case ENUMAUTHORIZEEDITORDATAINPUTMAPPINGDTOTYPE_ATTRIBUTE:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataInputMappingsAttributeInputMappingDTO); err != nil { // simple model
			return err
		}
	case ENUMAUTHORIZEEDITORDATAINPUTMAPPINGDTOTYPE_INPUT:
		if err := json.Unmarshal(data, &dst.AuthorizeEditorDataInputMappingsInputInputMappingDTO); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(AuthorizeEditorDataInputMappingDTO)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON`,
		},
	}
)
