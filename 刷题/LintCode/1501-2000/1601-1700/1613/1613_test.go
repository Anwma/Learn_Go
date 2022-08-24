package main

func HighestFrequency(ipLines []string) string {
	dist := map[string]int{}

	for _, v := range ipLines {
		if _, ok := dist[v]; ok {
			dist[v] = dist[v] + 1
		} else {
			dist[v] = 1
		}
	}

	var max int = 0
	var re string
	for k, v := range dist {
		if v > max {
			max = v
			re = k
		}
	}
	return re
}
