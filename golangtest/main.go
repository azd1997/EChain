/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2019/7/22 10:48
* @Description: The file is for
***********************************************************************/

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

type Circle struct {
	Vertex
	R float64
}

type Cylinder struct {
	Circle
	H float64
}

func ToCircle(u *Vertex, r *Cylinder) *Circle {
	if &r.Vertex == u {
		return &r.Circle
	}
	return nil
}

func main() {
	cy := Cylinder{Circle{}, 10}
	//cy := Cylinder{X: 3} 编译错误
	cy.X = 3
	fmt.Println(cy)  //{{{3 0} 0} 10}
	v := cy.Vertex
	v.Y = 4
	fmt.Println(cy, v) //cy = ?   // {{{3 0} 0} 10} {3 4}  这里v取得不是指针，而是复制的cy.Vertex的值
	u := &(cy.Vertex)
	u.Y = 4
	//c := Circle(u) 编译错误
	fmt.Println(cy, u) //cy = ?  // {{{3 4} 0} 10} &{3 4}  u取的是指针，所以cy的值也跟着改变了
	c := ToCircle(u, &cy)
	c.R = 5
	fmt.Println(cy, c, c.Vertex) //cy = ?   // {{{3 4} 5} 10} &{{3 4} 5} {3 4}
}
//---------------------
//作者：pmlpml
//来源：CSDN
//原文：https://blog.csdn.net/pmlpml/article/details/78326769
//版权声明：本文为博主原创文章，转载请附上博文链接！
