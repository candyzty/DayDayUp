/**
* @Author: redgr
* @Description:
* @File:  person
* @Version: 1.0.0
* @Date: 2020/7/26 22:12
 */

package models

type Person struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Grade float64 `json:"grade"`
}
