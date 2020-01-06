package golang_in_action

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/spf13/cast"
  "math/rand"
  "runtime"
  "testing"
  "time"
)


var Db *gorm.DB

/**
func init()  {
  var err error
  Db, err = gorm.Open("mysql",  "xxx:xxxxx@tcp(8.8.8.8:33061)/ocx?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic(err)
  }
}
*/

type Deploy struct {
  ID      int
  Name    string
}

func (d Deploy) TableName() string {
  return "deploy"
}

func GetDeploys() []Deploy {
  // 获取所有deploy信息
  var deploys []Deploy
  Db.Find(&deploys)
  return deploys
}

func AddDeploy(d *Deploy) error {
  // 验证 FirstResponse这个父协程退出(return)之后， 子协程runTask还会不会执行?
  if err := Db.Create(d).Error; err != nil {
    return err
  }
  return nil
}


func runTask(id int) string {
  //rand.Seed(time.Now().Unix())
  randI := rand.Intn(30)
  fmt.Println("睡眠 ", randI, "验证下FirstResponse函数这个协程已经摧毁了之后，其他的gor还会不会跑这个addDb")
  time.Sleep(time.Duration(randI) * time.Second)
  //time.Sleep(10 * time.Second)
  d := Deploy{Name: "xxxx" + cast.ToString(randI)}
  AddDeploy(&d)
  return fmt.Sprintf("the result is from %d", id)
}

func FirstResponse() string {
  numOfRunner := 10
  //ch := make(chan string)

  // fixme: 起了一个有长度的 buffer channel， 这样假如一旦写入了 ch <-ret， 这样一旦有其中一个 gor跑完了 runTask函数（也就是说函数有返回）， 那么就会马上触发   return <-ch（原来是一直阻塞的），则整个 FirstResponse函数都会返回， 则会退出 gor 的运行时， 所有的 gor都会退出
  ch := make(chan string, numOfRunner)
  for i := 0; i < numOfRunner; i++ {
    fmt.Println("runtime gor num:", runtime.NumGoroutine())
    go func(i int) {
      ret := runTask(i)
      ch <-ret
    }(i)
  }
  // fixme:
  time.Sleep(3 * time.Second)
  return <-ch
}

func TestFirstResponse(t *testing.T) {
  // fixme: 用了 buffer channel之后， 这里会输出2 ，所以证明必要一个channel 返回， 其他的gor都会退出，2是最少的协程数量， 因为main func一个，  返回的gor一个， 一共2个
  t.Log("Before:", runtime.NumGoroutine())
  t.Log(FirstResponse())
  //time.Sleep(3 * time.Second)
  time.Sleep(15 * time.Second)
  fmt.Println("总的main gor等待 15秒, 调用的函数等待 3 秒, 则cancel掉的协程只要执行时间不超过18秒的都会执行 AddDeploy 函数")
  t.Log("After:", runtime.NumGoroutine())
  t.Log("before  和  after 的协程数是一样的， 证明没有协程泄漏, 也就是说没有往close的 channel写东西， 也没有协程在做无意义的等待")
}
