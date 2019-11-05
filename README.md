# GildedRose-Refactoring-Kata
GildedRose-Refactoring-Kata

### 重构心得
* 减少使用下面判断方法，避免带来的臃肿
```
if a == name1 && a == name2 {

}

或者
switch a {
    b:
    c:
    d:
}
等方法
```
采用设计模式当中的策略模式
由调用方控制
> 策略模式的讲解 https://www.jianshu.com/p/3bcf55cf83d3

* 减少重复的if，比如代码片段中
```
if items[i].sellIn < 11 {
	if items[i].quality < 50 {
		items[i].quality = items[i].quality + 1
	}
}
if items[i].sellIn < 6 {
	if items[i].quality < 50 {
		items[i].quality = items[i].quality + 1
	}
}
```