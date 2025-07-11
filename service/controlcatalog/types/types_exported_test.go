// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/controlcatalog/types"
)

func ExampleMapping_outputUsage() {
	var union types.Mapping
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MappingMemberCommonControl:
		_ = v.Value // Value is types.CommonControlMappingDetails

	case *types.MappingMemberFramework:
		_ = v.Value // Value is types.FrameworkMappingDetails

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.FrameworkMappingDetails
var _ *types.CommonControlMappingDetails
