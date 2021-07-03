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
		jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescription{}).Type1()):            EventhubSpecCaptureDescriptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescriptionDestination{}).Type1()): EventhubSpecCaptureDescriptionDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecIdentity{}).Type1()):                    Namespace_SpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecNetworkRulesets{}).Type1()):             Namespace_SpecNetworkRulesetsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescription{}).Type1()):            EventhubSpecCaptureDescriptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescriptionDestination{}).Type1()): EventhubSpecCaptureDescriptionDestinationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecIdentity{}).Type1()):                    Namespace_SpecIdentityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecNetworkRulesets{}).Type1()):             Namespace_SpecNetworkRulesetsCodec{},
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
type EventhubSpecCaptureDescriptionCodec struct {
}

func (EventhubSpecCaptureDescriptionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EventhubSpecCaptureDescription)(ptr) == nil
}

func (EventhubSpecCaptureDescriptionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EventhubSpecCaptureDescription)(ptr)
	var objs []EventhubSpecCaptureDescription
	if obj != nil {
		objs = []EventhubSpecCaptureDescription{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescription{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EventhubSpecCaptureDescriptionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EventhubSpecCaptureDescription)(ptr) = EventhubSpecCaptureDescription{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EventhubSpecCaptureDescription

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescription{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EventhubSpecCaptureDescription)(ptr) = objs[0]
			} else {
				*(*EventhubSpecCaptureDescription)(ptr) = EventhubSpecCaptureDescription{}
			}
		} else {
			*(*EventhubSpecCaptureDescription)(ptr) = EventhubSpecCaptureDescription{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EventhubSpecCaptureDescription

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescription{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EventhubSpecCaptureDescription)(ptr) = obj
		} else {
			*(*EventhubSpecCaptureDescription)(ptr) = EventhubSpecCaptureDescription{}
		}
	default:
		iter.ReportError("decode EventhubSpecCaptureDescription", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EventhubSpecCaptureDescriptionDestinationCodec struct {
}

func (EventhubSpecCaptureDescriptionDestinationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EventhubSpecCaptureDescriptionDestination)(ptr) == nil
}

func (EventhubSpecCaptureDescriptionDestinationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EventhubSpecCaptureDescriptionDestination)(ptr)
	var objs []EventhubSpecCaptureDescriptionDestination
	if obj != nil {
		objs = []EventhubSpecCaptureDescriptionDestination{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescriptionDestination{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EventhubSpecCaptureDescriptionDestinationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EventhubSpecCaptureDescriptionDestination)(ptr) = EventhubSpecCaptureDescriptionDestination{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EventhubSpecCaptureDescriptionDestination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescriptionDestination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EventhubSpecCaptureDescriptionDestination)(ptr) = objs[0]
			} else {
				*(*EventhubSpecCaptureDescriptionDestination)(ptr) = EventhubSpecCaptureDescriptionDestination{}
			}
		} else {
			*(*EventhubSpecCaptureDescriptionDestination)(ptr) = EventhubSpecCaptureDescriptionDestination{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EventhubSpecCaptureDescriptionDestination

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EventhubSpecCaptureDescriptionDestination{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EventhubSpecCaptureDescriptionDestination)(ptr) = obj
		} else {
			*(*EventhubSpecCaptureDescriptionDestination)(ptr) = EventhubSpecCaptureDescriptionDestination{}
		}
	default:
		iter.ReportError("decode EventhubSpecCaptureDescriptionDestination", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type Namespace_SpecIdentityCodec struct {
}

func (Namespace_SpecIdentityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*Namespace_SpecIdentity)(ptr) == nil
}

func (Namespace_SpecIdentityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*Namespace_SpecIdentity)(ptr)
	var objs []Namespace_SpecIdentity
	if obj != nil {
		objs = []Namespace_SpecIdentity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecIdentity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (Namespace_SpecIdentityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*Namespace_SpecIdentity)(ptr) = Namespace_SpecIdentity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []Namespace_SpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*Namespace_SpecIdentity)(ptr) = objs[0]
			} else {
				*(*Namespace_SpecIdentity)(ptr) = Namespace_SpecIdentity{}
			}
		} else {
			*(*Namespace_SpecIdentity)(ptr) = Namespace_SpecIdentity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj Namespace_SpecIdentity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecIdentity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*Namespace_SpecIdentity)(ptr) = obj
		} else {
			*(*Namespace_SpecIdentity)(ptr) = Namespace_SpecIdentity{}
		}
	default:
		iter.ReportError("decode Namespace_SpecIdentity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type Namespace_SpecNetworkRulesetsCodec struct {
}

func (Namespace_SpecNetworkRulesetsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*Namespace_SpecNetworkRulesets)(ptr) == nil
}

func (Namespace_SpecNetworkRulesetsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*Namespace_SpecNetworkRulesets)(ptr)
	var objs []Namespace_SpecNetworkRulesets
	if obj != nil {
		objs = []Namespace_SpecNetworkRulesets{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecNetworkRulesets{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (Namespace_SpecNetworkRulesetsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*Namespace_SpecNetworkRulesets)(ptr) = Namespace_SpecNetworkRulesets{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []Namespace_SpecNetworkRulesets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecNetworkRulesets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*Namespace_SpecNetworkRulesets)(ptr) = objs[0]
			} else {
				*(*Namespace_SpecNetworkRulesets)(ptr) = Namespace_SpecNetworkRulesets{}
			}
		} else {
			*(*Namespace_SpecNetworkRulesets)(ptr) = Namespace_SpecNetworkRulesets{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj Namespace_SpecNetworkRulesets

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(Namespace_SpecNetworkRulesets{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*Namespace_SpecNetworkRulesets)(ptr) = obj
		} else {
			*(*Namespace_SpecNetworkRulesets)(ptr) = Namespace_SpecNetworkRulesets{}
		}
	default:
		iter.ReportError("decode Namespace_SpecNetworkRulesets", "unexpected JSON type")
	}
}
