export const isPangram = (data) => {
  const seen = new Set();
  for (const char of [...data.toLowerCase()]) {
    if (char.match(/[a-z]/)) {
      seen.add(char);
    }
  }
  return seen.size === 26;
};
