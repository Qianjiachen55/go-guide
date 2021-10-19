/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import "github.com/Qianjiachen55/go-guide/cobra/cmd"

func main() {
	cmd.Execute()
}

func demo(num int) int{
	//     if num == 1 {
	//         return 1
	//     }
	//     if num == 0 {
	//         return 0
	//     }
	//     if num == 2 {
	//         return 2
	//     }
	//     return demo(num - 1) + demo(num - 2)
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 2
	}
	first := 1
	second := 2
	var ans int
	var temp int
	for i := 3;i<=num;i++ {
		ans = first + second
		first.= second
		second = ans
	}
	return ans
}