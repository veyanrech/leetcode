package hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMinimumCost' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER edgeDeviceCost
 *  2. INTEGER inputPeripheralCost
 *  3. INTEGER bundleCost
 *  4. INTEGER x
 *  5. INTEGER y
 */

func getMinimumCost(edgeDeviceCost int32, inputPeripheralCost int32, bundleCost int32, x int32, y int32) int64 {
	// Write your code here
	var totalDevices int32 = x + y

	maxEdgeCost := edgeDeviceCost * x
	maxPeriCost := inputPeripheralCost * y

}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	edgeDeviceCostTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	edgeDeviceCost := int32(edgeDeviceCostTemp)

	inputPeripheralCostTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	inputPeripheralCost := int32(inputPeripheralCostTemp)

	bundleCostTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	bundleCost := int32(bundleCostTemp)

	xTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	x := int32(xTemp)

	yTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	y := int32(yTemp)

	result := getMinimumCost(edgeDeviceCost, inputPeripheralCost, bundleCost, x, y)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
