// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_optional_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("OptionalStruct",
			[]types.Field{
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), true},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), true},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_optional_CachedRef = types.RegisterPackage(&p)
}

// OptionalStruct

type OptionalStruct struct {
	_s          string
	__optionals bool
	_b          bool
	__optionalb bool

	ref *ref.Ref
}

func NewOptionalStruct() OptionalStruct {
	return OptionalStruct{

		ref: &ref.Ref{},
	}
}

type OptionalStructDef struct {
	S string
	B bool
}

func (def OptionalStructDef) New() OptionalStruct {
	return OptionalStruct{
		_s:          def.S,
		__optionals: true,
		_b:          def.B,
		__optionalb: true,
		ref:         &ref.Ref{},
	}
}

func (s OptionalStruct) Def() (d OptionalStructDef) {
	if s.__optionals {
		d.S = s._s
	}
	if s.__optionalb {
		d.B = s._b
	}
	return
}

var __typeRefForOptionalStruct types.TypeRef

func (m OptionalStruct) TypeRef() types.TypeRef {
	return __typeRefForOptionalStruct
}

func init() {
	__typeRefForOptionalStruct = types.MakeTypeRef(__genPackageInFile_struct_optional_CachedRef, 0)
	types.RegisterStruct(__typeRefForOptionalStruct, builderForOptionalStruct, readerForOptionalStruct)
}

func builderForOptionalStruct(values []types.Value) types.Value {
	i := 0
	s := OptionalStruct{ref: &ref.Ref{}}
	s.__optionals = bool(values[i].(types.Bool))
	i++
	if s.__optionals {
		s._s = values[i].(types.String).String()
		i++
	}
	s.__optionalb = bool(values[i].(types.Bool))
	i++
	if s.__optionalb {
		s._b = bool(values[i].(types.Bool))
		i++
	}
	return s
}

func readerForOptionalStruct(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(OptionalStruct)
	values = append(values, types.Bool(s.__optionals))
	if s.__optionals {
		values = append(values, types.NewString(s._s))
	}
	values = append(values, types.Bool(s.__optionalb))
	if s.__optionalb {
		values = append(values, types.Bool(s._b))
	}
	return values
}

func (s OptionalStruct) Equals(other types.Value) bool {
	return other != nil && __typeRefForOptionalStruct.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s OptionalStruct) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s OptionalStruct) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForOptionalStruct.Chunks()...)
	return
}

func (s OptionalStruct) ChildValues() (ret []types.Value) {
	if s.__optionals {
		ret = append(ret, types.NewString(s._s))
	}
	if s.__optionalb {
		ret = append(ret, types.Bool(s._b))
	}
	return
}

func (s OptionalStruct) S() (v string, ok bool) {
	if s.__optionals {
		return s._s, true
	}
	return
}

func (s OptionalStruct) SetS(val string) OptionalStruct {
	s.__optionals = true
	s._s = val
	s.ref = &ref.Ref{}
	return s
}

func (s OptionalStruct) B() (v bool, ok bool) {
	if s.__optionalb {
		return s._b, true
	}
	return
}

func (s OptionalStruct) SetB(val bool) OptionalStruct {
	s.__optionalb = true
	s._b = val
	s.ref = &ref.Ref{}
	return s
}
