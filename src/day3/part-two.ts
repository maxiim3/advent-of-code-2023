// Path: day-3/index.ts

import * as fs from "fs";

type DataCell = string
type Matrix = DataCell[][]

const file = fs.readFileSync("./javascript/day-3/input-sample.txt", { encoding: "utf-8" });

const lines = file.split("\n")
lines.pop()

const matrix: Matrix = lines.map(line => line.split(""))

function isANumber(data: string): { isNumber: boolean, value: number | null } {
    const toNumber = parseInt(data);
    if (Number.isInteger(toNumber)) {
        return {
            isNumber: true,
            value: toNumber
        }
    }
    return {
        isNumber: false,
        value: null
    }
}

function createCell([x, y, cellData]: [x: number, y: number, cellData?: string]) {
    const data = cellData || matrix?.[x]?.[y]
    return {
        x,
        y,
        data,
        isANumber: isANumber(data)
    }
}

const INLINE_SEPARATOR = "X"
const ratios: any[] = []
matrix.forEach((rowData, rowIndex) => {


    rowData.forEach((columnData, columnIndex) => {

        const cell = {
            ...createCell([rowIndex, columnIndex, columnData]),
            surround: {
                "TOP-LEFT": createCell([rowIndex - 1, columnIndex - 1]),
                "TOP": createCell([rowIndex - 1, columnIndex]),
                "TOP-RIGHT": createCell([rowIndex - 1, columnIndex + 1]),
                "LEFT": createCell([rowIndex, columnIndex - 1]),
                "RIGHT": createCell([rowIndex, columnIndex + 1]),
                "BOTTOM-LEFT": createCell([rowIndex + 1, columnIndex - 1]),
                "BOTTOM": createCell([rowIndex + 1, columnIndex]),
                "BOTTOM-RIGHT": createCell([rowIndex + 1, columnIndex + 1]),
            },
        }
        if (cell.data !== "*") {
            return;

        }
        const topLeft = cell.surround['TOP-LEFT'].isANumber.isNumber
        const top = cell.surround['TOP'].isANumber.isNumber
        const topRight = cell.surround['TOP-RIGHT'].isANumber.isNumber
        const bottomLeft = cell.surround['BOTTOM-LEFT'].isANumber.isNumber
        const bottom = cell.surround['BOTTOM'].isANumber.isNumber
        const bottomRight = cell.surround['BOTTOM-RIGHT'].isANumber.isNumber
        const right = cell.surround['RIGHT'].isANumber.isNumber
        const left = cell.surround['LEFT'].isANumber.isNumber

        let matchingGears = 0
        let gears = []

        if (left) {
            matchingGears += 1

            const buffer: any[] = []
            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex, columnIndex - i])
                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)

                }
            }

            gears.push(buffer)
        }
        if (right) {
            matchingGears += 1
            const buffer: any[] = []
            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }

            gears.push(buffer)
        }
        if (topLeft && top && topRight) {
            matchingGears += 1
            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex - 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            buffer.push(cell.surround.TOP.isANumber.value)

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex - 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)

        }
        if (topLeft && top && !topRight) {
            matchingGears += 1

            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex - 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            buffer.push(cell.surround.TOP.isANumber.value)

            gears.push(buffer)
        }
        if (topLeft && !top && topRight) {
            matchingGears += 2

            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex - 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            buffer.push(INLINE_SEPARATOR)

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex - 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)
        }
        if (topLeft && !top && !topRight) {
            matchingGears += 1

            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex - 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            gears.push(buffer)
        }
        if (!topLeft && top && topRight) {
            matchingGears += 1
            const buffer: any[] = []

            buffer.push(cell.surround.TOP.isANumber.value)

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex - 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)
        }
        if (!topLeft && !top && topRight) {
            matchingGears += 1
            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex - 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)
        }
        if (!topLeft && top && !topRight) {
            matchingGears += 1
            const buffer: any[] = []

            buffer.push(cell.surround.TOP.isANumber.value)
            gears.push(buffer)
        }
        if (bottomLeft && bottom && bottomRight) {
            matchingGears += 1
            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex + 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            buffer.push(cell.surround.BOTTOM.isANumber.value)

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex + 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)
        }
        if (bottomLeft && bottom && !bottomRight) {
            matchingGears += 1

            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex + 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            buffer.push(cell.surround.BOTTOM.isANumber.value)

            gears.push(buffer)
        }
        if (bottomLeft && !bottom && bottomRight) {
            matchingGears += 2

            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex + 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            buffer.push(INLINE_SEPARATOR)

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex + 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)
        }
        if (bottomLeft && !bottom && !bottomRight) {
            matchingGears += 1

            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {
                const newCell = createCell([rowIndex + 1, columnIndex - i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.unshift(newCell.isANumber.value)
                }
            }

            gears.push(buffer)
        }
        if (!bottomLeft && bottom && bottomRight) {
            matchingGears += 1
            const buffer: any[] = []

            buffer.push(cell.surround.BOTTOM.isANumber.value)

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex + 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)
        }
        if (!bottomLeft && !bottom && bottomRight) {
            matchingGears += 1
            const buffer: any[] = []

            for (let i = 1; i <= 5; i++) {

                const newCell = createCell([rowIndex + 1, columnIndex + i])

                if (!newCell.isANumber.isNumber) {
                    break
                } else if (newCell.isANumber.isNumber) {
                    buffer.push(newCell.isANumber.value)
                }

            }
            gears.push(buffer)
        }
        if (!bottomLeft && bottom && !bottomRight) {
            matchingGears += 1
            const buffer: any[] = []

            buffer.push(cell.surround.BOTTOM.isANumber.value)
            gears.push(buffer)
        }


        const pattern = `
${cell.surround["TOP-LEFT"].data} ${cell.surround.TOP.data} ${cell.surround["TOP-RIGHT"].data}
${cell.surround["LEFT"].data} ${cell.data} ${cell.surround["RIGHT"].data}
${cell.surround["BOTTOM-LEFT"].data} ${cell.surround["BOTTOM"].data} ${cell.surround["BOTTOM-RIGHT"].data}
        `

        if (matchingGears === 2) {
            if (gears.length === 1) {
                const [left, right] = gears[0].join("").split(INLINE_SEPARATOR)
                const ratio = Number(left) * Number(right)
                ratios.push(ratio)
            } else if (gears.length === 2) {
                const [left, right] = gears
                const ratio = Number(left.join("")) * Number(right.join(""))
                ratios.push(ratio)
            }


        }


    })

}
)

const sumOfGearRatios = ratios.reduce((acc, curr) => acc + curr)
console.log(sumOfGearRatios)
