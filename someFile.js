function addToArr(arr) {
    arr.push("Hello");
    return arr;
}
var numbers = Array(10).fill(1);
var strings = Array(10).fill("@");
var numStrings = Array(10).map(function (_, i) { return i % 2 === 0 ? "@" : 1; });
