package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type foodItem struct {
	menuId int
	count  int
}

func isDuplicateFoodmenuInEater(foodmenuId int, foodMenus []int) bool {
	for _, menuId := range foodMenus {
		if menuId == foodmenuId {
			return true
		}
	}
	return false
}

func calculateTop3Menus(eaters, menus []int) string {
	// Create a map to store the number of times each menu item has been consumed.
	counts := make(map[int]int)
	eaterMenuMap := make(map[int][]int)

	for i := 0; i < len(menus); i++ {
		eaterId := eaters[i]
		foodmenuId := menus[i]

		if isDuplicateFoodmenuInEater(foodmenuId, eaterMenuMap[eaterId]) {
			logrus.Errorf("Duplicate Food menu: %d, found in eater: %d", foodmenuId, eaterId)
			continue

		} else {
			eaterMenuMap[eaterId] = append(eaterMenuMap[eaterId], foodmenuId)
		}
		// Increment the count for the foodmenu_id.
		counts[foodmenuId]++
	}

	// Find the top 3 menu items consumed.
	menues := make([]foodItem, 0)
	logrus.Debug("foodItem Count Map", counts)
	for menuId, count := range counts {
		menues = append(menues, foodItem{
			menuId: menuId,
			count:  count,
		})
	}

	if len(menues) < 3 {
		logrus.Error("Please enter atleast 3 or more distinct menues")
		return ""
	}
	sort.Slice(menues, func(i, j int) bool {
		if menues[i].count == menues[j].count {
			return menues[i].menuId < menues[j].menuId
		} else {
			return menues[i].count > menues[j].count
		}
	})

	// Print the top 3 menu items consumed.
	output := fmt.Sprintf("First:%d Second:%d Third:%d", menues[0].menuId, menues[1].menuId, menues[2].menuId)
	return output
}

func main() {
	// Open the log file.
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)

	menus := make([]int, 0)
	eaters := make([]int, 0)
	// Iterate over the lines in the file.
	for scanner.Scan() {
		// Split the line into two parts: eater_id and foodmenu_id.
		line := scanner.Text()
		parts := strings.Split(line, ",")
		eater_id, err := strconv.Atoi(parts[0])
		if err != nil {
			logrus.Fatal(err)
		}
		foodmenuId, err := strconv.Atoi(parts[1])
		if err != nil {
			logrus.Fatal(err)
		}
		menus = append(menus, foodmenuId)
		eaters = append(eaters, eater_id)
	}

	output := calculateTop3Menus(eaters, menus)
	logrus.Info(output)
}
