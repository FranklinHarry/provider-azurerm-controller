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
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfiguration{}).Type1()):                          ToSiteVPNGatewaySpecConnectionConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoute{}).Type1()):                     ToSiteVPNGatewaySpecConnectionConfigurationRouteCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}).Type1()): ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTableCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}).Type1()):      ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPoolCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfiguration{}).Type1()):                          ToSiteVPNGatewaySpecConnectionConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoute{}).Type1()):                     ToSiteVPNGatewaySpecConnectionConfigurationRouteCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}).Type1()): ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTableCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}).Type1()):      ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPoolCodec{},
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
type ToSiteVPNGatewaySpecConnectionConfigurationCodec struct {
}

func (ToSiteVPNGatewaySpecConnectionConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr) == nil
}

func (ToSiteVPNGatewaySpecConnectionConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr)
	var objs []ToSiteVPNGatewaySpecConnectionConfiguration
	if obj != nil {
		objs = []ToSiteVPNGatewaySpecConnectionConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ToSiteVPNGatewaySpecConnectionConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr) = ToSiteVPNGatewaySpecConnectionConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ToSiteVPNGatewaySpecConnectionConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr) = objs[0]
			} else {
				*(*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr) = ToSiteVPNGatewaySpecConnectionConfiguration{}
			}
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr) = ToSiteVPNGatewaySpecConnectionConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ToSiteVPNGatewaySpecConnectionConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr) = obj
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfiguration)(ptr) = ToSiteVPNGatewaySpecConnectionConfiguration{}
		}
	default:
		iter.ReportError("decode ToSiteVPNGatewaySpecConnectionConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ToSiteVPNGatewaySpecConnectionConfigurationRouteCodec struct {
}

func (ToSiteVPNGatewaySpecConnectionConfigurationRouteCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr) == nil
}

func (ToSiteVPNGatewaySpecConnectionConfigurationRouteCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr)
	var objs []ToSiteVPNGatewaySpecConnectionConfigurationRoute
	if obj != nil {
		objs = []ToSiteVPNGatewaySpecConnectionConfigurationRoute{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoute{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ToSiteVPNGatewaySpecConnectionConfigurationRouteCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoute{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ToSiteVPNGatewaySpecConnectionConfigurationRoute

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoute{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr) = objs[0]
			} else {
				*(*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoute{}
			}
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoute{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ToSiteVPNGatewaySpecConnectionConfigurationRoute

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoute{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr) = obj
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfigurationRoute)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoute{}
		}
	default:
		iter.ReportError("decode ToSiteVPNGatewaySpecConnectionConfigurationRoute", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTableCodec struct {
}

func (ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTableCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr) == nil
}

func (ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTableCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr)
	var objs []ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable
	if obj != nil {
		objs = []ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTableCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr) = objs[0]
			} else {
				*(*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}
			}
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr) = obj
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable{}
		}
	default:
		iter.ReportError("decode ToSiteVPNGatewaySpecConnectionConfigurationRoutePropagatedRouteTable", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPoolCodec struct {
}

func (ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPoolCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr) == nil
}

func (ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPoolCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr)
	var objs []ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool
	if obj != nil {
		objs = []ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPoolCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr) = objs[0]
			} else {
				*(*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}
			}
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr) = obj
		} else {
			*(*ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool)(ptr) = ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool{}
		}
	default:
		iter.ReportError("decode ToSiteVPNGatewaySpecConnectionConfigurationVpnClientAddressPool", "unexpected JSON type")
	}
}
