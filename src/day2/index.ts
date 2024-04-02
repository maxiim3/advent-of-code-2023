import * as fs from "fs";

//12 red cubes, 13 green cubes, and 14 blue cubes
const MAX_GREEN = 13
const MAX_BLUE = 14
const MAX_RED = 12

const file = fs.readFileSync("src/day2/input.txt", { encoding: 'utf-8' });
const listOfGames = file.split("\n")
// remove the extra line
listOfGames.pop()

type Color = "red" | "green" | "blue"
type Cube = {
    color: Color,
    number: number
}
type Throw = Cube[]

class Game {
    gameID: number;
    throws: Throw[]
    blueCubes: number
    greenCubes: number
    redCubes: number
    isEligible: boolean;
    power: number

    constructor(line: string) {
        const [game, other] = line.split(": ");
        this.gameID = parseInt(game.replace("Game ", ""));
        this.blueCubes = 0
        this.redCubes = 0
        this.greenCubes = 0;
        this.power = 0
        this.isEligible = true;

        this.throws = other
            .split("; ")
            .map(singleThrow => {
                return singleThrow
                    .split(", ")
                    .map(item => {
                        let numberOfColor = parseInt(item);
                        let color: Color = item.replace(`${numberOfColor} `, "") as Color;
                        switch (color) {
                            case "blue":
                                if (numberOfColor > this.blueCubes) {
                                    this.blueCubes = numberOfColor
                                }

                                if (numberOfColor > MAX_BLUE) {
                                    this.isEligible = false;
                                }
                                break
                            case "red":
                                if (numberOfColor > this.redCubes) {
                                    this.redCubes = numberOfColor
                                }
                                if (numberOfColor > MAX_RED) {
                                    this.isEligible = false;
                                }
                                break
                            case "green":
                                if (numberOfColor > this.greenCubes) {
                                    this.greenCubes = numberOfColor
                                }

                                if (numberOfColor > MAX_GREEN) {
                                    this.isEligible = false;
                                }
                                break
                        }

                        return {
                            color,
                            number: numberOfColor
                        }
                    })
            })

        this.power = this.blueCubes * this.redCubes * this.greenCubes;

    }
}

const games = listOfGames.map(game => new Game(game));

console.log(`total games : ${games.length}`)
let eligibleGames = games.filter(game => game.isEligible)
console.log(`Valid Games : ${eligibleGames.length}`)
let total: number = 0
for (const { gameID } of eligibleGames) {
    total += gameID
}
console.log("total of eligible games", total)

console.log(listOfGames[2])

const totalPower = games.reduce((acc, currentValue) => acc + currentValue.power, 0)
console.log("total Power is : ", totalPower)
