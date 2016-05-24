// Code generated by thriftrw

package keyvalue

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/gen/testdata/services"
	"github.com/thriftrw/thriftrw-go/gen/testdata/unions"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type SetValueV2Args struct {
	Key   services.Key           `json:"key"`
	Value *unions.ArbitraryValue `json:"value"`
}

func (v *SetValueV2Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = v.Key.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Value == nil {
		return w, errors.New("field Value of SetValueV2Args is required")
	}
	w, err = v.Value.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SetValueV2Args) FromWire(w wire.Value) error {
	var err error
	keyIsSet := false
	valueIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Key, err = _Key_Read(field.Value)
				if err != nil {
					return err
				}
				keyIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Value, err = _ArbitraryValue_Read(field.Value)
				if err != nil {
					return err
				}
				valueIsSet = true
			}
		}
	}
	if !keyIsSet {
		return errors.New("field Key of SetValueV2Args is required")
	}
	if !valueIsSet {
		return errors.New("field Value of SetValueV2Args is required")
	}
	return nil
}

func (v *SetValueV2Args) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Key: %v", v.Key)
	i++
	fields[i] = fmt.Sprintf("Value: %v", v.Value)
	i++
	return fmt.Sprintf("SetValueV2Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *SetValueV2Args) MethodName() string {
	return "setValueV2"
}

func (v *SetValueV2Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type SetValueV2Result struct{}

func (v *SetValueV2Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SetValueV2Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SetValueV2Result) String() string {
	var fields [0]string
	i := 0
	return fmt.Sprintf("SetValueV2Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SetValueV2Result) MethodName() string {
	return "setValueV2"
}

func (v *SetValueV2Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var SetValueV2Helper = struct {
	IsException    func(error) bool
	Args           func(key services.Key, value *unions.ArbitraryValue) *SetValueV2Args
	WrapResponse   func(error) (*SetValueV2Result, error)
	UnwrapResponse func(*SetValueV2Result) error
}{}

func init() {
	SetValueV2Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SetValueV2Helper.Args = func(key services.Key, value *unions.ArbitraryValue) *SetValueV2Args {
		return &SetValueV2Args{Key: key, Value: value}
	}
	SetValueV2Helper.WrapResponse = func(err error) (*SetValueV2Result, error) {
		if err == nil {
			return &SetValueV2Result{}, nil
		}
		return nil, err
	}
	SetValueV2Helper.UnwrapResponse = func(result *SetValueV2Result) (err error) {
		return
	}
}
