package main

import (
	"math"
	"math/rand"
	"time"
)

// 1. 实现一个函数，输入为任意长度的int64数组，返回元素最大差值，例如输入arr=[5,8,10,1,3]，返回9。

func MaxDifference(arr []int64) int64 {
	val := math.Abs(float64(arr[1] - arr[0]))
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			temp := math.Abs(float64(arr[j] - arr[i]))
			if temp > val {
				val = temp
			}
		}
	}
	return int64(val)
}

//2. 实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。
func Shuffle(vals []int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]int, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

//3. 设计一个带失效时间的缓存数据结构，key和value都是string，并实现增删改查接口。

type Cache struct {
	data map[string]Value
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]Value)}
}

type Value struct {
	val         string
	expiredTime time.Time
}

func (c *Cache) Set(k, v string, et time.Time) {
	c.data[k] = Value{
		val:         v,
		expiredTime: et,
	}
}

func (c *Cache) Get(k string) string {
	v := c.data[k]
	if v.expiredTime.Before(time.Now()) {
		c.Del(k)
		return ""
	}

	return v.val
}

func (c *Cache) Del(k string) {
	delete(c.data, k)
}

// 4. 实现一个游戏算法：输入n和m，代表有n个选手围成一圈（选手编号为0至n-1），0号从1开始报数，报m的选手游戏失败从圆圈中退出，下一个人接着从1开始报数，如此反复，求最后的胜利者编号。
// 例如，n=3，m=2，那么失败者编号依次是1、0，最后的胜利者是2号。
// 这里考虑m，n都是正常的数据范围，其中：
// 1 <= n <= 10^5
// 1 <= m <= 10^6
// 算法要求考虑时间效率。

func findTheWinner(n, k int) int {
	winner := 1
	for i := 2; i <= n; i++ {
		winner = (k+winner)%i + 1
	}
	return winner
}
