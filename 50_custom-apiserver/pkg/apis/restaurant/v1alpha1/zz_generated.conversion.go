//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	restaurant "50_custom-apiserver/pkg/apis/restaurant"
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Pizza)(nil), (*restaurant.Pizza)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Pizza_To_restaurant_Pizza(a.(*Pizza), b.(*restaurant.Pizza), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*restaurant.Pizza)(nil), (*Pizza)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_restaurant_Pizza_To_v1alpha1_Pizza(a.(*restaurant.Pizza), b.(*Pizza), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PizzaList)(nil), (*restaurant.PizzaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PizzaList_To_restaurant_PizzaList(a.(*PizzaList), b.(*restaurant.PizzaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*restaurant.PizzaList)(nil), (*PizzaList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_restaurant_PizzaList_To_v1alpha1_PizzaList(a.(*restaurant.PizzaList), b.(*PizzaList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PizzaStatus)(nil), (*restaurant.PizzaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PizzaStatus_To_restaurant_PizzaStatus(a.(*PizzaStatus), b.(*restaurant.PizzaStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*restaurant.PizzaStatus)(nil), (*PizzaStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_restaurant_PizzaStatus_To_v1alpha1_PizzaStatus(a.(*restaurant.PizzaStatus), b.(*PizzaStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Topping)(nil), (*restaurant.Topping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Topping_To_restaurant_Topping(a.(*Topping), b.(*restaurant.Topping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*restaurant.Topping)(nil), (*Topping)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_restaurant_Topping_To_v1alpha1_Topping(a.(*restaurant.Topping), b.(*Topping), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ToppingList)(nil), (*restaurant.ToppingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ToppingList_To_restaurant_ToppingList(a.(*ToppingList), b.(*restaurant.ToppingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*restaurant.ToppingList)(nil), (*ToppingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_restaurant_ToppingList_To_v1alpha1_ToppingList(a.(*restaurant.ToppingList), b.(*ToppingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ToppingSpec)(nil), (*restaurant.ToppingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ToppingSpec_To_restaurant_ToppingSpec(a.(*ToppingSpec), b.(*restaurant.ToppingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*restaurant.ToppingSpec)(nil), (*ToppingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_restaurant_ToppingSpec_To_v1alpha1_ToppingSpec(a.(*restaurant.ToppingSpec), b.(*ToppingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*restaurant.PizzaSpec)(nil), (*PizzaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_restaurant_PizzaSpec_To_v1alpha1_PizzaSpec(a.(*restaurant.PizzaSpec), b.(*PizzaSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*PizzaSpec)(nil), (*restaurant.PizzaSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PizzaSpec_To_restaurant_PizzaSpec(a.(*PizzaSpec), b.(*restaurant.PizzaSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Pizza_To_restaurant_Pizza(in *Pizza, out *restaurant.Pizza, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PizzaSpec_To_restaurant_PizzaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PizzaStatus_To_restaurant_PizzaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Pizza_To_restaurant_Pizza is an autogenerated conversion function.
func Convert_v1alpha1_Pizza_To_restaurant_Pizza(in *Pizza, out *restaurant.Pizza, s conversion.Scope) error {
	return autoConvert_v1alpha1_Pizza_To_restaurant_Pizza(in, out, s)
}

func autoConvert_restaurant_Pizza_To_v1alpha1_Pizza(in *restaurant.Pizza, out *Pizza, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_restaurant_PizzaSpec_To_v1alpha1_PizzaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_restaurant_PizzaStatus_To_v1alpha1_PizzaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_restaurant_Pizza_To_v1alpha1_Pizza is an autogenerated conversion function.
func Convert_restaurant_Pizza_To_v1alpha1_Pizza(in *restaurant.Pizza, out *Pizza, s conversion.Scope) error {
	return autoConvert_restaurant_Pizza_To_v1alpha1_Pizza(in, out, s)
}

func autoConvert_v1alpha1_PizzaList_To_restaurant_PizzaList(in *PizzaList, out *restaurant.PizzaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]restaurant.Pizza, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_Pizza_To_restaurant_Pizza(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_PizzaList_To_restaurant_PizzaList is an autogenerated conversion function.
func Convert_v1alpha1_PizzaList_To_restaurant_PizzaList(in *PizzaList, out *restaurant.PizzaList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PizzaList_To_restaurant_PizzaList(in, out, s)
}

func autoConvert_restaurant_PizzaList_To_v1alpha1_PizzaList(in *restaurant.PizzaList, out *PizzaList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pizza, len(*in))
		for i := range *in {
			if err := Convert_restaurant_Pizza_To_v1alpha1_Pizza(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_restaurant_PizzaList_To_v1alpha1_PizzaList is an autogenerated conversion function.
func Convert_restaurant_PizzaList_To_v1alpha1_PizzaList(in *restaurant.PizzaList, out *PizzaList, s conversion.Scope) error {
	return autoConvert_restaurant_PizzaList_To_v1alpha1_PizzaList(in, out, s)
}

func autoConvert_v1alpha1_PizzaStatus_To_restaurant_PizzaStatus(in *PizzaStatus, out *restaurant.PizzaStatus, s conversion.Scope) error {
	out.Cost = in.Cost
	return nil
}

// Convert_v1alpha1_PizzaStatus_To_restaurant_PizzaStatus is an autogenerated conversion function.
func Convert_v1alpha1_PizzaStatus_To_restaurant_PizzaStatus(in *PizzaStatus, out *restaurant.PizzaStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PizzaStatus_To_restaurant_PizzaStatus(in, out, s)
}

func autoConvert_restaurant_PizzaStatus_To_v1alpha1_PizzaStatus(in *restaurant.PizzaStatus, out *PizzaStatus, s conversion.Scope) error {
	out.Cost = in.Cost
	return nil
}

// Convert_restaurant_PizzaStatus_To_v1alpha1_PizzaStatus is an autogenerated conversion function.
func Convert_restaurant_PizzaStatus_To_v1alpha1_PizzaStatus(in *restaurant.PizzaStatus, out *PizzaStatus, s conversion.Scope) error {
	return autoConvert_restaurant_PizzaStatus_To_v1alpha1_PizzaStatus(in, out, s)
}

func autoConvert_v1alpha1_Topping_To_restaurant_Topping(in *Topping, out *restaurant.Topping, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ToppingSpec_To_restaurant_ToppingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Topping_To_restaurant_Topping is an autogenerated conversion function.
func Convert_v1alpha1_Topping_To_restaurant_Topping(in *Topping, out *restaurant.Topping, s conversion.Scope) error {
	return autoConvert_v1alpha1_Topping_To_restaurant_Topping(in, out, s)
}

func autoConvert_restaurant_Topping_To_v1alpha1_Topping(in *restaurant.Topping, out *Topping, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_restaurant_ToppingSpec_To_v1alpha1_ToppingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_restaurant_Topping_To_v1alpha1_Topping is an autogenerated conversion function.
func Convert_restaurant_Topping_To_v1alpha1_Topping(in *restaurant.Topping, out *Topping, s conversion.Scope) error {
	return autoConvert_restaurant_Topping_To_v1alpha1_Topping(in, out, s)
}

func autoConvert_v1alpha1_ToppingList_To_restaurant_ToppingList(in *ToppingList, out *restaurant.ToppingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]restaurant.Topping)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ToppingList_To_restaurant_ToppingList is an autogenerated conversion function.
func Convert_v1alpha1_ToppingList_To_restaurant_ToppingList(in *ToppingList, out *restaurant.ToppingList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ToppingList_To_restaurant_ToppingList(in, out, s)
}

func autoConvert_restaurant_ToppingList_To_v1alpha1_ToppingList(in *restaurant.ToppingList, out *ToppingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Topping)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_restaurant_ToppingList_To_v1alpha1_ToppingList is an autogenerated conversion function.
func Convert_restaurant_ToppingList_To_v1alpha1_ToppingList(in *restaurant.ToppingList, out *ToppingList, s conversion.Scope) error {
	return autoConvert_restaurant_ToppingList_To_v1alpha1_ToppingList(in, out, s)
}

func autoConvert_v1alpha1_ToppingSpec_To_restaurant_ToppingSpec(in *ToppingSpec, out *restaurant.ToppingSpec, s conversion.Scope) error {
	out.Cost = in.Cost
	return nil
}

// Convert_v1alpha1_ToppingSpec_To_restaurant_ToppingSpec is an autogenerated conversion function.
func Convert_v1alpha1_ToppingSpec_To_restaurant_ToppingSpec(in *ToppingSpec, out *restaurant.ToppingSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ToppingSpec_To_restaurant_ToppingSpec(in, out, s)
}

func autoConvert_restaurant_ToppingSpec_To_v1alpha1_ToppingSpec(in *restaurant.ToppingSpec, out *ToppingSpec, s conversion.Scope) error {
	out.Cost = in.Cost
	return nil
}

// Convert_restaurant_ToppingSpec_To_v1alpha1_ToppingSpec is an autogenerated conversion function.
func Convert_restaurant_ToppingSpec_To_v1alpha1_ToppingSpec(in *restaurant.ToppingSpec, out *ToppingSpec, s conversion.Scope) error {
	return autoConvert_restaurant_ToppingSpec_To_v1alpha1_ToppingSpec(in, out, s)
}
