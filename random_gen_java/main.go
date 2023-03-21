package main

import (
	"bufio"
	"fmt"
	"github.com/pochard/commons/randstr"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var fileName string = "./tmp/gen_java.java"

var stringTmpl = []string{
	"String.valueOf(%d).equals(\"%d\")", "\"%s\".getBytes().toString()",
}

var floatTmpl = []string{
	"Float.valueOf(%f).equals(\"%f\")",
	"Float.valueOf(%f).intValue()",
	"new Float(%f).intValue()",
	"new Float(%f).floatValue()",
	"new Float(%f).doubleValue()",
	"new Float(%f).toString()",
	"new Float(%f).hashCode()",
}

var doubleTmpl = []string{
	"Double.valueOf(%2f).equals(\"%2f\")",
	"Double.valueOf(%2f).intValue()",
	"new Double(%2f).intValue()",
	"new Double(%2f).floatValue()",
	"new Double(%2f).doubleValue()",
	"new Double(%2f).toString()",
	"new Double(%2f).hashCode()",
}
var intTmpl = []string{
	"Integer.valueOf(%d).equals(\"%d\")",
	"Integer.valueOf(%d).intValue()",
	"new Integer(%d).intValue()",
	"new Integer(%d).floatValue()",
	"new Integer(%d).doubleValue()",
	"new Integer(%d).toString()",
	"new Integer(%d).hashCode()",
	"Integer num%d = %d;\n        num%d.equals(%d);",
}

var ifTmpl = []string{
	"if (%e) {\n%@\n%@\n%@\n}",
	"if (%e) {\n%@\n%@\n%@\n%@\n}",
	"if (%e) {\n%@\n%@\n%@\n%@} else {\n	%@\n%@\n}",
	"if (%e) {\n%@\n%@\n%@\n%@} else {\n	%@\n%@\n%@\n%@\n}",
}

func genIf() string {
	index := rand.Intn(len(ifTmpl))
	strTmpl := ifTmpl[index]
	out := replaceParamEx(strTmpl)
	//fmt.Println(out)
	return out
}

func main() {
	//w := genWriter()
	//defer w.Flush()
	rand.Seed(time.Now().UnixNano())
	//genString()
	//genFloat()
	//genDouble()
	//genInteger()
	callStr := genIf()
	fmt.Println(callStr)
}

func genWriter() *bufio.Writer {
	newPath := filepath.Dir(fileName)
	_, err := os.Stat(newPath)
	if os.IsNotExist(err) {
		os.MkdirAll(newPath, fs.ModePerm)
	}
	//fi, err := os.Open(fileName)
	fiW, err := os.Create(fileName)
	defer fiW.Close()
	// 创建 Reader
	w := bufio.NewWriter(fiW)
	return w
}

var intClassTmpl = []string{
	"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(float f) {\n        return new Float(f).intValue();\n    }\n}",
	"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(double f) {\n        return new Double(f).intValue();\n    }\n}",
	//"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(float f) {\n        return new Float(f).intValue();\n    }\n}",
}

var floatClassTmpl = []string{
	"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(float f) {\n        return new Float(f).floatValue();\n    }\n}",
	"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(double f) {\n        return new Double(f).floatValue();\n    }\n}",
	//"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(float f) {\n        return new Float(f).intValue();\n    }\n}",
}

var doubleClassTmpl = []string{
	"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(float f) {\n        return new Float(f).doubleValue();\n    }\n}",
	"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(double f) {\n        return new Double(f).doubleValue();\n    }\n}",
	//"package prePrePath.%s;\n\npublic final class %s {\n    public static int %s(float f) {\n        return new Float(f).intValue();\n    }\n}",
}

func genIntegerClass() string {
	//package
	str1 := randstr.RandomAlphabetic(rand.Intn(10) + 10)
	//newPath := filepath.Dir(str1)
	_, err := os.Stat(str1)
	if os.IsNotExist(err) {
		os.MkdirAll(str1, fs.ModePerm)
	}
	//java file
	str2 := randstr.RandomAlphabetic(rand.Intn(10) + 11)
	fiW, err := os.Create(fmt.Sprintf("./%s/%s.java", str1, str2))
	defer fiW.Close()
	str3 := randstr.RandomAlphabetic(rand.Intn(10) + 12)
	index := rand.Intn(len(intClassTmpl))
	str := intClassTmpl[index]
	content := fmt.Sprintf(str, str1, str2, str3)
	fiW.WriteString(content)

	return fmt.Sprintf("%s.%s.%s(%ff)", str1, str2, str3, rand.Float32()*9999)
}

func genFloatClass() string {
	//package
	str1 := randstr.RandomAlphabetic(rand.Intn(10) + 10)
	//newPath := filepath.Dir(str1)
	_, err := os.Stat(str1)
	if os.IsNotExist(err) {
		os.MkdirAll(str1, fs.ModePerm)
	}
	//java file
	str2 := randstr.RandomAlphabetic(rand.Intn(10) + 11)
	fiW, err := os.Create(fmt.Sprintf("./%s/%s.java", str1, str2))
	defer fiW.Close()
	str3 := randstr.RandomAlphabetic(rand.Intn(10) + 12)
	index := rand.Intn(len(floatClassTmpl))
	str := floatClassTmpl[index]
	content := fmt.Sprintf(str, str1, str2, str3)
	fiW.WriteString(content)

	return fmt.Sprintf("%s.%s.%s(%ff)", str1, str2, str3, rand.Float32()*9999)
}

func genDoubleClass() string {
	//package
	str1 := randstr.RandomAlphabetic(rand.Intn(10) + 10)
	//newPath := filepath.Dir(str1)
	_, err := os.Stat(str1)
	if os.IsNotExist(err) {
		os.MkdirAll(str1, fs.ModePerm)
	}
	//java file
	str2 := randstr.RandomAlphabetic(rand.Intn(10) + 11)
	fiW, err := os.Create(fmt.Sprintf("./%s/%s.java", str1, str2))
	defer fiW.Close()
	str3 := randstr.RandomAlphabetic(rand.Intn(10) + 12)
	index := rand.Intn(len(doubleClassTmpl))
	str := doubleClassTmpl[index]
	content := fmt.Sprintf(str, str1, str2, str3)
	fiW.WriteString(content)

	return fmt.Sprintf("%s.%s.%s(%ff)", str1, str2, str3, rand.Float32()*9999)
}

func genString() string {
	index := rand.Intn(len(stringTmpl))
	strTmpl := stringTmpl[index]
	out := replaceParam(strTmpl)
	//fmt.Println(out)
	return out
}

func genFloat() string {
	index := rand.Intn(len(floatTmpl))
	strTmpl := floatTmpl[index]
	out := replaceParam(strTmpl)
	//fmt.Println(out)
	return out
}

func genFloatByIndex(index int) string {
	//index := rand.Intn(len(floatTmpl))
	strTmpl := floatTmpl[index]
	out := replaceParam(strTmpl)
	//fmt.Println(out)
	return out
}

func genDouble() string {
	index := rand.Intn(len(doubleTmpl))
	strTmpl := doubleTmpl[index]
	out := replaceParam(strTmpl)
	//fmt.Println(out)
	return out
}

func genDoubleByIndex(index int) string {
	//index := rand.Intn(len(doubleTmpl))
	strTmpl := doubleTmpl[index]
	out := replaceParam(strTmpl)
	//fmt.Println(out)
	return out
}

func genIntegerNoLast(index int) string {
	//index := rand.Intn(len(intTmpl)) - 1
	strTmpl := intTmpl[index]
	out := replaceParam(strTmpl)
	//fmt.Println(out)
	return out
}

func genInteger() string {
	index := rand.Intn(len(intTmpl))
	strTmpl := intTmpl[index]
	out := replaceParam(strTmpl)
	//fmt.Println(out)
	return out
}

func replaceParam(tmpl string) string {
	ret := strings.ReplaceAll(tmpl, "%d", strconv.Itoa(rand.Intn(99999)))
	ret = strings.ReplaceAll(ret, "%f", fmt.Sprintf("%ff", rand.Float32()*99999))
	ret = strings.ReplaceAll(ret, "%2f", fmt.Sprintf("%fd", rand.Float32()*99999))
	ret = strings.ReplaceAll(ret, "%s", randstr.RandomAlphanumeric(len(ret)))

	return ret
}

func replaceParamEx(tmpl string) string {
	//"if (%e) {\n%@\n@\n@\n@}",
	index := rand.Intn(len(floatTmpl))
	var ret = ""
	if index < 2 {
		ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s == %s", genFloatByIndex(index), genDoubleByIndex(index)))
	} else if index < 3 {
		ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s >= %s", genFloatByIndex(index), genDoubleByIndex(index)))
	} else if index < 4 {
		ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s != %s", genFloatByIndex(index), genIntegerNoLast(index)))
	} else if index < 5 {
		ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s > %s && %s < %s", genFloatByIndex(index), genIntegerNoLast(index), genDoubleByIndex(index), genFloatByIndex(index)))
	} else {
		checkIfGenFile := rand.Intn(100)
		if checkIfGenFile > 96 {
			ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s < %s ", genDoubleClass(), genFloatClass()))
		} else if checkIfGenFile > 93 {
			ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s > %s ", genIntegerClass(), genFloatClass()))
		} else if checkIfGenFile > 60 {
			ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s >= %s", genDoubleByIndex(index), genIntegerNoLast(index)))
		} else {
			ret = strings.ReplaceAll(tmpl, "%e", fmt.Sprintf("%s > %s", genDoubleByIndex(index), genFloatByIndex(index)))
		}

	}
	ret = strings.ReplaceAll(ret, ";", "")
	for startIndex := strings.Count(tmpl, "%@"); startIndex > 0; startIndex-- {
		rate := rand.Intn(100)
		if rate < 20 {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genFloat()), 1)
		} else if rate < 40 {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genDouble()), 1)
		} else if rate < 60 {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genString()), 1)
		} else if rate < 65 {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genInteger()), 1)
		} else if rate < 80 {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genFloatClass()), 1)
		} else if rate < 85 {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genDoubleClass()), 1)
		} else if rate < 90 {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genIntegerClass()), 1)
		} else {
			ret = strings.Replace(ret, "%@", fmt.Sprintf("%s", genFloat()), 1)
		}
	}

	return ret
}
