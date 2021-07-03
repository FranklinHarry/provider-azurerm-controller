/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecAutoscaleConfiguration{}).Type1()):                GatewaySpecAutoscaleConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBackendHTTPSettingsConnectionDraining{}).Type1()): GatewaySpecBackendHTTPSettingsConnectionDrainingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecIdentity{}).Type1()):                              GatewaySpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecProbeMatch{}).Type1()):                            GatewaySpecProbeMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecRewriteRuleSetRewriteRuleUrl{}).Type1()):          GatewaySpecRewriteRuleSetRewriteRuleUrlCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSku{}).Type1()):                                   GatewaySpecSkuCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecWafConfiguration{}).Type1()):                      GatewaySpecWafConfigurationCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecAutoscaleConfiguration{}).Type1()):                GatewaySpecAutoscaleConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBackendHTTPSettingsConnectionDraining{}).Type1()): GatewaySpecBackendHTTPSettingsConnectionDrainingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecIdentity{}).Type1()):                              GatewaySpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecProbeMatch{}).Type1()):                            GatewaySpecProbeMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecRewriteRuleSetRewriteRuleUrl{}).Type1()):          GatewaySpecRewriteRuleSetRewriteRuleUrlCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSku{}).Type1()):                                   GatewaySpecSkuCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecWafConfiguration{}).Type1()):                      GatewaySpecWafConfigurationCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type GatewaySpecAutoscaleConfigurationCodec struct {
}

func (GatewaySpecAutoscaleConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecAutoscaleConfiguration)(ptr) == nil
}

func (GatewaySpecAutoscaleConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecAutoscaleConfiguration)(ptr)
	var objs []GatewaySpecAutoscaleConfiguration
	if obj != nil {
		objs = []GatewaySpecAutoscaleConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecAutoscaleConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecAutoscaleConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecAutoscaleConfiguration)(ptr) = GatewaySpecAutoscaleConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecAutoscaleConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecAutoscaleConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecAutoscaleConfiguration)(ptr) = objs[0]
			} else {
				*(*GatewaySpecAutoscaleConfiguration)(ptr) = GatewaySpecAutoscaleConfiguration{}
			}
		} else {
			*(*GatewaySpecAutoscaleConfiguration)(ptr) = GatewaySpecAutoscaleConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecAutoscaleConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecAutoscaleConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecAutoscaleConfiguration)(ptr) = obj
		} else {
			*(*GatewaySpecAutoscaleConfiguration)(ptr) = GatewaySpecAutoscaleConfiguration{}
		}
	default:
		iter.ReportError("decode GatewaySpecAutoscaleConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecBackendHTTPSettingsConnectionDrainingCodec struct {
}

func (GatewaySpecBackendHTTPSettingsConnectionDrainingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr) == nil
}

func (GatewaySpecBackendHTTPSettingsConnectionDrainingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr)
	var objs []GatewaySpecBackendHTTPSettingsConnectionDraining
	if obj != nil {
		objs = []GatewaySpecBackendHTTPSettingsConnectionDraining{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBackendHTTPSettingsConnectionDraining{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecBackendHTTPSettingsConnectionDrainingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr) = GatewaySpecBackendHTTPSettingsConnectionDraining{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecBackendHTTPSettingsConnectionDraining

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBackendHTTPSettingsConnectionDraining{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr) = objs[0]
			} else {
				*(*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr) = GatewaySpecBackendHTTPSettingsConnectionDraining{}
			}
		} else {
			*(*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr) = GatewaySpecBackendHTTPSettingsConnectionDraining{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecBackendHTTPSettingsConnectionDraining

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecBackendHTTPSettingsConnectionDraining{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr) = obj
		} else {
			*(*GatewaySpecBackendHTTPSettingsConnectionDraining)(ptr) = GatewaySpecBackendHTTPSettingsConnectionDraining{}
		}
	default:
		iter.ReportError("decode GatewaySpecBackendHTTPSettingsConnectionDraining", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecIdentityCodec struct {
}

func (GatewaySpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecIdentity)(ptr) == nil
}

func (GatewaySpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecIdentity)(ptr)
	var objs []GatewaySpecIdentity
	if obj != nil {
		objs = []GatewaySpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecIdentity)(ptr) = GatewaySpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecIdentity)(ptr) = objs[0]
			} else {
				*(*GatewaySpecIdentity)(ptr) = GatewaySpecIdentity{}
			}
		} else {
			*(*GatewaySpecIdentity)(ptr) = GatewaySpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecIdentity)(ptr) = obj
		} else {
			*(*GatewaySpecIdentity)(ptr) = GatewaySpecIdentity{}
		}
	default:
		iter.ReportError("decode GatewaySpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecProbeMatchCodec struct {
}

func (GatewaySpecProbeMatchCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecProbeMatch)(ptr) == nil
}

func (GatewaySpecProbeMatchCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecProbeMatch)(ptr)
	var objs []GatewaySpecProbeMatch
	if obj != nil {
		objs = []GatewaySpecProbeMatch{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecProbeMatch{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecProbeMatchCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecProbeMatch)(ptr) = GatewaySpecProbeMatch{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecProbeMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecProbeMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecProbeMatch)(ptr) = objs[0]
			} else {
				*(*GatewaySpecProbeMatch)(ptr) = GatewaySpecProbeMatch{}
			}
		} else {
			*(*GatewaySpecProbeMatch)(ptr) = GatewaySpecProbeMatch{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecProbeMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecProbeMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecProbeMatch)(ptr) = obj
		} else {
			*(*GatewaySpecProbeMatch)(ptr) = GatewaySpecProbeMatch{}
		}
	default:
		iter.ReportError("decode GatewaySpecProbeMatch", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecRewriteRuleSetRewriteRuleUrlCodec struct {
}

func (GatewaySpecRewriteRuleSetRewriteRuleUrlCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr) == nil
}

func (GatewaySpecRewriteRuleSetRewriteRuleUrlCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr)
	var objs []GatewaySpecRewriteRuleSetRewriteRuleUrl
	if obj != nil {
		objs = []GatewaySpecRewriteRuleSetRewriteRuleUrl{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecRewriteRuleSetRewriteRuleUrl{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecRewriteRuleSetRewriteRuleUrlCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr) = GatewaySpecRewriteRuleSetRewriteRuleUrl{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecRewriteRuleSetRewriteRuleUrl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecRewriteRuleSetRewriteRuleUrl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr) = objs[0]
			} else {
				*(*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr) = GatewaySpecRewriteRuleSetRewriteRuleUrl{}
			}
		} else {
			*(*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr) = GatewaySpecRewriteRuleSetRewriteRuleUrl{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecRewriteRuleSetRewriteRuleUrl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecRewriteRuleSetRewriteRuleUrl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr) = obj
		} else {
			*(*GatewaySpecRewriteRuleSetRewriteRuleUrl)(ptr) = GatewaySpecRewriteRuleSetRewriteRuleUrl{}
		}
	default:
		iter.ReportError("decode GatewaySpecRewriteRuleSetRewriteRuleUrl", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecSkuCodec struct {
}

func (GatewaySpecSkuCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecSku)(ptr) == nil
}

func (GatewaySpecSkuCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecSku)(ptr)
	var objs []GatewaySpecSku
	if obj != nil {
		objs = []GatewaySpecSku{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSku{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecSkuCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecSku)(ptr) = GatewaySpecSku{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecSku

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSku{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecSku)(ptr) = objs[0]
			} else {
				*(*GatewaySpecSku)(ptr) = GatewaySpecSku{}
			}
		} else {
			*(*GatewaySpecSku)(ptr) = GatewaySpecSku{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecSku

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecSku{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecSku)(ptr) = obj
		} else {
			*(*GatewaySpecSku)(ptr) = GatewaySpecSku{}
		}
	default:
		iter.ReportError("decode GatewaySpecSku", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySpecWafConfigurationCodec struct {
}

func (GatewaySpecWafConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySpecWafConfiguration)(ptr) == nil
}

func (GatewaySpecWafConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySpecWafConfiguration)(ptr)
	var objs []GatewaySpecWafConfiguration
	if obj != nil {
		objs = []GatewaySpecWafConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecWafConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySpecWafConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySpecWafConfiguration)(ptr) = GatewaySpecWafConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySpecWafConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecWafConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySpecWafConfiguration)(ptr) = objs[0]
			} else {
				*(*GatewaySpecWafConfiguration)(ptr) = GatewaySpecWafConfiguration{}
			}
		} else {
			*(*GatewaySpecWafConfiguration)(ptr) = GatewaySpecWafConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySpecWafConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySpecWafConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySpecWafConfiguration)(ptr) = obj
		} else {
			*(*GatewaySpecWafConfiguration)(ptr) = GatewaySpecWafConfiguration{}
		}
	default:
		iter.ReportError("decode GatewaySpecWafConfiguration", "unexpected JSON type")
	}
}
