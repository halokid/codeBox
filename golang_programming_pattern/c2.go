package main

import "log"

type Employee struct {
  Name    string
  Age     int
  Vacation   int
  Salary    int
}

// TODO: 这个一个处理好的map， map的最终结果就是把数据格式化、序列化, 而这个map的过程实际上
// TODO: 就是所谓的`业务逻辑`, map-reduce和map-filter的模式就是把`业务逻辑` 和 `控制逻辑`
// TODO: 分开的编程模式
func MapLogic() []Employee {
  var list = []Employee {
    {"Hao", 44, 0, 8000},
    {"Bob", 34, 10, 5000},
    {"Alice", 23, 5, 9000},
    {"Jack", 26, 0, 4000},
    {"Tom", 48, 9, 7500},
    {"Marry", 29, 0, 6000},
    {"Mike", 32, 8, 4000},
  }
  return list
}
//var list = []Employee {
//  {"Hao", 44, 0, 8000},
//  {"Bob", 34, 10, 5000},
//  {"Alice", 23, 5, 9000},
//  {"Jack", 26, 0, 4000},
//  {"Tom", 48, 9, 7500},
//  {"Marry", 29, 0, 6000},
//  {"Mike", 32, 8, 4000},
//}

// TODO: below is all the reduce & filter functions
// TODO: reduce做数据聚合， filter做数据过滤， 都是针对数据做处理
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
  count := 0
  for i, _ := range list {
    if fn(&list[i]) {
      count += 1
    }
  }
  return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
  var newList []Employee
  for i, _ := range list {
    if fn(&list[i]) {
      newList = append(newList, list[i])
    }
  }
  return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
  var sum = 0
  for i, _ := range list {
    sum += fn(&list[i])
  }
  return sum
}

func main() {
  // TODO: 统计有多少员工大于40岁
  old := EmployeeCountIf(list, func(e *Employee) bool {
    return e.Age > 40
  })
  log.Printf("old people: %d", old)

  // TODO: 统计多少员工薪水大于6000
  highPay := EmployeeCountIf(list, func(e *Employee) bool {
    return e.Salary > 6000
  })
  log.Printf("High salary people: %d", highPay)

  // TODO: 列出有没有休假的员工
  noVacation := EmployeeFilterIn(list, func(e *Employee) bool {
    return e.Vacation == 0
  })
  log.Printf("people no vacation: %v", noVacation)
}








