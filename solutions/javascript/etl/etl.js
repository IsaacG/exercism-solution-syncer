export const transform = (values) => {
  const transformed = {}
  for (const value in values) {
    const numeric = Number(value);
    for (const letter of values[value]) {
      transformed[letter.toLowerCase()] = numeric;
    };
  };
  return transformed;
};
