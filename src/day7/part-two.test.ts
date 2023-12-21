import { describe, expect, test } from "bun:test";
import {
    getHandsRanking,
    getLabelValue,
    HandRank,
    parseHand,
    regExPattern,
    sortLabelsInHand,
} from "./part-two.ts";

describe("When parsing a hand of 5 cards", () => {
    test("getLabelValue Should return the corresponding value", () => {
        expect(getLabelValue("2")).toEqual(2);
        expect(getLabelValue("3")).toEqual(3);
        expect(getLabelValue("4")).toEqual(4);
        expect(getLabelValue("5")).toEqual(5);
        expect(getLabelValue("6")).toEqual(6);
        expect(getLabelValue("7")).toEqual(7);
        expect(getLabelValue("8")).toEqual(8);
        expect(getLabelValue("9")).toEqual(9);
        expect(getLabelValue("J")).toEqual(1);
        expect(getLabelValue("Q")).toEqual(12);
        expect(getLabelValue("K")).toEqual(13);
        expect(getLabelValue("A")).toEqual(14);
    });
    describe("Sort labels in a hand", () => {
        test("it should return an Array", () => {
            expect(sortLabelsInHand(["2", "J", "3", "2", "Q"])).toBeArray();
        });
        test("The value should be sorted properly from highest to lowest", () => {
            expect(sortLabelsInHand(["2", "J", "3", "2", "Q"])).toStrictEqual([
                "Q",
                "3",
                "2",
                "2",
                "J",
            ]);
        });
    });
    describe("The Regex, ", () => {
        test("should return the uniq matching item", () => {
            expect("AAA89".match(regExPattern)).toEqual(["AAA"]);
        });
        test("should return the two matching items", () => {
            expect("AAA88".match(regExPattern)).toEqual(["AAA", "88"]);
        });
        test("J should be consider as a Joker, matching the value of the highest possibility", () => {
            expect("AAAJ8".match(regExPattern)).toEqual(["AAAA"]);
            expect("AAJKK".match(regExPattern)).toEqual(["AAA", "KK"]);
        });
    });

    describe("getHandsRanking is called", () => {
        test("It should return a Pair", () => {
            expect(getHandsRanking("HH")).toBe(HandRank.pair);
        });
    });

    describe("When parseHand is called", () => {
        test("A full house hand should return the number 4", () => {
            expect(parseHand(["A", "7", "A", "7", "A"])).toStrictEqual(4);
        });
        test("A hand with no match should return the number 0", () => {
            expect(parseHand(["9", "J", "2", "7", "A"])).toStrictEqual(1); // A Pair of "AA"
        });
    });
});
