package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i = i + 1 {
		for j := 0; j < len(podium); j = j + 1 {
			if i < j {
				tomp := podium[i]
				podium[i] = podium[j]
				podium[j] = tomp
			}
		}
	}
	return podium
}
