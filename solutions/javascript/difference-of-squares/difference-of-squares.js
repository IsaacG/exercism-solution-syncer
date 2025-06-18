export class Squares {
  constructor(number) {
    this.number = number;
  }

  get sumOfSquares() {
    let sum = 0;
    for (let i = 1; i <= this.number; i++)
      sum += i ** 2;
    return sum;
  }

  get squareOfSum() {
    let sum = 0;
    for (let i = 1; i <= this.number; i++)
      sum += i;
    return sum ** 2;
  }

  get difference() {
    return this.squareOfSum - this.sumOfSquares;
  }
}
