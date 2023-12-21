import * as fs from "fs";


function run() {
    let file = fs.readFileSync("./src/day4/input.txt", {encoding: "utf-8"})

    let lines = file.split("\n")
    lines.pop()

    const cards = lines.map((line) => {
        const [, numbers] = line.split(":");
        const [left, right] = numbers.split("|");

        const numberFound: number[] = []
        const winningNumbers = left
            .trim()
            .split(" ")
            .map(number => parseInt(number))
            .filter(number => !Number.isNaN(number))

        right
            .trim()
            .split(" ")
            .map(number => parseInt(number))
            .filter(number => {
                    if (Number.isNaN(number)) return

                    let include = winningNumbers.includes(number)
                    if (include) {
                        numberFound.push(number)
                    }
                }
            )


        if (numberFound.length === 0) return 0
        return Math.pow(2, numberFound.length - 1)
    })

    const result = cards.reduce((acc, curr) => acc + curr, 0)

    console.log(result)

}


run()
