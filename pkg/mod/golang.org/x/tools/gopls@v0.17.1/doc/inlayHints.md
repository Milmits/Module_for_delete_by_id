# Gopls: Inlay hints

Inlay hints are helpful annotations that the editor can optionally
display in-line in the source code, such as the names of parameters in
a function call. This document describes the inlay hints available
from `gopls`.

<!-- This portion is generated by doc/generate from golang.AllInlayHints. -->
<!-- BEGIN Hints: DO NOT MANUALLY EDIT THIS SECTION -->
## **assignVariableTypes**

`"assignVariableTypes"` controls inlay hints for variable types in assign statements:
```go
	i/* int*/, j/* int*/ := 0, len(r)-1
```


**Disabled by default. Enable it by setting `"hints": {"assignVariableTypes": true}`.**

## **compositeLiteralFields**

`"compositeLiteralFields"` inlay hints for composite literal field names:
```go
	{/*in: */"Hello, world", /*want: */"dlrow ,olleH"}
```


**Disabled by default. Enable it by setting `"hints": {"compositeLiteralFields": true}`.**

## **compositeLiteralTypes**

`"compositeLiteralTypes"` controls inlay hints for composite literal types:
```go
	for _, c := range []struct {
		in, want string
	}{
		/*struct{ in string; want string }*/{"Hello, world", "dlrow ,olleH"},
	}
```


**Disabled by default. Enable it by setting `"hints": {"compositeLiteralTypes": true}`.**

## **constantValues**

`"constantValues"` controls inlay hints for constant values:
```go
	const (
		KindNone   Kind = iota/* = 0*/
		KindPrint/*  = 1*/
		KindPrintf/* = 2*/
		KindErrorf/* = 3*/
	)
```


**Disabled by default. Enable it by setting `"hints": {"constantValues": true}`.**

## **functionTypeParameters**

`"functionTypeParameters"` inlay hints for implicit type parameters on generic functions:
```go
	myFoo/*[int, string]*/(1, "hello")
```


**Disabled by default. Enable it by setting `"hints": {"functionTypeParameters": true}`.**

## **parameterNames**

`"parameterNames"` controls inlay hints for parameter names:
```go
	parseInt(/* str: */ "123", /* radix: */ 8)
```


**Disabled by default. Enable it by setting `"hints": {"parameterNames": true}`.**

## **rangeVariableTypes**

`"rangeVariableTypes"` controls inlay hints for variable types in range statements:
```go
	for k/* int*/, v/* string*/ := range []string{} {
		fmt.Println(k, v)
	}
```


**Disabled by default. Enable it by setting `"hints": {"rangeVariableTypes": true}`.**

<!-- END Hints: DO NOT MANUALLY EDIT THIS SECTION -->