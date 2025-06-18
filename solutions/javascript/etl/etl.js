export const transform = (values) => {
  const transformed = {}
  Object.entries(values).forEach(([value, letters]) => {
    const numeric = Number(value);
    letters.forEach((letter) => {
      transformed[letter.toLowerCase()] = numeric;
    });
  });
  return transformed;
};
