package almanac

import "fmt"

func ExampleJulianDay() {
	var jd = julianDay(2022, 7, 12.5)
	fmt.Println(jd)

	// Output:
	// 2.459773e+06
}

// func ExampleTimeFromJD() {
//	var Time = timeFromJD(2459774.2261805558)
//	fmt.Println(Time)
//
//	// Output:
//	// {2022 7 13 17 25 42}
// }

func ExampleNewTime() {
	var t = NewTime(2022, 8, 11, 15, 34, 35)
	fmt.Println(t)
	// Output:
	// &{2022 8 11 15 34 35}
}

func ExampleNewDay() {
	var day = NewDay(&Time{year: 2022, month: 8, day: 12})
	fmt.Println(*day)

	// Output:
	// {8259 11 {2022 8 12 0 0 0} 1 5 32 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 丁酉 14 十五  望 8258.899803471168 {2022 8 12 9 35 43}  0 {0 0 0 0 0 0} 234 52 5 67 36 {[] [中元节、鬼节] [] false false false } 8259 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 14} {[] [] [] false false false }}
}

func ExampleNewMonth() {
	var month = NewMonth(2022, 8)
	fmt.Println(month)

	// Output:
	// &{2022 8 8248 {2022 8 1 12 0 0} 1 5 31 12038 8 2 2 {1949 9999 1948 当代 中国  公历纪元} [{8248 0 {2022 8 1 12 0 0} 1 1 0 0 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 7 丁未 丙戌 3 初四   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 223 41 -6 56 25 {[] [] [] false false false } 8248 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 3} {[] [建军节] [] false false false }} {8249 1 {2022 8 2 12 0 0} 1 2 0 0 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 7 丁未 丁亥 4 初五   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 224 42 -5 57 26 {[] [] [] false false false } 8249 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 4} {[] [] [] false false false }} {8250 2 {2022 8 3 12 0 0} 1 3 0 0 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 7 丁未 戊子 5 初六   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 225 43 -4 58 27 {[] [] [] false false false } 8250 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 5} {[] [] [] false false false }} {8251 3 {2022 8 4 12 0 0} 1 4 0 0 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 7 丁未 己丑 6 初七   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 226 44 -3 59 28 {[] [七夕(中国情人节,乞巧节,女儿节)] [] false false false } 8251 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 6} {[] [] [] false false false }} {8252 4 {2022 8 5 12 0 0} 1 5 0 0 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 7 丁未 庚寅 7 初八  上弦 8252.296201778412 {2022 8 5 19 6 31}  0 {0 0 0 0 0 0} 227 45 -2 60 29 {[] [] [] false false false } 8252 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 7} {[] [] [] false false false }} {8253 5 {2022 8 6 12 0 0} 1 6 0 0 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 7 丁未 辛卯 8 初九   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 228 46 -1 61 30 {[] [] [] false false false } 8253 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 8} {[] [] [] true false false }} {8254 6 {2022 8 7 12 0 0} 1 0 0 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 壬辰 9 初十 立秋  0 {0 0 0 0 0 0} 立秋 8254.353552472587 {2022 8 7 20 29 6} 229 47 0 62 31 {[] [立秋] [] false false false } 8254 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 9} {[] [] [] true false false }} {8255 7 {2022 8 8 12 0 0} 1 1 0 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 癸巳 10 十一   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 230 48 1 63 32 {[] [] [] false false false } 8255 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 10} {[] [] [中国男子节(爸爸节)] false false false }} {8256 8 {2022 8 9 12 0 0} 1 2 0 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 甲午 11 十二   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 231 49 2 64 33 {[] [] [] false false false } 8256 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 11} {[] [] [] false false false }} {8257 9 {2022 8 10 12 0 0} 1 3 0 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 乙未 12 十三   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 232 50 3 65 34 {[] [] [侗族吃新节] false false false } 8257 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 12} {[] [] [] false false false }} {8258 10 {2022 8 11 12 0 0} 1 4 0 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 丙申 13 十四   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 233 51 4 66 35 {[] [] [] false false false } 8258 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 13} {[] [] [] false false false }} {8259 11 {2022 8 12 12 0 0} 1 5 0 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 丁酉 14 十五  望 8258.899803471168 {2022 8 12 9 35 43}  0 {0 0 0 0 0 0} 234 52 5 67 36 {[] [中元节、鬼节] [] false false false } 8259 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 14} {[] [] [] false false false }} {8260 12 {2022 8 13 12 0 0} 1 6 0 1 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 戊戌 15 十六   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 235 53 6 68 37 {[] [] [] false false false } 8260 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 15} {[] [] [] true false false }} {8261 13 {2022 8 14 12 0 0} 1 0 0 2 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 己亥 16 十七   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 236 54 7 69 38 {[] [] [] false false false } 8261 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 16} {[] [] [] true false false }} {8262 14 {2022 8 15 12 0 0} 1 1 0 2 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 庚子 17 十八   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 237 55 8 70 39 {[] [末伏] [] false false false } 8262 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 17} {[] [] [] false false false }} {8263 15 {2022 8 16 12 0 0} 1 2 0 2 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 辛丑 18 十九   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 238 56 9 71 40 {[] [] [] false false false } 8263 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 18} {[] [] [] false false false }} {8264 16 {2022 8 17 12 0 0} 1 3 0 2 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 壬寅 19 二十   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 239 57 10 72 41 {[] [] [] false false false } 8264 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 19} {[] [] [] false false false }} {8265 17 {2022 8 18 12 0 0} 1 4 0 2 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 癸卯 20 廿一   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 240 58 11 73 42 {[] [] [] false false false } 8265 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 20} {[] [] [] false false false }} {8266 18 {2022 8 19 12 0 0} 1 5 0 2 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 甲辰 21 廿二  下弦 8266.025057239154 {2022 8 19 12 36 4}  0 {0 0 0 0 0 0} 241 59 12 74 43 {[] [] [] false false false } 8266 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 21} {[] [] [] false false false }} {8267 19 {2022 8 20 12 0 0} 1 6 0 2 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 乙巳 22 廿三   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 242 60 13 75 44 {[] [] [] false false false } 8267 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 22} {[] [] [] true false false }} {8268 20 {2022 8 21 12 0 0} 1 0 0 3 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 丙午 23 廿四   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 243 61 14 76 45 {[] [] [] false false false } 8268 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 23} {[] [] [] true false false }} {8269 21 {2022 8 22 12 0 0} 1 1 0 3 5 31 12038 8 2 2 7 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 丁未 24 廿五   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 244 62 15 77 46 {[] [] [] false false false } 8269 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 24} {[] [] [] false false false }} {8270 22 {2022 8 23 12 0 0} 1 2 0 3 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 戊申 25 廿六 处暑  0 {0 0 0 0 0 0} 处暑 8269.969550661894 {2022 8 23 11 16 9} 245 63 16 78 47 {[] [处暑] [] false false false } 8270 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 25} {[] [] [] false false false }} {8271 23 {2022 8 24 12 0 0} 1 3 0 3 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 己酉 26 廿七   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 246 64 17 79 48 {[] [] [] false false false } 8271 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 26} {[] [] [] false false false }} {8272 24 {2022 8 25 12 0 0} 1 4 0 3 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 庚戌 27 廿八   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 247 65 18 80 49 {[] [] [] false false false } 8272 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 27} {[] [] [] false false false }} {8273 25 {2022 8 26 12 0 0} 1 5 0 3 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 七 29  八 8 戊申 辛亥 28 廿九   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 248 66 19 81 50 {[] [] [] false false false } 8273 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 28} {[] [] [] false false false }} {8274 26 {2022 8 27 12 0 0} 1 6 0 3 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 八 30  九 8 戊申 壬子 0 初一  朔 8274.178545468692 {2022 8 27 16 17 6}  0 {0 0 0 0 0 0} 249 67 20 82 51 {[] [] [] false false false } 8274 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 29} {[] [] [] true false false }} {8275 27 {2022 8 28 12 0 0} 1 0 0 4 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 八 30  九 8 戊申 癸丑 1 初二   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 250 68 21 83 52 {[] [] [] false false false } 8275 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 1 30} {[] [] [] true false false }} {8276 28 {2022 8 29 12 0 0} 1 1 0 4 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 八 30  九 8 戊申 甲寅 2 初三   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 251 69 22 84 53 {[] [] [] false false false } 8276 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 2 1} {[] [] [] false false false }} {8277 29 {2022 8 30 12 0 0} 1 2 0 4 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 八 30  九 8 戊申 乙卯 3 初四   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 252 70 23 85 54 {[] [] [] false false false } 8277 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 2 2} {[] [] [] false false false }} {8278 30 {2022 8 31 12 0 0} 1 3 0 4 5 31 12038 8 2 2 8 {1949 9999 1948 当代 中国  公历纪元} {38 壬寅 壬寅 4720 八 30  九 8 戊申 丙辰 4 初五   0 {0 0 0 0 0 0}  0 {0 0 0 0 0 0} 253 71 24 86 55 {[] [] [] false false false } 8278 8011 7996 0 [8025 8040 8055 8070 8085 8099 8114 8130 8145 8160 8176 8192 8207 8223 8239 8254 8270 8285 8301 8316 8331 8346 8361 8376 8391] [8008 8038 8067 8097 8126 8156 8185 8215 8245 8274 8304 8333 8363 8392 8422] [30 29 30 29 30 29 30 30 29 30 29 30 29 30] [十一 十二 正 二 三 四 五 六 七 八 九 十 十一 十二]} {1444 2 3} {[] [] [] false false false }}]}
}

// func ExampleFormatCal() {
//	var Month = NewMonth(2022, 8)
//	fmt.Println(Month.FormatCal())
//
//	// Output:
//	// 2022年   8月
//	// Sun Mon Tue Wed Thu Fri Sat
//	//       1   2   3   4   5   6
//	//   7   8   9  10  11  12  13
//	//  14  15  16  17  18  19  20
//	//  21  22  23  24  25  26  27
//	//  28  29  30  31
// }