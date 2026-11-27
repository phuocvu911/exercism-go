export function decodedResistorValue(parts: string[]): string {
  const [first, second, multiplier] = parts.map(c => COLORS.indexOf(c))
  let res = (first*10 + second) * 10 ** multiplier

  let count = 0
  while (res > 0 && res%1000 === 0){
    count++
    res /= 1000
  }
  return `${res} ${PREFIX[count]}ohms`
}

const PREFIX: string[] = ['', 'kilo','mega','giga']
const COLORS: string[] = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
];
