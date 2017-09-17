package main

import (
   "fmt"
)

type person struct{
   Name string
   Age int
}

func (self *person) Get(){
   fmt.Println(self.Name)
}

func (self *person) Get1(){
   fmt.Println(self.Age)
}

func (self *person) Set(name string){
   self.Name = name
}

func (self *person) Set1(age int){
   self.Age = age
}

func main() {
   Me := &person{}
   Me.Set("yux")
   Me.Set1(22)
   Me.Get()
   Me.Get1()
   fmt.Printf("%+v",Me)
}