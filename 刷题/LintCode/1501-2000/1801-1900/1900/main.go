package main

import (
	"strconv"
)

func GeneSimilarity(gene1 string, gene2 string) string {
	// write your code here
	//1.由传入的gene1 gene2 生成两个新的字符串
	slicegene1, slicegene2 := []byte(gene1), []byte(gene2)
	//创建两个空切片 接收生成后的值
	tmp1, tmp2 := []byte{}, []byte{}
	//外层套一层结束
	for i := 0; i < len(gene1); i += 2 {
		//index out of range [4] with length 4
		for j := 0; j < int(slicegene1[i])-48; j++ {
			tmp1 = append(tmp1, slicegene1[i+1])
		}
	}

	for i := 0; i < len(gene2); i += 2 {
		for j := 0; j < int(slicegene2[i])-48; j++ {
			tmp2 = append(tmp2, slicegene2[i+1])
		}
	}

	//2.比对这两个生成的字符串 tmp1 tmp2
	count := 0
	for i := 0; i < len(tmp1); i++ {
		if tmp1[i] == tmp2[i] {
			count++
		}
	}

	//返回 相同的字符个数 / 总字符个数
	res := strconv.Itoa(count) + "/" + strconv.Itoa(len(tmp1))
	return res
}
