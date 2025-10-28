/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Get the first card in the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card} the first card in the deck
 */
export function getFirstCard(deck) {
  const [first, ...rest] = deck;
  return first
}

/**
 * Get the second card in the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card} the second card in the deck
 */
export function getSecondCard(deck) {
  const [first, second, ...rest] = deck;
  return second
}

/**
 * Switch the position of the first two cards in the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card[]} new deck with reordered cards
 */
export function swapTwoCards(deck) {
  const [first, second] = deck;
  return [second, first]
}

/**
 * Rotate (shift) the position of the three cards (by one place)
 *
 * @param {[Card, Card, Card]} deck
 *
 * @returns {[Card, Card, Card]} new deck with the 3 cards shifted by one position
 */
export function shiftThreeCardsAround(deck) {
  const [first, second, third] = deck;
  return [second, third, first]
}

export function swapNamedPile(piles) {
  // 🪄 Don't break the magic.
  // 🚨 Do NOT use piles.chosen or piles.disregarded.
  // 🚨 Do NOT touch the next line or Elyse will accidentally reveal the trick.
  const {chosen, disregarded} = piles;
  return {chosen: disregarded, disregarded: chosen};
}

/**
 * Grab the chosen pile from the available piles
 *
 * @param {{ chosen: Card[], disregarded: Card[] }} piles
 *
 * @returns {Card[]} the pile named chosen
 */
export function pickNamedPile(piles) {
  return piles.chosen;
}

/**
 * Put the top card of the given deck into a separate discard pile
 *
 * @param {Card[]} deck
 *
 * @returns {[Card, Card[]]} the top card of the given
 * deck and a new deck containing all the other cards
 */
export function discardTopCard(deck) {
  const [first, ...rest] = deck;
  return [first, rest]
}

/** @type {Card[]} **/
const FACE_CARDS = ['jack', 'queen', 'king'];

/**
 * Insert face cards into the given deck
 *
 * @param {Card[]} deck
 *
 * @returns {Card[]} new deck where the second,
 * third, and fourth cards are the face cards
 */
export function insertFaceCards(deck) {
  const [first, ...rest] = deck;
  return [first, ...FACE_CARDS, ...rest]
}