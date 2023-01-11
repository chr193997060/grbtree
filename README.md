# grbtree 

grbtree 是使用go实现的红黑树

使用示例
```go
    import (
        "fmt"
        
        "github.com/chr193997060/grbtree"
    )

    tree := grbtree.NewRBTree()
    tree.Add(1000, 1)
    tree.Add(2000, 2)
    v, err := tree.Get(1000)
    if err != nil {
        fmt.Println(err.Error())
    }
    if v.(int) != 1 {
        fmt.Println(v)
    }
    tree.PrintTree(5)
```
