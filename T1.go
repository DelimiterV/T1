package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var numbers []float64

const MAX = 11

func fill(tstr string) {
	var m float64
	var err error
	numbers = make([]float64, 0)
	ut := strings.Split(tstr, " ")
	for i := 0; i < len(ut); i++ {
		m, err = strconv.ParseFloat(ut[i], 6)
		if err == nil {
			numbers = append(numbers, m)
		} else {
			fmt.Println("<", ut[i], ">", err)
		}
	}
}

type rls struct {
	v [MAX]uint8
}
type myrules struct {
	n   int
	rez []rls
}

var krules []myrules
var grules myrules
var urules rls
var frules []rls

func mgen(cnt int, n int, prev uint8) {
	if cnt < n-1 {
		if prev != 0 {
			urules.v[cnt] = 0
			mgen(cnt+1, n, 0)
		}
		urules.v[cnt] = 1
		mgen(cnt+1, n, 1)
		urules.v[cnt] = 2
		mgen(cnt+1, n, 2)
		urules.v[cnt] = 3
		mgen(cnt+1, n, 3)
		urules.v[cnt] = 4
		mgen(cnt+1, n, 4)
	} else {
		frules = append(frules, urules)
	}
}
func generateRules(nelem int) {
	var mymax int
	if nelem < MAX {
		mymax = nelem
	}
	krules = make([]myrules, 0)
	for i := 2; i < mymax+1; i++ {
		grules.rez = make([]rls, 0)
		frules = make([]rls, 0)
		grules.n = i
		mgen(0, i, 7)
		grules.rez = append(grules.rez, frules[:]...)
		krules = append(krules, grules)
	}
}
func rCalculate(zdata []float64, zoper []uint8) (int, float64, string) {
	e := 0
	f := 0
	ez := 0
	ef := 0
	var uzstr string
	var r float64
	var data []float64
	var oper []uint8
	oper = make([]uint8, len(zdata)-1)
	copy(oper, zoper[:len(zdata)-1])
	data = make([]float64, len(zdata))
	copy(data, zdata)
	for i := 0; i < 3 && e == 0 && ez == 0; i++ {
		if len(data) > 1 {
			mlen := len(data)
			ef = 0
			for j := 0; j < mlen-1 && e == 0 && ez == 0 && ef == 0; j++ {
				f = 0
				switch {
				case oper[j] == 0:
					if i == 0 {
						if data[j] == 1.0 {
							r = 1
							f = 1
						} else {
							if data[j+1] == math.Trunc(data[j+1]) {
								if math.Trunc(data[j+1]) == 0 {
									r = 1
									f = 1
								} else {
									ku := int(math.Trunc(data[j+1]))
									if ku < 10 && ku > -10 {
										r = 1
										for oi := 0; oi < ku && e == 0; oi++ {
											r = r * data[j]
											if math.IsNaN(r) || math.IsInf(r, 0) {
												e = 1
											}
										}
										if e == 0 {
											f = 1
										}
									} else {
										e = 1
									}
								}
							} else {
								e = 1
							}
						}
					}
				case oper[j] == 1 || oper[j] == 2:
					if i == 1 {
						switch oper[j] {
						case 1:
							r = data[j] * data[j+1]
							//fmt.Println(data[j], "*", data[j+1], "=", r, ">", oper)
							if math.IsNaN(r) || math.IsInf(r, 0) {
								e = 1
							} else {
								f = 2
							}
						case 2:
							r = data[j] / data[j+1]
							//fmt.Println(data[j], "/", data[j+1], "=", r, ">", oper)
							if math.IsNaN(r) || math.IsInf(r, 0) {
								e = 1
							} else {
								f = 2
							}
						}
					}
				case oper[j] == 3 || oper[j] == 4:
					if i == 2 {
						switch oper[j] {
						case 3:
							r = data[j] + data[j+1]
							//fmt.Println("+")
							if math.IsNaN(r) || math.IsInf(r, 0) {
								e = 1
							} else {
								f = 2
							}
						case 4:
							r = data[j] - data[j+1]
							//fmt.Println("-")
							if math.IsNaN(r) || math.IsInf(r, 0) {
								e = 1
							} else {
								f = 2
							}
						}
					}
				case oper[j] < 0 || oper[j] > 4:
					fmt.Println("Error ", oper)
				}
				if f == 1 {
					data[j+1] = r
					oper = append(oper[:j], oper[j+1:]...)
					data = append(data[:j], data[j+1:]...)
					//fmt.Println(">>", data, " ", oper)
					ef = 1
				} else {
					if f == 2 {
						data[j+1] = r
						oper = append(oper[:j], oper[j+1:]...)
						data = append(data[:j], data[j+1:]...)
						//fmt.Println(">>", data, " ", oper)
						j--
					}
				}
				if len(data) < 2 {
					ez = 1
				}
				mlen = len(data)
			}
		}
	}
	if e != 1 {
		for ii := 0; ii < len(zoper); ii++ {
			switch zoper[ii] {
			case 0:
				uzstr += "^"
			case 1:
				uzstr += "*"
			case 2:
				uzstr += "/"
			case 3:
				uzstr += "+"
			case 4:
				uzstr += "-"
			}
			uzstr += " "
		}
	}
	return e, data[0], uzstr
}

type mctvt struct {
	r        float64
	activity string
}

