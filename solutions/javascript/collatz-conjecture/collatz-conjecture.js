export const steps = (num) => {
  if (num < 1)
    throw new Error('Only positive integers are allowed');
  let steps = 0;
  while (num > 1) {
    steps++;
    if (num % 2 === 0)
      num /= 2;
    else
      num = num * 3 + 1;
  };
  return steps;
};
