package xfn

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ndkimhao/go-xtd/generics"
	"github.com/ndkimhao/go-xtd/xmath"
)

type Placeholder int

func (p Placeholder) String() string {
	return fmt.Sprint("Placeholder(", int(p), ")")
}

var (
	P1 Placeholder = 1
	P2 Placeholder = 2
	P3 Placeholder = 3
	P4 Placeholder = 4
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
		return fmt.Sprint(i, "th")
	}
}

func Bind[Dest any](src any, args ...any) Dest {
	srcV := reflect.ValueOf(src)
	srcT := srcV.Type()
	if srcT.Kind() != reflect.Func {
		panic(fmt.Sprintf("invalid type: src [%s] is not a function type", srcT))
	}
	destT := generics.TypeOf[Dest]()
	if destT.Kind() != reflect.Func {
		panic(fmt.Sprintf("invalid type: dest [%s] is not a function type", destT))
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
	destIdx := 0
	for i, a := range args {
		if placeholder, ok := a.(Placeholder); ok {
			destPH := int(placeholder) - 1
			if destPH < 0 || destIn <= destPH {
				panic(fmt.Sprintf("invalid placeholder: "+
					"dest [%s] has %d inputs but provided %s at %s bind argument",
					destT, destIn, placeholder, _ith(i+1)))
			}
			if srcT.In(i) != destT.In(destPH) {
				panic(fmt.Sprintf("type mismatch: "+
					"src [%s] %s input is [%s] but provided [%s] via %s at %s bind argument",
					srcT, _ith(i+1), srcT.In(i), destT.In(destPH), placeholder, _ith(i+1)))
			}
			placeholders = append(placeholders, phInfo{fromArg: destPH, toArg: i})
			destIdx = xmath.Max(destIdx, destPH+1)
			continue
		}
		valA := reflect.ValueOf(a)
		if !valA.CanConvert(srcT.In(i)) {
			panic(fmt.Sprintf("type mismatch: "+
				"src [%s] %s input is [%s] but provided [%s] at %s bind argument",
				srcT, _ith(i+1), srcT.In(i), valA.Type(), _ith(i+1)))
		}
		values[i] = valA.Convert(srcT.In(i))
	}
	need := srcIn - len(args)
	have := destIn - destIdx
	if need > have {
		missing := need - have
		var sb strings.Builder
		for i := srcIn - missing; i < srcIn; i++ {
			sb.WriteString(srcT.In(i).String())
			if i < srcIn-1 {
				sb.WriteString(", ")
			}
		}
		panic(fmt.Sprintf("not enough inputs: "+
			"dest [%s] needs %d more inputs [%s]", destT, missing, sb.String()))
	}
	for i := len(args); i < srcIn; i++ {
		if srcT.In(i) != destT.In(destIdx) {
			panic(fmt.Sprintf("type mismatch: "+
				"src [%s] %s input is [%s] but provided [%s] via %s dest input [%s]",
				srcT, _ith(i+1), srcT.In(i), destT.In(destIdx), _ith(destIdx+1), destT))
		}
		placeholders = append(placeholders, phInfo{fromArg: destIdx, toArg: i})
		destIdx++
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
