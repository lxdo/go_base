package model

// 定义Student结构体 (首字母大写) 可以在其他包直接使用

type Student struct {
	Name  string
	Score float64
}

// 定义student结构体 (首字母小写) 不可在其他包直接使用
// 我们通过工厂模式来解决

type student struct {
	Name  string
	Score float64
}

// 定义一个函数  传入结构体student的字段值 返回指向结构体变量数据空间的指针变量
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

// 如果Score字段首字母小写(score)，则在其他包不可直接访问
// 我们可以给结构体student绑定一个方法用来返回score字段的值
func (s *student) GetScore() float64 {
	return s.Score  // score首字母小写 return s.score  等价于 (*s).score
}
