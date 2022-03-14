package main

// has done it in Java
func search(nums []int, target int) int {
	return 0
}


//public int search(int[] bitonic, int n) {
//int maxIndex = findMaxIndexFromBitonicArray(bitonic);
//if (maxIndex < 0) {
//return -1;
//}
//if (bitonic[maxIndex] == n){
//return maxIndex;
//}
//if (n < bitonic[0]) {
//return binarySearch(bitonic, maxIndex+1, bitonic.length - 1, n);
//}else{
//return binarySearch(bitonic, 0, maxIndex-1, n);
//}
//}
//
//private int findMaxIndexFromBitonicArray(int[] bitonic) {
//if (bitonic.length <= 1) {
//return bitonic.length - 1;
//}
//if (bitonic.length == 2) {
//if (bitonic[1] > bitonic[0]){
//return 1;
//}else {
//return 0;
//}
//}
//int start = 0;
//int end = bitonic.length-1;
//while (true){
//int middle = start + (end-start)/2;
//if (middle == bitonic.length-1 || middle == 0){
//return middle;
//}
//if (bitonic[middle] > bitonic[middle-1] && bitonic[middle] > bitonic[middle+1]){
//return middle;
//}
//if (bitonic[middle] < bitonic[middle-1] && bitonic[middle] > bitonic[middle+1] ){
//end = middle-1;
//continue;
//}
//if (bitonic[middle] > bitonic[middle-1] && bitonic[middle] < bitonic[middle+1] ){
//start = middle+1;
//continue;
//}
//}
//}
//
//private int binarySearch(int[] bitonic, int start, int end, int n) {
//if (end < 0 || end < start) {
//return -1;
//}
//int middle = start + (end - start) / 2;
//if (n == bitonic[middle]) {
//return middle;
//}
//if (n > bitonic[middle]) {
//return binarySearch(bitonic, middle+1, end, n);
//}
//return binarySearch(bitonic, start, middle-1, n);
//}
//
//func main() {
//
//}
