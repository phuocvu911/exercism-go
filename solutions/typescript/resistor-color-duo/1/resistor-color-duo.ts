export function decodedValue(parts:string[]): number {
  return COLORS.indexOf(parts[0])*10 + COLORS.indexOf(parts[1])
}

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