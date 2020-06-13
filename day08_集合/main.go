package main

import "fmt"

//1.声明
//func main() {
//	countryCapitalMap := make(map[string]string)
//	countryCapitalMap["France"] = "Paris"
//	countryCapitalMap["Italy"] = "Rome"
//	countryCapitalMap["Japan"] = "Tokyo"
//	countryCapitalMap["India"] = "New Delhi"
//
//	for country := range countryCapitalMap{
//		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
//	}
//}

//2. delete
//func main() {
//	countryCapitalMap := make(map[string]string)
//		countryCapitalMap["France"] = "Paris"
//		countryCapitalMap["Italy"] = "Rome"
//		countryCapitalMap["Japan"] = "Tokyo"
//		countryCapitalMap["India"] = "New Delhi"
//
//		for country := range countryCapitalMap{
//			fmt.Println("Capital of", country, "is", countryCapitalMap[country])
//		}
//		delete(countryCapitalMap, "France")
//		fmt.Println("删除元素后")
//	for country := range countryCapitalMap{
//		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
//	}
//}

//3. ok-idiom
//func main() {
//	m := make(map[string]int)
//	m["a"] = 10
//	x, ok := m["b"]
//	fmt.Println(x, ok)
//	x, ok = m["a"]
//	fmt.Println(x, ok)
//}

//4. len()
//func main() {
//	m := make(map[string]int)
//	m["a"] = 10
//	m["b"] = 20
//	fmt.Printf("length of %v is %d", m, len(m))
//}

//5.引用类型
func main() {
	personSalary := map[string]int{
		"Steve": 12000,
		"Jame":  15000,
	}
	personSalary["Mike"] = 9000
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["Mike"] = 890000000
	fmt.Println("New person salary", personSalary)

}
