package main

import (
	"encoding/json"
	"fmt"
)

//结构体内标识符首字母大写，对外可见
type Student struct {
	ID int
	Name string
}
type class struct {
	Title string
	StudentArr []Student
}

func newStudent(id int,name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}

func main()  {
	c1 := class{
		Title:      "101",
		StudentArr: make([]Student,0,20),
	}
	for i :=0;i<10;i++{
		tempStu := newStudent(i,fmt.Sprintf("stu%02d",i))
		c1.StudentArr = append(c1.StudentArr, tempStu)
	}

	fmt.Println(c1)
	//json 序列化
	data,err := json.Marshal(c1)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Printf("%s\n",data)
	c2 := class{}
	json.Unmarshal([]byte(data),&c2)
	fmt.Println(c2)
}