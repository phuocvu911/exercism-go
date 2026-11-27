export function decodedResistorValue(parts: string[]): string {
  let res = COLORS.indexOf(parts[0])*10 + COLORS.indexOf(parts[1])
  for (let i = 0; i < COLORS.indexOf(parts[2]); i++) {
    res *= 10
  }

  let count = 0
  while (res > 0){
    if (res%1000 === 0){
      count++
    } else{
      break
    }
    res /= 1000
  }
  return `${res} ${prefix[count]}ohms`
}

const prefix: string[] = ['', 'kilo','mega','giga']
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
