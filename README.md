


## go rep for alfred JSON model

### useage

```shell
go get -u github.com/luanruisong/alfred
```


```go

func main(){
	
	mgr := alfred.NewMgr()
	mgr.AddItem(alfred.NewDefItem("I am title","I am sub title"))
	mgr.AddItem(alfred.NewDefItem("I am title2","I am sub title2"))
	
	mgr.PrintJson()
	
}
```