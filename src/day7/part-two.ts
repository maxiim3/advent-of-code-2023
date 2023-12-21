import * as fs from "fs";
import { log } from "util";

/////////////////////// Utils ///////////////////////////
function readFile() {
    let file = fs.readFileSync("./src/day7/sample.txt", {encoding: "utf-8"})

    console.log(file)
    return file

}

/////////////////// Business logic ///////////////

export type Label = "A" | "K" | "Q" | "J" | "T" | "9" | "8" | "7" | "6" | "5" | "4" | "3" | "2"
export type HandOfCards = [Label, Label, Label, Label, Label]

/**
 * Converts the label of a card to a value
 * @param label
 */
export function getLabelValue(label: Label) {
    if (label === "A") return 14
    else if (label === "K") return 13
    else if (label === "Q") return 12
    else if (label === "T") return 10
    else if (label === "J") return 1
    else return parseInt(label)
}



/**
 * Set a value to the ranking of a hand
 */
export const HandRank = {
    pair: 1,
    doublePair: 2,
    brelan: 3,
    fullHouse: 4,
    carre: 5,
    clubDesCinq: 6

}

/**
 * Sort cards inside a hand in order to use the regex on concurrent repetitive values
 * @param hand
 */
export function sortLabelsInHand(hand: HandOfCards) {
    return hand.toSorted((a, b) => getLabelValue(a) > getLabelValue(b) ? -1 : 1)
}


/**
 * convert a hand to its ranking
 * @param match
 * @link HandRank
 */
export function getHandsRanking(match?: string) {
    switch (match?.length) {
        case 5:
            return HandRank.clubDesCinq;
        case 4:
            return HandRank.carre;
        case 3 :
            return HandRank.brelan;
        case 2 :
            return HandRank.pair;
            default:
        return null
    }
}

export const regExPattern = /(.)\1+/g

/**
 * Parse the hand and returns its value / Strength
 * @param hand
 */
export function parseHand(hand: HandOfCards): number {
    const sortedHand = sortLabelsInHand(hand)

    const matchedPattern = sortedHand
        .join("")
        .match(regExPattern)

    
    console.log("matchedPattern " ,hand, matchedPattern);
    
    if (!matchedPattern) {
        return 0
    }
   
    let [first, second] = matchedPattern

   // has a match : has J's, replace strongest by xJ's : has J matched, replace em to otherone 
    let rankScoreFirstHand = getHandsRanking(first)

    if(!rankScoreFirstHand) return 0

    let rankScoreSecondHand = getHandsRanking(second)

    if(!second || !rankScoreSecondHand){
        return rankScoreFirstHand;
    }
    
    switch (true) {
        case isAFullHouse(rankScoreFirstHand , rankScoreSecondHand ) :
            return HandRank.fullHouse;
        case isADoublePair(rankScoreFirstHand , rankScoreSecondHand ):
            return HandRank.doublePair;
    }
    return (rankScoreFirstHand || rankScoreSecondHand) || 0
}

function isAFullHouse(first:number, second:number) {
    return  (first === HandRank.pair && second === HandRank.brelan) || (second === HandRank.pair && first === HandRank.brelan)
}

function isADoublePair(first:number, second:number) {
    return  ( first === HandRank.pair ) && ( second === HandRank.pair );
}

////////////////// Main Implementation /////////////////////
function run() {

    // extract informations from the sample_file
    const file = readFile();
    const lines = file.split("\n")
    lines.pop(); // remove the extra line

    // Get the sets of [Hands of cards, bid][]
    const setsOfCards: [HandOfCards, number][] = lines
    .map((ln) => ln.split(" "))
    .map(([data, bid]) => [data .split("") as HandOfCards, parseInt(bid)])

    // Sort the the sets of Hand
    const sortedSetsByRank: [HandOfCards, number][] = setsOfCards
        .toSorted((
            current,
            next
        ) => {
            // Destructure hands / bid
            const [currentHand, currentBid] = current;
            const [nextHand, nextBid] = next;

            // Parse each
            const currentStrength = parseHand(currentHand)
            const nextStrength = parseHand(nextHand)

            // return sorted values
            if (currentStrength !== nextStrength) {
                return currentStrength > nextStrength ? 1 : -1

            } else {
                // console.log(currentHand, currentStrength, nextStrength, nextHand)
                // Parse and compare each label
                // Get the index where values of both hands do not match anymore
                let index = 0


                for (let i = 0; i < currentHand.length; i++) {
                    let currentValue = getLabelValue(currentHand[i])
                    let nextValue = getLabelValue(nextHand[i])

                    if (currentValue !== nextValue) {
                        break;
                    } else {
                        index++
                    }
                }
                // console.log(index)
                // console.log(getLabelValue(currentHand[index]) > getLabelValue(nextHand[index]))
                // compare those values at the index
                return getLabelValue(currentHand[index]) > getLabelValue(nextHand[index]) ? 1 : -1
            }
        })

    let total = 0
    for (let i = 0; i < sortedSetsByRank.length; i++) {
        const rank = i + 1
        total += sortedSetsByRank[i][1] * rank

    }
    // console.log(sortedSetsByRank, total)
}

run()

