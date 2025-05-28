/*
735.中 小行星碰撞

给定一个整数数组 asteroids，表示在同一行的小行星。数组中小行星的索引表示它们在空间中的相对位置。

对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。

找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。

示例 1：

输入：asteroids = [5,10,-5]
输出：[5,10]
解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
示例 2：

输入：asteroids = [8,-8]
输出：[]
解释：8 和 -8 碰撞后，两者都发生爆炸。
示例 3：

输入：asteroids = [10,2,-5]
输出：[10]
解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。

提示：

2 <= asteroids.length <= 104
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0
*/
package leetcode

import (
	"fmt"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	asteroids := []int{5, 10, -5}

	fmt.Println(asteroidCollision(asteroids))
}

func TestAsteroidCollision2(t *testing.T) {
	asteroids := []int{-5, -10, -5}

	fmt.Println(asteroidCollision(asteroids))
}

func asteroidCollision(asteroids []int) (st []int) {
	for _, aster := range asteroids {
		alive := true
		for alive && aster < 0 && len(st) > 0 && st[len(st)-1] > 0 { // aster < 0 && len(st) > 0 && st[len(st)-1] > 0 => 会碰撞
			alive = st[len(st)-1] < -aster // aster 是否存在
			if st[len(st)-1] <= -aster {   // 栈顶行星爆炸
				st = st[:len(st)-1]
			}
		}
		if alive {
			st = append(st, aster)
		}
	}
	return
}
