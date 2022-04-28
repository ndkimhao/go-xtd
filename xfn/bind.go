package xfn

import (
	"fmt"
	"reflect"

	"github.com/ndkimhao/go-xtd/generics"
)

type Placeholder int

func (p Placeholder) String() string {
	return fmt.Sprint("Placeholder(", p, ")")
}

var (
	P1 Placeholder = 1
	P2             = 2
	P3             = 3
	P4             = 4
)

func P(i int) Placeholder {
	if i < 1 {
		panic("invalid index")
	}
	return Placeholder(i)
}

type phInfo struct {
	fromArg int
	toArg   int
}

func _ith(i int) string {
	lastDigit := i % 10
	if lastDigit == 1 {
		return fmt.Sprint(i, "st")
	} else if lastDigit == 2 {
		return fmt.Sprint(i, "nd")
	} else if lastDigit == 3 {
		return fmt.Sprint(i, "rd")
	} else {
		return fmt.Sprint(i)
	}
}

func Bind[Dest any](src any, args ...any) Dest {
	srcV := reflect.ValueOf(src)
	srcT := srcV.Type()
	if srcT.Kind() != reflect.Func {
		panic(fmt.Sprintf("invaid type: src [%s] is not a func type", srcT))
	}
	destT := generics.TypeOf[Dest]()
	if destT.Kind() != reflect.Func {
		panic(fmt.Sprintf("invaid type: dest [%s] is not a func type", destT))
	}
	if srcT.NumOut() != destT.NumOut() {
		panic(fmt.Sprintf("incompatible output: "+
			"src [%s] has %d outputs but dest [%s] has %d",
			srcT, srcT.NumOut(), destT, destT.NumOut()))
	}
	for i := 0; i < srcT.NumOut(); i++ {
		if srcT.Out(i) != destT.Out(i) {
			panic(fmt.Sprintf("incompatible output: "+
				"%s output of src [%s] is [%s] but %s output of dest [%s] is [%s]",
				_ith(i+1), srcT, srcT.Out(i), _ith(i+1), destT, destT.Out(i)))
		}
	}

	srcIn := srcT.NumIn()
	destIn := destT.NumIn()
	if len(args) > srcIn {
		panic(fmt.Sprintf("too many arguments: "+
			"src [%s] has %d inputs but provided %d bind arguments",
			srcT, srcIn, len(args)))
	}
	values := make([]reflect.Value, srcIn)
	placeholders := []phInfo(nil)
	nextP := 0
	for i, a := range args {
		if placeholder, ok := a.(Placeholder); ok {
			pIdx := int(placeholder) - 1
			if pIdx < 0 || destIn <= pIdx {
				panic(fmt.Sprintf("invalid placeholder: "+
					"dest [%s] has %d inputs but provided %s by %s bind argument",
					destT, destIn, placeholder, _ith(i+1)))
			}
			if srcT.In(i) != destT.In(pIdx) {
				panic(fmt.Sprintf("type mismatch: "+
					"src [%s] %s input is [%s] but provided [%s] via %s",
					srcT, _ith(i+1), srcT.In(i), destT.In(pIdx), placeholder))
			}
			placeholders = append(placeholders, phInfo{fromArg: pIdx, toArg: i})
			nextP = pIdx + 1
			continue
		}
		valA := reflect.ValueOf(a)
		if !valA.CanConvert(srcT.In(i)) {
			panic(fmt.Sprintf("type mismatch: "+
				"src [%s] %s input is %s but provided %s",
				srcT, _ith(i+1), srcT.In(i), valA.Type()))
		}
		values[i] = valA.Convert(srcT.In(i))
	}
	if need, have := srcIn-len(args), destIn-nextP; need > have {
		panic(fmt.Sprintf("not enough inputs: "+
			"dest [%s] needs %d more inputs", destT, need-have))
	}
	for i := len(args); i < srcIn; i++ {
		placeholders = append(placeholders, phInfo{fromArg: nextP, toArg: i})
		nextP++
	}
	fOut := reflect.MakeFunc(destT, func(args []reflect.Value) (results []reflect.Value) {
		callValues := append([]reflect.Value(nil), values...)
		for _, pi := range placeholders {
			callValues[pi.toArg] = args[pi.fromArg]
		}
		return srcV.Call(callValues)
	})
	return fOut.Interface().(Dest)
}
