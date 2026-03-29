// описание пакета
package additor

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// описание функции
func Add[N Number](s1, s2 N) N {
	return s1 + s2
}