//+ - * / ^
func myCalculate(fll []float64, cnt int) (rez []mctvt) {
	var ruu float64
	var zzstr string
	var r mctvt
	r.r = 0.0
	y := 0
	rez = make([]mctvt, 0)
	if len(fll) == 1 {
		r.r = fll[0]
		r.activity = ""
		rez = append(rez, r)
	} else {
		tl := len(fll)
		for k := 0; k < len(krules[tl-2].rez); k++ {
			y, ruu, zzstr = rCalculate(fll, krules[tl-2].rez[k].v[:tl-1])
			if y == 0 {
				r.r = ruu
				r.activity = zzstr
				rez = append(rez, r)
			}
		}
	}
	return
}

type kakadu struct {
	vhod      []float64
	myfll     []float64
	calculate []float64
	activity  string
	rez       float64
}

var md [MAX]kakadu

type kb struct {
	n    float64
	rstr string
	flag int
}

var mr []kb

var eps float64 = 0.000001

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func expand(curfld float64, curindex int) string {
	var rd string = ""
	var td []string
	for l := curindex; l >= 0; l-- {
		if floatEquals(curfld, mr[l].n) && mr[l].flag == 0 {
			td = strings.Split(mr[l].rstr, " ")
			for m := 0; m < len(td); m++ {
				if len(td[m]) > 1 {
					mfl, e := strconv.ParseFloat(td[m], 6)
					if e == nil {
						rets := expand(mfl, l-1)
						if rets != "" {
							td[m] = rets
						}
					}
				}
			}
			rd = strings.Join(td, " ")
			mr[l].flag = 1
		}
	}
	return rd
}

func display_decission(cnt int) {
	var kunit kb
	var tg string
	mr = make([]kb, 0)
	for i := 0; i < cnt+1; i++ {
		tznak := strings.Split(md[i].activity, " ")
		kunit.n = md[i].rez
		tg = "( "
		for z := 0; z < len(md[i].calculate); z++ {
			tg += strconv.FormatFloat(md[i].calculate[z], 'f', 6, 64)
			if z < len(md[i].activity)-1 {
				tg += " " + tznak[z] + " "
			}
		}
		tg += " )"
		kunit.rstr = tg
		kunit.n = md[i].rez
		kunit.flag = 0
		mr = append(mr, kunit)
	}
	myrd := ""
	mytd := strings.Split(mr[len(mr)-1].rstr, " ")
	for hgu := 0; hgu < len(mytd); hgu++ {
		if len(mytd[hgu]) > 1 {
			mfl, e := strconv.ParseFloat(mytd[hgu], 6)
			if e == nil {
				rets := expand(mfl, cnt-1)
				if rets != "" {
					mytd[hgu] = rets
				}
			} else {
				fmt.Println(e)
				fmt.Println("!!!", mytd[hgu])
			}
		}
	}
	myrd = strings.Join(mytd, " ")
	fmt.Println(myrd)

}
func divideMe2(fll []float64, cnt int) float64 {
	var mfll []float64
	var rezt []mctvt
	flag := 0.0
	md[cnt].vhod = fll
	for i := 0; i < len(fll)-1 && flag == 0; i++ {
		for j := i + 1; j < len(fll) && flag == 0; j++ {
			if len(fll[i:j+1]) > 1 {
				md[cnt].myfll = fll[i : j+1]
				md[cnt].calculate = fll[i : j+1]
				rezt = myCalculate(fll[i:j+1], cnt)
				if len(fll) == j-i+1 {
					for k := 0; k < len(rezt) && flag == 0; k++ {
						if floatEquals(rezt[k].r, myrez) {
							md[cnt].activity = rezt[k].activity
							md[cnt].rez = rezt[k].r
							display_decission(cnt)
							fmt.Println("last:", rezt[k].activity, fll, "=", rezt[k].r)
							flag = 1
						}
					}
				} else {
					for k := 0; k < len(rezt) && flag == 0; k++ {
						mfll = make([]float64, 0)
						mfll = append(mfll, fll[:i]...)
						mfll = append(mfll, rezt[k].r)
						mfll = append(mfll, fll[j+1:]...)
						md[cnt].myfll = mfll
						md[cnt].activity = rezt[k].activity
						md[cnt].rez = rezt[k].r
						flag = divideMe2(mfll[:], cnt+1)
					}
				}
			} else {
				fmt.Println("dangerous", fll[:])
			}
		}
	}
	return flag
}

var teststrings []string

func readStrings(filename string) bool {
	rez := true
	ustr := ""
	var l, k int
	r := ""
	dat, e := ioutil.ReadFile(filename)
	if e == nil {
		l = 0
		for i := 0; i < len(dat); i++ {
			switch dat[i] {
			case 0x0d:
				k = i
			case 0x0a:
				if dat[l] != '/' {
					r = string(dat[l:k])
					ustr = strings.Trim(r, " ")
					teststrings = append(teststrings, ustr)
				}

				l = i + 1
			}
		}
	} else {
		rez = false
	}
	return rez
}

var myrez float64

func main() {
	teststrings = make([]string, 0)
	if readStrings("DecideMe.txt") {

		for k := 0; k < len(teststrings); k++ {
			fill(teststrings[k])
			fmt.Println("=======Prepare=======")
			fmt.Println(teststrings[k])
			generateRules(len(numbers) - 1)
			myrez = numbers[len(numbers)-1]
			fmt.Println("---------Start---------")
			if divideMe2(numbers[:len(numbers)-1], 0) == 0 {
				fmt.Println("Не найдено")
			} else {
				fmt.Println("Success :", myrez)
			}
		}

	}
}
